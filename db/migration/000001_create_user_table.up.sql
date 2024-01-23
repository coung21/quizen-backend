create table users (
    id int primary key auto_increment,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(500) not null,
    avatar json default null,
    is_verified boolean not null default false,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);
