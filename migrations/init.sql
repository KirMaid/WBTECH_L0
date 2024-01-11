CREATE TABLE orders
(
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255),
    entry VARCHAR(255),
    locale VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(255),
    sm_id INT,
    date_created TIMESTAMP,
    oof_shard VARCHAR(255)
)

CREATE TABLE deliveries
(
    order_uid VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(255),
    zip VARCHAR(255),
    city VARCHAR(255),
    address VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255),
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
)

CREATE TABLE payments
(
    order_uid VARCHAR(255) PRIMARY KEY,
    transaction VARCHAR(255),
    request_id VARCHAR(255),
    currency VARCHAR(255),
    provider VARCHAR(255),
    amount DECIMAL(10, 2),
    dt BIGINT,
    bank VARCHAR(255),
    delivery_cost DECIMAL(10, 2),
    goods_total DECIMAL(10, 2),
    custom_fee DECIMAL(10, 2),
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
)

CREATE TABLE items
(
    order_uid VARCHAR(255),
    chrt_id INT,
    track_number VARCHAR(255),
    price DECIMAL(10, 2),
    rid VARCHAR(255),
    name VARCHAR(255),
    sale DECIMAL(10, 2),
    size VARCHAR(255),
    total_price DECIMAL(10, 2),
    nm_id INT,
    brand VARCHAR(255),
    status INT,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
)