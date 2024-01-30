create table sessions (
    id char(36) default (uuid_to_bin(uuid())),
    user_id char(36) not null,
    refresh_token text not null,
    user_agent text not null,
    user_ip text not null,
    expires_at timestamp not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);