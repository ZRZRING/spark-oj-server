create table user_base
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
    create_at  timestamp    default now()                     not null,
    update_at  timestamp    default now()                     not null,
    delete_at  timestamp
);

alter table user_base
    owner to "spark-oj";

INSERT INTO public.user_base (username, password, nickname, role, solved, submitted, rating, cf_id, atc_id, company, department, major, class, email, tel, avatar, create_at, update_at, delete_at) VALUES ('admin', '21232f297a57a5a743894a0e4a801fc3', 'admin', 'user', 0, 0, 1500, null, null, null, null, null, null, null, null, null, '2025-05-28 17:09:26.969074', '2025-05-28 17:09:26.969074', null);
INSERT INTO public.user_base (username, password, nickname, role, solved, submitted, rating, cf_id, atc_id, company, department, major, class, email, tel, avatar, create_at, update_at, delete_at) VALUES ('test', 'e10adc3949ba59abbe56e057f20f883e', 'test', 'user', 0, 0, 1500, null, null, null, null, null, null, null, null, null, '2025-06-04 11:59:04.095885', '2025-06-04 11:59:04.095885', null);
INSERT INTO public.user_base (username, password, nickname, role, solved, submitted, rating, cf_id, atc_id, company, department, major, class, email, tel, avatar, create_at, update_at, delete_at) VALUES ('zrzring', 'f7f29b21dbc871560b63eb5dff8f58f4', 'zrzring', 'user', 0, 0, 1500, 'zrzring', 'zrzring', null, null, null, null, null, null, null, '2025-05-28 17:09:42.139981', '2025-05-28 17:09:42.139981', null);
