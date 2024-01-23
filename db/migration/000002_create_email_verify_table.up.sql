create table veriy_emails (
    id int primary key auto_increment,
    email varchar(255) not null,
    secret_code varchar(255) not null,
    is_used boolean default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    expires_at timestamp default (current_timestamp + interval 10 minute)
)