CREATE TYPE AUTH_T AS ENUM ('succeed', 'failed');

CREATE TABLE authorizations (
    "id"         SERIAL,
    "enabled"    BOOLEAN   NOT NULL,
    "token"      TEXT      NOT NULL,
    "error"      TEXT,
    "ip"         TEXT      NOT NULL,
    "name"       TEXT      NOT NULL,
    "status"     AUTH_T    NOT NULL,
    "timestamp"  TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX authorizations_token_idx ON authorizations("token");
