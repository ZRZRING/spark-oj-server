create table public.contest
(
    contest_id  integer generated always as identity (start with 1000)
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

create table public.problem
(
    problem_id   integer generated always as identity (start with 1000)
        primary key,
    create_at    timestamptz default now() not null,
    update_at    timestamptz default now() not null,
    delete_at    timestamptz,
    title        text                      not null,
    judge_type   text                      not null,
    time_limit   integer     default 1000  not null,
    memory_limit integer     default 256   not null,
    create_by    text                      not null,
    rating       integer,
    content      text
);

---

create table public.resource
(
    resource_uuid uuid        default gen_random_uuid() not null
        primary key,
    create_at     timestamptz default now()             not null,
    update_at     timestamptz default now()             not null,
    delete_at     timestamptz,
    path          text                                  not null,
    type          text                                  not null
);

---

create table public.submission
(
    submission_id integer generated always as identity (start with 10000000)
        primary key,
    create_at     timestamptz default now() not null,
    update_at     timestamptz default now() not null,
    delete_at     timestamptz,
    problem_id           text                      not null,
    contest_id           text,
    username      text                      not null,
    result        text                      not null,
    language      text                      not null,
    memory_cost   integer                   not null,
    time_cost     integer                   not null,
    code          text
);

create index submission_cid_index
    on public.submission (contest_id)
    where (contest_id IS NOT NULL);

---

create table public.user_base
(
    username  text                                          not null
        primary key,
    create_at timestamptz default now()                     not null,
    update_at timestamptz default now()                     not null,
    delete_at timestamptz,
    password  text                                          not null,
    nickname  text                                          not null,
    role      text        default 'user'::character varying not null,
    solved    integer     default 0                         not null,
    submitted integer     default 0                         not null,
    rating    integer     default 1500                      not null,
    extra     jsonb       default '{}'::jsonb
);

---

create table public.user_role
(
    role_id    text                            not null
        primary key,
    create_at  timestamptz default now()       not null,
    update_at  timestamptz default now()       not null,
    delete_at  timestamptz,
    permission jsonb       default '{}'::jsonb not null
);