package validates

type LoginInput struct {
	UserCode string `form:"user_code" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterInput struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username       string   `validate:"gt=0"`
	// 同上
	PasswordNew    string   `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat string   `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email          string   `validate:"email"`
}