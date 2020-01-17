CREATE DATABASE `sailing_user` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;

USE `sailing_user`;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20) unsigned                                           NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `uid`         varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户uid',
    `type`        enum('default','official','merchant') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default' COMMENT '用户类型',
    `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Email',
    `mobile` varchar(62) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
    `nickname`    varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '昵称',
    `username`    varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`    varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
    `salt`        char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci      NOT NULL DEFAULT '' COMMENT '盐值',
    `active` tinyint(1) NOT NULL DEFAULT '1' COMMENT '活动',
    `status` enum('normal','disable','deleted') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'normal' COMMENT '状态',
    `created_at`  timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at`  timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at`  timestamp(3)                                                  DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_uid_unique_index` (`uid`),
    UNIQUE KEY `user_mobile_unique_index` (`mobile`),
    UNIQUE KEY `user_nickname_unique_index` (`nickname`),
    UNIQUE KEY `user_username_unique_index` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户表';

INSERT INTO `user` (uid, username, password) VALUE ('10001', 'micro', '123456');