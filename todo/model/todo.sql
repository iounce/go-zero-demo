CREATE TABLE `todo` (   
  `id` bigint(20) NOT NULL AUTO_INCREMENT,  
  `title` varchar(128) NOT NULL DEFAULT '',   
  `content` varchar(255) NOT NULL DEFAULT '',   
  PRIMARY KEY (`id`)    
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;