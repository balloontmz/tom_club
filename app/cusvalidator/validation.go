package cusvalidator

import "gopkg.in/go-playground/validator.v9"

// //CheckIndentify for request
// type CheckIndentify struct {
// 	Key  string `query:"api_key" validate:"required,max=30,min=5"`
// 	Time int64  `query:"timestamp" validate:"required,max=9999999999,min=1000000000"`
// 	Sign string `query:"sign" validate:"required,min=44,max=64"`
// }

// //PostEmail post 体的验证
// type PostEmail struct {
// 	ToEmail    []string `form:"to_email[]" json:"to_email" validate:"required,max=30,min=1,dive,email"` // dive 进入切片内一层进行验证
// 	Subject    string   `form:"subject" json:"subject" validate:"required,max=30,min=5"`
// 	HTMLBody   string   `form:"html_body" json:"html_body" validate:"required"`
// 	Attachment string   `form:"attachment"`
// }

// //PostClient for request
// type PostClient struct {
// 	Key      string `form:"api_key" json:"api_key" validate:"required,max=30,min=5"`
// 	Secret   string `form:"secret" json:"secret" validate:"required,max=30,min=5"`
// 	EmailPre string `form:"email_pre" json:"email_pre" validate:"required,max=30,min=2"`
// }

// Validator 自定义验证器
type Validator struct {
	Validator *validator.Validate
}

// Validate
func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// NewPostEmailExample 创建一个 PostEmail 测试对象
// func NewPostEmailExample() PostEmail {
// 	var postemail PostEmail
// 	postemail.ToEmail = []string{"15111171986@163.com", "balloontmz@gmail.com"}
// 	postemail.Subject = "这是一个测试主题"
// 	postemail.HTMLBody = "这是一个测试网页体"
// 	return postemail
// }
