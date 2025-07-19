create table user_role
(
    name       varchar(255)                  not null
        primary key,
    permission jsonb     default '{}'::jsonb not null,
    create_at  timestamp default now()       not null,
    update_at  timestamp default now()       not null,
    delete_at  timestamp
);

alter table user_role
    owner to "spark-oj";

