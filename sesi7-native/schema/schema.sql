CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name varchar(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(20) NOT NULL
);