create table judger
(
    jid       varchar(255)            not null
        primary key,
    status    integer   default 0     not null,
    create_at timestamp default now() not null,
    update_at timestamp default now() not null,
    delete_at timestamp
);

alter table judger
    owner to "spark-oj";

