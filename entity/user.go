package entity

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Age      uint32 `json:"age"`
	Role     uint32 `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
