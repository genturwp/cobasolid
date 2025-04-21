package main

import (
	"cobasolid/entity"
	"cobasolid/repository"
	"cobasolid/service"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const dbUrl = "postgresql://gentur:rahasia@localhost:5432/cobasolid?sslmode=disable"

func dbConn(ctx context.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbUrl)

	if err != nil {
		return nil, err
	}
	config.MaxConns = 4
	return pgxpool.NewWithConfig(ctx, config)
}

func migrateDb() error {

	m, err := migrate.New("file://db/migrate", dbUrl)
	if err != nil {
		log.Println("error bung saat baca migration file")
		return err
	}
	if err = m.Up(); err != nil {

		if err != migrate.ErrNoChange {
			log.Println("error saat running sql migration")
			return err
		}
	}
	return nil
}

func main() {
	fmt.Println("hello world")
	ctx := context.Background()
	pool, err := dbConn(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalln("ping failed ", err)
	}
	err = migrateDb()
	if err != nil {
		log.Fatalln(err)
	}
	repo := repository.NewRepository(pool)
	userProfile := &entity.UserProfile{
		ProfileName: "profile1",
		Phone:       "phone",
		Email:       "profile1@gmail.com",
	}
	up, err := repo.UserProfileRepository.Save(ctx, userProfile)
	fmt.Printf("result dari userProfile = %v", up)
	svc := service.NewUserManagementService(repo)
	svc.UserRegistration(ctx, nil)
}
