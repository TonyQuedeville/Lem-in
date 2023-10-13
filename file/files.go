/*
Tony Quedeville : tquedevi
Zone01: Projet lem-in
05/01/2023
*/

package files

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	class "lemin/class"
)

/* Lecture fichier */
func ReadFile(filename string) []string {
	var lines []string
	txt, err := os.ReadFile("./examples/" + filename)

	if err != nil {
		check(err)
		os.Exit(1)
	} else {
		lines = strings.Split(string(txt), "\n")
	}

	return lines
}

/* Ecriture fichier */
func WriteFile(filename, txt string) {
	data := []byte(txt)
	err := ioutil.WriteFile(filename, data, 0o644) //(0777 pour unix)
	check(err)
}

/* Affichage des erreurs */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Json ----------------------------------------

/* Lecture fichier json */
func ReadFileJson(filename string) class.Farm {
	var ListFarm class.Farm
	path := "./static/json/"
	file, err := ioutil.ReadFile(path + filename)
	if err != nil {
		panic("Impossible de lire le fichier JSON")
	} else {
		json.Unmarshal(file, &ListFarm)
	}
	return ListFarm
}

//*/

/* Ecriture fichier json */
func WriteFileJson(filename string, data []byte) {
	path := "./static/json/"
	//fmt.Println(filename)

	err := ioutil.WriteFile(path+filename, data, 0o644)
	if err != nil {
		panic("Impossible d'écrire dans le fichier.json")
	}
}
