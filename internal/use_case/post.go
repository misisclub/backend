package usecase

import (
	"context"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"

	"github.com/google/uuid"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *models.CreatePostRequest) (*models.PostResponse, error) {
	post, err := s.postRepo.CreatePost(ctx, req)
	if err != nil {
		return nil, err
	}

	return &models.PostResponse{
		ID:          post.ID,
		Tag:         post.Tag,
		OwnerID:     post.OwnerID,
		Description: post.Description,
		ContentURL:  post.ContentURL,
		IsVideo:     post.IsVideo,
		FromOrg:     post.FromOrg,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}, nil
}

func (s *PostService) GetPost(ctx context.Context, id uuid.UUID) (*models.PostResponse, error) {
	post, err := s.postRepo.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.PostResponse{
		ID:          post.ID,
		Tag:         post.Tag,
		OwnerID:     post.OwnerID,
		Description: post.Description,
		ContentURL:  post.ContentURL,
		IsVideo:     post.IsVideo,
		FromOrg:     post.FromOrg,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}, nil
}

func (s *PostService) UpdatePost(ctx context.Context, id uuid.UUID,
	req *models.UpdatePostRequest,
) (*models.PostResponse, error) {
	post, err := s.postRepo.UpdatePost(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return &models.PostResponse{
		ID:          post.ID,
		Tag:         post.Tag,
		OwnerID:     post.OwnerID,
		Description: post.Description,
		ContentURL:  post.ContentURL,
		IsVideo:     post.IsVideo,
		FromOrg:     post.FromOrg,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}, nil
}

func (s *PostService) DeletePost(ctx context.Context, id uuid.UUID) error {
	return s.postRepo.DeletePost(ctx, id)
}
