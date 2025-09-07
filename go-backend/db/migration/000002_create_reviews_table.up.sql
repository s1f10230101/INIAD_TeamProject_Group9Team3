CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    spot_id UUID NOT NULL,
    user_id UUID NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (spot_id) REFERENCES spot(id) ON DELETE CASCADE
);
