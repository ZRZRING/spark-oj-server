create table file
(
    path      varchar(255)            not null
        primary key,
    type      varchar(255)            not null,
    create_at timestamp default now() not null,
    update_at timestamp default now() not null,
    delete_at timestamp
);

alter table file
    owner to "spark-oj";

