resource "mssql_role" "user" {
  name = "MyUser"
  password = "MyPassword"
}
