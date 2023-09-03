CREATE TABLE users (
	id serial PRIMARY KEY,
	name VARCHAR ( 100 ) NOT NULL,
	phone VARCHAR ( 20 ) UNIQUE NOT NULL,
	password VARCHAR ( 100 ) NOT NULL
);
