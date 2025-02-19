CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NULL,
    oauth_provider VARCHAR(255) NULL,
    oauth_id VARCHAR(255) NULL,
    verified BOOLEAN NOT NULL,
    created timestamp DEFAULT NOW(),
    updated timestamp DEFAULT NOW()
)