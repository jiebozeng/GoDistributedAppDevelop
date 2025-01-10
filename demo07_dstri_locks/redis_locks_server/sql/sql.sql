
# 红包信息表
CREATE TABLE `ts_redpack`  (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id  自增',
    `amount` bigint NOT NULL DEFAULT 0 COMMENT '红包金额',
    `num` int NOT NULL DEFAULT 0 COMMENT '红包数量',
    `valid_time` int NOT NULL DEFAULT 0 COMMENT '有效期',
    `status` int NOT NULL DEFAULT 0 COMMENT '状态:1可用,2已结束,3已取消,5已领完',
    `pro_num` int NOT NULL DEFAULT 0 COMMENT '已领取数量',
    `created_at` datetime NULL DEFAULT NULL,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

# 红包记录表
create table `ts_redpack_record`  (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id  自增',
    `redpack_id` bigint NOT NULL DEFAULT 0 COMMENT '红包id',
    `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户id',
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `amount` bigint NOT NULL DEFAULT 0 COMMENT '领取金额',
    `created_at` datetime NULL DEFAULT NULL,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
)ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
