-- +goose Up

ALTER TABLE fivehundred
ADD CONSTRAINT fivehundred_class_title_unique
UNIQUE (class,title);

-- +goose Down
ALTER TABLE fivehundred
DROP CONSTRAINT fivehundred_class_title_unique;