-- name: CreateSpot :one
INSERT INTO Spot (
    Id,
    Name,
    Description,
    Address
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetSpot :one
SELECT * FROM Spot
WHERE Id = $1;

-- name: ListSpots :many
SELECT * FROM Spot
ORDER BY Created_at DESC;

-- name: UpdateSpot :one
UPDATE Spot
SET 
    Name = $2,
    Description = $3,
    Address = $4
WHERE Id = $1
RETURNING *;

-- name: DeleteSpot :exec
DELETE FROM Spot
WHERE Id = $1;

-- name: SearchSpots :many
SELECT * FROM Spot
WHERE name LIKE $1 OR description LIKE $1
ORDER BY created_at DESC;
