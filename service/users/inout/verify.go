package inout

type SetVerifyInput struct {
	Token string `json:"token"`
	Email string `json:"email"`
} // @Name SetVerifyInput

type SetVerifyResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
} // @Name SetVerifyResponse
