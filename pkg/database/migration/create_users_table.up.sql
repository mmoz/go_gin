CREATE TABLE users (
    username VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    roles VARCHAR(255) NOT NULL,
    refreshtoken VARCHAR(255),
    istokenactive INT
);

