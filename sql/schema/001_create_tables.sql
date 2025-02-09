-- +goose Up
CREATE TABLE IF NOT EXISTS bakery_user(
    id TEXT NOT NULL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    birth_date TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    address1 TEXT NOT NULL,
    address2 TEXT NOT NULL DEFAULT '',
    gender TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS bakery_session(
    id TEXT NOT NULL PRIMARY KEY,
    refresh_token TEXT NOT NULL,
    access_token TEXT NOT NULL,
    username TEXT NOT NULL,
    user_id TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES bakery_user(id)
);

CREATE TABLE IF NOT EXISTS bakery_product(
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    category TEXT NOT NULL DEFAULT '',
    price REAL NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    issued INTEGER NOT NULL,
    has_stock INTEGER NOT NULL DEFAULT 0,
    discount REAL NOT NULL DEFAULT 0,
    rating REAL NOT NULL DEFAULT 0,
    images TEXT NOT NULL DEFAULT '',
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS bakery_order(
    id TEXT NOT NULL PRIMARY KEY,
    total_amount REAL NOT NULL,
    payment_method TEXT NOT NULL DEFAULT 'cash',
    status TEXT NOT NULL DEFAULT 'placed',
    user_id TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES bakery_user(id)
);

CREATE TABLE IF NOT EXISTS bakery_order_products(
    order_id TEXT NOT NULL,
    product_id TEXT NOT NULL REFERENCES bakery_product(id),
    product_name TEXT NOT NULL,
    product_price REAL NOT NULL,
    product_discount REAL NOT NULL,
    product_rating REAL NOT NULL,
    total_price REAL NOT NULL,
    quantity INTEGER NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES bakery_order(id),
    FOREIGN KEY (product_id) REFERENCES bakery_product(id)
);

-- +goose Down
DROP TABLE bakery_user;
DROP TABLE bakery_product;
DROP TABLE bakery_order;
DROP TABLE bakery_order_products;
DROP TABLE bakery_session;

