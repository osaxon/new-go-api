-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ? LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY due_date;

-- name: ListUsersTasks :many
SELECT * FROM tasks
WHERE user_id = ?
ORDER BY due_date;

-- name: CreateTask :one
INSERT INTO tasks (title, user_id, due_date, status, description)
VALUES (?, ?, ?, ?, ?)
    RETURNING *;

-- name: UpdateTask :exec
UPDATE tasks
SET title = ?,
    user_id = ?,
    due_date = ?,
    status = ?,
    description = ?
WHERE id = ?
    RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;
