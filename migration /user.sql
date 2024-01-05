CREATE SCHEMA IF NOT EXISTS "user";

CREATE TABLE IF NOT EXISTS "user"."profile" (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL CHECK (first_name <> ''),
    last_name VARCHAR(255) NOT NULL CHECK (last_name <> ''),
    email VARCHAR(255) UNIQUE NOT NULL CHECK (email <> ''),
    mobile VARCHAR(255) NOT NULL CHECK (mobile <> ''),
    role VARCHAR(255) NOT NULL DEFAULT 'user' CHECK  (role IN ('admin', 'user', 'manager')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "user"."credentials" (
    user_id UUID,
    password VARCHAR(255) NOT NULL CHECK (password <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES "user"."profile"(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "user"."tokens" (
    user_id UUID,
    forgot_token VARCHAR(255) NOT NULL CHECK (forgot_token <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES "user"."profile"(user_id) ON DELETE CASCADE
);


CREATE OR REPLACE FUNCTION "user"."func_updated_at"()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER "trg_profile_updated_at"
    BEFORE UPDATE ON "user"."profile"
    FOR EACH ROW
    EXECUTE PROCEDURE "user"."func_updated_at"();

CREATE TRIGGER "trg_credentials_updated_at"
    BEFORE UPDATE ON "user"."credentials"
    FOR EACH ROW
    EXECUTE PROCEDURE "user"."func_updated_at"();

CREATE TRIGGER "trg_tokens_updated_at"
    BEFORE UPDATE ON "user"."tokens"
    FOR EACH ROW
    EXECUTE PROCEDURE "user"."func_updated_at"();

CREATE TABLE IF NOT EXISTS "user"."audit" (
    ID SERIAL PRIMARY KEY,
    table_name VARCHAR(255) NOT NULL CHECK (table_name <> ''),
    table_id UUID NOT NULL,
    action VARCHAR(255) NOT NULL CHECK (action <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES "user"."profile"(user_id) ON DELETE SET NULL
);

CREATE OR REPLACE FUNCTION "user"."func_audit"()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO "user"."audit" (table_name, table_id, action, auth_id)
    VALUES (TG_TABLE_NAME, NEW.ID, TG_OP, NEW.auth_id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER "trg_profile_audit"
    AFTER INSERT OR UPDATE OR DELETE ON "user"."profile"
    FOR EACH ROW
    EXECUTE PROCEDURE "user"."func_audit"();

CREATE TRIGGER "trg_credentials_audit"
    AFTER INSERT OR UPDATE OR DELETE ON "user"."credentials"
    FOR EACH ROW
    EXECUTE PROCEDURE "user"."func_audit"();

CREATE TRIGGER "trg_tokens_audit"
    AFTER INSERT OR UPDATE OR DELETE ON "user"."tokens"
    FOR EACH ROW
    EXECUTE PROCEDURE "user"."func_audit"();