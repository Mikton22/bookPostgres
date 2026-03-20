CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(100) NOT NULL,
    review VARCHAR(1000) NOT NULL,
    year INT NOT NULL,
    read BOOLEAN NOT NULL,
    read_started DATE NOT NULL,
    read_finished DATE
);