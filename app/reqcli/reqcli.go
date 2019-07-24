package reqcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.taokezhushou.com/api/v1/all"
const appKey string = "b08fe2c272e19a9b"

//Client 封装http客户端
type Client struct {
	PlaceHolder interface{}
}

//NewClient 创建请求客户端
func NewClient() *Client {
	return &Client{}
}

//Result 封装返回结果
type Result interface{}

//AddTodo 没用
func (s *Client) AddTodo(todo *Result) error {
	url := fmt.Sprintf(baseURL + "/todos")
	fmt.Println(url)
	j, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

//Get 请求方法
func (s *Client) Get() (*Result, error) {
	url := fmt.Sprintf(baseURL+"?app_key=%s", appKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Result
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

//doRequest 请求统一方法
func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
