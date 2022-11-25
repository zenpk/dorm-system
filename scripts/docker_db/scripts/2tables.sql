USE `dorm`;
CREATE TABLE `dorms`
(
    `id`          bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `num`         varchar(10) UNIQUE NOT NULL,
    `building_id` bigint unsigned    NOT NULL,
    `gender`      varchar(10)        NOT NULL,
    `remain_cnt`  bigint unsigned    NOT NULL,
    `bed_cnt`     bigint unsigned    NOT NULL,
    `info`        varchar(200),
    `enabled`     tinyint            NOT NULL DEFAULT 1,
    `deleted`     datetime(3)
);
ALTER TABLE `dorms`
    ADD INDEX (`num`);
ALTER TABLE `dorms`
    ADD INDEX (`building_id`);
ALTER TABLE `dorms`
    ADD INDEX (`enabled`);
ALTER TABLE `dorms`
    ADD INDEX (`deleted`);

CREATE TABLE `buildings`
(
    `id`        bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `num`       varchar(10) UNIQUE NOT NULL,
    `info`      varchar(200),
    `image_url` varchar(200),
    `enabled`   tinyint            NOT NULL DEFAULT 1,
    `deleted`   datetime(3)
);
ALTER TABLE `buildings`
    ADD INDEX (`num`);
ALTER TABLE `buildings`
    ADD INDEX (`enabled`);
ALTER TABLE `buildings`
    ADD INDEX (`deleted`);

CREATE TABLE `orders`
(
    `id`          bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `building_id` bigint unsigned     NOT NULL,
    `dorm_id`     bigint unsigned     NOT NULL DEFAULT 0,
    `team_id`     bigint unsigned     NOT NULL,
    `code`        varchar(200) UNIQUE NOT NULL,
    `info`        varchar(200),
    `success`     tinyint             NOT NULL DEFAULT 0,
    `deleted`     datetime(3)
);
ALTER TABLE `orders`
    ADD INDEX (`building_id`);
ALTER TABLE `orders`
    ADD INDEX (`dorm_id`);
ALTER TABLE `orders`
    ADD INDEX (`team_id`);
ALTER TABLE `orders`
    ADD INDEX (`code`);
ALTER TABLE `orders`
    ADD INDEX (`deleted`);

CREATE TABLE `teams`
(
    `id`       bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `code`     varchar(200) UNIQUE    NOT NULL,
    `gender`   varchar(10)            NOT NULL,
    `owner_id` bigint unsigned UNIQUE NOT NULL,
    `deleted`  datetime(3)
);
ALTER TABLE `teams`
    ADD INDEX (`code`);
ALTER TABLE `teams`
    ADD INDEX (`owner_id`);
ALTER TABLE `teams`
    ADD INDEX (`deleted`);

CREATE TABLE `team_users`
(
    `id`      bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `team_id` bigint unsigned NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `deleted` datetime(3)
);
ALTER TABLE `team_users`
    ADD INDEX (`team_id`);
ALTER TABLE `team_users`
    ADD INDEX (`user_id`);
ALTER TABLE `team_users`
    ADD INDEX (`deleted`);

CREATE TABLE `accounts`
(
    `id`       bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `user_id`  bigint unsigned     NOT NULL,
    `username` varchar(200) UNIQUE NOT NULL,
    `password` varchar(200)        NOT NULL,
    `deleted`  datetime(3)
);
ALTER TABLE `accounts`
    ADD INDEX (`user_id`);
ALTER TABLE `accounts`
    ADD INDEX (`username`);
ALTER TABLE `accounts`
    ADD INDEX (`deleted`);

CREATE TABLE `users`
(
    `id`          bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `student_num` varchar(200) UNIQUE NOT NULL,
    `name`        varchar(20)         NOT NULL,
    `gender`      varchar(10)         NOT NULL,
    `role`        int                 NOT NULL DEFAULT 1,
    `deleted`     datetime(3)
);
ALTER TABLE `users`
    ADD INDEX (`name`);
ALTER TABLE `users`
    ADD INDEX (`gender`);
ALTER TABLE `users`
    ADD INDEX (`role`);
ALTER TABLE `users`
    ADD INDEX (`deleted`);

CREATE TABLE `tokens`
(
    `id`            bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `refresh_token` varchar(200)    NOT NULL,
    `user_id`       bigint unsigned NOT NULL,
    `create_time`   datetime(3)     NOT NULL,
    `exp_time`      datetime(3)     NOT NULL,
    `deleted`       datetime(3)
);
ALTER TABLE `tokens`
    ADD INDEX (`refresh_token`);
ALTER TABLE `tokens`
    ADD INDEX (`user_id`);
ALTER TABLE `tokens`
    ADD INDEX (`deleted`);