/*
Tony Quedeville : tquedevi
Zone01: Projet lem-in
03/01/2023
*/

package class

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Data ----------------------------------------------------------------
type Data struct {
	Typedata string
}

// Fourmi --------------------------------------------------------------
type Ant struct {
	Name     string // Id Fourmis (Lx)
	IdRoom   int    // position dans le chemin
	IdPath   int    // Index du chemin
	NameRoom string // Chambre qu'elle occupe
}

// Constructeur Ant
func NewAnt(name string) Ant {
	return Ant{
		Name:   name,
		IdRoom: 0,
		IdPath: -1,
	}
}

// Etape --------------------------------------------------------------
type Step struct {
	Id    int
	Ants []Ant
}

// Chemin --------------------------------------------------------------
type PathJson struct {
	Name  string
	Rooms []RoomJson // Ensemble des chambres
}

// Constructeur PathJson
func NewPathJson(name string, rooms []RoomJson) PathJson {
	return PathJson{
		Name:  name,
		Rooms: rooms,
	}
}

/* Edite le Set de chemins optimisés pour json */
func (f FarmJson) EditPathJson(farm Farm) []PathJson {
	for i, path := range farm.Paths {
		f.Paths = append(f.Paths, PathJson{})
		for _, room := range path.Rooms {
			//fmt.Println("room:", room.Name)
			rm, err := f.GetRoomByNameJson(room.Name)
			if err != nil {
				fmt.Println(err)
			} else {
				f.Paths[i].Name = "path_" + strconv.Itoa(i)
				f.Paths[i].Rooms = append(f.Paths[i].Rooms, rm)
			}
		}
		//fmt.Println()
	}
	return f.Paths
}

type Path struct {
	Name  string
	Rooms []*Room // Ensemble des chambres
}

// Constructeur Path
func NewPath(name string, rooms []*Room) Path {
	return Path{
		Name:  name,
		Rooms: rooms,
	}
}

func (f Farm) DisplayPath() {
	fmt.Println("nb paths: ", len(f.Paths))
	for i, path := range f.Paths {
		fmt.Print("  paths ", i, ": ")
		for _, room := range path.Rooms {
			fmt.Print(room.Name, " ")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------------------------------------------")
}

func DisplaySet(set []Path) {
	for j, path := range set {
		fmt.Print("    path ", j, ": ")
		for _, room := range path.Rooms {
			fmt.Print(room.Name, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func DisplaySets(sets [][]Path) {
	fmt.Println("nb sets:", len(sets))
	for i, set := range sets {
		fmt.Println("  set", i, ":")
		DisplaySet(set)
	}
	fmt.Println("--------------------------------------------------------------------")
}

// Chambre ----------------------------------------------------------------
type RoomJson struct {
	Name     string // nom de la chambre
	TypeRoom string // type de chambre (start, end ou chamber)
	X        int    // Coordonnée X
	Y        int    // Coordonnée Y
	NbAnt    int    // Nombre de fourmis
	Ants     []Ant  // Tableau de fourmis
}

// Constructeur RoomJson
func NewRoomJson(name, typeRoom string, x, y int) RoomJson {
	return RoomJson{
		Name:     name,
		TypeRoom: typeRoom,
		X:        x,
		Y:        y,
	}
}

type Room struct {
	Name     string  // nom de la chambre
	TypeRoom string  // type de chambre (start, end ou chamber)
	RoomLink []*Room // liste des chambres attenantes
	X        int     // Coordonnée X
	Y        int     // Coordonnée Y
	NbAnt    int     // Nombre de fourmis
	Ants     []Ant   // Tableau de fourmis
}

// Constructeur Room
func NewRoom(name, typeRoom string) *Room {
	return &Room{
		Name:     name,
		TypeRoom: typeRoom,
	}
}

// Ajout d'une liaison à la chambre
func (r *Room) AddLinkRoom(link *Room) {
	r.RoomLink = append(r.RoomLink, link)
}

// Colony --------------------------------------------------------------
type FarmJson struct {
	Name       string
	NbAnt      int        // Nombre de fourmis au départ
	Rooms      []RoomJson // Ensemble des chambres
	DestRooms  [][]string // liste des liens destination
	Paths      []PathJson // Liste de chemins optimisés
	NbStep     int        // Nombre d'étapes
	StepsByPath []int      // Nombre de fourmis par chemin optimisé
	Ants       []Ant      // Liste de fourmis
}

// Constructeur Farm
func NewFarmJson(name string, nbAnt int, datas []string) FarmJson {
	farm := FarmJson{}

	farm.Name = name

	// Fourmis
	farm.NbAnt = nbAnt
	var listAnts []Ant
	for i := 0; i < nbAnt; i++ {
		listAnts = append(listAnts, NewAnt("L"+strconv.Itoa(i+1)))
	}
	farm.Ants = listAnts

	typeRoom := "chamber"
	for _, data := range datas {
		if data == "##start" || data == "##end" {
			if data == "##start" {
				typeRoom = "start"
			}
			if data == "##end" {
				typeRoom = "end"
			}
		} else {
			dt := strings.Fields(string(data)) // Split avec ponctuation

			// Coordonnée x,y
			if len(dt) == 3 { // donnée de room (nom, coordonnées)
				x, _ := strconv.Atoi(dt[1])
				y, _ := strconv.Atoi(dt[2])

				farm.Rooms = append(farm.Rooms, NewRoomJson(dt[0], typeRoom, x, y))
				typeRoom = "chamber"
			} else { // données de liaison
				link := strings.Split(string(data), "-")
				if len(link) == 2 {
					farm.DestRooms = append(farm.DestRooms, link)
				}
			}
		}
	}
	//*/

	return farm
}

// Recuperation de la chambre par son nom
func (f *FarmJson) GetRoomByNameJson(name string) (RoomJson, error) {
	for i, room := range f.Rooms {
		if room.Name == name {
			return f.Rooms[i], nil
		}
	}

	return RoomJson{}, errors.New("no room found")
}

type Farm struct {
	Name       string
	NbAnt      int     // Nombre de fourmi
	Rooms      []*Room // Ensemble des chambres
	Paths      []Path  // Liste de chemins
	NbStep     int     // Nombre d'étapes
	StepsByPath []int   // Nombre de fourmis par chemin optimisé
	Ants       []Ant   // Liste de fourmis
}

// Recuperation de la chambre par son nom
func (f *Farm) GetRoomByName(name string) (*Room, error) {
	for i, room := range f.Rooms {
		if room.Name == name {
			return f.Rooms[i], nil
		}
	}

	return &Room{}, errors.New("no room found")
}

// Recuperation de la chambre par son type
func (f *Farm) GetRoomByType(typeRoom string) (*Room, error) {
	for i, room := range f.Rooms {
		if room.TypeRoom == typeRoom {
			return f.Rooms[i], nil
		}
	}

	return &Room{}, errors.New("no room found")
}

// Constructeur Farm
func NewFarm(name string, nbAnt int, datas []string) Farm {
	farm := Farm{}
	farm.Name = name

	// Fourmis
	farm.NbAnt = nbAnt
	var listAnts []Ant
	for i := 0; i < nbAnt; i++ {
		listAnts = append(listAnts, NewAnt("L"+strconv.Itoa(i+1)))
	}
	farm.Ants = listAnts

	// Rooms
	typeRoom := "chamber"
	for _, data := range datas {
		//fmt.Println(data)

		if data == "##start" || data == "##end" {
			if data == "##start" {
				typeRoom = "start"
			}
			if data == "##end" {
				typeRoom = "end"
			}
		} else {
			dt := strings.Fields(string(data)) // Fields = Split avec toute ponctuation

			if len(dt) == 3 { // donnée de room
				farm.Rooms = append(farm.Rooms, NewRoom(dt[0], typeRoom))
				typeRoom = "chamber"
			} else { // données de liaison
				link := strings.Split(string(data), "-")
				if len(link) == 2 {
					room1, err1 := farm.GetRoomByName(link[0])
					if err1 != nil {
						fmt.Println(err1)
					}
					room2, err2 := farm.GetRoomByName(link[1])
					if err2 != nil {
						fmt.Println(err2)
					}
					room1.AddLinkRoom(room2)
					room2.AddLinkRoom(room1)
				}
			}
		}
	}
	//*/

	return farm
}

// Recherche de tous les chemins
func (f *Farm) FindAllPaths(iRoom, end string, path *Path) {
	if iRoom == end { // Si on arrive à la fin du chemin
		temp := make([]*Room, len(path.Rooms))              // tableau temporaire pour créer une nouvelle instance
		copy(temp, path.Rooms)                              // Met le contenue path.Room dans temp
		f.Paths = append(f.Paths, NewPath(path.Name, temp)) //
		return
	}

	room, err := f.GetRoomByName(iRoom)
	if err == nil {
		for _, link := range room.RoomLink {
			if !contains(path.Rooms, link) {
				path.Rooms = append(path.Rooms, link)
				f.FindAllPaths(link.Name, end, path)
				path.Rooms = path.Rooms[:len(path.Rooms)-1]
			}
		}
	} else {
		fmt.Println("ERROR_____________", err)
	}

	// si aucun lien n'abouti au End
	//return
}

// Recherche de tous les Sets de chemin qui ne se croisent pas
func (f Farm) FindAllPathSets() [][]Path {
	var sets [][]Path

	for _, pathRef := range f.Paths {
		var set []Path
		set = append(set, pathRef)

		for _, path := range f.Paths {
			if !sliceEqual(pathRef.Rooms, path.Rooms) {
				found := false
				for _, subRef := range set {
					for _, room := range path.Rooms[1 : len(path.Rooms)-1] {
						found = contains(subRef.Rooms, room)
						if found {
							break
						}
					}
					if found {
						break
					}
				}
				if !found {
					set = append(set, path)
				}
			}
		}

		sets = append(sets, set)
	}

	return sets
}

// Choix du set
func (f *Farm) ChoiceSet(sets [][]Path) {
	var min int
	var iSet int
	var listStep []int
	var StepsByPath [][]int

	for iset, set := range sets {
		var mini int
		var maxi int
		var miniPath int
		var maxiPath int
		var maxiPaths []int

		// nb etapes par chemin
		var nbEtapes []int
		for i, path := range set {
			nb := len(path.Rooms) - 1
			nbEtapes = append(nbEtapes, nb)
			if nb <= mini {
				mini = nb
				miniPath = i
			}
			if nb > maxi {
				maxi = nb
				//maxiPath = i
			}
		}
		fmt.Println("nbEtapes: ", nbEtapes)

		// placement des fourmis
		for i := 0; i < f.NbAnt; i++ {
			nbEtapes[miniPath]++

			for j, nb := range nbEtapes {
				if nb <= nbEtapes[miniPath] {
					miniPath = j
				}
				if nb > maxi {
					maxi = nb
					maxiPaths = append(maxiPaths, j)
				}
			}
		}

		fmt.Println("nbEtapes: ", nbEtapes)

		// Recherche du mini des maxis
		maxiPath = 0
		for _, maxi := range maxiPaths {
			if maxi > maxiPath {
				maxiPath = maxi
			}
		}

		if iset > 0 {
			if maxiPath > min {
				min = maxiPath
				iSet = iset
			}
		} else {
			min = maxiPath
			f.NbStep = nbEtapes[maxiPath]
		}

		fmt.Println("maxiPath: ", maxiPath)
		fmt.Println("-------------")
		listStep = append(listStep, nbEtapes[maxiPath])
		StepsByPath = append(StepsByPath, nbEtapes)
	}

	//iSet++
	fmt.Println("iSet: ", iSet)
	f.Paths = sets[iSet]
	f.NbStep = listStep[iSet] - 1
	f.StepsByPath = StepsByPath[iSet]
}

/* Déplacement des fourmis */
func (f *Farm) MoveAnts() []Step {
	// initialisation de toutes les fourmis dans la chambre start
	start, err := f.GetRoomByType("start")
	end, _ := f.GetRoomByType("end")

	if err != nil {
		fmt.Println(err)
	} else {
		start.NbAnt = f.NbAnt // initialisation de toutes les fourmis dans la chambre start
	}

	nbAntByPath := make([]int, len(f.StepsByPath)) // Mesure du nb de fourmis en cours de déplacement par chemin
	for i, p := range f.Paths {
		nbAntByPath[i] = len(p.Rooms) - 1
	}

	// Liste d'étapes
	var ListStep []Step

	fmt.Println("nbStep:", f.NbStep)
	for iStep := 1; iStep <= f.NbStep; iStep++ {
		fmt.Print("Step:", iStep, " -- ")

		isEnd := false

		var step Step
		step.Id = iStep

		for iant := range f.Ants {
			end.NbAnt = 0

			if f.Ants[iant].IdPath == -1 {
				// choisi un chemin si la premiere room d'un chemin est libre
				for iPath, path := range f.Paths {
					if path.Rooms[1].NbAnt == 0 &&
						nbAntByPath[iPath] < f.StepsByPath[iPath] {
						if len(path.Rooms) == 2 && isEnd {
							continue
						}
						nbAntByPath[iPath]++
						f.Ants[iant].IdPath = iPath
						f.Ants[iant].IdRoom = 0
						break
					}
				}
			}

			// si la fourmi n'as pas trouver de chemin libre -> fourmi suivante
			if f.Ants[iant].IdPath == -1 {
				continue
			}

			var path = f.Paths[f.Ants[iant].IdPath].Rooms // chemin de la fourmi

			// Test si la fourmi est arrivée
			if path[f.Ants[iant].IdRoom].TypeRoom == "end" {
				continue
			}

			path[f.Ants[iant].IdRoom].NbAnt-- // decremente le nb de fourmi dans la room actuelle
			f.Ants[iant].IdRoom++             // On passe à la salle suivante
			path[f.Ants[iant].IdRoom].NbAnt++ // increment le nb de fourmi dans la room suivante
			f.Ants[iant].NameRoom = path[f.Ants[iant].IdRoom].Name

			if path[f.Ants[iant].IdRoom].TypeRoom == "end" && len(path) == 2 {
				isEnd = true
			}
			fmt.Print(f.Ants[iant].Name, "-", f.Ants[iant].NameRoom, " ")

			step.Ants = append(step.Ants, Ant{Name: f.Ants[iant].Name , NameRoom: f.Ants[iant].NameRoom})
		}

		ListStep = append(ListStep, step)

		//fmt.Println("room end:", end.NbAnt)
		fmt.Println("")
	}

	fmt.Println("--------------------------------------------------------------------")
	return ListStep
}

// Fonctions ----------------------------------------------------------------------------

/* Verifie si une room est contenue dans un chemin */
func contains(rooms []*Room, r *Room) bool {
	for _, room := range rooms {
		if room == r {
			return true
		}
	}
	return false
}

func sliceEqual(a, b []*Room) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
