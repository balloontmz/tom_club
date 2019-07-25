package assistant

import (
	"context"
	"net/url"
)

var goodsAPI = &apiConfig{
	host: "https://api.taokezhushou.com",
	path: "/api/v1/all",
}

//GoodsResponse represents a Distance Matrix API response.
type GoodsResponse struct {
	APIName string `json:"api_name"`
}

//GoodsRequest  test
type GoodsRequest struct {
	APIKey string
	Page   string
}

//Goods makes a Distance Matrix API request
func (c *Client) Goods(ctx context.Context, r *GoodsRequest) (*GoodsResponse, error) {

	var response GoodsResponse

	if err := c.getJSON(ctx, goodsAPI, r, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *GoodsRequest) params() url.Values {
	q := make(url.Values)

	if r.APIKey != "" {
		q.Set("apikey", r.APIKey)
	}
	if r.Page != "" {
		q.Set("page", r.Page)
	}
	return q
}
