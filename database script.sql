-- Adminer 4.2.1 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `event_sequence`;
CREATE TABLE `event_sequence` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(45) DEFAULT NULL,
  `job_id` bigint(20) DEFAULT NULL,
  `event_type` varchar(45) DEFAULT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `expired_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `job_id` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `job_accepted`;
CREATE TABLE `job_accepted` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `job_id` bigint(20) DEFAULT NULL,
  `driver_id` int(11) DEFAULT NULL,
  `accepted_location` varchar(75) CHARACTER SET latin1 DEFAULT NULL,
  `accepted_lat` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `accepted_lon` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `expired_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `job_id` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;


DROP TABLE IF EXISTS `job_completed`;
CREATE TABLE `job_completed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(45) DEFAULT NULL,
  `job_id` bigint(20) DEFAULT NULL,
  `address` varchar(75) DEFAULT NULL,
  `complete_loc` varchar(45) DEFAULT NULL,
  `complete_lat` varchar(45) DEFAULT NULL,
  `complete_lon` varchar(45) DEFAULT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `expired_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `job_id` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `job_created`;
CREATE TABLE `job_created` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(45) DEFAULT NULL,
  `job_id` bigint(20) DEFAULT NULL,
  `module` varchar(15) DEFAULT NULL,
  `passenger_id` int(11) DEFAULT NULL,
  `pickup_name` varchar(75) DEFAULT NULL,
  `pickup_phone` int(11) DEFAULT NULL,
  `pickup_location` varchar(75) DEFAULT NULL,
  `pickup_lat` varchar(45) DEFAULT NULL,
  `pickup_lon` varchar(45) DEFAULT NULL,
  `drop_location` varchar(45) DEFAULT NULL,
  `drop_lat` varchar(45) DEFAULT NULL,
  `drop_lon` varchar(45) DEFAULT NULL,
  `pickup_time` bigint(20) DEFAULT NULL,
  `prebooking` tinyint(1) DEFAULT NULL COMMENT '0-false , 1-true',
  `promocode` varchar(45) DEFAULT NULL,
  `payment_method` varchar(5) DEFAULT NULL,
  `created_date` date DEFAULT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `expired_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `trip_created_trip_id_idx` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `job_details`;
CREATE TABLE `job_details` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_id` bigint(20) NOT NULL,
  `order_details` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `job_payment_type`;
CREATE TABLE `job_payment_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_id` bigint(20) NOT NULL,
  `payment_method` tinyint(2) NOT NULL COMMENT '1- Cash, 2 - Card, 3 - Points',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1 - Primary 2- Secondary',
  PRIMARY KEY (`id`),
  KEY `job_id` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `job_region`;
CREATE TABLE `job_region` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_id` bigint(20) DEFAULT NULL,
  `region_id` int(11) DEFAULT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `expired_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `trip_created_trip_id_trip_region_trip_id_idx` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `job_rejected`;
CREATE TABLE `job_rejected` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_id` bigint(20) NOT NULL,
  `driver_id` int(11) NOT NULL,
  `rejection_type` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `rejected_lat` int(11) NOT NULL,
  `rejected_lon` int(11) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `expired_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `job_id` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


-- 2018-10-19 05:02:28
