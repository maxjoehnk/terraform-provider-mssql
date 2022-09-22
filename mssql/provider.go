package mssql

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/url"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username to authenticate against the mssql server with. Requires the dbcreator role to create new databases.",
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address the mssql server is reachable on.",
			},
			"port": {
				Type:     schema.TypeInt,
				Default:  1433,
				Optional: true,
				Description: `The port the mssql server is listening on.

Defaults to 1433`,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mssql_database": resourceDatabase(),
			"mssql_role":     resourceRole(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%d", d.Get("host"), d.Get("port")),
	}

	return sql.Open("sqlserver", u.String())
}
