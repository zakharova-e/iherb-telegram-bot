begin;

create SCHEMA if not exists users;

create table if not exists users.users(
    id integer generated always as identity,
    user_name varchar(1000) not null,
    telegram_user_id bigint not null,
    telegram_user_name VARCHAR(1000),
    language_code VARCHAR(5)
);

commit;