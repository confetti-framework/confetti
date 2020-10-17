package model

type User struct {
	id    int
	email string
}

func (u User) Find() User {
	return User{id: 1, email: "fake@test.nl"}
}

func NewUser(id int, email string) User {
	return User{id: id, email: email}
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
