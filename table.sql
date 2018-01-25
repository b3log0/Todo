CREATE TABLE `board` (
    `bid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id` VARCHAR(64) NOT NULL,
    `name` VARCHAR(50) NULL,
    `desc` VARCHAR(100) NULL,
    `created` DATE NULL
);

CREATE TABLE `group` (
    `gid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id` VARCHAR(64) NOT NULL,
    `name` VARCHAR(200) NULL,
    `created` DATE NULL
);

CREATE TABLE `card` (
    `cid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `id` VARCHAR(64) NOT NULL,
    `name` VARCHAR(200) NULL,
    `desc` TEXT NULL,
    `idList` VARCHAR(64),
    `created` DATE NULL
);
