package models

import "time"

var schemaUser = `
CREATE TABLE users (
	user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	photo_profile TEXT NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password TEXT NOT NULL,
	address TEXT NOT NULL,
	phone_number VARCHAR(15) UNIQUE NOT NULL,
	display_name VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	birthday DATE NOT NULL,
	gender CHAR(1) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NULL
);
`

type User struct {
	User_id       string `db:"user_id" form:"user_id" json:"user_id"`
	Photo_profile string `db:"photo_profile" form:"photo_profile" json:"photo_profile"`
	Email         string `db:"email" form:"email" json:"email"`
	Password      string `db:"password" form:"password" json:"password"`
	Address       string `db:"address" form:"address" json:"address"`
	Phone_number  string `db:"phone_number" form:"phone_number" json:"phone_number"`
	Display_name  string `db:"display_name" form:"display_name" json:"display_name"`
	First_name    string `db:"first_name" form:"first_name" json:"first_name"`
	Last_name     string `db:"last_name" form:"last_name" json:"last_name"`
	Birthday      string `db:"birthday" form:"birthday" json:"birthday"`
	Gender        string `db:"gender" form:"gender" json:"gender"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
