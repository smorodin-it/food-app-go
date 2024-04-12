create table inventory
(
    inventory_id varchar(36)  not null,
    user_id      varchar(36)  not null,
    name         varchar(255) not null,
    weight       int,

    primary key (inventory_id),

    constraint fk_user
        foreign key (user_id)
            references users (user_id)
)