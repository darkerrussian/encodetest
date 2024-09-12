CREATE TABLE persons (
                         id SERIAL PRIMARY KEY,
                         email VARCHAR(255) NOT NULL,
                         phone VARCHAR(20) NOT NULL,
                         first_name VARCHAR(50) NOT NULL,
                         last_name VARCHAR(50) NOT NULL
);