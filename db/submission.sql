-- public.submission definition

-- Drop table

DROP TABLE public.submission;

CREATE TABLE public.submission
(
    sid         integer                 NOT NULL
        GENERATED ALWAYS AS IDENTITY (START WITH 10000000),
    pid         integer                 NOT NULL,
    cid         varchar(255),
    username    varchar(255)            NOT NULL,
    "result"    varchar(255)            NOT NULL,
    "language"  varchar(255)            NOT NULL,
    memory_cost integer                 NOT NULL,
    time_cost   integer                 NOT NULL,
    code        text                    NOT NULL,
    create_at   timestamp DEFAULT now() NOT NULL,
    update_at   timestamp DEFAULT now() NOT NULL,
    delete_at   timestamp,
    CONSTRAINT submission_pkey PRIMARY KEY (sid)
);