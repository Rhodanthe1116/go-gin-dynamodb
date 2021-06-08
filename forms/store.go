package forms

type StoreSignup struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
	UUID string `json:"uuid"`
}

type StoreLogin struct {
    Phone     string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StoreProfile struct {
	Name     string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
    QrCode string `json:"qrcode" binding:"required"`
    UUID string `json:"uuid"`
}

type StoreToken struct {
    Token string `json:"token" binding:"required"`
}
