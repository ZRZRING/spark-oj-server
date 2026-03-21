create table public.contest
(
    cid         integer generated always as identity (start with 1000)
        primary key,
    title       varchar(255)                                                               not null,
    password    varchar(255),
    description text,
    start_time  timestamp default (date_trunc('hour'::text, now()) + '01:00:00'::interval) not null,
    end_time    timestamp default (date_trunc('hour'::text, now()) + '06:00:00'::interval) not null,
    lock_time   timestamp,
    create_by   varchar(255)                                                               not null,
    practice    boolean   default true                                                     not null,
    problems    jsonb     default '{}'::jsonb,
    create_at   timestamp default now()                                                    not null,
    update_at   timestamp default now()                                                    not null,
    delete_at   timestamp
);

alter table public.contest
    owner to "spark-oj";

---

create table public.problem
(
    pid          integer generated always as identity (start with 1000)
        primary key,
    title        varchar(255)            not null,
    judge_type   integer                 not null,
    time_limit   integer   default 1000  not null,
    memory_limit integer   default 1024  not null,
    create_by    varchar(255)            not null,
    rating       integer,
    create_at    timestamp default now() not null,
    update_at    timestamp default now() not null,
    delete_at    timestamp,
    content      text,
    analysis     jsonb     default '{}'::jsonb
);

alter table public.problem
    owner to "spark-oj";

---

create table public.resource
(
    uuid      uuid      default gen_random_uuid() not null
        constraint file_pkey
            primary key,
    path      varchar(255)                        not null,
    type      varchar(255)                        not null,
    create_at timestamp default now()             not null,
    update_at timestamp default now()             not null,
    delete_at timestamp
);

alter table public.resource
    owner to "spark-oj";

---

create table public.submission
(
    sid         integer generated always as identity (start with 10000000)
        primary key,
    title       varchar(255)            not null,
    pid         varchar(255)            not null,
    username    varchar(255)            not null,
    cid         varchar(255),
    result      varchar(255)            not null,
    language    varchar(255)            not null,
    memory_cost integer                 not null,
    time_cost   integer                 not null,
    code        text                    not null,
    create_at   timestamp default now() not null,
    update_at   timestamp default now() not null,
    delete_at   timestamp
);

alter table public.submission
    owner to "spark-oj";

create index submission_cid_index
    on public.submission (cid)
    where (cid IS NOT NULL);

---

create table public.user_base
(
    username   varchar(255)                                   not null
        primary key,
    password   varchar(255)                                   not null,
    nickname   varchar(255)                                   not null,
    role       varchar(255) default 'user'::character varying not null,
    solved     integer      default 0                         not null,
    submitted  integer      default 0                         not null,
    rating     integer      default 1500                      not null,
    cf_id      varchar(255),
    atc_id     varchar(255),
    company    varchar(255),
    department varchar(255),
    major      varchar(255),
    class      varchar(255),
    email      varchar(255),
    tel        varchar(255),
    avatar     varchar(255),
    extra      jsonb        default '{}'::jsonb,
    create_at  timestamp    default now()                     not null,
    update_at  timestamp    default now()                     not null,
    delete_at  timestamp
);

alter table public.user_base
    owner to "spark-oj";

---

create table public.user_role
(
    rid        varchar(255)                  not null
        primary key,
    permission jsonb     default '{}'::jsonb not null,
    create_at  timestamp default now()       not null,
    update_at  timestamp default now()       not null,
    delete_at  timestamp
);

alter table public.user_role
    owner to "spark-oj";

