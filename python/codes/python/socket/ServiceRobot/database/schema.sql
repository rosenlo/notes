DROP TABLE IF EXISTS `chat_records`;

CREATE TABLE `chat_records` (
 `user_id` int(4) NOT NULL AUTO_INCREMENT COMMENT '用户自增id',
 `username` varchar(24) NOT NULL COMMENT '用户名',
 `sender` varchar(24) NOT NULL COMMENT '发送者',
 `recevier` varchar(24) NOT NULL COMMENT '接受者',
 `text` varchar(256) NOT NULL COMMENT '文本内容',
 `send_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
 PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '聊天记录';
