create table users (
	id serial primary key,
	name varchar(225) not null,
	email varchar(255) unique not null
)