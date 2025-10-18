-- name: CreateOrg :one
insert into org (
    user_id,
    name,
    specs,
    description,
    website_url,
    logo_url,
    video_url,
    admin_contact,
    inn,
    ogrn
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
returning *;

-- name: GetOrgByID :one
select * from org where id = $1;

-- name: UpdateOrg :one
update org set
    name           = coalesce(sqlc.narg('name'), name),
    specs          = coalesce(sqlc.narg('specs'), specs),
    description    = coalesce(sqlc.narg('description'), description),
    website_url    = coalesce(sqlc.narg('website_url'), website_url),
    logo_url       = coalesce(sqlc.narg('logo_url'), logo_url),
    video_url      = coalesce(sqlc.narg('video_url'), video_url),
    admin_contact  = coalesce(sqlc.narg('admin_contact'), admin_contact),
    inn            = coalesce(sqlc.narg('inn'), inn),
    ogrn           = coalesce(sqlc.narg('ogrn'), ogrn)
where id = $1
returning *;

-- name: DeleteOrg :exec
delete from org where id = $1;
