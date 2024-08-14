CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `avatar` varchar(255),
  `first_name` varchar(255),
  `last_name` varchar(255),
  `username` varchar(255) UNIQUE NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255),
  `birth_of_date` date,
  `phone_number` varchar(255),
  `created_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `addresses` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `title` varchar(255),
  `address_line_1` varchar(255),
  `address_line_2` varchar(255),
  `country` varchar(255),
  `city` varchar(255),
  `postal_code` varchar(255),
  `landmark` varchar(255),
  `phone_number` varchar(255),
  `created_at` timestamp,
  `deleted_at` timestamp
);

ALTER TABLE `addresses` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


CREATE TABLE `products_skus` (
  `id` integer PRIMARY KEY,
  `product_id` integer,
  `sku` varchar(255),
  `price` varchar(255),
  `quantity` integer,
  `cover` varchar(255),
  `created_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `description` varchar(255),
  `cover` varchar(255),
  `category_id` integer,
  `created_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `categories` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `description` varchar(255),
  `created_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `products_categories` (
  `products_category_id` integer,
  `categories_id` integer,
  PRIMARY KEY (`products_category_id`, `categories_id`)
);

ALTER TABLE `products_skus` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

-- ALTER TABLE `products_categories` ADD FOREIGN KEY (`products_category_id`) REFERENCES `products` (`category_id`);
ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `products_categories` (`products_category_id`);

ALTER TABLE `products_categories` ADD FOREIGN KEY (`categories_id`) REFERENCES `categories` (`id`);