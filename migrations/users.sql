create table users
(
    user_id       varchar(36) not null,
    email         varchar(50)  not null unique,
    password_hash varchar(128) not null,
    refresh_token varchar(36) unique,
    created_at    timestamp with time zone default now(),
    updated_at    timestamp with time zone default now(),

    primary key (user_id)
);

