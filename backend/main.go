package main

import (
	"bytes"
	// "database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	// "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
type Token42 struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	CreatedAt        int    `json:"created_at"`
	SecretValidUntil int    `json:"secret_valid_until"`
	ExpiresDate      time.Time
}

var token = Token42{}

// func SetGin() {
// 	// Initialisez le routeur Gin
// 	r := gin.Default()

// 	// Endpoint pour récupérer le contenu du fichier depuis la base de données
// 	r.GET("/api/fichier", func(c *gin.Context) {
// 		// Connexion à la base de données PostgreSQL
// 		connStr := "user=postgres dbname=postgres password=admin sslmode=disable"
// 		db, err := sql.Open("postgres", connStr)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la connexion à la base de données"})
// 			return
// 		}
// 		defer db.Close()

// 		// Récupération du contenu du fichier depuis la base de données
// 		var contenu string
// 		err = db.QueryRow("SELECT contenu FROM fichiers WHERE id = $1", 1).Scan(&contenu)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du fichier depuis la base de données"})
// 			return
// 		}

// 		// Réponse JSON contenant le contenu du fichier
// 		c.JSON(http.StatusOK, gin.H{"contenu": contenu})
// 	})

// 	// Démarrez le serveur
// 	err := r.Run(":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func SetDatabase() {
// 	connStr := "user=postgres dbname=postgres password=admin sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Lecture du contenu du fichier texte
// 	content, err := os.ReadFile("prez.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Insertion du contenu dans la base de données
// 	_, err = db.Exec("INSERT INTO fichiers (contenu) VALUES ($1)", content)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Fichier enregistré dans la base de données.")
// }

func (t *Token42) CheckToken() bool {
	if t.AccessToken == "" {
		return false
	}
	if t.ExpiresDate.Before(time.Now()) {
		return false
	}
	return true
}

func (t *Token42) GrabToken() error {
	url := "https://api.intra.42.fr/oauth/token?grant_type=client_credentials"
	requestUrl := fmt.Sprintf(
		"%s&client_id=%s&client_secret=%s",
		url,
		os.Getenv("BACKEND_42_UID"),
		os.Getenv("BACKEND_42_SECRET"),
	)

	resp, err := http.Post(requestUrl, "", bytes.NewBuffer([]byte("")))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, t)
	if err != nil {
		return err
	}
	t.ExpiresDate = time.Now().Add(time.Second * time.Duration(t.ExpiresIn))
	return nil
}

func (t *Token42) RefreshToken() error {
	if t.CheckToken() {
		return nil
	}
	err := token.GrabToken()
	return err
}

// arr := make([]string, 0)
// arr = append(arr, element)
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

func meHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := token.RefreshToken()
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
	meReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

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

func main() {
	// SetDatabase()
	// SetGin()
	token.AccessToken = ""
	http.HandleFunc("/me", meHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
