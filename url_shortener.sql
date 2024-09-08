CREATE DATABASE IF NOT EXISTS `url_shortener` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT ENCRYPTION = 'N';
use `url_shortener`;

CREATE TABLE `url` (
  `id` char(36) NOT NULL,
  `short_url` varchar(256) NOT NULL,
  `original_url` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `short_url_UNIQUE` (`short_url`),
  UNIQUE KEY `original_url_UNIQUE` (`original_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci