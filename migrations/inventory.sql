create table inventory
(
    inventory_id varchar(36)  not null,
    user_id      varchar(36)  not null,
    name         varchar(255) not null,
    weight       int,
    created_at   timestamp without time zone default now(),
    updated_at   timestamp without time zone default now(),

    primary key (inventory_id),

    constraint fk_user
        foreign key (user_id)
            references users (user_id)
)