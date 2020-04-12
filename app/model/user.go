package model

type User struct {
	id int32
}

func (u User) Id() int32 {
	return u.id
}

// Todo move to one Abstract Query builder
func Query() User {
	return User{id: 123}
}
