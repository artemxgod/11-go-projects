package model

type Movie struct {
	ID       string    `json:"id"`
	Isnb     string    `json:"isnb"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Movies []Movie