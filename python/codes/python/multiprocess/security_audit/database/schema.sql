DROP TABLE IF EXISTS `group_relation`;
CREATE TABLE `group_relation` (
  `gid` int(4) NOT NULL AUTO_INCREMENT,
  `user_group` varchar(64) NOT NULL COMMENT '用户组',
  `host_group` varchar(64) NOT NULL COMMENT '主机组',
  PRIMARY KEY (`user_group`,`host_group`),
  UNIQUE KEY `gid` (`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='群组关系表';


DROP TABLE IF EXISTS `host_list`;
CREATE TABLE `host_list` (
  `host_ip` varchar(64) NOT NULL COMMENT '主机ip',
  `hostname` varchar(64) NOT NULL COMMENT '主机名',
  `host_group` varchar(64) NOT NULL COMMENT '主机组',
  PRIMARY KEY (`hostname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='主机列表';

DROP TABLE IF EXISTS `host_users`;
CREATE TABLE `host_users` (
  `hostname` VARCHAR(64) NOT NULL COMMENT '主机名',
  `hostuser` VARCHAR(64) NOT NULL COMMENT '主机用户',
  `hostpwd` VARCHAR(64) NOT NULL COMMENT '主机用户密码',
  PRIMARY KEY (`hostuser`),
  FOREIGN KEY(`hostname`) REFERENCES host_list(`hostname`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '主机用户表';


DROP TABLE IF EXISTS `records`;
CREATE TABLE `records` (
  `hostname` VARCHAR(64) NOT NULL COMMENT '主机名',
  `hostuser` VARCHAR(64) NOT NULL COMMENT '主机用户',
  `records` VARCHAR(128) NOT NULL COMMENT '历史记录',
  FOREIGN KEY(`hostname`) REFERENCES host_list(`hostname`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY(`hostuser`) REFERENCES host_users(`hostuser`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户操作记录';

DROP TABLE IF EXISTS `user2group`;
CREATE TABLE `user2group` (
  `user_id` int(4) NOT NULL AUTO_INCREMENT COMMENT '用户自增id',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `after_login` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '上次登录时间',
  `user_group` varchar(64) NOT NULL COMMENT '用户组',
  PRIMARY KEY (`username`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户信息';
