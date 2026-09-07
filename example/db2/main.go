package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/db2dialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull"`
	Email     string    `bun:"email,notnull,unique"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func main() {
	ctx := context.Background()

	// Connection string for IBM DB2, see github.com/ibmdb/go_ibm_db for details.
	dsn := os.Getenv("DB2_DSN")
	if dsn == "" {
		dsn = "HOSTNAME=localhost;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=password"
	}

	sqldb, err := sql.Open("go_ibm_db", dsn)
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, db2dialect.New())
	defer db.Close()

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	var tableName string
	err = db.NewRaw("SELECT TABNAME FROM SYSCAT.TABLES WHERE TABSCHEMA = CURRENT SCHEMA AND TABNAME = 'USERS' FETCH FIRST 1 ROW ONLY").
		Scan(ctx, &tableName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == sql.ErrNoRows {
		if _, err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx); err != nil {
			panic(err)
		}
	}

	user := &User{
		Name:  "Alice",
		Email: fmt.Sprintf("alice+%d@example.com", time.Now().UnixNano()),
	}
	if _, err := db.NewInsert().Model(user).Exec(ctx); err != nil {
		panic(err)
	}

	var users []User
	if err := db.NewSelect().
		Model(&users).
		OrderExpr(`"user"."id" ASC`).
		Limit(10).
		Scan(ctx); err != nil {
		panic(err)
	}
	fmt.Printf("users: %v\n", users)
}
