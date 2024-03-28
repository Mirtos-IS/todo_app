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
    `title` VARCHAR(64) NULL,
    `todo_list_uid` INTEGER NULL,
    `is_marked` NUMBER(1) DEFAULT 0,
    FOREIGN KEY (`todo_list_uid`) REFERENCES todo_lists(uid)
);

CREATE TABLE `todo_lists` (
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(64) NULL,
    `user_uid` INTEGER NULL,
    FOREIGN KEY (`user_uid`) REFERENCES users(uid)
);

INSERT INTO `todo_lists`(`name`) VALUES('Test')
