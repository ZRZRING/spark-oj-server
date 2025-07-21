create table file
(
    uuid      uuid      default gen_random_uuid() not null primary key,
    path      varchar(255)                        not null,
    type      varchar(255)                        not null,
    create_at timestamp default now()             not null,
    update_at timestamp default now()             not null,
    delete_at timestamp
);

alter table file
    owner to "spark-oj";

