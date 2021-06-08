package forms

type RecordStoreId struct {
	StoreId   string `json:"store_id" binding:"required"`
}
