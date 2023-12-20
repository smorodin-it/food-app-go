create table meals
(
    meal_id      varchar(36) not null,
    user_id      varchar(36)  not null,
    meal_name    varchar(255) not null,
    total_weight int          not null,
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now(),

    primary key (meal_id),

    constraint fk_user
        foreign key (user_id)
            references users (user_id)
)