CREATE TABLE favorites (
	user_id UUID,
	product_id UUID,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NULL,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO favorites (user_id, product_id)
SELECT u.user_id, p.product_id
FROM users u
CROSS JOIN products p;