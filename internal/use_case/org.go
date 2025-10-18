package usecase

import (
	"context"
	"echo-template/db"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"

	"github.com/google/uuid"
)

type OrgService struct {
	clientRepo *repository.OrgRepository
}

func NewOrgService(repo *repository.OrgRepository) *OrgService {
	return &OrgService{clientRepo: repo}
}

func (s *OrgService) CreateOrg(ctx context.Context, c *models.OrgCreate) (*db.Org, error) {
	params := db.CreateOrgParams{
		UserID:       c.UserID,
		Name:         c.Name,
		Specs:        c.Specs,
		Description:  c.Description,
		WebsiteUrl:   c.WebsiteURL,
		LogoUrl:      c.LogoURL,
		VideoUrl:     c.VideoURL,
		AdminContact: c.AdminContact,
		Inn:          c.Inn,
		Ogrn:         c.Ogrn,
	}
	org, err := s.clientRepo.CreateOrg(ctx, params)
	if err != nil {
		return &db.Org{}, err
	}

	return org, nil
}

func (s *OrgService) DeleteOrg(ctx context.Context, id uuid.UUID) error {
	return s.clientRepo.DeleteOrg(ctx, id)
}

func (s *OrgService) UpdateOrg(ctx context.Context, c *models.OrgUpdate, id uuid.UUID,
) (*db.Org, error) {
	params := db.UpdateOrgParams{
		ID:           id,
		Name:         c.Name,
		Specs:        c.Specs,
		Description:  c.Description,
		WebsiteUrl:   c.WebsiteURL,
		LogoUrl:      c.LogoURL,
		VideoUrl:     c.VideoURL,
		AdminContact: c.AdminContact,
		Inn:          c.Inn,
		Ogrn:         c.Ogrn,
	}
	resp, err := s.clientRepo.UpdateOrg(ctx, params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *OrgService) GetOrgByID(ctx context.Context, id uuid.UUID) (*db.Org, error) {
	resp, err := s.clientRepo.GetOrgByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
