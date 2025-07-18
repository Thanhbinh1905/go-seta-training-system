-- Folder Management

-- name: CreateFolder :one
INSERT INTO folders (
  name, owner_id
)
VALUES ($1, $2)
RETURNING *;

-- name: GetFolder :one
SELECT * FROM folders WHERE folder_id = $1;

-- name: UpdateFolder :exec
UPDATE folders
SET name = $2
WHERE folder_id = $1;

-- name: DeleteFolder :exec
DELETE FROM folders
WHERE folder_id = $1;

-- name: DeleteNotesByFolderID :exec
DELETE FROM notes
WHERE folder_id = $1;

-- Note Management

-- name: CreateNote :one
INSERT INTO notes (
  title, body, folder_id, owner_id
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetNote :one
SELECT * FROM notes WHERE note_id = $1;

-- name: UpdateNote :exec
UPDATE notes
SET title = $2, body = $3
WHERE note_id = $1;

-- name: DeleteNote :exec
DELETE FROM notes WHERE note_id = $1;

-- Sharing API

-- name: ShareFolder :exec
INSERT INTO folder_shares (
  folder_id, user_id, access
)
VALUES ($1, $2, $3)
ON CONFLICT (folder_id, user_id) DO UPDATE
SET access = EXCLUDED.access;

-- name: RevokeFolderShare :exec
DELETE FROM folder_shares
WHERE folder_id = $1 AND user_id = $2;

-- name: ShareNote :exec
INSERT INTO note_shares (
  note_id, user_id, access
)
VALUES ($1, $2, $3)
ON CONFLICT (note_id, user_id) DO UPDATE
SET access = EXCLUDED.access;

-- name: RevokeNoteShare :exec
DELETE FROM note_shares
WHERE note_id = $1 AND user_id = $2;
