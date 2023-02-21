package endpoint

type UserRequest struct {
	Id int `json:"id"`
}

type UserResponse struct {
	Name string `json:"name"`
}
