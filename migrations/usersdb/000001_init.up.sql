CREATE EXTENSION if not exists "uuid-ossp";


create table if not exists users (

    Uuid uuid default uuid_generate_v4() unique,
    Name varchar(30),
    Login varchar(30) unique not null,
    Password varchar(30) not null,
    Email varchar(256)

);

create table if not exists sessions (

    Id serial primary key,
    Token varchar(300),
    Life_time bigint,
    Owner uuid not null,
    Addr varchar(16),
    Device varchar(100),
    Browser varchar(100),

);

create index uuid_index on users using hash(Uuid);
create index token_index on sessions using hash (Token);
