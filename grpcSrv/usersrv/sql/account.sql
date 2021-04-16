
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账户ID',
  `mobile` varchar(11) NOT NULL COMMENT '账户编号,手机号码 ',
  `password` varchar(64) NOT NULL COMMENT '账号密码',
  `nickname` varchar(64) DEFAULT 'no nickname' COMMENT '账户名称,用来说明账户的简短描述,账户对应的名称或者命名',
  `headurl` varchar(64)  DEFAULT 'no url' COMMENT 'url',
  `address` varchar(64) DEFAULT 'no address' COMMENT '用户地址',
  `desc` varchar(64) DEFAULT 'no desc' COMMENT '详情描述',
  `birthday` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '用户生日',
  `gender` tinyint(2) DEFAULT 1 COMMENT  '用户性别',
  `role` tinyint(2) NOT NULL DEFAULT 0 COMMENT '用户角色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `mobile_idx` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=171 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
