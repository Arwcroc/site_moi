package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TwitchStreamData struct {
	UserID       string    `json:"user_id"`
	UserLogin    string    `json:"user_login"`
	UserName     string    `json:"user_name"`
	GameName     string    `json:"game_name"`
	Type         string    `json:"type"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
}

type TwitchUserData struct {
	ProfileImageURL string    `json:"profile_image_url"`
}

type TwitchUserResponse struct {
	Data []TwitchUserData `json:"data"`
}

type TwitchStreamResponse struct {
	Data []TwitchStreamData `json:"data"`
}

type TwitchReqResponse struct {
	User TwitchUserData `json:"user"`
	Stream TwitchStreamData `json:"stream"`
	HasStream bool `json:"has_stream"`
}

func doTwitchRequest(w http.ResponseWriter, url string) ([]byte, error) {
	// Créer une nouvelle requête GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil, err
	}

	// En-têtes (le cas échéant)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", clientTwitch.Token.AccessToken))
	req.Header.Add("Client-Id", fmt.Sprintf("%s", clientTwitch.Uid))

	// Effectuer la requête HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	// Lire la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil, err
	}

	// Vérifier le code de statut de la réponse
	if resp.StatusCode >= 400 {
		return body, errors.New(fmt.Sprintf("Error on request: %d", resp.StatusCode))
	}

	return body, nil
}

func twitchHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	broadcaster := req.URL.Query().Get("user_id")
	if broadcaster == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := clientTwitch.RefreshToken()
	if err != nil {
		fmt.Println("DEBUG :: on token fetch")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	baseUrl := "https://api.twitch.tv/helix"
	streamDataUrl := fmt.Sprintf("%s/streams?user_login=%s", baseUrl, broadcaster)
	streamDataBody, err := doTwitchRequest(w, streamDataUrl)
	if err != nil {
		fmt.Println("DEBUG :: on streams fetch")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	streamData := TwitchStreamResponse{}
	err = json.Unmarshal(streamDataBody, &streamData)
	if err != nil {
		fmt.Println("DEBUG :: on streams unmarshal")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userDataUrl := fmt.Sprintf("%s/users?login=%s", baseUrl, broadcaster)
	userDataBody, err := doTwitchRequest(w, userDataUrl)
	if err != nil {
		fmt.Println("DEBUG :: on users fetch")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userData := TwitchUserResponse{}
	err = json.Unmarshal(userDataBody, &userData)
	if err != nil {
		fmt.Println("DEBUG :: on users unmarshal")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(userData.Data) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hasStream := true
	if len(streamData.Data) <= 0 {
		streamData.Data = append(streamData.Data, TwitchStreamData{})
		hasStream = false
	}
	responseData := TwitchReqResponse{
		User: userData.Data[0],
		Stream: streamData.Data[0],
		HasStream: hasStream,
	}

	responseBytes, err := json.Marshal(responseData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(responseBytes)
}
