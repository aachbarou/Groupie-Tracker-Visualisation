package Handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Parssing_Error error

// Get_JSONData fetches JSON data from a specified URL and decodes it into the provided interface.
func Get_JSONData(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func SearchArtists(query string) []Artist {
	var filtered []Artist
	Exist := false
	for _, artist := range Data.Artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) {
			filtered = append(filtered, artist)
			Exist = true
		}
		if !Exist {
			for i := range artist.Members {
				if strings.Contains(strings.ToLower(artist.Members[i]), strings.ToLower(query)) {
					filtered = append(filtered, artist)
					Exist = true
					break
				}
			}
		}
	}
	if !Exist {
		for i, artistLoc := range Data.Locations {
			for _, Loc := range artistLoc.Locations {
				if strings.Contains(strings.ToLower(Loc), strings.ToLower(query)) {
					filtered = append(filtered, Data.Artists[i])
					break
				}
			}
		}
	}
	if !Exist {
		for i, artistdate := range Data.Dates {
			for _, Date := range artistdate.Dates {
				if strings.Contains(strings.ToLower(Date), strings.ToLower(query)) {
					filtered = append(filtered, Data.Artists[i])
					break
				}
			}
		}
	}
	return filtered
}

func FilterArists(query string, Start int, End int, Fstart int, Fend int) []Artist {
	var filtered []Artist
	Exist := false
	for i, artist := range Data.Artists {
		Album_year, _ := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
		if (artist.CreationDate >= Start && artist.CreationDate <= End) && (Album_year >= Fstart && Album_year <= Fend) {
			Exist = true
		}
		if Exist {
			for _, artistLoc := range Data.Locations[i].Locations {
				if strings.Contains(strings.ToLower(artistLoc), strings.ToLower(strings.ReplaceAll(query, ", ", "-"))) {
					filtered = append(filtered, Data.Artists[i])
					break
				}
			}
		}
		Exist = false
	}
	return filtered
}

func Startup() {
	for {
		if Get_JSONData(url, &Data.Artists) != nil {
			Parssing_Error = errors.New("filled to Get  Data  Json Data &url")
			return
		}
		for _, Artist := range Data.Artists {
			var (
				temp1 Locations
				temp2 Dates
			)
			if Get_JSONData(Artist.Locations, &temp1) != nil || Get_JSONData(Artist.Dates, &temp2) != nil {
				Parssing_Error = errors.New("filed To Parssing Data  (Dates  || Locations)")
				break
			}
			Data.Locations = append(Data.Locations, temp1)
			Data.Dates = append(Data.Dates, temp2)
		}
		Data.Founded = Data.Artists
		time.Sleep(1 * time.Hour)
	}

}
