package models

type User struct {
	User_ID         int `goqu:"skipinsert"`
	Username        string
	Password        string
	Email           string
	First_Name      string
	Last_Name       string
	Created_By      int
	Datetime_Create string `goqu:"skipinsert"`
	Updated_By      int
	Datetime_Update string `goqu:"skipinsert"`
}
