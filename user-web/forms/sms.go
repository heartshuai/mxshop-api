package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻,自定义validate
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`
}
