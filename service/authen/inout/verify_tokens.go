package inout

type VerifyTokenResponse struct {
	Verify bool   `json:"verify"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
