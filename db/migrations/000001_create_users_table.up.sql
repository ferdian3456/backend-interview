CREATE TABLE IF NOT EXISTS users(
    id char(36) PRIMARY KEY,
    username varchar(22) UNIQUE NOT NULL,
    email varchar(80) UNIQUE NOT NULL,
    contact_phone varchar(15),
    password varchar(60) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
)