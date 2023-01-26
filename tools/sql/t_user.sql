create table t_user
(
    id       serial,
    username text not null primary key,
    encrypted_password text not null
);
