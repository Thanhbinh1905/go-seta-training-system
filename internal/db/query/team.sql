-- name: CreateTeam :one
INSERT INTO teams (
  team_name
)
VALUES ($1)
RETURNING *;

-- name: AddTeamMember :exec
INSERT INTO team_members (
  team_id,
  user_id
)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveTeamMember :exec
DELETE FROM team_members
WHERE team_id = $1 AND user_id = $2;

-- name: AddTeamManager :exec
INSERT INTO team_managers (
  team_id,
  user_id
)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveTeamManager :exec
DELETE FROM team_managers
WHERE team_id = $1 AND user_id = $2;

-- name: GetTeamMembers :many
SELECT u.user_id, u.username, u.email, u.role
FROM team_members tm
JOIN users u ON u.user_id = tm.user_id
WHERE tm.team_id = $1;

-- name: GetTeamManagers :many
SELECT u.user_id, u.username, u.email, u.role
FROM team_managers tm
JOIN users u ON u.user_id = tm.user_id
WHERE tm.team_id = $1;
