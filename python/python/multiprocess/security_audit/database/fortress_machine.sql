-- MySQL dump 10.13  Distrib 5.7.14, for osx10.11 (x86_64)
--
-- Host: localhost    Database: fortress_machine
-- ------------------------------------------------------
-- Server version	5.7.14

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `group_relation`
--

DROP TABLE IF EXISTS `group_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group_relation` (
  `gid` int(4) NOT NULL AUTO_INCREMENT,
  `user_group` varchar(64) NOT NULL COMMENT '用户组',
  `host_group` varchar(64) NOT NULL COMMENT '主机组',
  PRIMARY KEY (`user_group`,`host_group`),
  UNIQUE KEY `gid` (`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='群组关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_relation`
--

LOCK TABLES `group_relation` WRITE;
/*!40000 ALTER TABLE `group_relation` DISABLE KEYS */;
INSERT INTO `group_relation` VALUES (3,'admin','group1');
/*!40000 ALTER TABLE `group_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `host_list`
--

DROP TABLE IF EXISTS `host_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `host_list` (
  `host_ip` varchar(64) NOT NULL COMMENT '主机ip',
  `hostname` varchar(64) NOT NULL COMMENT '主机名',
  `host_group` varchar(64) NOT NULL COMMENT '主机组',
  PRIMARY KEY (`hostname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='主机列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `host_list`
--

LOCK TABLES `host_list` WRITE;
/*!40000 ALTER TABLE `host_list` DISABLE KEYS */;
INSERT INTO `host_list` VALUES ('192.168.186.1','JasonLuoMac','group1');
/*!40000 ALTER TABLE `host_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `host_users`
--

DROP TABLE IF EXISTS `host_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `host_users` (
  `hostname` varchar(64) NOT NULL COMMENT '主机名',
  `hostuser` varchar(64) NOT NULL COMMENT '主机用户',
  `hostpwd` varchar(64) NOT NULL COMMENT '主机用户密码',
  PRIMARY KEY (`hostuser`),
  KEY `hostname` (`hostname`),
  CONSTRAINT `host_users_ibfk_1` FOREIGN KEY (`hostname`) REFERENCES `host_list` (`hostname`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='主机用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `host_users`
--

LOCK TABLES `host_users` WRITE;
/*!40000 ALTER TABLE `host_users` DISABLE KEYS */;
INSERT INTO `host_users` VALUES ('JasonLuoMac','jason','qianyishi,.');
/*!40000 ALTER TABLE `host_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `records`
--

DROP TABLE IF EXISTS `records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `records` (
  `hostname` varchar(64) NOT NULL COMMENT '主机名',
  `hostuser` varchar(64) NOT NULL COMMENT '主机用户',
  `records` varchar(128) NOT NULL COMMENT '历史记录',
  KEY `hostname` (`hostname`),
  KEY `hostuser` (`hostuser`),
  CONSTRAINT `records_ibfk_1` FOREIGN KEY (`hostname`) REFERENCES `host_list` (`hostname`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `records_ibfk_2` FOREIGN KEY (`hostuser`) REFERENCES `host_users` (`hostuser`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户操作记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `records`
--

LOCK TABLES `records` WRITE;
/*!40000 ALTER TABLE `records` DISABLE KEYS */;
/*!40000 ALTER TABLE `records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user2group`
--

DROP TABLE IF EXISTS `user2group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user2group` (
  `user_id` int(4) NOT NULL AUTO_INCREMENT COMMENT '用户自增id',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `after_login` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '上次登录时间',
  `user_group` varchar(64) NOT NULL COMMENT '用户组',
  PRIMARY KEY (`username`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user2group`
--

LOCK TABLES `user2group` WRITE;
/*!40000 ALTER TABLE `user2group` DISABLE KEYS */;
INSERT INTO `user2group` VALUES (3,'rosen','2016-10-31 12:29:39','0000-00-00 00:00:00','admin');
/*!40000 ALTER TABLE `user2group` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-10-31 20:46:23
