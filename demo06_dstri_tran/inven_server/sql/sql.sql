
# 商品表
CREATE TABLE `ts_product`  (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id  自增',
    `product_name` varchar(64) NOT NULL DEFAULT '' COMMENT '商品名称',
    `inven_num` int NOT NULL DEFAULT 0 COMMENT '商品库存',
    `status` int NOT NULL DEFAULT 2 COMMENT '状态:1审核中,2已上架,3以下架',
    `price` decimal(10, 2) NOT NULL DEFAULT 0 COMMENT '价格',
    `created_at` datetime NULL DEFAULT NULL,
    `updated_at` datetime NULL DEFAULT NULL,
    `deleted_at` datetime NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

insert into ts_product(product_name,inven_num,status,price,created_at,updated_at) values('测试商品1',1000,2,90.00,now(),now());
insert into ts_product(product_name,inven_num,status,price,created_at,updated_at) values('测试商品2',1000,2,80.00,now(),now());