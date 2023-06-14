package internal

//go:generate go install github.com/golang/mock/mockgen@v1.6.0
//go:generate mockgen -source ./service.go -package internal_test -destination ./mocks/mocks.go Repository

import "context"

type Repository interface {
	SaveJournal(ctx context.Context, journal Journal) (*Journal, error)
	RetrieveAllJournals(ctx context.Context) (*[]Journal, error)
}

type Service interface {
	AddJournal(ctx context.Context, journal Journal) (*Journal, error)
	GetAllJournals(ctx context.Context) (*[]Journal, error)
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
