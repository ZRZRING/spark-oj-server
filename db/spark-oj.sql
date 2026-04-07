create table if not exists public.contest
(
    contest_id  integer generated always as identity ( start with 1000 )                     not null
        primary key,
    create_at   timestamptz default now()                                                    not null,
    update_at   timestamptz default now()                                                    not null,
    delete_at   timestamptz,
    title       text                                                                         not null,
    password    text,
    description text,
    start_time  timestamptz default (date_trunc('hour'::text, now()) + '01:00:00'::interval) not null,
    end_time    timestamptz default (date_trunc('hour'::text, now()) + '06:00:00'::interval) not null,
    lock_time   timestamptz,
    create_by   text                                                                         not null,
    practice    boolean     default true                                                     not null,
    problems    jsonb       default '{}'::jsonb
);

---

create table if not exists public.problem
(
    problem_id   integer generated always as identity ( start with 1000 ) not null
        primary key,
    create_at    timestamptz default now()                                not null,
    update_at    timestamptz default now()                                not null,
    delete_at    timestamptz,
    title        text                                                     not null,
    judge_type   text                                                     not null,
    time_limit   integer     default 1000                                 not null,
    memory_limit integer     default 256                                  not null,
    create_by    text                                                     not null,
    rating       integer,
    content      text
);

---

create table if not exists public.submission
(
    submission_id integer generated always as identity ( start with 100000000 ) not null
        primary key,
    create_at     timestamptz default now()                                     not null,
    update_at     timestamptz default now()                                     not null,
    delete_at     timestamptz,
    problem_id    integer                                                       not null,
    contest_id    integer,
    username      text                                                          not null,
    result        text                                                          not null,
    language      text                                                          not null,
    memory_cost   integer                                                       not null,
    time_cost     integer                                                       not null,
    code          text
);

create index if not exists submission_cid_index
    on public.submission (contest_id)
    where (contest_id IS NOT NULL);

---

create table if not exists public.user_base
(
    username  text                       not null
        primary key,
    create_at timestamptz default now()  not null,
    update_at timestamptz default now()  not null,
    delete_at timestamptz,
    password  text                       not null,
    nickname  text                       not null,
    role      text        default 'user' not null,
    solved    jsonb       default '{}'::jsonb,
    submitted jsonb       default '{}'::jsonb,
    rating    integer     default 1500   not null,
    extra     jsonb       default '{}'::jsonb
);