package games

type Game struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Href string `json:"href" db:"href"`
}
