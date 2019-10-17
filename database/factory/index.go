package factory

import (
	"library-api/database"
)

// dropAndCreateTable 清空表
func dropAndCreateTable(table interface{}) {
	database.DB.DropTable(table)
	database.DB.CreateTable(table)
}

// Mock -
func Mock() {
	UsersTableSeeder(true)
}
