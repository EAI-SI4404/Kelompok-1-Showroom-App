package dto

type PaymentRequest struct {
	PaymentTypeID string   `json:"payment_type_id"`
	Title         string   `json:"title"`
	Phone         string   `json:"phone"`
	ExpireTime    string   `json:"expire_time"`
	ProductList   []string `json:"product_list"`
}
