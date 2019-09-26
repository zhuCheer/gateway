-- 导出  表 gateway.qi_nodes 结构
CREATE TABLE IF NOT EXISTS `qi_nodes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `site_id` int(11) NOT NULL,
  `addr` varchar(100) NOT NULL,
  `weight` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- 正在导出表  gateway.qi_nodes 的数据：~2 rows (大约)
DELETE FROM `qi_nodes`;
/*!40000 ALTER TABLE `qi_nodes` DISABLE KEYS */;
INSERT INTO `qi_nodes` (`id`, `site_id`, `addr`, `weight`, `created_at`, `updated_at`) VALUES
	(33, 23, '127.0.0.1:8081', 5, '2019-09-25 18:05:20', '2019-09-25 18:05:20'),
	(36, 23, '127.0.0.1:8082', 10, '2019-09-26 03:17:09', '2019-09-26 03:17:09');


-- 导出  表 gateway.qi_sites 结构
CREATE TABLE IF NOT EXISTS `qi_sites` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `domain` varchar(100) NOT NULL,
  `scheme` varchar(8) NOT NULL DEFAULT 'http',
  `balance` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COMMENT='站点信息';

-- 正在导出表  gateway.qi_sites 的数据：~1 rows (大约)
DELETE FROM `qi_sites`;
/*!40000 ALTER TABLE `qi_sites` DISABLE KEYS */;
INSERT INTO `qi_sites` (`id`, `name`, `domain`, `scheme`, `balance`, `created_at`, `updated_at`) VALUES
	(23, '网关测试', 'www.qiproxy.cn', 'http', 'roundrobin', '2019-09-25 07:03:48', '2019-09-26 03:17:31');

