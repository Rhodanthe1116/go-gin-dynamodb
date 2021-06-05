package forms

type RecordStoreId struct {
	StoreId   string `json:"store_id" binding:"required"`
}

type RecordRecord struct {
    Phone string `json:"phone" binding:"required"`
    UserId string `json:"user_id" binding:"required"`
    StoreId string `json:"store_id" binding:"required"`
    Time int64 `json:"time" binding:"required"`
}
