package users

type  User struct {
	Id int64
	FirstName string `json: "id"`
	LastName string `json: "first_name"`
	Email string `json: "last_name"`
	DateCreated string `json: "date_created"`
}
