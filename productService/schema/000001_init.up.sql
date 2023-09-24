CREATE TABLE products 
(
    id serial not null unique,
    code varchar(255) not null,
    name varchar(255) not null,
    price varchar(255) not null,
    description varchar(255) not null
);
