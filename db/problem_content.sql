create table problem_content
(
    uuid uuid default gen_random_uuid() not null primary key,
    cid  varchar(255)                   not null,
    pid  varchar(255)                   not null
);

alter table problem_content
    owner to "spark-oj";

