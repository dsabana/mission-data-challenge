package internal

type Journal struct {
	Id   *string `json:"id,omitempty" db:"id"`
	Name string  `json:"name" db:"journal_name"`
}

type Entry struct {
	Id        *string `json:"id,omitempty" db:"id"`
	JournalID string  `json:"journal_id" db:"journal_id"`
	Content   string  `json:"content" db:"content"`
}
