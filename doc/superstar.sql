/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50723
Source Host           : localhost:3306
Source Database       : superstar

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2018-12-07 16:55:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_star
-- ----------------------------
DROP TABLE IF EXISTS `t_star`;
CREATE TABLE `t_star` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `name_zh` varchar(50) DEFAULT NULL COMMENT '球星中文名',
  `name_en` varchar(50) DEFAULT NULL COMMENT '球星英文名',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_star
-- ----------------------------
INSERT INTO `t_star` VALUES ('1', null, '中国', 'china', '2018-12-07 09:46:20', '2018-12-07 09:46:20');
INSERT INTO `t_star` VALUES ('2', null, '韩国', 'korea', '2018-12-07 09:46:59', '2018-12-07 09:46:59');
INSERT INTO `t_star` VALUES ('3', null, '美国', 'American', '2018-12-07 09:47:00', '2018-12-07 09:47:00');
INSERT INTO `t_star` VALUES ('6', 'www.google.com', '谷歌', 'Google', '2018-12-07 14:19:18', '2018-12-07 14:19:18');
