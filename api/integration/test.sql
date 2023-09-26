-- +migrate Up
CREATE TABLE IF NOT EXISTS user (
    user_id VARCHAR(128) NOT NULL PRIMARY KEY,
    user_name VARCHAR(128)

);
-- +migrate Down
DROP TABLE IF EXISTS user;
