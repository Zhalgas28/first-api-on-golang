create table users
(
    id  integer primary key generated always as identity,
    username      varchar not null,
    password_hash varchar not null,
    email         varchar
);

create table news
(
    id integer primary key generated always as identity,
    title varchar(64) not null,
    description varchar not null,
    photo_url varchar,
    created_time timestamp not null,
    author_id integer,

    constraint news_author_id FOREIGN KEY (author_id) REFERENCES users (id)
)