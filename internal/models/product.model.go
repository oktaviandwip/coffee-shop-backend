package models

import (
	"mime/multipart"
	"time"
)

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
	product_type VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NULL
);
`

type Product struct {
	Product_id      string                `db:"product_id" form:"product_id" json:"product_id"`
	PhotoUpload     *multipart.FileHeader `form:"photo_product"`
	Photo_product   string                `db:"photo_product" json:"photo_product"`
	Product_name    string                `db:"product_name" form:"product_name" json:"product_name"`
	Price           int                   `db:"price" form:"price" json:"price"`
	Description     string                `db:"description" form:"description" json:"description" valid:"stringlength(150|1000)~Description min 150 characters"`
	Size            []string              `db:"size" form:"size" json:"size"`
	Delivery_method []string              `db:"delivery_method" form:"delivery_method" json:"delivery_method"`
	Start_hour      string                `db:"start_hour" form:"start_hour" json:"start_hour"`
	End_hour        string                `db:"end_hour" form:"end_hour" json:"end_hour"`
	Stock           int                   `db:"stock" form:"stock" json:"stock"`
	Product_type    string                `db:"product_type" form:"product_type" json:"product_type"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Products []Product
