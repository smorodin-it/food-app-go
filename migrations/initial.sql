create table if not exists users
(
    id            varchar(36)  not null,
    email         varchar(50)  not null,
    password_hash varchar(128) not null,
    created_at    timestamp with time zone default now(),
    updated_at    timestamp with time zone default now(),

    primary key (id),
    unique (id),
    unique (email)
)