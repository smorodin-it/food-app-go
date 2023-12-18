create table meals
(
    id           varchar(36)  not null primary key,
    user_id      varchar(36)  not null,
    name         varchar(255) not null,
    total_weight int          not null,
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
)