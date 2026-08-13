-- name: InsertData :many
INSERT INTO fivehundred (title, contents)
VALUES ($1, $2)
RETURNING title, contents;
