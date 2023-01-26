create table t_dict
(
    key   text not null primary key,
    value text not null
);

insert into t_dict (key, value) values ('k', 'v');

select * from t_dict;