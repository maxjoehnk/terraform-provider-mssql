---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mssql Provider"
subcategory: ""
description: |-
  
---

# mssql Provider



## Example Usage

```terraform
provider "mssql" {
  host = "localhost"
  username = "sa"
  password = "password"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) The address the mssql server is reachable on.
- `password` (String)
- `username` (String) The username to authenticate against the mssql server with. Requires the dbcreator role to create new databases.

### Optional

- `port` (Number) The port the mssql server is listening on.

Defaults to 1433
