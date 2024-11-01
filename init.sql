/*
 * Copyright (c) 2024 flowerinsnow
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
DROP SCHEMA IF EXISTS `www`;
CREATE SCHEMA `www`;
DROP USER IF EXISTS 'www';
CREATE USER 'www'@'%';
ALTER USER 'www' IDENTIFIED BY 'my-password';
GRANT ALL PRIVILEGES ON `www`.* TO 'www'@'%';
USE `www`;
CREATE TABLE `access_log` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `remote_address` VARCHAR(255),
    `method` VARCHAR(7) NOT NULL,
    `host` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `referer` VARCHAR(255),
    `user_agent` VARCHAR(255),
    `params` TEXT,
    `blocked` BOOL NOT NULL DEFAULT FALSE,
    `time` DATETIME NOT NULL
);
CREATE TABLE `blacklist_address` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `address` VARCHAR(255) UNIQUE,
    `time` DATETIME NOT NULL,
    INDEX(`address`)
);
