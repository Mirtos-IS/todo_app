CREATE TABLE `users` (
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `username` VARCHAR(64) NULL,
    `password` VARCHAR(64) NULL,
    `business_name` VARCHAR(64) DEFAULT '',
    `created_at` DATE NULL,
    `updated_at` DATE NULL
);

CREATE TABLE `todo_items` (
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR(64) NULL
);

-- CREATE TABLE `todo_lists` (
--     `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
--     `title` VARCHAR(64) NULL,
--     `author` VARCHAR(64) NULL,
--     `content` VARCHAR(64) NULL,
--     `like_count` INTEGER DEFAULT 0,
--     `created_at` DATE NULL,
--     `updated_at` DATE NULL
-- );
