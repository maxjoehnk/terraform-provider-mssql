# Terraform Microsoft SQL Server Provider

## Usage
```hcl
provider "mssql" {
  host = "localhost"
  username = "sa"
  password = "password"
}

resource "mssql_database" "db" {
  name = "MyDatabase"
  owner = "${mssql_role.user.name}"
}

resource "mssql_role" "user" {
  name = "MyUser"
  password = "MyPassword"
}
```