create table measurements
(
    measurement_id     varchar(36) not null,
    user_id            varchar(36) not null,
    measurement_weight int         not null,
    created_at         timestamp without time zone default now(),
    updated_at         timestamp without time zone default now(),

    primary key (measurement_id),

    constraint fk_user
        foreign key (user_id)
            references users (user_id)
)