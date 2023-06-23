package handler

type UserResponse struct {
	ID    uint   `json:"id" copier:"must"`
	Name  string `json:"name" copier:"must"`
	Email string `json:"email" copier:"must"`
}
