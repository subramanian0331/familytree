CREATE TABLE user_profile (
id SERIAL PRIMARY KEY,
firstname VARCHAR(100),
lastname  VARCHAR(100),
nickname VARCHAR(100),
email VARCHAR(100),
password_hash VARCHAR(200));