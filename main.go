package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/maxjoehnk/terraform-provider-mssql/mssql"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mssql.Provider,
	})
}
