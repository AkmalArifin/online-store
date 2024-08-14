CREATE TABLE `users` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `avatar` varchar(255),
  `first_name` varchar(255),
  `last_name` varchar(255),
  `username` varchar(255) UNIQUE NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255),
  `birth_of_date` datetime,
  `phone_number` varchar(255),
  `created_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `addresses` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `title` varchar(255),
  `address_line_1` varchar(255),
  `address_line_2` varchar(255),
  `country` varchar(255),
  `city` varchar(255),
  `postal_code` varchar(255),
  `landmark` varchar(255),
  `phone_number` varchar(255),
  `created_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `products_skus` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `product_id` integer,
  `sku` varchar(255),
  `price` varchar(255),
  `quantity` integer,
  `cover` varchar(255),
  `created_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `description` varchar(255),
  `cover` varchar(255),
  `category_id` integer,
  `created_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `categories` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `description` varchar(255),
  `created_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `wishlist` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `product_id` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `cart` (
  `id` integer AUTO_INCREMENT,
  `user_id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`, `user_id`)
);

CREATE TABLE `cart_item` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `cart_id` integer,
  `product_sku_id` integer,
  `quantity` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `order_details` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `payment_id` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `order_item` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `order_id` integer,
  `products_sku_id` integer,
  `quantity` integer,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `payment_details` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `order_id` integer,
  `amount` integer,
  `provider` varchar(255),
  `status` varchar(255),
  `created_at` datetime,
  `updated_at` datetime
);

ALTER TABLE `addresses` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

CREATE TABLE `products_categories` (
  `products_category_id` integer,
  `categories_id` integer,
  PRIMARY KEY (`products_category_id`, `categories_id`)
);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `products_categories` (`products_category_id`);

ALTER TABLE `products_categories` ADD FOREIGN KEY (`categories_id`) REFERENCES `categories` (`id`);


ALTER TABLE `products_skus` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

CREATE TABLE `wishlist_users` (
  `wishlist_user_id` integer,
  `users_id` integer,
  PRIMARY KEY (`wishlist_user_id`, `users_id`)
);

ALTER TABLE `wishlist` ADD FOREIGN KEY (`user_id`) REFERENCES `wishlist_users` (`wishlist_user_id`);

ALTER TABLE `wishlist_users` ADD FOREIGN KEY (`users_id`) REFERENCES `users` (`id`);


CREATE TABLE `wishlist_products` (
  `wishlist_product_id` integer,
  `products_id` integer,
  PRIMARY KEY (`wishlist_product_id`, `products_id`)
);

ALTER TABLE `wishlist` ADD FOREIGN KEY (`product_id`) REFERENCES `wishlist_products` (`wishlist_product_id`);

ALTER TABLE `wishlist_products` ADD FOREIGN KEY (`products_id`) REFERENCES `products` (`id`);

ALTER TABLE `cart` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `cart_item` ADD FOREIGN KEY (`cart_id`) REFERENCES `cart` (`id`);

CREATE TABLE `cart_item_products_skus` (
  `cart_item_product_sku_id` integer,
  `products_skus_id` integer,
  PRIMARY KEY (`cart_item_product_sku_id`, `products_skus_id`)
);

ALTER TABLE `cart_item` ADD FOREIGN KEY (`product_sku_id`) REFERENCES `cart_item_products_skus` (`cart_item_product_sku_id`);

ALTER TABLE `cart_item_products_skus` ADD FOREIGN KEY (`products_skus_id`) REFERENCES `products_skus` (`id`);


ALTER TABLE `order_item` ADD FOREIGN KEY (`order_id`) REFERENCES `order_details` (`id`);

ALTER TABLE `order_details` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `order_details` ADD FOREIGN KEY (`payment_id`) REFERENCES `payment_details` (`id`);

CREATE TABLE `order_item_products_skus` (
  `order_item_products_sku_id` integer,
  `products_skus_id` integer,
  PRIMARY KEY (`order_item_products_sku_id`, `products_skus_id`)
);

ALTER TABLE `order_item` ADD FOREIGN KEY (`products_sku_id`) REFERENCES `order_item_products_skus` (`order_item_products_sku_id`);

ALTER TABLE `order_item_products_skus` ADD FOREIGN KEY (`products_skus_id`) REFERENCES `products_skus` (`id`);

