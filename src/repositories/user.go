package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (repository UsersRepository) Find(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	lines, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE name LIKE ? or nickname LIKE ?",
		nameOrNickname,
		nameOrNickname,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository UsersRepository) Get(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err := lines.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository UsersRepository) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}
