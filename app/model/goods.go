package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

//Goods 商品
//Some        *time.Time
//Some        sql.NullInt64
//IgnoreMe        int     `gorm:"-"`               // ignore this field
type Goods struct {
	gorm.Model
	GoodsID         sql.NullInt64 `gorm:"column:goods_id;unique_index" json:"-"` // 此处注意如果该变量为存在 valid 需要设置为 true
	GoodsPic        string        `gorm:"column:goods_pic" json:"goods_pic"`
	GoodsLongPic    string        `gorm:"column:goods_long_pic" json:"goods_long_pic"`
	GoodsTitle      string        `gorm:"column:goods_title" json:"goods_title"`
	GoodsShortTitle string        `gorm:"column:goods_short_title" json:"goods_short_title"`
	GoodsIntro      string        `gorm:"column:goods_intro" json:"goods_intro"`
	CateID          sql.NullInt64 `gorm:"column:goods_cate_id" json:"-"`
	GoodsPrice      string        `gorm:"column:goods_price" json:"goods_price"`
	SaleNum         int           `gorm:"column:goods_sale_num" json:"goods_sale_num"`
	CommissionRate  float64       `gorm:"column:commission_rate" json:"commission_rate"`
	CouponID        string        `gorm:"column:coupon_id" json:"coupon_id"`
	CouponAmount    string        `gorm:"column:coupon_apply_amount" json:"coupon_apply_amount"`
	StartTime       *time.Time    `gorm:"column:coupon_start_time" json:"-"`
	EndTime         *time.Time    `gorm:"column:coupon_end_time" json:"-"`
	IsMall          int           `gorm:"column:is_tmall;type:tinyint(4);default:0" json:"is_tmall"`
	JuHuaSuan       int           `gorm:"column:juhuasuan;type:tinyint(4);default:0" json:"juhuasuan"`
	TaoQiangGou     int           `gorm:"column:taoqianggou;type:tinyint(4);default:0" json:"taoqianggou"`
	YunFeiXian      int           `gorm:"column:yunfeixian;type:tinyint(4);default:0" json:"yunfeixian"`
	JinPai          int           `gorm:"column:jinpai;type:tinyint(4);default:0" json:"jinpai"`
	JiYouJia        int           `gorm:"column:jiyoujia;type:tinyint(4);default:0" json:"jiyoujia"`
	HaiTao          int           `gorm:"column:haitao;type:tinyint(4);default:0" json:"haitao"`
	DSR             float64       `gorm:"column:dsr" json:"dsr"`
	ChaoShi         int           `gorm:"column:chaoshi;type:tinyint(4);default:0" json:"chaoshi"`
	PaiXiaJian      string        `gorm:"column:paixiajian" json:"paixiajian"`
	BuyPrice        string        `gorm:"column:buy_price" json:"buy_price"`
}

//TableName Set Goods's table name to be `goods`
func (Goods) TableName() string {
	return "goods"
}
