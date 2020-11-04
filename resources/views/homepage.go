package views

type book struct {
	Title string `json:"title"`
}

func Book(title string) *book {
	return &book{
		Title: title,
	}
}
