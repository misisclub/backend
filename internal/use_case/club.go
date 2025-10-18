package usecase

import (
	"context"
	"echo-template/db"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"

	"github.com/google/uuid"
)

type ClubService struct {
	clubRepo *repository.ClubRepository
}

func NewClubService(repo *repository.ClubRepository) *ClubService {
	return &ClubService{clubRepo: repo}
}

func (s *ClubService) CreateClub(ctx context.Context, c *models.ClubCreate) (*db.Club, error) {
	params := db.CreateClubParams{
		OwnerID:      c.OwnerID,
		Name:         c.Name,
		Specs:        c.Specs,
		Description:  c.Description,
		TgUrl:        c.TgURL,
		LogoUrl:      c.LogoURL,
		AdminContact: c.AdminContact,
		FromOrg:      c.FromOrg,
	}
	club, err := s.clubRepo.CreateClub(ctx, params)
	if err != nil {
		return &db.Club{}, err
	}

	return club, nil
}

func (s *ClubService) DeleteClub(ctx context.Context, id uuid.UUID) error {
	return s.clubRepo.DeleteClub(ctx, id)
}

func (s *ClubService) UpdateClub(ctx context.Context, c *models.ClubUpdate, id uuid.UUID) (*db.Club, error) {
	params := db.UpdateClubParams{
		ID:           id,
		Name:         c.Name,
		Specs:        c.Specs,
		Description:  c.Description,
		TgUrl:        c.TgURL,
		LogoUrl:      c.LogoURL,
		AdminContact: c.AdminContact,
		FromOrg:      c.FromOrg,
	}
	resp, err := s.clubRepo.UpdateClub(ctx, params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ClubService) GetClubByID(ctx context.Context, id uuid.UUID) (*db.GetClubByIDRow, error) {
	resp, err := s.clubRepo.GetClubByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
