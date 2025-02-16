CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NULL,
    isManualRegist BOOLEAN NOT NULL,
    oauth_provider VARCHAR(255),
    oauth_id VARCHAR(255),
    created timestamp DEFAULT NOW(),
    updated timestamp DEFAULT NOW()
)