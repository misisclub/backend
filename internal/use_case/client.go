package usecase

import (
	"context"
	"echo-template/db"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"
	"echo-template/pkg/hash"

	tokenjwt "echo-template/pkg/token_jwt"

	"github.com/google/uuid"
)

type ClientService struct {
	secretKey  string
	clientRepo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{clientRepo: repo}
}

func (s *ClientService) SignUpClient(ctx context.Context, c *models.ClientSignUp) (*models.SignSuccess, error) {
	pswHsh, err := hash.GenerateHash(c.Password)
	if err != nil {
		return nil, err
	}
	params := db.CreateClientParams{
		FirstName:    c.FirstName,
		SecondName:   c.SecondName,
		FamilyName:   c.FamilyName,
		Age:          c.Age,
		PhoneNumber:  c.PhoneNumber,
		PasswordHash: pswHsh,
	}
	client, err := s.clientRepo.CreateClient(ctx, params)
	if err != nil {
		return &models.SignSuccess{}, err
	}

	token, err := tokenjwt.GenerateJWT(client.ID.String(), s.secretKey)
	if err != nil {
		return &models.SignSuccess{}, err
	}

	return &models.SignSuccess{
		Token: token,
		ID:    client.ID,
	}, nil
}

func (s *ClientService) SignInClient(ctx context.Context, c *models.ClientSignIn) (*models.SignSuccess, error) {
	client, err := s.clientRepo.GetClientByPhone(ctx, c.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if err := hash.ComparePassword(c.Password, client.PasswordHash); err != nil {
		return nil, err
	}

	token, err := tokenjwt.GenerateJWT(client.ID.String(), s.secretKey)
	if err != nil {
		return nil, err
	}

	return &models.SignSuccess{
		Token: token,
		ID:    client.ID,
	}, nil
}

func (s *ClientService) DeleteClient(ctx context.Context, id uuid.UUID) error {
	return s.clientRepo.DeleteClient(ctx, id)
}

func (s *ClientService) UpdateClient(ctx context.Context, c *models.ClientUpdate, id uuid.UUID,
) (*models.ClientModel, error) {
	params := db.UpdateClientParams{
		ID:          id,
		FirstName:   c.FirstName,
		SecondName:  c.SecondName,
		FamilyName:  c.FamilyName,
		Age:         c.Age,
		Profession1: c.Profession1,
		Profession2: c.Profession2,
		Company:     c.Company,
		Future:      c.Future,
		University:  c.University,
	}
	cc, err := s.clientRepo.UpdateClient(ctx, params)
	if err != nil {
		return nil, err
	}
	resp := models.ClientModel{
		ID:          cc.ID,
		FirstName:   cc.FirstName,
		SecondName:  cc.SecondName,
		FamilyName:  cc.FamilyName,
		Age:         cc.Age,
		PhoneNumber: cc.PhoneNumber,
		Profession1: cc.Profession1,
		Profession2: cc.Profession2,
		Company:     cc.Company,
		University:  cc.University,
	}
	return &resp, nil
}

func (s *ClientService) GetClientByID(ctx context.Context, id uuid.UUID) (*models.ClientModel, error) {
	c, err := s.clientRepo.GetClientByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := &models.ClientModel{
		ID:          c.ID,
		FirstName:   c.FirstName,
		SecondName:  c.SecondName,
		FamilyName:  c.FamilyName,
		Age:         c.Age,
		PhoneNumber: c.PhoneNumber,
		Profession1: c.Profession1,
		Profession2: c.Profession2,
		Company:     c.Company,
		Future:      c.Future,
		University:  c.University,
	}
	return resp, nil
}
