drop table if exists posts cascade;
drop table if exists comments;

create table posts (
	id int primary key auto_increment,
	content varchar(255),
	author varchar(255)
);

create table comments (
	id int primary key auto_increment,
	content varchar(255),
	author varchar(255),
	post_id int references posts(id)
);