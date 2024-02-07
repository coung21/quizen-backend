create table flashcards (
    id char(36) default (uuid_to_bin(uuid())),
    study_set_id char(36) not null,
    term varchar(255) not null,
    definition varchar(255) not null,
    image json,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create table study_sets (
    id char(36) default (uuid_to_bin(uuid())),
    user_id char(36) not null,
    set_name varchar(255) not null,
    description text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);