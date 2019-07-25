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
	Total int      `json:"total"`
	Data  []*Goods `json:"data"`
}

//Goods 商品
type Goods struct {
	GoodsID         int     `json:"goods_id"`
	GoodsPic        string  `json:"goods_pic"`
	GoodsLongPic    string  `json:"goods_long_pic"`
	GoodsTitle      string  `json:"goods_title"`
	GoodsShortTitle string  `json:"goods_short_title"`
	GoodsIntro      string  `json:"goods_intro"`
	CateID          int     `json:"goods_cate_id"`
	GoodsPrice      string  `json:"goods_price"`
	SaleNum         int     `json:"goods_sale_num"`
	CommissionRate  float64 `json:"commission_rate"`
	CouponID        string  `json:"coupon_id"`
	CouponAmount    string  `json:"coupon_apply_amount"`
	StartTime       string  `json:"coupon_start_time"`
	EndTime         string  `json:"coupon_end_time"`
	IsMall          int     `json:"is_tmall"`
	JuHuaSuan       int     `json:"juhuasuan"`
	TaoQiangGou     int     `json:"taoqianggou"`
	YunFeiXian      int     `json:"yunfeixian"`
	JinPai          int     `json:"jinpai"`
	JiYouJia        int     `json:"jiyoujia"`
	HaiTao          int     `json:"haitao"`
	DSR             float64 `json:"dsr"`
	ChaoShi         int     `json:"chaoshi"`
	PaiXiaJian      string  `json:"paixiajian"`
	BuyPrice        string  `json:"buy_price"`
}

//GoodsRequest  test
type GoodsRequest struct {
	APIKey string
	Page   string
}

//Goods makes a Distance Matrix API request
func (c *Client) Goods(ctx context.Context, r *GoodsRequest) (*GoodsResponse, error) {

	var response struct {
		commonResponse
		GoodsResponse
	}

	if err := c.getJSON(ctx, goodsAPI, r, &response); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.GoodsResponse, nil
}

func (r *GoodsRequest) params() url.Values {
	q := make(url.Values)

	if r.APIKey != "" {
		q.Set("app_key", r.APIKey)
	}
	if r.Page != "" {
		q.Set("page", r.Page)
	}
	return q
}
