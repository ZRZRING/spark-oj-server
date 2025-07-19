create table contest
(
    cid         varchar(255)            not null
        primary key,
    title       varchar(255)            not null,
    password    varchar(255),
    problems    jsonb                   not null,
    description text,
    start_time  timestamp               not null,
    end_time    timestamp               not null,
    lock_time   timestamp               not null,
    create_by   varchar(255)            not null,
    create_at   timestamp default now() not null,
    update_at   timestamp default now() not null,
    delete_at   timestamp
);

alter table contest
    owner to "spark-oj";

