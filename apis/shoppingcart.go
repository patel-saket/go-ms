package apis

type ShoppingCart struct {
	Id          string   `json:"id"`
	CartItems   []string `json:"cartItems"`
	CreatedDate string   `json:"createdDate"`
}
