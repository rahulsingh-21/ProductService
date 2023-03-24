package input

type CatalougeInput struct {
	Id            string `json:"id" bson:"id"`
	Title         string `json:"title" bson:"title"`
	Description   string `json:"description" bson:"description"`
	Category      string `json:"product_category" bson:"product_category"`
	Brand         string `json:"brand" bson:"brand"`
	Availablility int    `json:"availability" bson:"availability"`
	Price         int64  `json:"price" bson:"price"`
	SalePrice     int64  `json:"sale_price" bson:"sale_price"`
}

type OrderInput []struct {
	Id            string `json:"id" bson:"id"`
	Title         string `json:"title" bson:"title"`
	Description   string `json:"description" bson:"description"`
	Category      string `json:"product_category" bson:"product_category"`
	Brand         string `json:"brand" bson:"brand"`
	Availablility int    `json:"availability" bson:"availability"`
	Price         int64  `json:"price" bson:"price"`
	SalePrice     int64  `json:"sale_price" bson:"sale_price"`
}
