#创建数据库
CREATE DATABASE `godist_segmentsids` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';
use godist_segmentsids;
#创建表
CREATE TABLE id_generator (
    `id`            BIGINT NOT NULL AUTO_INCREMENT,
    `max_id`        BIGINT NOT NULL COMMENT '当前最大id',
    `step`          BIGINT NOT NULL COMMENT '号段的步长',
    `biz_type`      int NOT NULL COMMENT '业务类型',
    `version`       BIGINT NOT NULL COMMENT '版本号',
    `created_at`    DATETIME,
    `updated_at`    DATETIME,
    `deleted_at`    DATETIME,
    PRIMARY KEY (`id`)
);

#初始化两条记录 初始max_id 为1 1 步长为1000 500 业务类型 1 2 版本号为 1 1
INSERT INTO id_generator (max_id, step, biz_type, version,created_at,updated_at) VALUES (1, 1000, 1, 1,NOW(),NOW());
INSERT INTO id_generator (max_id, step, biz_type, version,created_at,updated_at) VALUES (1, 500, 2, 1,NOW(),NOW());