package repositorys

import (
	"fmt"
	"time"

	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/models"
	"github.com/gofiber/fiber/v2"
)

type PostgreSQLRepository struct {
	db *database.Database
}

func NewPostgreSQLRepository(db *database.Database) *PostgreSQLRepository {
	return &PostgreSQLRepository{db}
}

func (r *PostgreSQLRepository) GetEmail(ctx *fiber.Ctx, email string) (int64, error) {
	var count int64
	query := `SELECT COUNT(*) FROM users where email=$1`
	err := r.db.GetContext(ctx.Context(), &count, query, email)
	if err != nil {
		return 1, err
	}
	return count, nil
}

func (r *PostgreSQLRepository) CreateUser(ctx *fiber.Ctx, user models.User) error {
	query := `INSERT INTO users (email,password,first_name,last_name ) VALUES ($1,$2,$3,$4)`
	_, err := r.db.ExecContext(ctx.Context(), query, user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLRepository) GetUserByEmail(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	query := `SELECT * FROM users WHERE email=$1`
	err := r.db.GetContext(ctx.Context(), &user, query, user.Email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func (r *PostgreSQLRepository) UpdateToken(userID int64, token string, refreshToken string, updatedAt time.Time) error {
	// Update the tokens
	query := `UPDATE users SET 
			token=$1, refresh_token=$2, updated_at=$3 
			WHERE id = $4`
	_, err := r.db.Exec(query, token, refreshToken, updatedAt, userID)
	if err != nil {
		return fmt.Errorf("failed to update token: %w", err)
	}

	return nil
}
