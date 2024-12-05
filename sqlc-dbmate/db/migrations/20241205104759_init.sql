-- migrate:up
create table dbmate (
    id integer,
    name varchar(255),
    email varchar(255) not null
);
-- migrate:down
DROP TABLE dbmate;