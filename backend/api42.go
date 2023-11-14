package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserData42 struct {
	Login     string `json:"login"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     struct {
		Link     string `json:"link"`
		Versions struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Small  string `json:"small"`
			Micro  string `json:"micro"`
		} `json:"versions"`
	} `json:"image"`
	CursusUsers []struct {
		ID     int         `json:"id"`
		Grade  interface{} `json:"grade"`
		Level  float64     `json:"level"`
		Cursus struct {
			ID int `json:"id"`
		} `json:"cursus"`
	} `json:"cursus_users"`
	ProjectsUsers []struct {
		FinalMark int  `json:"final_mark"`
		Validated bool `json:"validated?"`
		Project   struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"project"`
	} `json:"projects_users"`
	Titles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"titles"`
	TitlesUsers []struct {
		ID       int  `json:"id"`
		UserID   int  `json:"user_id"`
		TitleID  int  `json:"title_id"`
		Selected bool `json:"selected"`
	} `json:"titles_users"`
}

type Project struct {
	Name      string `json:"name"`
	Mark      int    `json:"mark"`
	Validated bool   `json:"validated?"`
}

type Title struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
}

type UserDataParsed struct {
	User struct {
		Login     string `json:"login"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Grade     string
		Level     float64
		Image     struct {
			Link     string `json:"link"`
			Versions struct {
				Large  string `json:"large"`
				Medium string `json:"medium"`
				Small  string `json:"small"`
				Micro  string `json:"micro"`
			} `json:"versions"`
		} `json:"image"`
	}
	Projects []Project `json:"projects"`
	Titles   []Title   `json:"titles"`
}

func (u *UserData42) Parse() UserDataParsed {
	ret := UserDataParsed{}
	ret.User.Login = u.Login
	ret.User.FirstName = u.FirstName
	ret.User.LastName = u.LastName
	for _, cursusUser := range u.CursusUsers {
		if cursusUser.Cursus.ID == 21 {
			if cursusUser.Grade != nil {
				ret.User.Grade = cursusUser.Grade.(string)
			}
			ret.User.Level = cursusUser.Level
		}
	}
	ret.User.Image = u.Image
	ret.Projects = make([]Project, 0)
	for _, projectUsers := range u.ProjectsUsers {
		project := Project{
			Name:      projectUsers.Project.Name,
			Mark:      projectUsers.FinalMark,
			Validated: projectUsers.Validated,
		}
		ret.Projects = append(ret.Projects, project)
	}
	selectedId := 0
	for _, titleUser := range u.TitlesUsers {
		if titleUser.Selected == true {
			selectedId = titleUser.TitleID
			break
		}
	}
	ret.Titles = make([]Title, 0)
	for _, titleRaw := range u.Titles {
		title := Title{
			ID:       titleRaw.ID,
			Name:     titleRaw.Name,
			Selected: selectedId == titleRaw.ID,
		}
		ret.Titles = append(ret.Titles, title)
	}
	return ret
}

func meHandler42(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := client42.RefreshToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	url := "https://api.intra.42.fr/v2/users/tefroiss"

	// Créer une nouvelle requête GET
	meReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Paramètres (le cas échéant)
	// Vous pouvez les ajouter à la requête en utilisant req.URL.Query().Add("parametre1", "valeur1")

	// En-têtes (le cas échéant)
	meReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client42.Token.AccessToken))

	// Effectuer la requête HTTP
	client := &http.Client{}
	resp, err := client.Do(meReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Vérifier le code de statut de la réponse
	if resp.StatusCode >= 400 {
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return
	}
	myData := UserData42{}
	err = json.Unmarshal(body, &myData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	myDataBytes, err := json.Marshal(myData.Parse())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(myDataBytes)
}
