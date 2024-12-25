

CREATE TABLE `ts_user`  (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id  自增',
    `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户id',
    `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
    `user_pwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
    `user_mobile` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号码',
    `created_at` datetime NULL DEFAULT NULL,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;



insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('1','test1','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('2','test2','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('3','test3','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('4','test4','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('5','test5','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('6','test6','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('7','test7','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('8','test8','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('9','test9','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('10','test10','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('11','test11','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('12','test12','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('13','test13','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('14','test14','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());
insert into ts_user(user_id,user_name,user_pwd,user_mobile,created_at,updated_at) values('15','test15','E10ADC3949BA59ABBE56E057F20F883E','13800138000',now(),now());