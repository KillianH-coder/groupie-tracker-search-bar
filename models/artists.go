package models

type Artist struct {
	CreationDate         int      `json:"creationDate"`
	FirstAlbum           string   `json:"firstAlbum"`
	ID                   int      `json:"id"`
	Image                string   `json:"image"`
	Members              []string `json:"members"`
	Name                 string   `json:"name"`
	Relations            string   `json:"relations"`
	ResolvedConcertDates []string
	ResolvedLocations    []string
	ResolvedRelations    Relation
}
