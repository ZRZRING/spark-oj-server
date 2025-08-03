-- public.submission definition

-- Drop table

DROP TABLE public.submission;

CREATE TABLE public.submission
(
    sid         int4                    NOT NULL
        GENERATED ALWAYS AS IDENTITY (START WITH 10000000),
    pid         varchar(255)            NOT NULL,
    uid         varchar(255)            NOT NULL,
    cid         varchar(255)            NULL,
    "result"    varchar(255)            NOT NULL,
    "language"  varchar(255)            NOT NULL,
    memory_cost int4                    NOT NULL,
    time_cost   int4                    NOT NULL,
    code        text                    NOT NULL,
    create_at   timestamp DEFAULT now() NOT NULL,
    update_at   timestamp DEFAULT now() NOT NULL,
    delete_at   timestamp               NULL,
    CONSTRAINT submission_pkey PRIMARY KEY (sid)
);