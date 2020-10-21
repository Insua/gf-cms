package auth

type LoginRequest struct {
	Name     string `p:"name"`
	Password string `p:"password"`
}

var (
	LoginRequestRules = map[string]string{
		"Name":     "required",
		"Password": "required|login:name",
	}
	LoginRequestMessages = map[string]interface{}{
		"Name": map[string]string{
			"required": "用户名不能为空",
		},
		"Password": map[string]string{
			"required": "密码不能为空",
		},
	}
)
