CREATE TABLE position (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT
);
COMMIT;

CREATE TABLE permission (
    id SERIAL PRIMARY KEY,
    code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT
);
COMMIT;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    position_id INT REFERENCES position(id)
);
COMMIT;

CREATE TABLE position_permission (
    position_id INT REFERENCES position(id) ON DELETE CASCADE,
    permission_id INT REFERENCES permission(id) ON DELETE CASCADE,
    PRIMARY KEY (position_id, permission_id)
);
COMMIT;