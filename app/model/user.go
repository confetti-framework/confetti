package model

type User struct {
	id    int
	email string
}

func NewUser(id int, fullName string) User {
	return User{id: id, email: fullName}
}

func (u User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":    u.id,
		"email": u.Email(),
	}
}

func (u User) Id() int {
	return u.id
}

func (u User) Email() interface{} {
	return u.email
}
