CREATE DATABASE IF NOT EXISTS golang_test;
USE golang_test;
CREATE TABLE `game` (
  `id` int NOT NULL,
  `mdate` int DEFAULT NULL,
  `stadium` varchar(255) DEFAULT NULL,
  `team1` varchar(255) DEFAULT NULL,
  `team2` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
