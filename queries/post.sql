-- name: CreatePost :one
INSERT INTO post (
    tag,
    owner_id,
    description,
    image_url,
    from_org
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPost :one
SELECT * FROM post
WHERE id = $1 LIMIT 1;

-- name: GetPostByOwnerId :one
SELECT * FROM post
WHERE owner_id = $1 LIMIT 1;


-- name: UpdatePost :one
UPDATE post
SET 
    tag = $2,
    description = $3,
    image_url = $4,
    from_org = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = $1;

