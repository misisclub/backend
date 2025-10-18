package usecase

import (
	"context"
	"echo-template/db"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"

	"github.com/google/uuid"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context,
	req models.SubscriptionCreate,
) (*models.SubscriptionModel, error) {
	subscription, err := s.repo.CreateSubscription(ctx, db.CreateSubscriptionParams{
		UserID: req.UserID,
		ClubID: req.ClubID,
	})
	if err != nil {
		return nil, err
	}

	return &models.SubscriptionModel{
		ID:     subscription.ID,
		UserID: subscription.UserID,
		ClubID: subscription.ClubID,
	}, nil
}

func (s *SubscriptionService) GetSubscriptionByID(ctx context.Context,
	id uuid.UUID,
) (*models.SubscriptionModel, error) {
	subscription, err := s.repo.GetSubscriptionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.SubscriptionModel{
		ID:     subscription.ID,
		UserID: subscription.UserID,
		ClubID: subscription.ClubID,
	}, nil
}

func (s *SubscriptionService) GetSubscriptionsByUserID(ctx context.Context,
	userID uuid.UUID,
) ([]models.SubscriptionModel, error) {
	subscriptions, err := s.repo.GetSubscriptionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]models.SubscriptionModel, len(subscriptions))
	for i, sub := range subscriptions {
		result[i] = models.SubscriptionModel{
			ID:     sub.ID,
			UserID: sub.UserID,
			ClubID: sub.ClubID,
		}
	}

	return result, nil
}

func (s *SubscriptionService) GetSubscriptionsByClubID(ctx context.Context,
	clubID uuid.UUID,
) ([]models.SubscriptionModel, error) {
	subscriptions, err := s.repo.GetSubscriptionsByClubID(ctx, clubID)
	if err != nil {
		return nil, err
	}

	result := make([]models.SubscriptionModel, len(subscriptions))
	for i, sub := range subscriptions {
		result[i] = models.SubscriptionModel{
			ID:     sub.ID,
			UserID: sub.UserID,
			ClubID: sub.ClubID,
		}
	}

	return result, nil
}

func (s *SubscriptionService) GetSubscriptionByUserAndClub(ctx context.Context,
	userID, clubID uuid.UUID,
) (*models.SubscriptionModel, error) {
	subscription, err := s.repo.GetSubscriptionByUserAndClub(ctx, userID, clubID)
	if err != nil {
		return nil, err
	}

	return &models.SubscriptionModel{
		ID:     subscription.ID,
		UserID: subscription.UserID,
		ClubID: subscription.ClubID,
	}, nil
}

func (s *SubscriptionService) DeleteSubscription(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSubscription(ctx, id)
}

func (s *SubscriptionService) DeleteSubscriptionByUserAndClub(ctx context.Context, userID, clubID uuid.UUID) error {
	return s.repo.DeleteSubscriptionByUserAndClub(ctx, userID, clubID)
}
