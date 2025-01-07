create table meals_ingredients
(
    id                varchar(36) not null,
    meal_id           varchar(36) not null,
    ingredient_id     varchar(36) not null,
    ingredient_weight int         not null,
    created_at        timestamp without time zone default now(),
    updated_at        timestamp without time zone default now(),

    primary key (id),

    constraint fk_meal
        foreign key (meal_id)
            references meals (meal_id),

    constraint fk_ingredient
        foreign key (ingredient_id)
            references ingredients (ingredient_id)
)