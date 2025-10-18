package repository

import (
	"context"
	"echo-template/db"
	"echo-template/internal/models"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// FeedRepository handles feed data operations
type FeedRepository struct {
	db *pgxpool.Pool
}

// Feed repository errors
var (
	ErrFeedNotFound      = errors.New("feed not found")
	ErrInvalidTag        = errors.New("invalid tag provided")
	ErrDatabaseError     = errors.New("database operation failed")
	ErrInvalidPagination = errors.New("invalid pagination parameters")
)

// NewFeedRepository creates a new feed repository
func NewFeedRepository(db *pgxpool.Pool) *FeedRepository {
	return &FeedRepository{db: db}
}

// GetSubscriptionFeed retrieves posts from subscribed clubs for a user
func (r *FeedRepository) GetSubscriptionFeed(ctx context.Context, userID uuid.UUID, tag *string, limit, offset int32) ([]*models.PostResponse, error) {
	query := db.New(r.db)

	posts, err := query.GetSubscriptionFeed(ctx, db.GetSubscriptionFeedParams{
		UserID: userID,
		Tag:    tag,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*models.PostResponse{}, nil
		}
		return nil, err
	}

	var result []*models.PostResponse
	for _, post := range posts {
		result = append(result, &models.PostResponse{
			ID:          post.ID,
			Tag:         post.Tag,
			OwnerID:     post.OwnerID,
			Description: post.Description,
			ContentURL:  post.ContentUrl,
			IsVideo:     post.IsVideo,
			FromOrg:     post.FromOrg,
			CreatedAt:   post.CreatedAt.Time,
			UpdatedAt:   post.UpdatedAt.Time,
		})
	}

	return result, nil
}

// GetEveryoneFeed retrieves all posts ordered by creation date
func (r *FeedRepository) GetEveryoneFeed(ctx context.Context, tag *string, limit, offset int32) ([]*models.PostResponse, error) {
	// Validate pagination parameters
	if limit <= 0 || offset < 0 {
		return nil, ErrInvalidPagination
	}

	query := db.New(r.db)

	posts, err := query.GetEveryoneFeed(ctx, db.GetEveryoneFeedParams{
		Tag:    tag,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*models.PostResponse{}, nil
		}
		// Wrap database errors with context
		return nil, errors.Join(ErrDatabaseError, err)
	}

	var result []*models.PostResponse
	for _, post := range posts {
		result = append(result, &models.PostResponse{
			ID:          post.ID,
			Tag:         post.Tag,
			OwnerID:     post.OwnerID,
			Description: post.Description,
			ContentURL:  post.ContentUrl,
			IsVideo:     post.IsVideo,
			FromOrg:     post.FromOrg,
			CreatedAt:   post.CreatedAt.Time,
			UpdatedAt:   post.UpdatedAt.Time,
		})
	}

	return result, nil
}

// GetFeedByClubID retrieves posts from a specific club
func (r *FeedRepository) GetFeedByClubID(ctx context.Context, clubID uuid.UUID, tag *string, limit, offset int32) ([]*models.PostResponse, error) {
	query := db.New(r.db)

	posts, err := query.GetFeedByClubID(ctx, db.GetFeedByClubIDParams{
		OwnerID: clubID,
		Tag:     tag,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*models.PostResponse{}, nil
		}
		return nil, err
	}

	var result []*models.PostResponse
	for _, post := range posts {
		result = append(result, &models.PostResponse{
			ID:          post.ID,
			Tag:         post.Tag,
			OwnerID:     post.OwnerID,
			Description: post.Description,
			FromOrg:     post.FromOrg,
			CreatedAt:   post.CreatedAt.Time,
			UpdatedAt:   post.UpdatedAt.Time,
		})
	}

	return result, nil
}

// GetFeedCount retrieves total count of all posts
func (r *FeedRepository) GetFeedCount(ctx context.Context, tag *string) (int64, error) {
	query := db.New(r.db)

	count, err := query.GetFeedCount(ctx, tag)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetSubscriptionFeedCount retrieves total count of posts from subscribed clubs for a user
func (r *FeedRepository) GetSubscriptionFeedCount(ctx context.Context, userID uuid.UUID, tag *string) (int64, error) {
	query := db.New(r.db)

	count, err := query.GetSubscriptionFeedCount(ctx, db.GetSubscriptionFeedCountParams{
		UserID: userID,
		Tag:    tag,
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetFeedByClubIDCount retrieves total count of posts from a specific club
func (r *FeedRepository) GetFeedByClubIDCount(ctx context.Context, clubID uuid.UUID, tag *string) (int64, error) {
	query := db.New(r.db)

	count, err := query.GetFeedByClubIDCount(ctx, db.GetFeedByClubIDCountParams{
		OwnerID: clubID,
		Tag:     tag,
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}
