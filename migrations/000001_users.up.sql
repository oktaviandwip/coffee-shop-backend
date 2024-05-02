CREATE TABLE users (
	user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	photo_profile TEXT NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password TEXT NOT NULL,
	role VARCHAR(255) NOT NUll,
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

INSERT INTO users (photo_profile, email, password, role, address, phone_number, display_name, first_name, last_name, birthday, gender) 
VALUES
('http://localhost:3001/user/image/zulaikha.jpg', 'zulaikha17@gmail.com', '1234abcd', 'user', 'Iskandar Street no. 67 Block A Near Bus Stop', '+62813456782', 'Zulaikha', 'Zulaikha', 'Nirmala', '1990-04-03', 'f'),
('http://localhost:3001/user/image/okta.jpg', 'okta@gmail.com', '1234abcd', 'user', 'Halim, Jakarta Timur', '+6277888999910', 'Okta', 'Okta', 'Dwi', '1996-10-02', 'm'),
('http://localhost:3001/user/image/vian.jpg', 'vian@gmail.com', '1234abcd', 'user', 'Halim, Jakarta Timur', '+6277888999911', 'Vian', 'Vian', 'Putra', '1996-10-02', 'm'),
('http://localhost:3001/user/image/dwi.jpg', 'dwi@gmail.com', '1234abcd', 'user', 'Halim, Jakarta Timur', '+6277888999912', 'Dwi', 'Dwi', 'Putra', '1996-11-02', 'm'),
('http://localhost:3001/user/image/putra.jpg', 'putra@gmail.com', '1234abcd', 'user', 'Halim, Jakarta Timur', '+6277888999913', 'Putra', 'Putra', 'Dwi', '1996-12-02', 'm');