CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE SCHEMA IF NOT EXISTS "user";

CREATE TABLE IF NOT EXISTS "user"."identity" (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL CHECK (email <> ''),
    password VARCHAR(255) NOT NULL CHECK (password <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "user"."tokens" (
    ID SERIAL PRIMARY KEY,
    forgot_token VARCHAR(255) NOT NULL CHECK (forgot_token <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    auth_id UUID,
    FOREIGN KEY (auth_id) REFERENCES "user"."identity"(ID) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "user"."profile" (
    ID SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL CHECK (first_name <> ''),
    last_name VARCHAR(255) NOT NULL CHECK (last_name <> ''),
    mobile VARCHAR(255) NOT NULL CHECK (mobile <> ''),
    admin_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    auth_id UUID,
    FOREIGN KEY (auth_id) REFERENCES "user"."identity"(ID) ON DELETE CASCADE,
    FOREIGN KEY (admin_id) REFERENCES "user"."identity"(ID) ON DELETE SET NULL
);