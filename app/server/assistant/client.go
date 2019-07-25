package assistant

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/gommon/log"
	"golang.org/x/time/rate"
)

// ClientOption is the type of constructor options for NewClient(...).
type ClientOption func(*Client) error

// Client may be used to make requests to the Google Maps WebService APIs
type Client struct {
	httpClient        *http.Client
	apiKey            string
	baseURL           string
	requestsPerSecond int
	rateLimiter       *rate.Limiter
}

var defaultRequestsPerSecond = 50

// NewClient constructs a new Client which can make requests to the Google Maps
// WebService APIs.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{requestsPerSecond: defaultRequestsPerSecond}
	WithHTTPClient(&http.Client{})(c)
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if c.apiKey == "" {
		return nil, errors.New("maps: API Key or Maps for Work credentials missing")
	}

	// 限制访问次数
	if c.requestsPerSecond > 0 {
		c.rateLimiter = rate.NewLimiter(rate.Limit(c.requestsPerSecond), c.requestsPerSecond)
	}

	return c, nil
}

// WithHTTPClient configures a Maps API client with a http.Client to make requests
// 设置客户端
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithAPIKey 设置api key
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = apiKey
		return nil
	}
}

// WithBaseURL 设置基础url
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// WithRateLimit configures the rate limit for back end requests. Default is to
// 设置访问速度
func WithRateLimit(requestsPerSecond int) ClientOption {
	return func(c *Client) error {
		c.requestsPerSecond = requestsPerSecond
		return nil
	}
}

type apiConfig struct {
	host string
	path string
}

// 实现了该接口的结构即属于该接口
type apiRequest interface {
	params() url.Values
}

func (c *Client) awaitRateLimiter(ctx context.Context) error {
	if c.rateLimiter == nil {
		return nil
	}
	return c.rateLimiter.Wait(ctx)
}

func (c *Client) get(ctx context.Context, config *apiConfig, apiReq apiRequest) (*http.Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	host := config.host
	if c.baseURL != "" {
		host = c.baseURL
	}
	req, err := http.NewRequest("GET", host+config.path, nil)
	if err != nil {
		return nil, err
	}
	// log.Print("在此处", req)
	q, err := c.generateAuthQuery(config.path, apiReq.params())
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = q
	log.Print("请求体为：", req)
	return c.do(ctx, req)
}

func (c *Client) post(ctx context.Context, config *apiConfig, apiReq interface{}) (*http.Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	host := config.host
	if c.baseURL != "" {
		host = c.baseURL
	}

	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", host+config.path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	q, err := c.generateAuthQuery(config.path, url.Values{})
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = q
	return c.do(ctx, req)
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(req.WithContext(ctx))
}

func (c *Client) getJSON(ctx context.Context, config *apiConfig, apiReq apiRequest, resp interface{}) error {
	httpResp, err := c.get(ctx, config, apiReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	// b, _ := ioutil.ReadAll(httpResp.Body)
	// var data interface{}
	// err = json.Unmarshal(b, &data)
	// log.Print("返回结果为：", data)
	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (c *Client) postJSON(ctx context.Context, config *apiConfig, apiReq interface{}, resp interface{}) error {
	httpResp, err := c.post(ctx, config, apiReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

type binaryResponse struct {
	statusCode  int
	contentType string
	data        io.ReadCloser
}

func (c *Client) getBinary(ctx context.Context, config *apiConfig, apiReq apiRequest) (binaryResponse, error) {
	httpResp, err := c.get(ctx, config, apiReq)
	if err != nil {
		return binaryResponse{}, err
	}

	return binaryResponse{httpResp.StatusCode, httpResp.Header.Get("Content-Type"), httpResp.Body}, nil
}

func (c *Client) generateAuthQuery(path string, q url.Values) (string, error) {
	if c.apiKey != "" {
		q.Set("app_key", c.apiKey)
		return q.Encode(), nil
	}
	return "", errors.New("maps: API Key missing")
}
