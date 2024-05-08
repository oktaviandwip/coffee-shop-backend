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

INSERT INTO products (photo_product, product_name, price, description, size, delivery_method, start_hour, end_hour, stock, product_type)
VALUES
('http://localhost:3001/user/image/veggie_tomato_mix.jpg', 'Veggie Tomato Mix', 34000, 'Veggie with Tomato Mix', '{250 gr, 300 gr, 500 gr}', '{Home Delivery, Dine In}', '13:00:00', '17:00:00', 50, 'food'),
('http://localhost:3001/user/image/hazelnut_latte.jpg', 'Hazelnut Latte', 25000, 'Latte with Hazelnut', '{R, L, XL}', '{Home Delivery, Take Away}', '13:00:00', '19:00:00', 60, 'coffee'),
('http://localhost:3001/user/image/summer_fried_rice.jpg', 'Summer Fried Rice', 32000, 'Fried Rice with Seafood', '{250 gr, 300 gr, 500 gr}', '{Home Delivery, Dine In, Take Away}', '15:00:00', '17:00:00', 50, 'food'),
('http://localhost:3001/user/image/creamy_ice_latte.jpg', 'Creamy Ice Latte', 32000, 'Ice Latte with Cream', '{R, L, XL}', '{Home Delivery, Take Away}', '13:00:00', '17:00:00', 100, 'non coffee');