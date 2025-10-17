-- name: CreateClient :one
insert into client (
    first_name,
    second_name,
    family_name,
    age,
    phone_number,
    profession_1,
    profession_2,
    company,
    university,
    password_hash
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
returning *;

-- name: GetClientByPhone :one
select * from client
where phone_number = $1;

-- name: GetClientByID :one
select * from client
where id = $1;

-- name: UpdateClient :one
update client
set
    first_name   = coalesce(sqlc.narg('first_name'), first_name),
    second_name  = coalesce(sqlc.narg('second_name'), second_name),
    family_name  = coalesce(sqlc.narg('family_name'), family_name),
    age          = coalesce(sqlc.narg('age'), age),
    phone_number = coalesce(sqlc.narg('phone_number'), phone_number),
    profession_1 = coalesce(sqlc.narg('profession_1'), profession_1),
    profession_2 = coalesce(sqlc.narg('profession_2'), profession_2),
    company      = coalesce(sqlc.narg('company'), company),
    university   = coalesce(sqlc.narg('university'), university)
where
    id = $1
returning *;

-- name: DeleteClient :exec
delete from client
where id = $1;

