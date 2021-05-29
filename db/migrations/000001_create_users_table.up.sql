create sequence if not exists users_id_seq;


create table if not exists users
(
    id         int8 default nextval('users_id_seq') primary key,
    created_at timestamptz(6),
    updated_at timestamptz(6),
    deleted_at timestamptz(6),
    name       text,
    email      text
);

