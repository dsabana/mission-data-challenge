INSERT INTO entry (
    journal_id,
    content
) VALUES ($1, $2)
RETURNING *;