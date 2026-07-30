-- name: InsertData :many
INSERT INTO fivehundred_test (slug, title, contents)
VALUES ($1, $2, $3)
RETURNING title, contents;
