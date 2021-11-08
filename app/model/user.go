package model

type User struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Find(u *User) ([]User, error) {
	var users []User
	err := db.Where(u).Find(&users).Error

	return users, err
}

func Select(u *User, id string) (User, error) {
	var user User
	err := db.First(&user, id).Error

	return user, err
}
