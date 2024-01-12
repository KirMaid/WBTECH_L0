-- Таблица orders
INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id,
                    delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('some_order_uid', 'some_track_number', 'some_entry', 'some_locale',
        'some_internal_signature', 'some_customer_id', 'some_delivery_service',
        'some_shardkey', 1, NOW(), 'some_oof_shard');

-- Таблица deliveries
INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email)
VALUES ('some_order_uid', 'some_name', 'some_phone', 'some_zip', 'some_city',
        'some_address', 'some_region', 'some_email');

-- Таблица payments
INSERT INTO payments (order_uid, transaction, request_id, currency, provider, amount, dt, bank,
                      delivery_cost, goods_total, custom_fee)
VALUES ('some_order_uid', 'some_transaction', 'some_request_id', 'some_currency',
        'some_provider', 100.00, EXTRACT(EPOCH FROM NOW()), 'some_bank',
        10.00, 90.00, 10.00);

-- Таблица items
INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price,
                   nm_id, brand, status)
VALUES ('some_order_uid', 1, 'some_track_number', 100.00, 'some_rid', 'some_name',
        10.00, 'some_size', 90.00, 1, 'some_brand', 1);

INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price,
                   nm_id, brand, status)
VALUES ('some_order_uid', 1, 'another_track_number', 100.00, 'another_rid', 'another_name',
        10.00, 'another_size', 90.00, 1, 'another_brand', 1);