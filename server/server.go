/*
	Tony Quedeville : tquedevi
	03/01/2023
	Zone01 : Projet Lem-in
	Github : https://zone01normandie.org/git/tquedevi/lem-in
	Tuto package websocket : https://tutorialedge.net/golang/go-websocket-tutorial/
*/

package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	class "lemin/class"
	//file "lemin/file"

	"github.com/gorilla/websocket"
)

var Datas class.Data
var Listfarm class.Farm

/*----------------------------------------------------------*/

// Configuration de la connexion
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Connexion au client. (Point de terminaison connexion TCP)
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connecté ! ")

	err = ws.WriteMessage(1, []byte(Listfarm.Name))
	if err != nil {
		log.Println(err)
	}

	reader(ws) // Pas besoin car le client ne fait pas de requete sur application
}

// Ecoute permanante de l'url // Pas besoin car le client ne fait pas de requete sur application
func reader(conn *websocket.Conn) {
	defer conn.Close()
	for {
		// Recuperation des messages et des données qu'il contient.
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erreur conn.ReadMessage() ! ", err)
			return
		}
		json.Unmarshal(data, &Datas)
		fmt.Println(Datas.Typedata)

		// Type de données : farm
		if Datas.Typedata == "farm" {
			json.Unmarshal(data, &Listfarm)
			fmt.Println("Listfarm : ", Listfarm)
		}

		// Confirmation de reception des données vers le client
		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Println("Erreur conn.WriteMessage() ! ", err)
			return
		}
	}
}
//*/

/*----------------------------------------------------------*/

/* Lancement du serveur */
func StartAppWeb() {
	port := ":8080" // 67 test erreur

	f := http.FileServer(http.Dir("static"))
	s := http.StripPrefix("/static/", f)
	http.Handle("/static/", s)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsEndpoint)

	fmt.Print("Start server : http://localhost", port, "/\n")
	log.Fatal(http.ListenAndServe(port, nil))
}

// Page d'acceuil
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Erreur 404
	if r.URL.Path != "/" {
		t, _ := template.ParseFiles("./template/error404.html")
		t.Execute(w, nil)
		return
	}

	// Ecriture du fichier farm.json

	index, err := template.ParseFiles("./template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // erreur 500
	} else {
		index.Execute(w, nil)
	}
}

//*/
