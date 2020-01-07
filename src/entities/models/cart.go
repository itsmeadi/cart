package models

type Cart struct {
	ID        int64 `json:"id" db:"id"`
	UserID    int64 `json:"user_id" db:"user_id"`
	Status    int   `json:"status" db:"status"`
	CreatedAt int64 `json:"created_at" db:"created_at"`
	UpdatedAt int64 `json:"updated_at" db:"updated_at"`
}

type CartItems struct {
	ID        int64 `json:"id" db:"id"`
	CartID    int64 `json:"cart_id" db:"cart_id"`
	ProductID int64 `json:"product_id" db:"product_id"`
	Qty       int64 `json:"qty" db:"qty"`
	Status    int   `json:"status" db:"status"`
	CreatedAt int64 `json:"created_at" db:"created_at"`
	UpdatedAt int64 `json:"updated_at" db:"updated_at"`
}

type CartXCartItems struct {
	CartID     int64 `json:"cart_id" db:"cart_id"`
	UserID     int64 `json:"user_id" db:"user_id"`
	ProductID  int64 `json:"product_id" db:"product_id"`
	Qty        int64 `json:"qty" db:"qty"`
	ItemStatus int   `json:"item_status" db:"item_status"`
	CreatedAt  int64 `json:"created_at" db:"created_at"`
	UpdatedAt  int64 `json:"updated_at" db:"updated_at"`
}

type CartDetail struct {
	CartID     int64 `json:"cart_id" db:"cart_id"`
	UserID     int64 `json:"user_id" db:"user_id"`
	TotalPrice float64
	Products   []CartProduct `json:"products"`
}
type CartProduct struct {
	Product
	Qty        int64 `json:"qty" db:"qty"`
	TotalPrice float64
}
