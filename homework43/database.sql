
create table users (id uuid primary key, name varchar, age int, phone varchar);

create table cards (id uuid primary key, number varchar, user_id uuid references users(id));

create table station (id uuid primary key, name varchar);

create table terminal (id uuid primary key, station_id uuid references station(id));



CREATE TABLE transaction (
    id uuid primary key,
    card_id uuid references cards(id),
    amount int,
    terminal_id uuid references terminal(id),
    transaction_type VARCHAR NOT NULL CHECK (transaction_type IN ('credit', 'debit'))
);

