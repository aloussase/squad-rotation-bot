create table squad_members (
    id serial primary key,
    full_name text not null,
    avatar_url text
);

create sequence squad_rotation
    start 1
    increment 1
    minvalue 1
    no maxvalue
    cache 1;