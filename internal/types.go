package internal

type Journal struct {
	Id   *string `json:"id,omitempty" db:"id"`
	Name string  `json:"name" db:"journal_name"`
}
