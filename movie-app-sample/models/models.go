package models

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastNAme  string `json:"lastname"`
}

func CreateMovies() []Movie {

	var movies []Movie
	movies = append(movies, Movie{ID: "1", Isbn: "333489", Title: "Movie id one", Director: &Director{FirstName: "John", LastNAme: "Glom"}})
	movies = append(movies, Movie{ID: "2", Isbn: "335890", Title: "Movie id two", Director: &Director{FirstName: "Kevin", LastNAme: "Blond"}})
	return movies
}
