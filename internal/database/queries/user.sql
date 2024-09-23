-- name: GetUser :one
select * from users
where id = ? limit 1;

-- name: GetUserByEmail :one
select * from users
where email = ? limit 1;

-- name: ListUsers :many
select * from users
order by name;

-- name: CreateUser :one
insert into users (name, email, password, bio) values (?, ?, ?, ?)
returning *;

-- name: UpdateUser :exec
update users
set name = ?,
    email = ?,
    password = ?,
    bio = ?
where id = ?
returning *;

-- name: DeleteUser :exec
delete from users
where id = ?;