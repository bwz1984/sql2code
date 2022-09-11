CREATE TABLE `t_student` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `age` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `height` int unsigned NOT NULL DEFAULT '0' COMMENT '身高',
  `weight` int unsigned NOT NULL DEFAULT '0' COMMENT '体重',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='学生表'