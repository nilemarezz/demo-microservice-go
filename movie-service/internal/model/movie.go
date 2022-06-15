package model

type Movie struct {
	Id          int    `db:"movie_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ScreenDate  string `db:"screen_date"`
	Cast        []Cast `db:"cast"`
}

func (m *Movie) SetCast(cast []Cast) {
	m.Cast = cast
}

type Cast struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
