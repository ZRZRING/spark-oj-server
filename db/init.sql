drop table if exists user_base;
create table if not exists user_base
(
    username   varchar(255) not null primary key,
    password   varchar(255) not null,
    nickname   varchar(255) not null,
    role       varchar(255) not null default 'user',
    solved     integer      not null default 0,
    submitted  integer      not null default 0,
    rating     integer      not null default 1500,
    cf_id      varchar(255),
    atc_id     varchar(255),
    company    varchar(255),
    department varchar(255),
    major      varchar(255),
    class      varchar(255),
    email      varchar(255),
    tel        varchar(255),
    avatar     varchar(255),
    create_at  timestamp    not null default now(),
    update_at  timestamp    not null default now(),
    delete_at  timestamp
);


drop table if exists user_role;
create table if not exists user_role
(
    name       varchar(255) not null primary key,
    permission jsonb        not null default '{}'::JSONB,
    create_at  timestamp    not null default now(),
    update_at  timestamp    not null default now(),
    delete_at  timestamp
);



drop table if exists problem;
create table if not exists problem
(
    pid          varchar(255) not null primary key,
    title        varchar(255) not null default 'Unknown Problem',
    judge_type         integer      not null default 0,
    time         integer      not null default 1000,
    memory       integer      not null default 1024,
    content_type integer      not null default 0,
    content      text,
    create_by    varchar(255) not null default 'Unknown User',
    rating       integer,
    config       jsonb,
    create_at    timestamp    not null default now(),
    update_at    timestamp    not null default now(),
    delete_at    timestamp
);



drop table if exists file;
create table if not exists file
(
    path      varchar(255) not null primary key,
    type      varchar(255) not null,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);



drop table if exists contest;
create table if not exists contest
(
    cid         varchar(255) not null primary key,
    title       varchar(255) not null,
    password    varchar(255),
    problems    jsonb        not null,
    description text,
    start_time  timestamp    not null,
    end_time    timestamp    not null,
    lock_time   timestamp    not null,
    create_by   varchar(255) not null,
    create_at   timestamp    not null default now(),
    update_at   timestamp    not null default now(),
    delete_at   timestamp
);



drop table if exists submission;
create table if not exists submission
(
    sid       integer      not null primary key,
    pid       varchar(255) not null,
    uid       varchar(255) not null,
    cid       varchar(255),
    result    varchar(255) not null,
    language  varchar(255) not null,
    memory    integer      not null,
    time      integer      not null,
    code      text         not null,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);



drop table if exists judger;
create table if not exists judger
(
    jid       varchar(255) not null primary key,
    status    integer      not null default 0,
    create_at timestamp    not null default now(),
    update_at timestamp    not null default now(),
    delete_at timestamp
);