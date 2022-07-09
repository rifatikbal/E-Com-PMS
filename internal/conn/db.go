package conn

import (
	"database/sql"
	"fmt"
	// mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rifatikbal/E-Com-PMS/internal/config"
)

var db *DB

type DB struct {
	RawSQL *sql.DB
	GormDB *gorm.DB
}

func ConnectDB(cfg *config.DBCfg) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name)

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	sqldb.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := sqldb.Ping(); err != nil {
		return err
	}

	gormDB, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: sqldb,
		}))
	if err != nil {
		return err
	}

	db = &DB{
		RawSQL: sqldb,
		GormDB: gormDB,
	}

	return nil
}

func GetDB() *DB {
	return db
}
