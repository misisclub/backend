package usecase

import (
	"context"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"
	"errors"

	"github.com/google/uuid"
)

// FeedService handles feed business logic
type FeedService struct {
	feedRepo *repository.FeedRepository
}

// NewFeedService creates a new feed service
func NewFeedService(feedRepo *repository.FeedRepository) *FeedService {
	return &FeedService{
		feedRepo: feedRepo,
	}
}

// GetSubscriptionFeed retrieves posts from subscribed clubs for a user
func (s *FeedService) GetSubscriptionFeed(ctx context.Context, userID uuid.UUID, tag *string, limit, offset int32) (*models.FeedResponse, error) {
	// Set default limit if not provided
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	// Ensure offset is not negative
	if offset < 0 {
		offset = 0
	}

	posts, err := s.feedRepo.GetSubscriptionFeed(ctx, userID, tag, limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.feedRepo.GetSubscriptionFeedCount(ctx, userID, tag)
	if err != nil {
		return nil, err
	}

	// Convert []*PostResponse to []PostResponse
	var postResponses []models.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, *post)
	}

	return &models.FeedResponse{
		Posts: postResponses,
		Total: total,
	}, nil
}

// GetEveryoneFeed retrieves all posts ordered by creation date
func (s *FeedService) GetEveryoneFeed(ctx context.Context, tag *string, limit, offset int32) (*models.FeedResponse, error) {
	// Set default limit if not provided
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	// Ensure offset is not negative
	if offset < 0 {
		offset = 0
	}

	posts, err := s.feedRepo.GetEveryoneFeed(ctx, tag, limit, offset)
	if err != nil {
		// Handle specific repository errors
		if errors.Is(err, repository.ErrInvalidPagination) {
			return nil, errors.New("invalid pagination parameters")
		}
		if errors.Is(err, repository.ErrDatabaseError) {
			return nil, errors.New("failed to retrieve posts from database")
		}
		return nil, err
	}

	total, err := s.feedRepo.GetFeedCount(ctx, tag)
	if err != nil {
		return nil, errors.New("failed to retrieve total post count")
	}

	// Convert []*PostResponse to []PostResponse
	var postResponses []models.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, *post)
	}

	return &models.FeedResponse{
		Posts: postResponses,
		Total: total,
	}, nil
}

// GetFeedByClubID retrieves posts from a specific club
func (s *FeedService) GetFeedByClubID(ctx context.Context, clubID uuid.UUID, tag *string, limit, offset int32) (*models.FeedResponse, error) {
	// Set default limit if not provided
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	// Ensure offset is not negative
	if offset < 0 {
		offset = 0
	}

	posts, err := s.feedRepo.GetFeedByClubID(ctx, clubID, tag, limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.feedRepo.GetFeedByClubIDCount(ctx, clubID, tag)
	if err != nil {
		return nil, err
	}

	// Convert []*PostResponse to []PostResponse
	var postResponses []models.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, *post)
	}

	return &models.FeedResponse{
		Posts: postResponses,
		Total: total,
	}, nil
}
