BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    firstname       TEXT   DEFAULT '' NOT NULL,
    lastname       TEXT   DEFAULT '' NOT NULL,
    email      VARCHAR(100) NOT NULL UNIQUE,
    phone      VARCHAR(20),
    is_active  BOOLEAN DEFAULT TRUE,
    subscribe VARCHAR,
    blocks VARCHAR,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS relations
(
    id         SERIAL PRIMARY KEY,
    requester_id     INTEGER NOT NULL REFERENCES users(id),
    addressee_id     INTEGER NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;



