create table ingredients
(
    ingredient_id   varchar(36) not null,
    user_id         varchar(36)  not null,
    ingredient_name varchar(255) not null,
    manufacturer    varchar(255) not null,
    barcode         varchar(13) unique,
    proteins        int,
    carbs           int,
    fats            int,
    calories        int,
    created_at      timestamp without time zone default now(),
    updated_at      timestamp without time zone default now(),

    primary key (ingredient_id),

    constraint fk_user
        foreign key (user_id)
            references users (user_id)
);

