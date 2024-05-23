CREATE EXTENSION if not exists "uuid-ossp";


create table if not exists users (

    Id serial primary key,
    Uuid uuid default uuid_generate_v4() unique,
    Name varchar(30),
    Login varchar(30) unique not null,
    Password varchar(30) not null,
    Email varchar(256)

);

create table if not exists refresh_tokens (

    Id serial primary key,
    Token varchar(300),
    Time_end_of_life bigint,
    Owner_uuid uuid not null references users(uuid) on delete cascade

);