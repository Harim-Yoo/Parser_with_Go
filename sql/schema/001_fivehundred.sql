-- +goose Up
CREATE TABLE public.fivehundred (
  id uuid primary key default gen_random_uuid(),
  class varchar,
  title text not null,
  contents text not null,
  created_at timestamptz default now(),
  updated_at timestamptz,
  is_deleted boolean default false
);

-- +goose Down
DROP TABLE public.fivehundred;
