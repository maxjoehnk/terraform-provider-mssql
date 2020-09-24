package mssql

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/maxjoehnk/terraform-provider-mssql/mssql/connector"
)

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDatabaseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	db := m.(connector.MssqlConnector)
	name := d.Get("name").(string)
	if err := db.CreateDatabase(ctx, name); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(name)

	owner := d.Get("owner")
	if owner != nil {
		if err := db.SetDatabaseOwner(ctx, name, owner.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	db := m.(connector.MssqlConnector)
	name := d.Id()
	hasTable, err := db.HasDatabase(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if !hasTable {
		return nil
	}
	if err = d.Set("name", name); err != nil {
		return diag.FromErr(err)
	}
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	db := m.(connector.MssqlConnector)
	name := d.Id()
	if err := db.DropDatabase(ctx, name); err != nil {
		return diag.FromErr(err)
	}
}
