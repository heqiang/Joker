package swag

type User struct {
	UserName  string `json:"username" gorm:"type:varchar(200);column:username"  binding:"required,min=3,max=20"`
	PassWord  string `json:"password" gorm:"type:varchar(200);column:password"  binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
