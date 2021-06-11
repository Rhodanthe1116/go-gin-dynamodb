package forms

type RecordStoreId struct {
	StoreId   string `json:"store_id" binding:"required"`
}

type Record struct {
    UserId string `json:"user_id"`
    StoreId string `json:"store_id"`
    Time int64 `json:"time"`
    StoreName string `json:"store_name"`
}
