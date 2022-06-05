package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const posterURL = "http://www.omdbapi.com/?"
const apiKey = "&apikey=57013694"

type Movie struct {
	Title  string
	Poster string
}

func main() {

	movieUrl := getMovieTitleUrl(os.Args[1:])
	movie, err := getMovieDataFromUrl(movieUrl)
	if err != nil {
		log.Fatal(err)
	}
        posterUrlLen := len(movie.Poster)
        if !(posterUrlLen >= 4) {
             log.SetPrefix("Error: ")
             log.SetFlags(0)
             log.Fatal("Movie Not Found!")
        }

        fmt.Printf("Success Movie Found!\n")
	fmt.Printf("Movie Title -> %s\nMovie Poster URL -> \n%s\n", 
                    movie.Title,
                    movie.Poster)

	imgExtension := movie.Poster[posterUrlLen-4:]
	filename := strings.Join(strings.Split(movie.Title, " "), "_") + imgExtension

	err = downloadFile(movie.Poster, filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s downloaded to current directory\n", filename)

}

func getMovieTitleUrl(terms []string) string {
	q := url.QueryEscape(strings.Join(terms, " "))
	url := posterURL + "t=" + q + apiKey
	return fmt.Sprintf("%s", url)
}

func getMovieDataFromUrl(url string) (*Movie, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Movie

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()

	return &result, nil
}

func downloadFile(posterURL, posterFileName string) error {
	//Get the response bytes from the url
	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Received Not OK Status Code")
	}
	//Create a empty file
	file, err := os.Create(posterFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
