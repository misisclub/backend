package repository

import (
	"context"
	"echo-template/db"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrOrgNotFound = errors.New("client not found")

type OrgRepository struct {
	db *pgxpool.Pool
}

func NewOrgRepository(db *pgxpool.Pool) *OrgRepository {
	if db == nil {
		panic("Database connection is nil in org repository")
	}
	return &OrgRepository{db: db}
}

func (r *OrgRepository) CreateOrg(ctx context.Context, params db.CreateOrgParams) (*db.Org, error) {
	q := db.New(r.db)
	client, err := q.CreateOrg(ctx, params)
	return &client, err
}

func (r *OrgRepository) GetOrgByID(ctx context.Context, id uuid.UUID) (*db.Org, error) {
	q := db.New(r.db)
	client, err := q.GetOrgByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrOrgNotFound
	}
	return &client, err
}

func (r *OrgRepository) DeleteOrg(ctx context.Context, id uuid.UUID) error {
	q := db.New(r.db)
	err := q.DeleteOrg(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrgRepository) UpdateOrg(ctx context.Context, p db.UpdateOrgParams) (*db.Org, error) {
	q := db.New(r.db)
	client, err := q.UpdateOrg(ctx, p)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrOrgNotFound
	}
	return &client, nil
}
