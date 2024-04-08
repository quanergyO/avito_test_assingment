CREATE TABLE banners (
     id SERIAL PRIMARY KEY,
     tag_ids INT[],
     feature_id INT,
     content JSONB,
     is_active BOOLEAN,
     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(10)
);
