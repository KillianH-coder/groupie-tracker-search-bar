package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
	"strings"
	"time"
)

func httpClient(url string) (string, int, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", 0, fmt.Errorf("create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", res.StatusCode, fmt.Errorf("read response body: %w", err)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return string(body), res.StatusCode, fmt.Errorf("request failed with status %s", res.Status)
	}

	return string(body), res.StatusCode, nil
}

func GetArtists() []models.Artist {
	body, _, err := httpClient("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return []models.Artist{}
	}

	var artists []models.Artist
	err = json.Unmarshal([]byte(body), &artists)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return []models.Artist{}
	}
	return artists
}

func GetLocations() []models.Location {
	body, _, err := httpClient("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return []models.Location{}
	}

	var locationResp models.LocationsResponse

	err = json.Unmarshal([]byte(body), &locationResp)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return []models.Location{}
	}

	return locationResp.Index
}

func GetArtistRelationsByID(id string) (models.Relation, int) {
	body, _, err := httpClient("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return models.Relation{}, 500
	}

	var relation models.Relation
	err = json.Unmarshal([]byte(body), &relation)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return models.Relation{}, 500
	}
	return relation, 200
}

func GetArtistByID(id string) (models.Artist, int) {
	body, _, err := httpClient("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return models.Artist{}, 500
	}

	var artist models.Artist
	err = json.Unmarshal([]byte(body), &artist)

	if artist.ID == 0 {
		return models.Artist{}, 404
	}

	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return models.Artist{}, 500
	}

	artistRelation, status := GetArtistRelationsByID(id)
	switch status {
	case 200:
		artist.ResolvedRelations = artistRelation
		artist.ResolvedLocations = make([]string, 0, len(artistRelation.DatesLocations))
		artist.ResolvedConcertDates = make([]string, 0, len(artistRelation.DatesLocations))
		newMap := make(map[string][]string)
		for location, dates := range artistRelation.DatesLocations {
			if location == "" || len(dates) == 0 {
				fmt.Println("Invalid location or dates:", location, dates)
				continue
			}

			originalName := location
			splittedLocation := strings.Split(location, "-")
			cityName := splittedLocation[0]
			splittedCity := strings.Split(cityName, "_")

			for i, word := range splittedCity {
				splittedCity[i] = strings.Title(word)
			}

			cityName = strings.Join(splittedCity, " ")
			splittedLocation[0] = cityName
			splittedLocation[1] = strings.ToUpper(splittedLocation[1])
			splittedLocation[1] = strings.Replace(splittedLocation[1], "_", " ", -1)
			location = strings.Join(splittedLocation, ", ")

			if artist.ResolvedRelations.DatesLocations == nil {
				artist.ResolvedRelations.DatesLocations = make(map[string][]string)
			}
			//artist.ResolvedRelations.DatesLocations[location] = dates
			newMap[location] = dates
			delete(artist.ResolvedRelations.DatesLocations, originalName)
		}
		artist.ResolvedRelations.DatesLocations = newMap
	case 404:
		return models.Artist{}, 404
	default:
		return models.Artist{}, 500
	}
	fmt.Println(artist)
	return artist, 200
}
