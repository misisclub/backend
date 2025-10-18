package repository

import (
	"context"
	"echo-template/db"
	"echo-template/internal/models"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostRepository handles post data operations
type PostRepository struct {
	db *pgxpool.Pool
}

// Post repository errors
var (
	ErrPostNotFound = errors.New("post not found")
	ErrPostExists   = errors.New("post already exists for this owner")
)

// NewPostRepository creates a new post repository
func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

// CreatePost creates a new post
func (r *PostRepository) CreatePost(ctx context.Context, req *models.CreatePostRequest) (*models.Post, error) {
	query := db.New(r.db)

	post, err := query.CreatePost(ctx, db.CreatePostParams{
		Tag:         req.Tag,
		OwnerID:     req.OwnerID,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
		FromOrg:     req.FromOrg,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, ErrPostExists
			}
		}
		return nil, err
	}

	return &models.Post{
		ID:          post.ID,
		Tag:         post.Tag,
		OwnerID:     post.OwnerID,
		Description: post.Description,
		ImageURL:    post.ImageUrl,
		FromOrg:     post.FromOrg,
		CreatedAt:   post.CreatedAt.Time,
		UpdatedAt:   post.UpdatedAt.Time,
	}, nil
}

// GetPost retrieves a post by ID
func (r *PostRepository) GetPost(ctx context.Context, id uuid.UUID) (*models.Post, error) {
	query := db.New(r.db)

	post, err := query.GetPost(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	return &models.Post{
		ID:          post.ID,
		Tag:         post.Tag,
		OwnerID:     post.OwnerID,
		Description: post.Description,
		ImageURL:    post.ImageUrl,
		FromOrg:     post.FromOrg,
		CreatedAt:   post.CreatedAt.Time,
		UpdatedAt:   post.UpdatedAt.Time,
	}, nil
}

// UpdatePost updates an existing post
func (r *PostRepository) UpdatePost(ctx context.Context, id uuid.UUID, req *models.UpdatePostRequest) (
	*models.Post, error,
) {
	query := db.New(r.db)

	post, err := query.UpdatePost(ctx, db.UpdatePostParams{
		ID:          id,
		Tag:         req.Tag,
		Description: req.Description,
		ImageUrl:    req.ImageURL,
		FromOrg:     req.FromOrg,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}

	return &models.Post{
		ID:          post.ID,
		Tag:         post.Tag,
		OwnerID:     post.OwnerID,
		Description: post.Description,
		ImageURL:    post.ImageUrl,
		FromOrg:     post.FromOrg,
		CreatedAt:   post.CreatedAt.Time,
		UpdatedAt:   post.UpdatedAt.Time,
	}, nil
}

// DeletePost deletes a post by ID
func (r *PostRepository) DeletePost(ctx context.Context, id uuid.UUID) error {
	query := db.New(r.db)

	err := query.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
