-- MySQL dump 10.13  Distrib 5.1.73, for redhat-linux-gnu (x86_64)
--
-- Host: localhost    Database: pingan
-- ------------------------------------------------------
-- Server version	5.1.73

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
  PRIMARY KEY (`USER_ID`),
  KEY `I_username` (`USERNAME`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=200000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `commodity_info`
--

DROP TABLE IF EXISTS `commodity_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `commodity_info` (
  `GOODS_ID` bigint(20) NOT NULL COMMENT '商品ID' AUTO_INCREMENT,
  `COMMODITY_ID` varchar(63) DEFAULT NULL COMMENT '商品类别唯一标识(条形码)',
  `NAME` varchar(255) NOT NULL COMMENT '商品名称',
  `DESCRIBE` varchar(255) NOT NULL COMMENT '商品描述',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `IN_PRICE` float DEFAULT 0 COMMENT '进货单价,单位: 元',
  `OUT_PRICE` float DEFAULT 0 COMMENT '售卖单价,单位: 元',
  `STATUS` tinyint DEFAULT 1 COMMENT '商品状态 1 激活 2 删除',
  `REMARK` varchar(255) NOT NULL COMMENT '商品备注',
  PRIMARY KEY (`GOODS_ID`),
  KEY `I_commodity_id` (`COMMODITY_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_list` (
  `ORDER_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `CREATE_USER_ID` bigint NOT NULL COMMENT '订单创建人用户ID',
  `TOTAL` float DEFAULT 0 COMMENT '应收总计,单位: 元',
  `TOTAL_FACT` float DEFAULT 0 COMMENT '实收,单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  `STATUS` tinyint DEFAULT 1 COMMENT '订单状态',
  PRIMARY KEY (`ORDER_ID`),
  KEY `I_CREATETIME` (`CREATETIME`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000000000000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `order_details`
--

DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_details` (
  `ORDER_DETAILS_ID` bigint NOT NULL AUTO_INCREMENT COMMENT '订单详情ID',
  `ORDER_ID` bigint(20) NOT NULL COMMENT '订单ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `GOODS_ID` bigint NOT NULL COMMENT '商品ID',
  `GOODS_NUM` int NOT NULL DEFAULT 1 COMMENT '商品数量', 
  `TOTAL` float DEFAULT 0 COMMENT '应收总计,单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  PRIMARY KEY (`ORDER_DETAILS_ID`),
  KEY `I_ORDER_ID` (`ORDER_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20000000000000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `repertory`
--

DROP TABLE IF EXISTS `repository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository` (
  `STORE_ID` bigint NOT NULL AUTO_INCREMENT COMMENT '库存ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `ITEM_COST` float DEFAULT 0 COMMENT '进货总价,单位: 元',
  `ITEM_COST_FACT` float DEFAULT 0 COMMENT '进货总价(实付),单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  PRIMARY KEY (`STORE_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000000000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `repository_details`
--

DROP TABLE IF EXISTS `repository_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `repository_details` (
  `STORE_DETAILS_ID` bigint NOT NULL AUTO_INCREMENT COMMENT '库存详情ID',
  `STORE_ID` bigint NOT NULL COMMENT '库存ID',
  `CREATETIME` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `GOODS_ID` bigint NOT NULL COMMENT '商品ID',
  `GOODS_NUM` int NOT NULL DEFAULT 1 COMMENT '商品数量', 
  `TOTAL_COST` float DEFAULT 0 COMMENT '进货总价,单位: 元',
  `TOTAL_COST_FACT` float DEFAULT 0 COMMENT '进货总价(实付),单位: 元',
  `REMARK` varchar(255) NOT NULL COMMENT '备注',
  PRIMARY KEY (`STORE_DETAILS_ID`),
  KEY `I_STORE_ID` (`STORE_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2000000000000 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

-- Dump completed on 2017-04-25 17:43:36
