package models

import "time"

var schemaProduct = `
CREATE TABLE products (
	product_id UUID DEFAULT UUID_GENERATE_V4() PRIMARY KEY,
	photo_product TEXT NOT NULL,
	product_name VARCHAR(255) NOT NULL,
	price INT NOT NULL,
	description TEXT NOT NULL,
	size VARCHAR(255)[],
	delivery_method VARCHAR(255)[],
	start_hour TIME NOT NULL,
	end_hour TIME NOT NULL,
	stock INT NOT NULL,
	type VARCHAR(255) NOT NULL
);
`

type Product struct {
	Product_id      string   `db:"product_id" form:"product_id" json:"product_id"`
	Photo_product   string   `db:"photo_product" form:"photo_product" json:"photo_product"`
	Product_name    string   `db:"product_name" form:"product_name" json:"product_name"`
	Price           int      `db:"price" form:"price" json:"price"`
	Description     string   `db:"description" form:"description" json:"description"`
	Size            []string `db:"size" form:"size" json:"size"`
	Delivery_method []string `db:"delivery_method" form:"delivery_method" json:"delivery_method"`
	Start_hour      string   `db:"start_hour" form:"start_hour" json:"start_hour"`
	End_hour        string   `db:"end_hour" form:"end_hour" json:"end_hour"`
	Stock           int      `db:"stock" form:"stock" json:"stock"`
	Product_type    string   `db:"product_type" form:"product_type" json:"product_type"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
