DROP TABLE IF EXISTS "order_items";
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS "delivery";
DROP TABLE IF EXISTS "client";

CREATE TABLE IF NOT EXISTS "client" (
	client_id varchar(255) PRIMARY KEY NOT NULL,
	cpf varchar(80),
	name varchar(100),
	email varchar(120) unique,
	username varchar(80) unique,
	password varchar(255),
	telephone varchar(20),
	recover_token varchar(255),
	created_at date,
	updated_at date DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "delivery" (
	delivery_id varchar(255) PRIMARY KEY NOT NULL,
	country varchar(20) DEFAULT 'BR' NOT NULL,
	state varchar(100) DEFAULT 'BM' NOT NULL,
	address varchar(120) NOT NULL,
	number_home varchar(80) NOT NULL,
	observation text,
	token varchar(255) NOT NULL,
	status varchar(255) DEFAULT 'ANALYZING' NOT NULL,
	client_id varchar(255) REFERENCES client(client_id) NOT NULL,
	created_at date,
	updated_at date DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "order" (
	order_id varchar(255) PRIMARY KEY NOT NULL,
	billing_type varchar(20) DEFAULT 'PIX',
	price float NOT NULL,
	discount float,
	price_delivery float NOT NULL,
	is_delivery boolean DEFAULT TRUE,
	status varchar(255) DEFAULT 'ANALYZING',
	client_id varchar(255) REFERENCES client(client_id) NOT NULL,
	delivery_id varchar(255) REFERENCES delivery(delivery_id) NOT NULL,
	created_at date,
	updated_at date DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "order_items" (
	order_items_id varchar(255) PRIMARY KEY NOT NULL,
	product_id varchar(255),
	quantity integer,
	price float,
	created_at date,
	updated_at date DEFAULT NOW(),
	order_id varchar(255) NOT NULL,
	CONSTRAINT fk_order_items FOREIGN KEY(order_id) REFERENCES "order"(order_id)
);