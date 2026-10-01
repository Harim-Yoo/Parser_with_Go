-- name: InsertData :many
INSERT INTO fivehundred (title, contents)
VALUES ($1, $2)
RETURNING title, contents;

-- name: UpdateSlug :exec
UPDATE fivehundred 
SET slug=$1,
  updated_at=now()
WHERE title LIKE $2 || '%';

-- name: UpdateClass :exec 
UPDATE fivehundred
SET class=$1,
  updated_at=now()
WHERE slug = ANY($2::text[]);

-- name: InsertBookData :exec
INSERT INTO fivehundred (class, title, contents, slug, class_order)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (class, title)
DO UPDATE SET
  contents = EXCLUDED.contents,
  slug = EXCLUDED.slug,
  class_order = EXCLUDED.class_order,
  updated_at = now();