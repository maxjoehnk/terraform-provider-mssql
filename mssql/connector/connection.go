package connector

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
)

type MssqlConnector struct {
	ConnectionString string
}

func New(host string, port int, user string, password string) MssqlConnector {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%d", host, port),
	}
	return MssqlConnector{ConnectionString: u.String()}
}

func (c MssqlConnector) Execute(ctx context.Context, command string) error {
	conn, err := sql.Open("sqlserver", c.ConnectionString)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err = conn.ExecContext(ctx, command); err != nil {
		return err
	}
}

func (c MssqlConnector) QueryOne(ctx context.Context, query string, row interface{}) error {
	conn, err := sql.Open("sqlserver", c.ConnectionString)
	if err != nil {
		return err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	if err = rows.Scan(row); err != nil {
		return err
	}
}
