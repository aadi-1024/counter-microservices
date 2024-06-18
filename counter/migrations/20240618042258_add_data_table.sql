-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE data (
    UserId INT PRIMARY KEY,
    Value INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE data;
-- +goose StatementEnd
