package postgre

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type dbLogger struct { }

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func Connect() *pg.DB{

	dbHost := viper.GetString("database.db_host")
	dbPort := viper.GetString("database.db_port")
	dbUser := viper.GetString("database.db_user")
	dbPass := viper.GetString("database.db_pass")
	dbName := viper.GetString("database.db_name")
	dbSslMode := viper.GetString("database.db_sslmode")

	parse := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName,dbSslMode)
	opt, err := pg.ParseURL(parse)

	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	if db == nil {
		logrus.Printf("Failed to connect database \n")
	}

	logrus.Printf("Success connected to DB \n")

	return db
}