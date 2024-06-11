package logout

type LogoutReq struct {
	Secret string `json:"token" binding:"required"`
}
