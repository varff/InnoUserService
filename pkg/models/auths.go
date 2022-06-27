package models

type LoginModel struct {
	Phone int32  `json:"phone"`
	Pass  string `json:"password"`
}

type RegisterModel struct {
	Phone int32  `json:"phone"`
	Pass  string `json:"password"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
