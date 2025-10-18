-- name: GetSubscriptionFeed :many
SELECT DISTINCT p.*
FROM post p
INNER JOIN subscription s ON p.owner_id = s.club_id
WHERE s.user_id = sqlc.arg('user_id')
  AND (sqlc.narg('tag')::text IS NULL OR p.tag = sqlc.narg('tag'))
ORDER BY p.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetEveryoneFeed :many
SELECT *
FROM post
WHERE (sqlc.narg('tag')::text IS NULL OR tag = sqlc.narg('tag'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetFeedByClubID :many
SELECT *
FROM post
WHERE owner_id = sqlc.arg('owner_id')
  AND (sqlc.narg('tag')::text IS NULL OR tag = sqlc.narg('tag'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetFeedCount :one
SELECT COUNT(*)
FROM post
WHERE (sqlc.narg('tag')::text IS NULL OR tag = sqlc.narg('tag'));

-- name: GetSubscriptionFeedCount :one
SELECT COUNT(DISTINCT p.id)
FROM post p
INNER JOIN subscription s ON p.owner_id = s.club_id
WHERE s.user_id = sqlc.arg('user_id')
  AND (sqlc.narg('tag')::text IS NULL OR p.tag = sqlc.narg('tag'));

-- name: GetFeedByClubIDCount :one
SELECT COUNT(*)
FROM post
WHERE owner_id = sqlc.arg('owner_id')
  AND (sqlc.narg('tag')::text IS NULL OR tag = sqlc.narg('tag'));
