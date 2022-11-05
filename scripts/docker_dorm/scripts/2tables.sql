USE `dorm`;
CREATE TABLE `dorms`
(
    `id`          bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `num`         varchar(10) UNIQUE NOT NULL,
    `building_id` bigint unsigned    NOT NULL,
    `gender`      varchar(10)        NOT NULL,
    `remain_cnt`  bigint unsigned    NOT NULL,
    `bed_cnt`     bigint unsigned    NOT NULL,
    `enabled`     tinyint            NOT NULL DEFAULT 1,
    `info`        varchar(200)
);

ALTER TABLE `dorms`
    ADD INDEX (`num`);

ALTER TABLE `dorms`
    ADD INDEX (`building_id`);

ALTER TABLE `dorms`
    ADD INDEX (`enabled`);

CREATE TABLE `buildings`
(
    `id`      bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `num`     varchar(10) UNIQUE NOT NULL,
    `enabled` tinyint            NOT NULL DEFAULT 1,
    `info`    varchar(200)
);

ALTER TABLE `buildings`
    ADD INDEX (`num`);

ALTER TABLE `buildings`
    ADD INDEX (`enabled`);

CREATE TABLE `orders`
(
    `id`           bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `dorm_id`      bigint unsigned        NOT NULL,
    `student_id_1` bigint unsigned UNIQUE NOT NULL,
    `student_id_2` bigint unsigned,
    `student_id_3` bigint unsigned,
    `student_id_4` bigint unsigned
);

ALTER TABLE `orders`
    ADD INDEX (`dorm_id`);

CREATE TABLE `team`
(
    `id`            bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `student_num_1` varchar(200) UNIQUE NOT NULL,
    `student_num_2` varchar(200),
    `student_num_3` varchar(200),
    `student_num_4` varchar(200)
);

ALTER TABLE `team`
    ADD INDEX (`student_num_1`);

CREATE TABLE `user_credentials`
(
    `id`       bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `username` varchar(200) UNIQUE NOT NULL,
    `password` varchar(200)        NOT NULL
);
ALTER TABLE `user_credentials`
    ADD INDEX (`username`);

CREATE TABLE `user_infos`
(
    `id`            bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `credential_id` bigint unsigned UNIQUE NOT NULL,
    `username`      varchar(200) UNIQUE    NOT NULL,
    `student_num`   varchar(200) UNIQUE    NOT NULL,
    `name`          varchar(20)            NOT NULL,
    `gender`        varchar(10)            NOT NULL,
    `dorm_id`       bigint unsigned        NOT NULL DEFAULT 0
);

ALTER TABLE `user_infos`
    ADD INDEX (`credential_id`);

ALTER TABLE `user_infos`
    ADD INDEX (`username`);

ALTER TABLE `user_infos`
    ADD INDEX (`student_num`);

ALTER TABLE `user_infos`
    ADD INDEX (`name`);

ALTER TABLE `user_infos`
    ADD INDEX (`gender`);

ALTER TABLE `user_infos`
    ADD INDEX (`dorm_id`);
