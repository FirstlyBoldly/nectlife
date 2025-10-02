package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	slog.Info("Initializing Database seeder.")

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, config.Data.POSTGRES_DB_URL)
	if err != nil {
		panic(err)
	}

	defer pool.Close()
	db := postgres.New(pool)

	if err = db.FlushUsers(ctx); err != nil {
		panic(err)
	}

	if err = db.FlushCourses(ctx); err != nil {
		panic(err)
	}

	if err = db.FlushRoles(ctx); err != nil {
		panic(err)
	}

	var course postgres.Course
	if course, err = db.CreateCourse(ctx, postgres.CreateCourseParams{
		DepartmentID: pgtype.Int4{
			Int32: 123,
			Valid: false,
		},
		KeyName: "know_nothing",
	}); err != nil {
		panic(err)
	}

	var role postgres.Role
	if role, err = db.CreateRole(ctx, postgres.CreateRoleParams{
		KeyName: "superuser",
		Description: pgtype.Text{
			String: "Highest executive holder.",
		},
	}); err != nil {
		panic(err)
	}

	for i := range 5 {
		password, err := bcrypt.GenerateFromPassword(
			[]byte(fmt.Sprintf("password#%d", i+1)),
			bcrypt.DefaultCost,
		)
		if err != nil {
			panic(err)
		}

		if _, err = db.CreateUser(ctx, postgres.CreateUserParams{
			CourseID:     course.ID,
			RoleID:       role.ID,
			StudentID:    fmt.Sprintf("tk19000%d", i+1),
			FirstName:    fmt.Sprintf("User #%d", i+1),
			LastName:     "Test",
			Email:        fmt.Sprintf("user%d@test.com", i+1),
			PasswordHash: string(password),
		}); err != nil {
			panic(err)
		}
	}

	slog.Info("Database seeding completed!")
}
