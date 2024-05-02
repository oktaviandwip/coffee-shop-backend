package models

import "time"

var schemaFavorite = `
CREATE TABLE favorites (
	user_id UUID,
	product_id UUID,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NULL,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE ON UPDATE CASCADE
);
`

type Favorite struct {
	User_id    string `db:"user_id" form:"user_id" json:"user_id"`
	Product_id string `db:"product_id" form:"product_id" json:"product_id"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Favorites []Favorite
