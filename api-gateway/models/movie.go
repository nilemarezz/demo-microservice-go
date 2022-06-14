package models

// Movie model info
// @Description Movie information
// @Description with id, name, description, screen_date and cast
type Movie struct {
	Id          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"movie name"`
	Description string `json:"description" example:"movie description"`
	ScreenDate  string `json:"screen_date" example:"01-01-1999"`
	Cast        []Cast `json:"cast"`
}

type Cast struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
