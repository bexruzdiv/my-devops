CREATE TABLE IF NOT EXISTS sms_send (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    phone_number VARCHAR(15) NOT NULL,
    text VARCHAR(160) NOT NULL,
    sent_at TIMESTAMP DEFAULT NULL
);