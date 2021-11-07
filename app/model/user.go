package model

type User struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Find(u *User) []User {
	var users []User
	db.Where(u).Find(&users)

	return users
}

func Select(u *User, id string) User {
	var user User
	db.First(&user, id)

	return user
}
