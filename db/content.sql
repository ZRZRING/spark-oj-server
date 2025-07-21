create table content
(
    uuid      uuid      default gen_random_uuid() not null primary key,
    type      integer                             not null,
    payload   jsonb,
    create_at timestamp default now()             not null,
    update_at timestamp default now()             not null,
    delete_at timestamp
);

alter table content
    owner to "spark-oj";

