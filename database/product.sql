CREATE TABLE `products` (
  `id` int(6) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `quantities` int(11) NOT NULL DEFAULT 0,
  `description` text DEFAULT NULL,
  `created_at` timestamp DEFAULT now() NOT NULL,
  `updated_at` timestamp DEFAULT now() NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ID_UNIQUE` (`id`)
);



INSERT INTO `products` (id, title, quantities) VALUES
(1, 'Áo thun', 10),
(2, 'Áo sơ mi', 30),
(3, 'Áo croptop', 5),
(4, 'Quần jean', 9),
(5, 'Quần đen', 23),
(6, 'Quần short', 33),
(7, 'Quần dài', 2)
