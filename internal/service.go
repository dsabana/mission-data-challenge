package internal

//go:generate go install github.com/golang/mock/mockgen@v1.6.0
//go:generate mockgen -source ./service.go -package internal_test -destination ./mocks/mocks.go Repository

import "context"

type Repository interface {
	SaveJournal(ctx context.Context, journal Journal) (*Journal, error)
	RetrieveAllJournals(ctx context.Context) (*[]Journal, error)
	SaveEntry(ctx context.Context, journalID string, entry Entry) (*Entry, error)
	RetrieveAllEntries(ctx context.Context, journalID string) (*[]Entry, error)
}

type Service interface {
	AddJournal(ctx context.Context, journal Journal) (*Journal, error)
	GetAllJournals(ctx context.Context) (*[]Journal, error)
	AddEntry(ctx context.Context, journalID string, entry Entry) (*Entry, error)
	GetAllEntries(ctx context.Context, journalID string) (*[]Entry, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddJournal(ctx context.Context, journal Journal) (*Journal, error) {
	return s.r.SaveJournal(ctx, journal)
}

func (s *service) GetAllJournals(ctx context.Context) (*[]Journal, error) {
	return s.r.RetrieveAllJournals(ctx)
}

func (s *service) AddEntry(ctx context.Context, journalID string, entry Entry) (*Entry, error) {
	return s.r.SaveEntry(ctx, journalID, entry)
}

func (s *service) GetAllEntries(ctx context.Context, journalID string) (*[]Entry, error) {
	return s.r.RetrieveAllEntries(ctx, journalID)
}
