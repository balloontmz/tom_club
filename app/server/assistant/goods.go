package assistant

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	"tom_club/app/model"
)

var goodsAPI = &apiConfig{
	host: "https://api.taokezhushou.com",
	path: "/api/v1/all",
}

//Response 响应内容
type Response struct {
	commonResponse
	GoodsResponse
}

//GoodsResponse represents a Distance Matrix API response.
type GoodsResponse struct {
	Total int `json:"total"`
	Data  []*struct {
		model.Goods
		temp
	} `json:"data"`
}

//Goods 商品
// type Goods struct {
// 	GoodsID         sql.NullInt64 `json:"-"`
// 	GoodsPic        string        `json:"goods_pic"`
// 	GoodsLongPic    string        `json:"goods_long_pic"`
// 	GoodsTitle      string        `json:"goods_title"`
// 	GoodsShortTitle string        `json:"goods_short_title"`
// 	GoodsIntro      string        `json:"goods_intro"`
// 	CateID          sql.NullInt64 `json:"-"`
// 	GoodsPrice      string        `json:"goods_price"`
// 	SaleNum         int           `json:"goods_sale_num"`
// 	CommissionRate  float64       `json:"commission_rate"`
// 	CouponID        string        `json:"coupon_id"`
// 	CouponAmount    string        `json:"coupon_apply_amount"`
// 	StartTime       *time.Time    `json:"-"`
// 	EndTime         *time.Time    `json:"-"`
// 	IsMall          int           `json:"is_tmall"`
// 	JuHuaSuan       int           `json:"juhuasuan"`
// 	TaoQiangGou     int           `json:"taoqianggou"`
// 	YunFeiXian      int           `json:"yunfeixian"`
// 	JinPai          int           `json:"jinpai"`
// 	JiYouJia        int           `json:"jiyoujia"`
// 	HaiTao          int           `json:"haitao"`
// 	DSR             float64       `json:"dsr"`
// 	ChaoShi         int           `json:"chaoshi"`
// 	PaiXiaJian      string        `json:"paixiajian"`
// 	BuyPrice        string        `json:"buy_price"`
// }

type temp struct {
	St       string `json:"coupon_start_time"`
	Et       string `json:"coupon_end_time"`
	TGoodsID int64  `json:"goods_id"`
	TCateID  int64  `json:"goods_cate_id"`
}

//GoodsRequest  test
type GoodsRequest struct {
	APIKey string
	Page   string
}

//Goods makes a Distance Matrix API request
func (c *Client) Goods(ctx context.Context, r *GoodsRequest) (*GoodsResponse, error) {

	var response Response

	if err := c.getJSON(ctx, goodsAPI, r, &response); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}
	response.preReturn()

	return &response.GoodsResponse, nil
}

// 响应数据返回前
func (response Response) preReturn() {
	for _, v := range response.GoodsResponse.Data {
		// 将字符串转换为时间格式
		st, err := time.Parse("2006-01-02 15:04:05", v.St)
		v.StartTime = &time.Time{}
		v.EndTime = &time.Time{}
		if err == nil {
			*v.StartTime = st
		}
		et, err := time.Parse("2006-01-02 15:04:05", v.Et)
		if err == nil {
			*v.EndTime = et
		}
		// 将int转换为数据库格式
		v.GoodsID = sql.NullInt64{Int64: v.TGoodsID, Valid: false}
		v.CateID = sql.NullInt64{Int64: v.TCateID, Valid: false}
	}
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
