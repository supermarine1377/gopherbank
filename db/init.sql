create database if not exists gopherbank character set utf8mb4 collate utf8mb4_bin;

use gopherbank;

drop table if exists users;
create table if not exists users
(
  id int unsigned not null primary key,
  name varchar(128) not null
) character set utf8mb4 collate utf8mb4_bin;