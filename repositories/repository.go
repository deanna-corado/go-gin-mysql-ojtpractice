package repositories

import (
	"database/sql"
	"go-gin-mysql/models"
)

// medyo lito pa dito
type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetMovies() ([]models.Movie, error) {
	query := "SELECT id, title, director FROM movies"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Director)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}
