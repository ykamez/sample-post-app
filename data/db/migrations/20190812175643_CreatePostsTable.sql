
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table posts (
  id              serial primary key,
  title           varchar(255)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table posts;

