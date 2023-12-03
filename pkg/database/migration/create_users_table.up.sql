CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) ,
    password VARCHAR(255) NOT NULL,
    roles VARCHAR(255) NOT NULL,
    refreshtoken VARCHAR(255),
    istokenactive INT
);

