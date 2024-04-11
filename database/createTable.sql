CREATE TABLE `users` (
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `username` VARCHAR(64) NULL,
    `password` VARCHAR(64) NULL,
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

ALTER TABLE `users` DROP COLUMN `created_at`
ALTER TABLE `users` DROP COLUMN `updated_at`
ALTER TABLE `todo_lists` ADD COLUMN `user_uid`

UPDATE `todo_lists` SET user_uid = 1 WHERE user_uid = NULL
