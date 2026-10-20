package handlers

import (
	"groupie-tracker/models"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type indexPageData struct {
	Artists     []models.Artist
	Suggestions []models.Artist
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	artists := GetArtists()
	var result []models.Artist

	locations := GetLocations()
	for i := range artists {
		if artists[i].ID-1 < len(locations) {
			artists[i].ResolvedLocations = locations[artists[i].ID-1].Locations
		}
	}

	query := r.URL.Query()

	value, exists := query["search"]
	if exists {
		for _, artist := range artists {
			if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(value[0])) {
				result = append(result, artist)
				continue
			}

			if strings.ToLower(value[0]) == strconv.Itoa(artist.CreationDate) {
				result = append(result, artist)
				continue
			}

			if strings.ToLower(value[0]) == artist.FirstAlbum {
				result = append(result, artist)
				continue
			}

			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), strings.ToLower(value[0])) {
					result = append(result, artist)
					continue
				}
			}

			for _, location := range artist.ResolvedLocations {
				var isAlreadyHere bool
				for _, elem := range result {
					if elem.ID == artist.ID {
						isAlreadyHere = true
					}
				}
				if strings.Contains(strings.ToLower(location), strings.ToLower(value[0])) && !isAlreadyHere {
					result = append(result, artist)
					continue
				}
			}
		}
	} else {
		result = artists
	}

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, indexPageData{
		Artists:     result,
		Suggestions: artists,
	})
}
