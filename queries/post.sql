-- name: CreatePost :one
INSERT INTO post (
    tag,
    owner_id,
    description,
    from_org,
    is_video, 
    content_url
) VALUES (
    $1, $2, $3, $4, $5, $6
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
    tag = coalesce(sqlc.narg('tag'), tag),
    description = coalesce(sqlc.narg('description'), description),
    content_url = coalesce(sqlc.narg('content_url'), content_url),
    is_video = coalesce(sqlc.narg('is_video'), is_video),
    from_org = coalesce(sqlc.narg('from_org'), from_org),
    updated_at = now()
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = $1;

