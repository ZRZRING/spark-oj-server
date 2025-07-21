create sequence submission_id_seq start with 1000000;

create table submission
(
    sid       integer                 not null primary key,
    pid       varchar(255)            not null,
    uid       varchar(255)            not null,
    cid       varchar(255),
    result    varchar(255)            not null,
    language  varchar(255)            not null,
    memory    integer                 not null,
    time      integer                 not null,
    code      text                    not null,
    create_at timestamp default now() not null,
    update_at timestamp default now() not null,
    delete_at timestamp
);

alter table submission
    owner to "spark-oj";

