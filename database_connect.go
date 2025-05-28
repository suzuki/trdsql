//go:build !cgo

package trdsql

import (
	"database/sql"

	// MySQL driver.
	_ "github.com/go-sql-driver/mysql"

	// PostgreSQL driver.
	_ "github.com/lib/pq"

	// SQL Server driver.
	_ "github.com/microsoft/go-mssqldb"

	// SQLite3 driver.
	_ "modernc.org/sqlite"
)

var DefaultDriver = "sqlite"

// Connect is connects to the database.
// Currently supported drivers are sqlite3, mysql, postgres, mssql.
// Set quote character and maxBulk depending on the driver type.
func Connect(driver, dsn string) (*DB, error) {
	sqlDB, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	db := &DB{
		DB:     sqlDB,
		driver: driver,
		dsn:    dsn,
	}
	debug.Printf("driver: %s, dsn: %s", driver, dsn)

	switch driver {
	case "sqlite3", "sqlite3_ext", "sqlite":
		db.quote = "`"
		db.maxBulk = 1000
	case "mysql":
		db.quote = "`"
		db.maxBulk = 1000
	case "postgres":
		db.quote = "\""
	case "mssql", "sqlserver":
		db.quote = "["
		db.quoteEnd = "]"
		db.maxBulk = 1000
	default:
		db.quote = "\""
	}

	return db, nil
}
