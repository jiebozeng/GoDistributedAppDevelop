
# 订单表
CREATE TABLE `ts_order`  (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id  自增',
    `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户id',
    `product_id` bigint NOT NULL DEFAULT 0 COMMENT '商品id',
    `num` int NOT NULL DEFAULT 0 COMMENT '红包数量',
    `status` int NOT NULL DEFAULT 0 COMMENT '状态:1未支付,2已支付,3已取消,5已发货,6已签收,7已关闭',
    `amount` decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '总价',
    `created_at` datetime NULL DEFAULT NULL,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

