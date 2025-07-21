create table problem
(
    pid          varchar(255)            not null primary key,
    title        varchar(255)            not null,
    judge_type   integer   default 0     not null,
    time_limit   integer   default 1000  not null,
    memory_limit integer   default 1024  not null,
    create_by    varchar(255)            not null,
    rating       integer,
    config       jsonb,
    create_at    timestamp default now() not null,
    update_at    timestamp default now() not null,
    delete_at    timestamp
);

alter table problem
    owner to "spark-oj";

INSERT INTO public.problem (pid, title, judge_type, time_limit, memory_limit, create_by, rating, create_at, update_at, delete_at) VALUES ('P1000', 'A+B problem', 1, 1000, 1024, 'zrzring', 800, '2025-07-18 06:48:33.658444', '2025-07-18 06:48:33.658444', null);
INSERT INTO public.problem (pid, title, judge_type, time_limit, memory_limit, create_by, rating, create_at, update_at, delete_at) VALUES ('P1001', '八皇后问题', 1, 2000, 1024, 'zrzring', 1200, '2025-07-18 14:50:17.000000', '2025-07-18 06:50:22.145422', null);
INSERT INTO public.problem (pid, title, judge_type, time_limit, memory_limit, create_by, rating, create_at, update_at, delete_at) VALUES ('P1002', '爬楼梯', 1, 1000, 1024, 'zrzring', 1000, '2025-07-18 06:50:51.424871', '2025-07-18 06:50:51.424871', null);
INSERT INTO public.problem (pid, title, judge_type, time_limit, memory_limit, create_by, rating, create_at, update_at, delete_at) VALUES ('P1003', '01背包', 1, 1000, 1024, 'zrzring', 900, '2025-07-18 06:51:37.577754', '2025-07-18 06:51:37.577754', null);
