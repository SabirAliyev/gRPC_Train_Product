CREATE TABLE product (
    id                  SERIAL PRIMARY KEY,
    name                varchar(100),
    description         varchar(255) DEFAULT '',
    value               float
);