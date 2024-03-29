-- name: GetHero :one
SELECT * FROM hero
WHERE id = $1 LIMIT 1;

-- name: ListHeros :many
SELECT * FROM hero
ORDER BY name;

-- name: CreateHero :one
INSERT INTO hero (
  id,
  name
) VALUES (
  $1,
  $2
)
RETURNING *;

-- name: UpdateHero :exec
UPDATE hero
  set name = $2
WHERE id = $1;

-- name: DeleteHero :exec
DELETE FROM hero
WHERE id = $1;
