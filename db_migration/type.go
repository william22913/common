package dbmigration

import "database/sql"

type DBMigration interface {
	DbMigratePostresql(
		db *sql.DB,
		fileMigratePath string,
		schemaName string,
	) (
		int,
		error,
	)
}
