package entity

type User struct {
    Id          string       `json:"id"`
    UserName    string       `json:"username"`
    Name        string       `json:"name"`
    Picture     string       `json:"picture_url"`
    Permissions []Permission `json:"permissions,omitempty"`
}
