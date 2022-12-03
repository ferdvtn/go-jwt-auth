package domains

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	Name     string `json:"name"`
}
