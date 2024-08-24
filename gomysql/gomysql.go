package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func GetAlbums() ([]Album, error) {
	fmt.Println("Hello mysql")

	db, err := sql.Open("mysql", "root:strong_password@tcp(db-mysql:3306)/yourdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	}
	log.Println("Hello MySql")
	// defer the close till after the main function has finished
	// executing

	albums, err := albumsAll(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	defer db.Close()
	return albums, nil
}

func InsertAlbums(albums []Album) error {
	fmt.Println("Hello mysql")

	db, err := sql.Open("mysql", "root:strong_password@tcp(db-mysql:3306)/yourdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	}
	log.Println("Hello MySql")
	// defer the close till after the main function has finished
	// executing

	err2 := albumsInsert(db, albums)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Albums found: %v\n", albums)

	defer db.Close()
	return nil

}

func albumsByArtist(name string, db *sql.DB) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("Select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf(`albumsByArtist %q: %v`, name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil

}

func albumsAll(db *sql.DB) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("Select * from album")
	if err != nil {
		return nil, fmt.Errorf(`albumsByArtist %v`, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", err)
	}
	return albums, nil

}

func albumsInsert(db *sql.DB, albums []Album) error {

	query := "INSERT INTO `album` (`Title`, `Artist`, `Price`) VALUES (?, ?, ?)"
	for _, album := range albums {
		insertResult, err := db.ExecContext(context.Background(), query, album.Title, album.Artist, album.Price)
		log.Println(insertResult)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
