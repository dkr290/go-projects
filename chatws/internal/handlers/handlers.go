package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var (
	views = jet.NewSet(
		jet.NewOSFileSystemLoader("./html"),
		jet.InDevelopmentMode(),
	)

	upgradeConnection = websocket.Upgrader{

		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	wsChan  = make(chan WsJsonPayload)
	clients = make(map[WebSocketConnection]string)
)

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

type WebSocketConnection struct {
	*websocket.Conn
}

// define response from the websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WsJsonPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

// upgrades connection to websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to the endpoint")

	var resp WsJsonResponse

	resp.Message = `<em><small>Connected to the server</small></em>`

	conn := WebSocketConnection{
		Conn: ws,
	}

	clients[conn] = ""

	err = ws.WriteJSON(resp)
	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn)

}

func ListenForWs(conn *WebSocketConnection) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsJsonPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			//do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ListenToWsChannel() {
	var response WsJsonResponse
	for {

		e := <-wsChan
		response.Action = "Go here"
		response.Message = fmt.Sprintf("Some message, and action was %s", e.Action)
	}

}

func BroadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocker error")
			_ = client.Close()
			delete(clients, client)
		}

	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {

	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
