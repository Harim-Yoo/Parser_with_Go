-- +goose Up
CREATE TABLE fivehundred_test (
  id bigint generated always as identity primary key,
  slug varchar not null,
  title text not null,
  contents text not null,
  created_at timestamptz default now(),
  updated_at timestamptz,
  is_deleted boolean default false
);

-- +goose Down
DROP TABLE fivehundred_test;
