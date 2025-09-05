-- name: CreateReview :one
INSERT INTO reviews (
    id,
    spot_id,
    user_id,
    rating,
    comment,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetReviewsBySpotID :many
SELECT * FROM reviews
WHERE spot_id = $1
ORDER BY created_at DESC;
