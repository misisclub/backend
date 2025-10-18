-- name: CreateClub :one
insert into club (
    owner_id,
    name,
    specs,
    description,
    tg_url,
    logo_url,
    admin_contact,
    from_org
) values (
    $1, $2, $3, $4, $5, $6, $7, $8
)
returning *;

-- name: GetClubByID :one
select 
    c.id, c.owner_id, c.name, c.specs, c.description, c.tg_url, c.logo_url, c.admin_contact, c.from_org,
    o.name as org_name
from club c
left join org o on c.owner_id = o.user_id and c.from_org = true
where c.id = $1;

-- name: UpdateClub :one
update club set
    name = coalesce(sqlc.narg('name'), name),
    specs = coalesce(sqlc.narg('specs'), specs),
    description = coalesce(sqlc.narg('description'), description),
    tg_url = coalesce(sqlc.narg('tg_url'), tg_url),
    logo_url = coalesce(sqlc.narg('logo_url'), logo_url),
    admin_contact = coalesce(sqlc.narg('admin_contact'), admin_contact),
    from_org = coalesce(sqlc.narg('from_org'), from_org)
where id = $1
returning *;

-- name: DeleteClub :exec
delete from club where id = $1;

-- name: GetClubsByOwner :many
select 
    c.id, c.owner_id, c.name, c.specs, c.description, c.tg_url, c.logo_url, c.admin_contact, c.from_org,
    o.name as org_name
from club c
left join org o on c.owner_id = o.user_id and c.from_org = true
where c.owner_id = $1;
