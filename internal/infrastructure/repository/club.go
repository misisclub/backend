package repository

import (
	"context"
	"echo-template/db"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrClubNotFound = errors.New("club not found")

type ClubRepository struct {
	db *pgxpool.Pool
}

func NewClubRepository(db *pgxpool.Pool) *ClubRepository {
	if db == nil {
		panic("Database connection is nil in club repository")
	}
	return &ClubRepository{db: db}
}

func (r *ClubRepository) CreateClub(ctx context.Context, params db.CreateClubParams) (*db.Club, error) {
	q := db.New(r.db)
	club, err := q.CreateClub(ctx, params)
	return &club, err
}

func (r *ClubRepository) GetClubByID(ctx context.Context, id uuid.UUID) (*db.GetClubByIDRow, error) {
	q := db.New(r.db)
	club, err := q.GetClubByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrClubNotFound
	}
	return &club, err
}

func (r *ClubRepository) DeleteClub(ctx context.Context, id uuid.UUID) error {
	q := db.New(r.db)
	err := q.DeleteClub(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClubRepository) UpdateClub(ctx context.Context, p db.UpdateClubParams) (*db.Club, error) {
	q := db.New(r.db)
	club, err := q.UpdateClub(ctx, p)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrClubNotFound
	}
	return &club, nil
}
