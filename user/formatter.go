package user

type UserFormatter struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return formatter
}
