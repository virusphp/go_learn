package dto

type UserDTO struct {
	Nickname   string `json:"nickname"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Pic        string `json:"pic"`
	Otoritas   uint32 `json:"otoritas"`
	Status     string `gorm:"column:status"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ListUserDTO struct {
	Search *string `json:"search"`
	Limit  string  `json:"limit"`
	Page   string  `json:"page"`
	Order  string  `json:"order"`
}

type Response_login struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}
