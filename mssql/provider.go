package mssql

import (
	"database/sql"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"net/url"
	_ "github.com/denisenkom/go-mssqldb"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host": {
				Type: schema.TypeString,
				Required: true,
			},
			"port": {
				Type: schema.TypeInt,
				Default: 1433,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mssql_database": resourceDatabase(),
			"mssql_role": resourceRole(),
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
		ConfigureFunc: providerConfigure,
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
