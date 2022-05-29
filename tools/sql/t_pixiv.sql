create table t_pixiv(
    id serial primary key,
    picture_id int not null,
    author_id int not null,
    tags text not null,
    is_r18 bool not null,
    origin_url text not null,
    relative_path text not null,
    create_time text not null
);
