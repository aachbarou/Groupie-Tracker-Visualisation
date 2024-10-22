package Handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var (
	templates, errT       = template.ParseFiles("Templates/artists.html")
	error_template, errE  = template.ParseFiles("Templates/Error.html")
	dtailesTemplate, errD = template.ParseFiles("Templates/detailes.html")
	url                   = "https://groupietrackers.herokuapp.com/api/artists"
)

var Data All_Data

func Handle_SearchBar(w http.ResponseWriter, r *http.Request) {
	defer func() {
		panic := recover()
		if panic != nil {
			Error_handle(w, 416)
		}
	}()
	if r.Method == "GET" {
		query := r.FormValue("search")
		Searched := SearchArtists(query)
		Data.Founded = Searched
		if len(Searched) == 0 {
			templates.Execute(w, Data)
			return
		}
		templates.Execute(w, Data)
		return
	} else {
		Error_handle(w, 504)
	}
}

// Handles HTTP requests to the root URL (/).
func HAndleHOme(w http.ResponseWriter, r *http.Request) {
	Parssing := CheckParseFile(w, r, errT, errE, errD)
	if !Parssing {
		return
	}
	if r.Method != "GET" {
		Error_handle(w, 405)
		return
	}
	if r.URL.Path != "/" {
		Error_handle(w, 404)
		return
	}
	if  len(Data.Artists) == 0 || Parssing_Error != nil {
		Error_handle(w  ,  500)
		fmt.Println(Parssing_Error)
		return
	}

	Data.Founded = Data.Artists
	var buf bytes.Buffer
	err := templates.Execute(&buf, Data)
	if err != nil {
		Error_handle(w, 500)
		return
	} else {
		templates.Execute(w, Data)
		return
	}
}

// Handles HTTP requests to the URL (/detailes).
func Handle_Detailes(w http.ResponseWriter, r *http.Request) {
	Parssing := CheckParseFile(w, r, errT, errE, errD)
	if !Parssing {
		fmt.Println("Eroor  In Stylshets Files ")
	}
	Id, ER := strconv.Atoi(r.FormValue("id"))
	if ER != nil || (Id < 0 || Id > 52) {
		Error_handle(w, 404)
		return
	}
	url1 := "https://groupietrackers.herokuapp.com/api/dates/" + r.FormValue("id")
	url2 := "https://groupietrackers.herokuapp.com/api/locations/" + r.FormValue("id")
	url3 := "https://groupietrackers.herokuapp.com/api/relation/" + r.FormValue("id")
	var Data Artist_Info
	if Get_JSONData(url1, &Data.D_ates) != nil ||
		Get_JSONData(url2, &Data.L_ocations) != nil ||
		Get_JSONData(url3, &Data.R_elation) != nil {
		Error_handle(w, 404)
		return
	}
	err := dtailesTemplate.Execute(w, Data)
	if err != nil {
		Error_handle(w, 500)
		return
	}
}

func Handle_Filters(w http.ResponseWriter, r *http.Request) {
	defer func() {
		panic := recover()
		if panic != nil {
			Error_handle(w, 416)
		}
	}()
	if r.Method == "GET" {
		r.ParseForm()
		if len(r.Form) == 0 {
			templates.Execute(w, Data)
			return
		}
		Numbers := r.Form["number"]
		query := r.FormValue("queryfilter")
		Y_start, err := strconv.Atoi(r.FormValue("yearstart"))
		Y_end, erre := strconv.Atoi(r.FormValue("yearend"))
		F_start, _ := strconv.Atoi(r.Form["FralbumStart"][0])
		F_end, _ := strconv.Atoi(r.Form["FralbumEnd"][0])

		if err != nil || erre != nil {
			Error_handle(w, 400)
			return
		}
		filtered := FilterArists(query, Y_start, Y_end, F_start, F_end)
		if len(filtered) == 0 {
			Data.Founded = nil
			templates.Execute(w, Data)
			return
		}
		if len(Numbers) != 0 {
			Data.Founded = nil
			for _, artist := range filtered {
				for _, num := range Numbers {
					num2, _ := strconv.Atoi(num)
					if len(artist.Members) == num2 {
						Data.Founded = append(Data.Founded, artist)
						break
					}
				}
			}
		} else {
			Data.Founded = filtered
		}

		templates.Execute(w, Data)
		return
	} else {
		Error_handle(w, 400)
		return
	}
}
