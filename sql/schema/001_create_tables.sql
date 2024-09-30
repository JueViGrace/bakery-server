-- +goose Up
CREATE TABLE IF NOT EXISTS bakery_user(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    birth_date TIMESTAMP NOT NULL DEFAULT NOW(),
    phone VARCHAR(25) NOT NULL,
    role VARCHAR(10) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS bakery_product(
    id UUID NOT NULL PRIMARY KEY,
    price NUMERIC NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    category VARCHAR(255) NOT NULL DEFAULT '',
    stock INT NOT NULL DEFAULT 0,
    image TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS bakery_order(
    id UUID NOT NULL PRIMARY KEY,
    total_amount REAL NOT NULL,
    payment_method VARCHAR(20) NOT NULL DEFAULT 'cash',
    status VARCHAR(30) NOT NULL DEFAULT 'placed',
    user_id UUID NOT NULL REFERENCES bakery_user(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS bakery_order_products(
    order_id UUID NOT NULL REFERENCES bakery_order(id),
    product_id UUID NOT NULL REFERENCES bakery_product(id),
    price REAL NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (order_id, product_id)
);

-- +goose Down
DROP TABLE bakery_user CASCADE;
DROP TABLE bakery_product CASCADE;
DROP TABLE bakery_order CASCADE;
DROP TABLE bakery_order_products;
