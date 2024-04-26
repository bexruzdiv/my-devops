CREATE TABLE IF NOT EXISTS sms_payments (
    id UUID primary key,
    shipper_id uuid not null,
    order_id uuid not null,
    ext_order_id integer,
    phone varchar(20) not null,
    operator_id uuid not null,
    operator_name varchar not null,
    created_at timestamp default current_timestamp,
    payment_type varchar not null,
    order_amount numeric(9,2) not null
);