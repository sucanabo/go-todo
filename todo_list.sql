-- -------------------------------------------------------------
-- TablePlus 5.6.8(524)
--
-- https://tableplus.com/
--
-- Database: todo_list
-- Generation Time: 2023-12-27 17:21:10.2380
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `image` json DEFAULT NULL,
  `description` text,
  `status` enum('doing','done','deleted') DEFAULT 'doing',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `todo_user_like_items` (
  `item_id` int NOT NULL,
  `user_id` int NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `todo_items` (`id`, `title`, `image`, `description`, `status`, `created_at`, `updated_at`) VALUES
(1, 'task 1', NULL, 'this is task 1', 'doing', '2023-12-27 03:21:17', '2023-12-27 03:21:17'),
(2, 'task 2', NULL, 'this is task 2222', 'done', '2023-12-27 03:22:12', '2023-12-27 03:22:19'),
(3, 'task 3', NULL, 'this is task 333', 'done', '2023-12-27 03:22:12', '2023-12-27 03:22:48'),
(4, 'task 4', NULL, 'this is task 444', 'deleted', '2023-12-27 03:22:12', '2023-12-27 03:22:42'),
(5, 'first task', NULL, 'no des', 'doing', '2023-12-27 10:07:26', '2023-12-27 10:07:26'),
(6, 'third task', NULL, 'no ddd', 'doing', '2023-12-27 10:15:40', '2023-12-27 10:15:40'),
(7, 'third task', NULL, 'no ddd', NULL, '2023-12-27 10:16:42', '2023-12-27 10:16:42'),
(8, 'four task', NULL, 'no ddd', NULL, '2023-12-27 10:17:01', '2023-12-27 10:17:01');

INSERT INTO `todo_user_like_items` (`item_id`, `user_id`, `created_at`) VALUES
(1, 1, '2023-12-27 08:08:16'),
(2, 1, '2023-12-27 08:08:42'),
(1, 2, '2023-12-27 08:08:16');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;