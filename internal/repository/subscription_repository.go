package repository

import (
	"context"
	"fmt"
	"subscriptions-service/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, sub model.Subscription) (uuid.UUID, error) {
	sub.ID = uuid.New()
	_, err := r.db.Exec(ctx,
		"INSERT INTO subscriptions (id, service_name, price, user_id, start_date) VALUES ($1,$2,$3,$4,$5)",
		sub.ID, sub.Service, sub.Price, sub.UserID, sub.StartDate)
	return sub.ID, err
}
func (r *Repository) Get(ctx context.Context, id uuid.UUID) (model.Subscription, error) {

	var s model.Subscription

	query := `
	SELECT id,service_name,price,user_id,start_date,end_date
	FROM subscriptions
	WHERE id=$1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&s.ID,
		&s.Service,
		&s.Price,
		&s.UserID,
		&s.StartDate,
		&s.EndDate,
	)

	return s, err
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {

	_, err := r.db.Exec(ctx,
		"DELETE FROM subscriptions WHERE id=$1",
		id)

	return err
}
func (r *Repository) Sum(ctx context.Context, userID uuid.UUID, serviceName, from, to string) (int, error) {
	query := `SELECT COALESCE(SUM(price), 0) FROM subscriptions WHERE 1=1`
	args := []interface{}{}
	i := 1

	if userID != uuid.Nil {
		query += ` AND user_id = $` + fmt.Sprint(i)
		args = append(args, userID)
		i++
	}
	if serviceName != "" {
		query += ` AND service_name = $` + fmt.Sprint(i)
		args = append(args, serviceName)
		i++
	}
	if from != "" {
		query += ` AND start_date >= $` + fmt.Sprint(i)
		args = append(args, from)
		i++
	}
	if to != "" {
		query += ` AND start_date <= $` + fmt.Sprint(i)
		args = append(args, to)
		i++
	}

	var total int
	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	return total, err
}
func (r *Repository) List(ctx context.Context) ([]model.Subscription, error) {

	rows, err := r.db.Query(ctx,
		"SELECT id,service_name,price,user_id,start_date,end_date FROM subscriptions")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []model.Subscription

	for rows.Next() {

		var s model.Subscription

		err := rows.Scan(
			&s.ID,
			&s.Service,
			&s.Price,
			&s.UserID,
			&s.StartDate,
			&s.EndDate,
		)

		if err != nil {
			return nil, err
		}

		list = append(list, s)
	}

	return list, nil
}
func (r *Repository) ListFiltered(ctx context.Context, userID uuid.UUID, serviceName string) ([]model.Subscription, error) {
	query := "SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	i := 1

	if userID != uuid.Nil {
		query += " AND user_id = $" + fmt.Sprint(i)
		args = append(args, userID)
		i++
	}

	if serviceName != "" {
		query += " AND service_name = $" + fmt.Sprint(i)
		args = append(args, serviceName)
		i++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Subscription
	for rows.Next() {
		var s model.Subscription
		if err := rows.Scan(&s.ID, &s.Service, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}
