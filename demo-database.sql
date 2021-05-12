create database demo;

use demo;

create table users (
	id int auto_increment primary key,
    username varchar(50),
    password varchar(50)
    );
    
INSERT INTO `demo`.`users` (`username`, `password`) VALUES ('root', 'root');