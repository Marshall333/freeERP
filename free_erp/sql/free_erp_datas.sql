-- MySQL dump 10.13  Distrib 5.6.35, for macos10.12 (x86_64)
--
-- Host: localhost    Database: free_erp
-- ------------------------------------------------------
-- Server version	5.6.35

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
-- Table structure for table `commodity_info`
--

DROP TABLE IF EXISTS `commodity_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `commodity_info` (
  `GOODS_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `COMMODITY_ID` varchar(63) DEFAULT NULL COMMENT '商品类别唯一标识(条形码)',
  `NAME` varchar(255) NOT NULL COMMENT '商品名称',
  `DESCRIBE` varchar(255) NOT NULL COMMENT '商品描述',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `IN_PRICE` float DEFAULT '0' COMMENT '进货单价,单位: 元',
  `OUT_PRICE` float DEFAULT '0' COMMENT '售卖单价,单位: 元',
  `STATUS` tinyint(4) DEFAULT '1' COMMENT '商品状态 1 激活 2 删除',
  `REMARK` varchar(255) NOT NULL COMMENT '商品备注',
  PRIMARY KEY (`GOODS_ID`),
  KEY `I_commodity_id` (`COMMODITY_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commodity_info`
--

LOCK TABLES `commodity_info` WRITE;
/*!40000 ALTER TABLE `commodity_info` DISABLE KEYS */;
INSERT INTO `commodity_info` VALUES (100001,'6956367338895','王老吉(罐装)','','2019-01-11 06:57:15',3,4,1,''),(100002,'6935284455588','卫龙亲嘴烧','','2019-01-11 06:59:18',0.3,1,1,''),(100003,'6901668053787','奥利奥87g草莓双心脆','','2019-01-11 07:01:23',3,5.5,1,'');
/*!40000 ALTER TABLE `commodity_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_details`
--

DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_details` (
  `ORDER_DETAILS_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '订单详情ID',
  `ORDER_ID` bigint(20) NOT NULL COMMENT '订单ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `GOODS_ID` bigint(20) NOT NULL COMMENT '商品ID',
  `GOODS_NUM` int(11) NOT NULL DEFAULT '1' COMMENT '商品数量',
  `TOTAL` float DEFAULT '0' COMMENT '应收总计,单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  PRIMARY KEY (`ORDER_DETAILS_ID`),
  KEY `I_ORDER_ID` (`ORDER_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_details`
--

LOCK TABLES `order_details` WRITE;
/*!40000 ALTER TABLE `order_details` DISABLE KEYS */;
INSERT INTO `order_details` VALUES (10000000000001,10000000000000,'2019-01-11 07:07:08',100003,5,27.5,''),(10000000000002,10000000000000,'2019-01-11 07:07:22',100001,6,24,''),(10000000000003,10000000000000,'2019-01-11 07:07:53',100002,3,3,''),(10000000000004,10000000000001,'2019-01-11 07:09:18',100001,30,120,''),(10000000000005,10000000000001,'2019-01-11 07:09:30',100002,38,38,'');
/*!40000 ALTER TABLE `order_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_list`
--

DROP TABLE IF EXISTS `order_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_list` (
  `ORDER_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `CREATE_USER_ID` bigint(20) NOT NULL COMMENT '订单创建人用户ID',
  `TOTAL` float DEFAULT '0' COMMENT '应收总计,单位: 元',
  `TOTAL_FACT` float DEFAULT '0' COMMENT '实收,单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  `STATUS` tinyint(4) DEFAULT '1' COMMENT '订单状态',
  PRIMARY KEY (`ORDER_ID`),
  KEY `I_CREATETIME` (`CREATETIME`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000000000002 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_list`
--

LOCK TABLES `order_list` WRITE;
/*!40000 ALTER TABLE `order_list` DISABLE KEYS */;
INSERT INTO `order_list` VALUES (10000000000000,'2019-01-08 09:54:18',200001,54.5,54,'',1),(10000000000001,'2019-01-08 09:54:41',200001,158,158,'',1);
/*!40000 ALTER TABLE `order_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `repository`
--

DROP TABLE IF EXISTS `repository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository` (
  `STORE_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '库存ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `ITEM_COST` float DEFAULT '0' COMMENT '进货总价,单位: 元',
  `ITEM_COST_FACT` float DEFAULT '0' COMMENT '进货总价(实付),单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  PRIMARY KEY (`STORE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `repository`
--

LOCK TABLES `repository` WRITE;
/*!40000 ALTER TABLE `repository` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `repository_details`
--

DROP TABLE IF EXISTS `repository_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_details` (
  `STORE_DETAILS_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '库存详情ID',
  `STORE_ID` bigint(20) NOT NULL COMMENT '库存ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `GOODS_ID` bigint(20) NOT NULL COMMENT '商品ID',
  `GOODS_NUM` int(11) NOT NULL DEFAULT '1' COMMENT '商品数量',
  `TOTAL_COST` float DEFAULT '0' COMMENT '进货总价,单位: 元',
  `TOTAL_COST_FACT` float DEFAULT '0' COMMENT '进货总价(实付),单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  PRIMARY KEY (`STORE_DETAILS_ID`),
  KEY `I_STORE_ID` (`STORE_ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `repository_details`
--

LOCK TABLES `repository_details` WRITE;
/*!40000 ALTER TABLE `repository_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `repository_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `USER_ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `USERNAME` varchar(255) NOT NULL COMMENT '用户名',
  `PASSWORD` varchar(255) NOT NULL COMMENT '密码md5,32位小写',
  `MOBILE` varchar(255) DEFAULT NULL COMMENT '手机号码',
  `ROLE` varchar(63) DEFAULT NULL COMMENT '权限 SUPER_ADMIN 超级管理员 ADMIN 管理员 SALE 销售 REPERTORY 仓库管理',
  `AUTH` varchar(255) DEFAULT NULL COMMENT '权限详情,各权限以英文逗号隔开',
  `SESSIONID` varchar(255) DEFAULT NULL,
  `SESSION_TTL` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`USER_ID`),
  KEY `I_username` (`USERNAME`) USING BTREE,
  KEY `I_sessionid` (`SESSIONID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','202cb962ac59075b964b07152d234b70',NULL,NULL,NULL,'',1547196452);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-01-11 19:20:33
