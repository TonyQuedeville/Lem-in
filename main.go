/*
	Tony Quedeville : tquedevi
	03/01/2023
	Zone01 : Projet Lem-in
	Github : https://zone01normandie.org/git/tquedevi/lem-in
	lien utile : https://www.youtube.com/watch?v=aD20O2oQ1DQ
--------------------------------------------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	class "lemin/class"
	file "lemin/file"
	server "lemin/server"
)

/*-------------------------------- Main -------------------------------------*/

// recupere les arguments, gere les erreurs
func main() {

	// Arguments ------------------------------------------------------------
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Veuillez specifier le fichier example.txt SVP !")
		os.Exit(0)
	}

	nomFichier := args[1]

	// Initialisation des structures ----------------------------------------
	datas := file.ReadFile(nomFichier)

	// Nombre de fourmi
	nbAnt, err := strconv.Atoi(datas[0])
	if err != nil || nbAnt == 0 {
		fmt.Println("Erreur nombre de fourmi : ", datas[0])
		os.Exit(0)
	}
	fmt.Println("nbAnt:", nbAnt)

	// Données de la fourmilliaire
	datas = datas[1:]
	farm := class.NewFarm(nomFichier, nbAnt, datas) // Instance de la fourmiliaire
	farmJson := class.NewFarmJson(nomFichier, nbAnt, datas) // Instance de la fourmiliaire pour la visualisation

	// Recherche de tous les chemins
	roomStart, err := farm.GetRoomByType("start")
	if err != nil {
		fmt.Println("Erreur Room Start !")
	}
	//fmt.Println(roomStart.Name)

	roomEnd, err := farm.GetRoomByType("end")
	if err != nil {
		fmt.Println("Erreur Room End !")
	}
	//fmt.Println(roomEnd.Name)

	path := class.NewPath("", []*class.Room{roomStart})
	farm.FindAllPaths(roomStart.Name, roomEnd.Name, &path)
	if len(farm.Paths) == 0 {
		fmt.Println("Erreur ! Not Path possible.")
		return
	}
	farm.DisplayPath()

	// Tri des Sets de chemins qui ne se croise pas
	sets := farm.FindAllPathSets()
	class.DisplaySets(sets)

	// Choix du Set
	farm.ChoiceSet(sets)	
	farm.DisplayPath()	

	// Deplacement des fourmis
	listStepJson := farm.MoveAnts()

	// Visualisation ------------------------------------------------------------------

	// Ecriture du fichier farm.json
	farmJson.Paths = farmJson.EditPathJson(farm)
	farmJson.NbStep = farm.NbStep
	farmJson.StepsByPath = farm.StepsByPath
	data, err := json.Marshal(farmJson)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteFileJson("farm.json", data)

	// Ecriture du fichier step.json
	dataStep, errStep := json.Marshal(listStepJson)
	if errStep != nil {
		fmt.Println(err)
	}
	file.WriteFileJson("step.json", dataStep)

	// Lancement server
	server.StartAppWeb()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
