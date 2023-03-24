package response

import "time"

type Catalouge struct {
	Id            string    `json:"id" bson:"id"`
	Title         string    `json:"title" bson:"title"`
	Description   string    `json:"description" bson:"description"`
	Category      string    `json:"product_category" bson:"product_category"`
	Brand         string    `json:"brand" bson:"brand"`
	Availablility int       `json:"availability" bson:"availability"`
	Price         int64     `json:"price" bson:"price"`
	SalePrice     int64     `json:"sale_price" bson:"sale_price"`
	CreatedAt     time.Time `json:"-" bson:"created_at"`
	UpdatedAt     time.Time `json:"-" bson:"updated_at"`
	IsActive      int       `json:"-" bson:"isActive"`
}
