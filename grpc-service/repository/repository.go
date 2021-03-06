package repository

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	entities "github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

type sqlRepo struct {
	DB     *sql.DB
	Logger log.Logger
}

func NewSQL(db *sql.DB, log log.Logger) *sqlRepo {
	return &sqlRepo{db, log}
}

func (repo *sqlRepo) CreateUser(ctx context.Context, user entities.User) (int64, error) {

	stmt, err := repo.DB.Prepare("INSERT INTO USER (first_name, age, pass) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.Name, user.Age, user.Pass)
	if err != nil {
		return 0, err
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	repo.Logger.Log(res, "rows affected")

	return userId, nil
}

func (repo *sqlRepo) GetUser(ctx context.Context, userId int64) (entities.User, error) {

	stmt, err := repo.DB.Query("SELECT first_name, id, age FROM USER WHERE ID = ?", userId)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{}
	for stmt.Next() {
		err := stmt.Scan(&user.Name, &user.Age)
		if err != nil {
			level.Error(repo.Logger).Log("error")
		}
	}
	return user, nil
}
