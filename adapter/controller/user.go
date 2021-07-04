package controller

type RequestUserPost struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type RequestUserUpdate struct {
	Id       string `json:"id"`
	Username string `json:"id"`
}

type RequestUserDelete struct {
	Id string `json:"id"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"id"`
}
