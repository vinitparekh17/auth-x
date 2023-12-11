-- Table: users
CREATE DATABASE IF NOT EXISTS "auth";

USE "auth";

CREATE SCHEMA IF NOT EXISTS "user";

CREATE TABLE IF NOT EXISTS "user.identity" (
    UID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL CHECK (email <> ''),
    password VARCHAR(255) NOT NULL CHECK (password <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW()
)

CREATE TABLE IF NOT EXISTS "user.tokens" (
    UUID UUID PRIMARY KEY REFERENCES "user.identity"(UID) CASCADE ON DELETE CASCADE,
    forgot_token VARCHAR(255) NOT NULL CHECK (token <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW()
)

CREATE TABLE IF NOT EXISTS "user.profile" (
    UID UUID PRIMARY KEY REFERENCES "user.identity"(UID) CASCADE ON DELETE CASCADE,
    user_name VARCHAR(255) NOT NULL CHECK (first_name <> ''),
    mobile VARCHAR(255) NOT NULL CHECK (mobile <> ''),
    admin_id UUID REFERENCES "user.identity"(UID) ON DELETE SET NULL DEFAULT NULL CHECK (admin_id <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW()
)