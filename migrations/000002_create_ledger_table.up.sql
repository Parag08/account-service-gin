CREATE TYPE credit_type_enum AS ENUM('credit', 'debit');

CREATE TABLE ledger (
    id BIGSERIAL PRIMARY KEY,
    uuid uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_uuid uuid NOT NULL,
    transaction_reference uuid,
    transaction_note text,
    transaction_currency_code text NOT NULL,
    transaction_country_id integer NOT NULL,
    credit_type credit_type_enum,
    credit_changed integer NOT NULL DEFAULT 0,
    credits_used integer NOT NULL DEFAULT 0,
    available_amount int DEFAULT 0, 
    exausted boolean DEFAULT false, 
    type VARCHAR(32), 
    priority int,
    expires_at timestamp without time zone, 
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    CONSTRAINT ledger_user_uuid_user_country_id_user_currency_code_transac_key UNIQUE (user_uuid, transaction_reference, created_at)
);