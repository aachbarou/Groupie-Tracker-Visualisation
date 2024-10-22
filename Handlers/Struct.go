package Handler

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Dates        string   `json:"concertDates"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Artist_Info struct {
	L_ocations Locations
	D_ates     Dates
	R_elation  Relation
}
type Relation struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}
type All_Data struct {
	Artists   []Artist
	Locations []Locations
	Dates     []Dates
	Founded   []Artist
}
