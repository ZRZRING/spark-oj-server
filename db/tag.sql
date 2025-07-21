create table tag
(
    tag       varchar(255)            not null
        primary key,
    approved  boolean   default false not null,
    create_at timestamp default now() not null,
    update_at timestamp default now() not null,
    delete_at timestamp               not null
);