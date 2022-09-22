package mssql

import (
	"database/sql"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
		Description: `Manage a role

Updating a role is currently not supported and will drop the user.
`,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The logon name of the role",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password used by the role",
			},
		},
	}
}

func resourceRoleCreate(d *schema.ResourceData, m interface{}) error {
	db := m.(*sql.DB)
	name := d.Get("name").(string)
	password := d.Get("password").(string)
	_, err := db.Query(fmt.Sprintf("CREATE LOGIN %s WITH PASSWORD = '%s', CHECK_POLICY = OFF, CHECK_EXPIRATION = OFF", name, password))
	if err != nil {
		return err
	}
	_, err = db.Query(fmt.Sprintf("CREATE USER %s FOR LOGIN %s", name, name))
	if err != nil {
		return err
	}
	row := db.QueryRow(fmt.Sprintf("SELECT principal_id FROM master.sys.server_principals WHERE name = '%s'", name))
	var id int
	if err = row.Scan(&id); err != nil {
		return err
	}

	d.SetId(fmt.Sprint(id))
	return err
}

func resourceRoleRead(d *schema.ResourceData, m interface{}) error {
	db := m.(*sql.DB)
	row := db.QueryRow(fmt.Sprintf("SELECT name FROM master.sys.server_principals WHERE principal_id = %s", d.Id()))
	var name string
	err := row.Scan(&name)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	if err := d.Set("name", name); err != nil {
		return err
	}
	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRoleDelete(d *schema.ResourceData, m interface{}) error {
	db := m.(*sql.DB)
	row := db.QueryRow(fmt.Sprintf("SELECT name FROM master.sys.server_principals WHERE principal_id = %s", d.Id()))
	var name string
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	dropUserQuery := fmt.Sprintf("DROP USER %s", name)
	print(dropUserQuery)
	_, err = db.Query(dropUserQuery)
	if err != nil {
		return err
	}
	_, err = db.Query(fmt.Sprintf("DROP LOGIN %s", name))
	return err
}
