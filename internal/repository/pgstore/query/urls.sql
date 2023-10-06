-- name: CreateUrl :one
INSERT INTO urls(
    org_url,
    short_url,
    user_id
)VALUES(
    $1,$2,$3
) RETURNING *;

-- name: GetUrlById :one
SELECT * FROM urls
WHERE id=$1 LIMIT 1;

-- name: GetUrlByShort :one
SELECT * FROM urls
WHERE short_url=$1 LIMIT 1;

-- name: ListUserUrls :many
SELECT * FROM urls
WHERE user_id=$1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: IncrementClick :exec
UPDATE urls
SET clicks = clicks+1
WHERE id=$1;

-- name: UpdateOrgUrl :one
UPDATE urls
SET
    org_url=$2,
    updated_at=$3
WHERE id=$1
RETURNING *;

-- name: DeleteUrl :exec
DELETE FROM urls WHERE id=$1;