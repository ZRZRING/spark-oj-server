create table public.contest
(
    cid         integer generated always as identity (start with 1000)
        primary key,
    create_at   timestamp default now()                                                    not null,
    update_at   timestamp default now()                                                    not null,
    delete_at   timestamp,
    title       varchar(255)                                                               not null,
    password    varchar(255),
    description text,
    start_time  timestamp default (date_trunc('hour'::text, now()) + '01:00:00'::interval) not null,
    end_time    timestamp default (date_trunc('hour'::text, now()) + '06:00:00'::interval) not null,
    lock_time   timestamp,
    create_by   varchar(255)                                                               not null,
    practice    boolean   default true                                                     not null,
    problems    jsonb     default '{}'::jsonb
);

---

create table public.problem
(
    pid          integer generated always as identity (start with 1000)
        primary key,
    create_at    timestamp default now() not null,
    update_at    timestamp default now() not null,
    delete_at    timestamp,
    title        varchar(255)            not null,
    judge_type   varchar(255)            not null,
    time_limit   integer   default 1000  not null,
    memory_limit integer   default 256   not null,
    create_by    varchar(255)            not null,
    rating       integer,
    content      text
);

---

create table public.resource
(
    uuid      uuid      default gen_random_uuid() not null
        primary key,
    create_at timestamp default now()             not null,
    update_at timestamp default now()             not null,
    delete_at timestamp,
    path      varchar(255)                        not null,
    type      varchar(255)                        not null
);

---

create table public.submission
(
    sid         integer generated always as identity (start with 10000000)
        primary key,
    create_at   timestamp default now() not null,
    update_at   timestamp default now() not null,
    delete_at   timestamp,
    pid         varchar(255)            not null,
    cid         varchar(255),
    username    varchar(255)            not null,
    result      varchar(255)            not null,
    language    varchar(255)            not null,
    memory_cost integer                 not null,
    time_cost   integer                 not null,
    code        text
);

create index submission_cid_index
    on public.submission (cid)
    where (cid IS NOT NULL);

create index submission_create_at_index
    on public.submission (create_at);

---

create table public.user_base
(
    username   varchar(255)                                   not null
        primary key,
    create_at  timestamp    default now()                     not null,
    update_at  timestamp    default now()                     not null,
    delete_at  timestamp,
    password   varchar(255)                                   not null,
    nickname   varchar(255)                                   not null,
    role       varchar(255) default 'user'::character varying not null,
    solved     integer      default 0                         not null,
    submitted  integer      default 0                         not null,
    rating     integer      default 1500                      not null,
    cf_id      varchar(255),
    atc_id     varchar(255),
    extra      jsonb        default '{}'::jsonb
);

---

create table public.user_role
(
    rid        varchar(255)                  not null
        primary key,
    create_at  timestamp default now()       not null,
    update_at  timestamp default now()       not null,
    delete_at  timestamp,
    permission jsonb     default '{}'::jsonb not null
);