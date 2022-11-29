CREATE TABLE IF NOT EXISTS cakes (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title varchar(200) NOT NULL DEFAULT '',
    description TEXT,
    rating FLOAT NOT NULL DEFAULT 0,
    image TEXT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW()
);