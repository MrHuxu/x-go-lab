CREATE TABLE `user`(
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` TEXT NOT NULL,
  `age` INT,
  `created_at` timestamp NOT NULL DEFAULT '1971-01-01 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);