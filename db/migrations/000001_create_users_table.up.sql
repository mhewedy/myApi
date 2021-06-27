create sequence if not exists users_id_seq;

create table if not exists users
(
    id         int8 default nextval('users_id_seq') primary key,
    username   text,
    password   text,
    active     bool,
    created_at timestamptz(6),
    updated_at timestamptz(6),
    deleted_at timestamptz(6)
);

