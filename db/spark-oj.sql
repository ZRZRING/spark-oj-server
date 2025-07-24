CREATE TABLE public."content"
(
    "uuid"    uuid      DEFAULT gen_random_uuid() NOT NULL,
    "type"    int4                                NOT NULL,
    payload   jsonb                               NULL,
    create_at timestamp DEFAULT now()             NOT NULL,
    update_at timestamp DEFAULT now()             NOT NULL,
    delete_at timestamp                           NULL,
    CONSTRAINT content_pkey PRIMARY KEY (uuid)
);



CREATE TABLE public.contest
(
    cid         varchar(255)            NOT NULL,
    title       varchar(255)            NOT NULL,
    "password"  varchar(255)            NULL,
    description text                    NULL,
    start_time  timestamp               NOT NULL,
    end_time    timestamp               NOT NULL,
    lock_time   timestamp               NOT NULL,
    create_by   varchar(255)            NOT NULL,
    create_at   timestamp DEFAULT now() NOT NULL,
    update_at   timestamp DEFAULT now() NOT NULL,
    delete_at   timestamp               NULL,
    CONSTRAINT contest_pkey PRIMARY KEY (cid)
);



CREATE TABLE public.file
(
    "uuid"    uuid      DEFAULT gen_random_uuid() NOT NULL,
    "path"    varchar(255)                        NOT NULL,
    "type"    varchar(255)                        NOT NULL,
    create_at timestamp DEFAULT now()             NOT NULL,
    update_at timestamp DEFAULT now()             NOT NULL,
    delete_at timestamp                           NULL,
    CONSTRAINT file_pkey PRIMARY KEY (uuid)
);



CREATE TABLE public.judger
(
    jid       varchar(255)            NOT NULL,
    status    int4      DEFAULT 0     NOT NULL,
    create_at timestamp DEFAULT now() NOT NULL,
    update_at timestamp DEFAULT now() NOT NULL,
    delete_at timestamp               NULL,
    CONSTRAINT judger_pkey PRIMARY KEY (jid)
);



CREATE TABLE public.problem
(
    pid          varchar(255)            NOT NULL,
    title        varchar(255)            NOT NULL,
    judge_type   int4      DEFAULT 0     NOT NULL,
    time_limit   int4      DEFAULT 1000  NOT NULL,
    memory_limit int4      DEFAULT 1024  NOT NULL,
    create_by    varchar(255)            NOT NULL,
    rating       int4                    NULL,
    create_at    timestamp DEFAULT now() NOT NULL,
    update_at    timestamp DEFAULT now() NOT NULL,
    delete_at    timestamp               NULL,
    CONSTRAINT problem_pkey PRIMARY KEY (pid)
);



CREATE TABLE public.problem_content
(
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    cid    varchar(255)                   NOT NULL,
    pid    varchar(255)                   NOT NULL,
    CONSTRAINT problem_content_pkey PRIMARY KEY (uuid)
);



CREATE TABLE public.submission
(
    sid        int4                    NOT NULL,
    pid        varchar(255)            NOT NULL,
    uid        varchar(255)            NOT NULL,
    cid        varchar(255)            NULL,
    "result"   varchar(255)            NOT NULL,
    "language" varchar(255)            NOT NULL,
    memory     int4                    NOT NULL,
    "time"     int4                    NOT NULL,
    code       text                    NOT NULL,
    create_at  timestamp DEFAULT now() NOT NULL,
    update_at  timestamp DEFAULT now() NOT NULL,
    delete_at  timestamp               NULL,
    CONSTRAINT submission_pkey PRIMARY KEY (sid)
);



CREATE TABLE public.tag
(
    tag       varchar(255)            NOT NULL,
    approved  bool      DEFAULT false NOT NULL,
    create_at timestamp DEFAULT now() NOT NULL,
    update_at timestamp DEFAULT now() NOT NULL,
    delete_at timestamp               NOT NULL,
    CONSTRAINT tag_pkey PRIMARY KEY (tag)
);



CREATE TABLE public.user_base
(
    username   varchar(255)                                   NOT NULL,
    "password" varchar(255)                                   NOT NULL,
    nickname   varchar(255)                                   NOT NULL,
    "role"     varchar(255) DEFAULT 'user'::character varying NOT NULL,
    solved     int4         DEFAULT 0                         NOT NULL,
    submitted  int4         DEFAULT 0                         NOT NULL,
    rating     int4         DEFAULT 1500                      NOT NULL,
    cf_id      varchar(255)                                   NULL,
    atc_id     varchar(255)                                   NULL,
    company    varchar(255)                                   NULL,
    department varchar(255)                                   NULL,
    major      varchar(255)                                   NULL,
    "class"    varchar(255)                                   NULL,
    email      varchar(255)                                   NULL,
    tel        varchar(255)                                   NULL,
    avatar     varchar(255)                                   NULL,
    create_at  timestamp    DEFAULT now()                     NOT NULL,
    update_at  timestamp    DEFAULT now()                     NOT NULL,
    delete_at  timestamp                                      NULL,
    CONSTRAINT user_base_pkey PRIMARY KEY (username)
);



CREATE TABLE public.user_role
(
    rid          varchar(255)                  NOT NULL,
    "permission" jsonb     DEFAULT '{}'::jsonb NOT NULL,
    create_at    timestamp DEFAULT now()       NOT NULL,
    update_at    timestamp DEFAULT now()       NOT NULL,
    delete_at    timestamp                     NULL,
    CONSTRAINT user_role_pkey PRIMARY KEY (rid)
);



CREATE TABLE public.notice
(
    uuid         uuid      default gen_random_uuid() NOT NULL,
    "permission" jsonb     DEFAULT '{}'::jsonb       NOT NULL,
    create_at    timestamp DEFAULT now()             NOT NULL,
    update_at    timestamp DEFAULT now()             NOT NULL,
    delete_at    timestamp                           NULL,
    CONSTRAINT user_role_pkey PRIMARY KEY (uuid)
);