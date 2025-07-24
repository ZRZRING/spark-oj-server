-- lduoj.contest_balloons definition

CREATE TABLE `contest_balloons`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `solution_id` bigint                   DEFAULT NULL,
    `sent`        tinyint(1)      NOT NULL DEFAULT '0',
    `send_time`   datetime                 DEFAULT NULL,
    `created_at`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.contest_cate definition

CREATE TABLE `contest_cate`
(
    `id`          bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `title`       varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'unamed',
    `description` text COLLATE utf8mb4_unicode_ci,
    `hidden`      tinyint(1)                              NOT NULL DEFAULT '0',
    `order`       bigint                                  NOT NULL DEFAULT '0',
    `parent_id`   bigint                                  NOT NULL DEFAULT '0',
    `created_at`  datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `contest_cate_order_index` (`order`),
    KEY `contest_cate_parent_id_index` (`parent_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.contest_notices definition

CREATE TABLE `contest_notices`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `contest_id` bigint                                           DEFAULT NULL,
    `title`      varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'unamed',
    `content`    text COLLATE utf8mb4_unicode_ci,
    `user_id`    bigint                                           DEFAULT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `contest_notices_contest_id_index` (`contest_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.contest_problems definition

CREATE TABLE `contest_problems`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `contest_id` bigint          NOT NULL,
    `index`      int             NOT NULL,
    `problem_id` bigint          NOT NULL,
    `solved`     int             NOT NULL DEFAULT '0',
    `accepted`   int             NOT NULL DEFAULT '0',
    `submitted`  int             NOT NULL DEFAULT '0',
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `contest_problems_contest_id_index` (`contest_id`),
    KEY `contest_problems_problem_id_index` (`problem_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.contest_users definition

CREATE TABLE `contest_users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `contest_id` bigint          NOT NULL,
    `user_id`    bigint          NOT NULL,
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `contest_users_contest_id_index` (`contest_id`),
    KEY `contest_users_user_id_index` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.contests definition

CREATE TABLE `contests`
(
    `id`                bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `judge_type`        varchar(5) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT 'acm',
    `enable_discussing` tinyint(1)                              NOT NULL DEFAULT '0',
    `enable_tagging`    tinyint(1)                              NOT NULL DEFAULT '0',
    `title`             varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'unamed',
    `description`       text COLLATE utf8mb4_unicode_ci,
    `allow_lang`        bigint                                  NOT NULL DEFAULT '15' COMMENT '按位标记允许的语言',
    `start_time`        datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `end_time`          datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `lock_rate`         double(8, 2)                            NOT NULL DEFAULT '0.00' COMMENT '0~1,封榜比例',
    `public_rank`       tinyint(1)                              NOT NULL DEFAULT '0',
    `access`            varchar(10) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT 'public' COMMENT 'public,password,private',
    `password`          varchar(40) COLLATE utf8mb4_unicode_ci           DEFAULT NULL,
    `user_id`           bigint                                           DEFAULT NULL,
    `num_members`       int                                     NOT NULL DEFAULT '0',
    `sections`          json                                             DEFAULT NULL,
    `hidden`            tinyint(1)                              NOT NULL DEFAULT '0',
    `order`             bigint                                  NOT NULL DEFAULT '0',
    `cate_id`           bigint                                  NOT NULL DEFAULT '0',
    `created_at`        datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `ban_code_editor`   tinyint                                 NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `contests_start_time_index` (`start_time`),
    KEY `contests_end_time_index` (`end_time`),
    KEY `contests_user_id_index` (`user_id`),
    KEY `contests_hidden_index` (`hidden`),
    KEY `contests_order_index` (`order`),
    KEY `contests_cate_id_index` (`cate_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1000
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.discussions definition

CREATE TABLE `discussions`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `problem_id`     bigint          NOT NULL,
    `discussion_id`  bigint                                  DEFAULT NULL,
    `reply_username` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `username`       varchar(60) COLLATE utf8mb4_unicode_ci  DEFAULT NULL,
    `content`        text COLLATE utf8mb4_unicode_ci,
    `top`            int             NOT NULL                DEFAULT '0',
    `hidden`         tinyint(1)      NOT NULL                DEFAULT '0',
    `created_at`     datetime        NOT NULL                DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     datetime        NOT NULL                DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `discussions_problem_id_index` (`problem_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.failed_jobs definition

CREATE TABLE `failed_jobs`
(
    `id`         bigint unsigned                     NOT NULL AUTO_INCREMENT,
    `connection` text COLLATE utf8mb4_unicode_ci     NOT NULL,
    `queue`      text COLLATE utf8mb4_unicode_ci     NOT NULL,
    `payload`    longtext COLLATE utf8mb4_unicode_ci NOT NULL,
    `exception`  longtext COLLATE utf8mb4_unicode_ci NOT NULL,
    `failed_at`  datetime                            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.group_contests definition

CREATE TABLE `group_contests`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `contest_id` bigint          NOT NULL,
    `group_id`   bigint          NOT NULL,
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `order`      int             NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `group_contests_contest_id_index` (`contest_id`),
    KEY `group_contests_group_id_index` (`group_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.group_users definition

CREATE TABLE `group_users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `group_id`   bigint          NOT NULL,
    `user_id`    bigint          NOT NULL,
    `identity`   tinyint         NOT NULL DEFAULT '0',
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `archive`    json                     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `group_users_group_id_index` (`group_id`),
    KEY `group_users_user_id_index` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.`groups` definition

CREATE TABLE `groups`
(
    `id`             bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `type`           tinyint                                 NOT NULL DEFAULT '0',
    `hidden`         tinyint(1)                              NOT NULL DEFAULT '1',
    `private`        tinyint(1)                              NOT NULL DEFAULT '1',
    `unlock_contest` tinyint(1)                              NOT NULL DEFAULT '0',
    `name`           varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'unamed',
    `description`    text COLLATE utf8mb4_unicode_ci,
    `teacher`        varchar(255) COLLATE utf8mb4_unicode_ci          DEFAULT NULL COMMENT 'teacher''s name',
    `class`          varchar(255) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `user_id`        bigint                                           DEFAULT NULL,
    `num_members`    int                                     NOT NULL DEFAULT '0',
    `num_problems`   int                                     NOT NULL DEFAULT '0',
    `created_at`     datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `archive_cite`   json                                             DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `groups_class_index` (`class`),
    KEY `groups_user_id_index` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.migrations definition

CREATE TABLE `migrations`
(
    `id`        int unsigned                            NOT NULL AUTO_INCREMENT,
    `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `batch`     int                                     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 16
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.notices definition

CREATE TABLE `notices`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `title`      varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'unamed',
    `content`    text COLLATE utf8mb4_unicode_ci,
    `state`      tinyint                                 NOT NULL DEFAULT '1' COMMENT '0:hidden,1:normal,2:top',
    `order`      int                                     NOT NULL DEFAULT '0',
    `hidden`     tinyint(1)                              NOT NULL DEFAULT '0',
    `user_id`    bigint                                  NOT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.password_resets definition

CREATE TABLE `password_resets`
(
    `email`      varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `token`      varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `password_resets_email_index` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.permissions definition

CREATE TABLE `permissions`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `name`       varchar(125) COLLATE utf8mb4_unicode_ci NOT NULL,
    `guard_name` varchar(125) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `permissions_name_guard_name_unique` (`name`, `guard_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 63
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.personal_access_tokens definition

CREATE TABLE `personal_access_tokens`
(
    `id`             bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `tokenable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `tokenable_id`   bigint unsigned                         NOT NULL,
    `name`           varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `token`          varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL,
    `abilities`      text COLLATE utf8mb4_unicode_ci,
    `last_used_at`   timestamp                               NULL DEFAULT NULL,
    `expires_at`     timestamp                               NULL DEFAULT NULL,
    `created_at`     timestamp                               NULL DEFAULT NULL,
    `updated_at`     timestamp                               NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
    KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`, `tokenable_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.problems definition

CREATE TABLE `problems`
(
    `id`            bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `type`          tinyint                                 NOT NULL DEFAULT '0' COMMENT '0:编程,1:代码填空',
    `title`         varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'unamed',
    `description`   text COLLATE utf8mb4_unicode_ci,
    `input`         text COLLATE utf8mb4_unicode_ci,
    `output`        text COLLATE utf8mb4_unicode_ci,
    `hint`          text COLLATE utf8mb4_unicode_ci,
    `source`        varchar(255) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `tags`          json                                             DEFAULT NULL,
    `samples`       json                                             DEFAULT NULL,
    `fill_in_blank` text COLLATE utf8mb4_unicode_ci,
    `language`      int                                     NOT NULL DEFAULT '0' COMMENT '代码填空的语言',
    `spj`           tinyint(1)                              NOT NULL DEFAULT '0',
    `spj_language`  int                                     NOT NULL DEFAULT '14',
    `time_limit`    int                                     NOT NULL DEFAULT '1000' COMMENT 'MS',
    `memory_limit`  int                                     NOT NULL DEFAULT '1024' COMMENT 'MB',
    `hidden`        tinyint(1)                              NOT NULL DEFAULT '1',
    `user_id`       bigint                                           DEFAULT NULL,
    `level`         int                                     NOT NULL DEFAULT '0' COMMENT '0:null,1:easy,2:middle,3:difficult',
    `solved`        int                                     NOT NULL DEFAULT '0',
    `accepted`      int                                     NOT NULL DEFAULT '0',
    `submitted`     int                                     NOT NULL DEFAULT '0',
    `created_at`    datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `problems_type_index` (`type`),
    KEY `problems_hidden_index` (`hidden`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1000
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.roles definition

CREATE TABLE `roles`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `name`       varchar(125) COLLATE utf8mb4_unicode_ci NOT NULL,
    `guard_name` varchar(125) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `roles_name_guard_name_unique` (`name`, `guard_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.settings definition

CREATE TABLE `settings`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `key`        varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `value`      varchar(255) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.solutions definition

CREATE TABLE `solutions`
(
    `id`           bigint unsigned                       NOT NULL AUTO_INCREMENT,
    `problem_id`   bigint                                NOT NULL,
    `contest_id`   bigint                                NOT NULL DEFAULT '-1',
    `user_id`      bigint                                         DEFAULT NULL,
    `judge_type`   varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'oi' COMMENT 'acm,oi',
    `result`       int                                   NOT NULL DEFAULT '0',
    `time`         int                                   NOT NULL DEFAULT '0',
    `memory`       double(8, 2)                          NOT NULL DEFAULT '0.00',
    `language`     int                                   NOT NULL DEFAULT '0',
    `submit_time`  datetime                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `judge_time`   datetime                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `judge_result` json                                           DEFAULT NULL,
    `pass_rate`    double(8, 2)                          NOT NULL DEFAULT '0.00',
    `error_info`   text COLLATE utf8mb4_unicode_ci,
    `wrong_data`   text COLLATE utf8mb4_unicode_ci,
    `ip`           varchar(16) COLLATE utf8mb4_unicode_ci         DEFAULT NULL,
    `ip_loc`       varchar(64) COLLATE utf8mb4_unicode_ci         DEFAULT NULL,
    `judger`       varchar(60) COLLATE utf8mb4_unicode_ci         DEFAULT NULL,
    `code_length`  int                                   NOT NULL DEFAULT '0',
    `code`         text COLLATE utf8mb4_unicode_ci,
    `sim_rate`     int                                   NOT NULL DEFAULT '0' COMMENT '0~100',
    `sim_sid`      bigint                                         DEFAULT NULL,
    `sim_report`   json                                           DEFAULT NULL,
    `created_at`   datetime                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `solutions_problem_id_index` (`problem_id`),
    KEY `solutions_contest_id_index` (`contest_id`),
    KEY `solutions_user_id_index` (`user_id`),
    KEY `solutions_result_index` (`result`),
    KEY `solutions_language_index` (`language`),
    KEY `solutions_ip_index` (`ip`),
    KEY `solutions_submit_time_index` (`submit_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1000
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.tag_marks definition

CREATE TABLE `tag_marks`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `problem_id` bigint          NOT NULL,
    `user_id`    bigint          NOT NULL,
    `tag_id`     bigint          NOT NULL,
    `created_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `tag_marks_problem_id_index` (`problem_id`),
    KEY `tag_marks_user_id_index` (`user_id`),
    KEY `tag_marks_tag_id_index` (`tag_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.tag_pool definition

CREATE TABLE `tag_pool`
(
    `id`         bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `name`       varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
    `user_id`    bigint                                           DEFAULT NULL,
    `hidden`     tinyint(1)                              NOT NULL DEFAULT '0',
    `created_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.users definition

CREATE TABLE `users`
(
    `id`                bigint unsigned                        NOT NULL AUTO_INCREMENT,
    `username`          varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
    `email`             varchar(100) COLLATE utf8mb4_unicode_ci         DEFAULT NULL,
    `email_verified_at` timestamp                              NULL     DEFAULT NULL,
    `password`          varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
    `nick`              varchar(50) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `school`            varchar(50) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `class`             varchar(50) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `revise`            tinyint(1)                             NOT NULL DEFAULT '1' COMMENT '允许修改个人信息',
    `locked`            tinyint(1)                             NOT NULL DEFAULT '0' COMMENT '锁定：禁止用户访问网站',
    `solved`            int                                    NOT NULL DEFAULT '0',
    `accepted`          int                                    NOT NULL DEFAULT '0',
    `submitted`         int                                    NOT NULL DEFAULT '0',
    `api_token`         varchar(65) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `remember_token`    varchar(100) COLLATE utf8mb4_unicode_ci         DEFAULT NULL,
    `created_at`        datetime                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        datetime                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_username_unique` (`username`),
    UNIQUE KEY `users_api_token_unique` (`api_token`),
    KEY `users_nick_index` (`nick`),
    KEY `users_school_index` (`school`),
    KEY `users_class_index` (`class`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.model_has_permissions definition

CREATE TABLE `model_has_permissions`
(
    `permission_id` bigint unsigned                         NOT NULL,
    `model_type`    varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `model_id`      bigint unsigned                         NOT NULL,
    PRIMARY KEY (`permission_id`, `model_id`, `model_type`),
    KEY `model_has_permissions_model_id_model_type_index` (`model_id`, `model_type`),
    CONSTRAINT `model_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.model_has_roles definition

CREATE TABLE `model_has_roles`
(
    `role_id`    bigint unsigned                         NOT NULL,
    `model_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `model_id`   bigint unsigned                         NOT NULL,
    PRIMARY KEY (`role_id`, `model_id`, `model_type`),
    KEY `model_has_roles_model_id_model_type_index` (`model_id`, `model_type`),
    CONSTRAINT `model_has_roles_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


-- lduoj.role_has_permissions definition

CREATE TABLE `role_has_permissions`
(
    `permission_id` bigint unsigned NOT NULL,
    `role_id`       bigint unsigned NOT NULL,
    PRIMARY KEY (`permission_id`, `role_id`),
    KEY `role_has_permissions_role_id_foreign` (`role_id`),
    CONSTRAINT `role_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
    CONSTRAINT `role_has_permissions_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;