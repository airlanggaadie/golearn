CREATE TABLE IF NOT EXISTS versions(
    id bigserial NOT NULL PRIMARY KEY,
    app VARCHAR (50) NOT NULL,
    "version" VARCHAR (50) NOT NULL,
    code SMALLINT NULL,
    description VARCHAR (255) NOT NULL,
    is_active bool NOT NULL DEFAULT true,
    created_at timestamp(0) NOT NULL DEFAULT now(),
    updated_at timestamp(0) NOT NULL DEFAULT now()
);