CREATE TABLE `board` (
    `bid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `status` INTEGER NOT NULL,
    `id` VARCHAR(64) NULL,
    `name` VARCHAR(50) NOT NULL,
    `desc` VARCHAR(100) NULL,
    `created` DATE NULL
);

CREATE TABLE `group` (
    `gid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `status` INTEGER NOT NULL,
    `id` VARCHAR(64) NULL,
    `name` VARCHAR(200) NOT NULL,
    `bid` INTEGER NOT NULL,
    `created` DATE NULL
);

CREATE TABLE `card` (
    `cid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id` VARCHAR(64) NULL,
    `status` INTEGER NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `desc` TEXT NULL,
    `gid` INTEGER NOT NULL,
    `created` DATE NULL
);
