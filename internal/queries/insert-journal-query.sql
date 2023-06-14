INSERT INTO journal (
    journal_name
) VALUES ($1)
RETURNING *;