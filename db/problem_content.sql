create table problem_content
(
    id  varchar(255) not null,
    cid varchar(255) not null,
    pid varchar(255) not null
);

alter table problem_content
    owner to "spark-oj";

