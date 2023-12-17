create table ingredients
(
    id           varchar(36)  not null primary key,
    user_id      varchar(36)  not null,
    name         varchar(255) not null,
    manufacturer varchar(255) not null,
    barcode      varchar(13),
    proteins     int,
    carbs        int,
    fats         int,
    calories     int,
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
);

