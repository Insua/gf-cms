package migrate

import (
	_ "github.com/go-sql-driver/mysql"
)

func Up() {
	dbMigrate("up")
}
