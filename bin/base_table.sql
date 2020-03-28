SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

--
-- Database: `dc_es_log`
--

-- --------------------------------------------------------

--
-- 表的结构 `t_es_base`
--

DROP TABLE IF EXISTS `t_es_base`;
CREATE TABLE IF NOT EXISTS `t_es_base` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `docid` varchar(58) NOT NULL,
  `ename` varchar(58) NOT NULL,
  `app_id` int(11) UNSIGNED NOT NULL,
  `esdoc` text NOT NULL,
  `logtime` int(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `docid` (`docid`),
  KEY `logtime` (`logtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
COMMIT;

