-- name: CreateSubscription :one
insert into subscription (
    user_id,
    club_id
) values (
    $1, $2
)
returning *;

-- name: GetSubscriptionByID :one
select * from subscription
where id = $1;

-- name: GetSubscriptionsByUserID :many
select * from subscription
where user_id = $1;

-- name: GetSubscriptionsByClubID :many
select * from subscription
where club_id = $1;

-- name: GetSubscriptionByUserAndClub :one
select * from subscription
where user_id = $1 and club_id = $2;

-- name: DeleteSubscription :exec
delete from subscription
where id = $1;

-- name: DeleteSubscriptionByUserAndClub :exec
delete from subscription
where user_id = $1 and club_id = $2;
