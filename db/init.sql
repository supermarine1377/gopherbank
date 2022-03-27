create database if not exists gopherbank character set utf8mb4 collate utf8mb4_bin;

use gopherbank;

drop table if exists users;
create table if not exists users
(
  id int unsigned not null primary key auto_increment,
  name varchar(128) not null,
  balance int unsigned not null,
  is_deleted boolean not null
) character set utf8mb4 collate utf8mb4_bin;

drop table if exists transactions;
create table if not exists transactions
(
  id int unsigned not null primary key auto_increment,
  amount int unsigned not null,
  user_id int unsigned not null,
  foreign key (user_id) references users(id)
) character set utf8mb4 collate utf8mb4_bin;