package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/pusher/pusher-http-go"
)

// first thing is to authenticate
func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {

	// is this user authenticated

	// get authentication info from the session

	userID := repo.App.Session.GetInt(r.Context(), "userID")
	u, _ := repo.DB.GetUserById(userID)
	params, _ := ioutil.ReadAll(r.Body)

	presenceData := pusher.MemberData{
		UserID: strconv.Itoa(userID),
		UserInfo: map[string]string{
			"name": u.FirstName,
			"id":   strconv.Itoa(userID),
		},
	}

	response, err := app.WsClient.AuthenticatePresenceChannel(params, presenceData)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)

}
