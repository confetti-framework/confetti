package model

type User struct {
	id       int
	fullName string
}

func NewUser(id int, fullName string) User {
	return User{id: id, fullName: fullName}
}

func (u User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":        u.id,
		"full_name": u.FullName(),
	}
}

func (u User) Id() int {
	return u.id
}

func (u User) FullName() interface{} {
	return u.fullName
}
