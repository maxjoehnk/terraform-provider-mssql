package connector

import (
	"context"
	"database/sql"
	"fmt"
)

func (c MssqlConnector) CreateDatabase(ctx context.Context, name string) error {
	return c.Execute(ctx, fmt.Sprintf("CREATE DATABASE %s", name))
}

func (c MssqlConnector) SetDatabaseOwner(ctx context.Context, db string, owner string) error {
	return c.Execute(ctx, fmt.Sprintf("ALTER AUTHORIZATION ON DATABASE::%s TO %s", db, owner))
}

type DatabaseSchemaRow struct {
	name string
}

func (c MssqlConnector) HasDatabase(ctx context.Context, db string) (bool, error) {
	var row DatabaseSchemaRow
	err := c.QueryOne(ctx, fmt.Sprintf("SELECT name FROM sys.databases where name = '%s'", db), row)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c MssqlConnector) DropDatabase(ctx context.Context, db string) error {
	return c.Execute(ctx, fmt.Sprintf("DROP DATABASE %s", db))
}
