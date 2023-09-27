CREATE TABLE storeroom (
    id serial primary key,
    code text unique,
    quantity int
);
