CREATE TABLE `t_url_data` (
	`f_id` int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
	`f_did` varchar(100) NOT NULL COMMENT '短网址',
	`f_value` varchar(1000) NOT NULL COMMENT '长网址',
	`f_create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`f_update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY (`f_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '短网址映射表'