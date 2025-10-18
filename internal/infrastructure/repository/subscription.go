package repository

import (
	"context"
	"echo-template/db"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrSubscriptionNotFound = errors.New("subscription not found")

type SubscriptionRepository struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepository(db *pgxpool.Pool) *SubscriptionRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) CreateSubscription(ctx context.Context, params db.CreateSubscriptionParams) (*db.Subscription, error) {
	q := db.New(r.db)
	subscription, err := q.CreateSubscription(ctx, params)
	return &subscription, err
}

func (r *SubscriptionRepository) GetSubscriptionByID(ctx context.Context, id uuid.UUID) (*db.Subscription, error) {
	q := db.New(r.db)
	subscription, err := q.GetSubscriptionByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrSubscriptionNotFound
	}
	return &subscription, err
}

func (r *SubscriptionRepository) GetSubscriptionsByUserID(ctx context.Context, userID uuid.UUID) ([]db.Subscription, error) {
	q := db.New(r.db)
	return q.GetSubscriptionsByUserID(ctx, userID)
}

func (r *SubscriptionRepository) GetSubscriptionsByClubID(ctx context.Context, clubID uuid.UUID) ([]db.Subscription, error) {
	q := db.New(r.db)
	return q.GetSubscriptionsByClubID(ctx, clubID)
}

func (r *SubscriptionRepository) GetSubscriptionByUserAndClub(ctx context.Context, userID, clubID uuid.UUID) (*db.Subscription, error) {
	q := db.New(r.db)
	subscription, err := q.GetSubscriptionByUserAndClub(ctx, db.GetSubscriptionByUserAndClubParams{
		UserID: userID,
		ClubID: clubID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrSubscriptionNotFound
	}
	return &subscription, err
}

func (r *SubscriptionRepository) DeleteSubscription(ctx context.Context, id uuid.UUID) error {
	q := db.New(r.db)
	err := q.DeleteSubscription(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) DeleteSubscriptionByUserAndClub(ctx context.Context, userID, clubID uuid.UUID) error {
	q := db.New(r.db)
	err := q.DeleteSubscriptionByUserAndClub(ctx, db.DeleteSubscriptionByUserAndClubParams{
		UserID: userID,
		ClubID: clubID,
	})
	if err != nil {
		return err
	}
	return nil
}
