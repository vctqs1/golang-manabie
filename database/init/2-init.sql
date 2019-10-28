USE `golang_manabie`;

CREATE TABLE IF NOT EXISTS `products` (
  `id` int(6) NOT NULL,
  `title` varchar(255) NOT NULL,
  `quantities` int(11) NOT NULL DEFAULT 0,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



INSERT INTO `products` (`id`, `title`, `quantities`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Ao thun', 10, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40'),
(2, 'Ao somi', 30, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40'),
(3, 'Ao croptop', 5, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40'),
(4, 'Quan jean', 9, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40'),
(5, 'Quan den', 23, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40'),
(6, 'Quan short', 33, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40'),
(7, 'Quan dai', 2, '', '2019-10-28 13:04:40', '2019-10-28 13:04:40');


ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `ID_UNIQUE` (`id`);


ALTER TABLE `products`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;
