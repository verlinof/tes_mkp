-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS studios (
    id SERIAL PRIMARY KEY,
    cinema_id INT NOT NULL REFERENCES cinemas(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    total_seats INT NOT NULL DEFAULT 50,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_studios_cinema_id ON studios(cinema_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS studios;
-- +goose StatementEnd
