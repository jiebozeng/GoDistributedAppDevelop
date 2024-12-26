-- 创建用户表模板
CREATE TABLE user_template (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户ID',
    `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
    `user_pwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
    `user_mobile` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号码',
    `user_email` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
    `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- 使用模板生成具体表
CREATE TABLE `user_1` LIKE `user_template`;
CREATE TABLE `user_2` LIKE `user_template`;
CREATE TABLE `user_3` LIKE `user_template`;
CREATE TABLE `user_4` LIKE `user_template`;
CREATE TABLE `user_5` LIKE `user_template`;
CREATE TABLE `user_6` LIKE `user_template`;
CREATE TABLE `user_7` LIKE `user_template`;
CREATE TABLE `user_8` LIKE `user_template`;
CREATE TABLE `user_9` LIKE `user_template`;
CREATE TABLE `user_10` LIKE `user_template`;
