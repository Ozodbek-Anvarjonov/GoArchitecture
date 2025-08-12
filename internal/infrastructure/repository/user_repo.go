package repository

import (
	"Architecture/internal/domain"
	"context"
	"database/sql"
	"errors"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Get(ctx context.Context) ([]domain.User, error) {
	query := `SELECT * FROM users ORDER BY id DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user := new(domain.User)
	query := `SELECT * FROM users WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return user, err
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	var query = `
		INSERT INTO users (first_name, last_name, email)
		VALUES ($1, $2, $3) RETURNING id;
	`

	err := r.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users SET first_name = $1, last_name = $2, email = $3 where id = $4;
	`

	res, err := r.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1;`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}
