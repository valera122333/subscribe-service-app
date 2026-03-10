package service

import (
	"context"
	"subscriptions-service/internal/model"
	"subscriptions-service/internal/repository"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, sub model.Subscription) (uuid.UUID, error) {
	id, err := s.repo.Create(ctx, sub)
	return id, err
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (model.Subscription, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]model.Subscription, error) {
	return s.repo.List(ctx)
}
func (s *Service) Total(ctx context.Context, userID uuid.UUID, serviceName, from, to string) (int, error) {
	return s.repo.Sum(ctx, userID, serviceName, from, to)
}
func (s *Service) ListFiltered(ctx context.Context, userID uuid.UUID, serviceName string) ([]model.Subscription, error) {
	return s.repo.ListFiltered(ctx, userID, serviceName)
}
