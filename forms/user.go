package forms

type UserSignup struct {
	Phone   string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	UUID string `json:"uuid"`
}

type UserLogin struct {
    Phone     string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserToken struct {
    Token string `json:"token" binding:"required"`
}
