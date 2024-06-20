create table users(
    id uuid primary key,
    name varchar not null,
    age int,
    is_working boolean not null,
    phone varchar
)