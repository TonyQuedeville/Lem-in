/*
06/01/2023
@author: 
    Tony Quedeville (tquedevi)
Zone01
Projet Lem-in
*/
/*------------------------------------------------------------------------------------------*/

function indexApp(){
    const canvas = document.getElementById('canvas')
    /**
     * @type {CanvasRenderingContext2D}
     */
    const ctx = canvas.getContext('2d')
    let RoomsData = []
    let Rooms = []

    async function fetchFarm() {
        const response = await fetch('./static/json/farm.json');
        if (!response.ok) {
            console.error(`Error fetch farm.json : ${response.status}`)
        }
        const farmData = await response.json();
        return farmData
    }

    fetchFarm().then(farmData => {
        document.getElementById("example").textContent = farmData.Name
        document.getElementById("NbAnt").textContent = farmData.NbAnt
        document.getElementById("NbStep").textContent = "0/" + farmData.NbStep

        // Parcours des données Rooms
        for (let room of farmData.Rooms){
            RoomsData.push(room)
        }
        
        // Instanciations et affichage des Chambres
        RoomsData.forEach(roomData => {
            const room = new Room(roomData.Name, roomData.TypeRoom, roomData.X, roomData.Y, roomData.RoomLink, roomData.RoomDest)
            room.addRoom()
            Rooms.push(room)
        })

        // Affichage des liens
        for (let link of farmData.DestRooms){
            let room
            let roomLink

            Rooms.forEach(el => {
                if (el.name == link[0]) {room = el}
                if (el.name == link[1]) {roomLink = el}
            })
            drawLine(ctx, room.left, room.top, roomLink.left, roomLink.top)
        }

        // Affichage des Chemins optimisés
        document.getElementById("nbPath").textContent = "path : " + farmData.Paths.length
        let ipath = 1
        
        for (let path of farmData.Paths){
            let ligResult = ""
            let x,y
            for (let iRoom in path.Rooms){
                ligResult = ligResult + " " + path.Rooms[iRoom].Name
                
                const room = path.Rooms[iRoom]

                if(iRoom > 0) {
                    drawLine(ctx, x * 80 , y * 80 , room.X * 80 , room.Y * 80 , "yellow")
                }
                
                x = room.X
                y = room.Y
            }

            const antsbypath = farmData.StepsByPath[ipath-1] - path.Rooms.length + 1
            addPathResult(ipath++, antsbypath, ligResult)
        }

        // Instanciations et affichage des fourmis
        var Ants = []
        const room = Room.getRoomByType("start", Rooms)
        for(i=1; i<=farmData.NbAnt; i++){
            const antObj = new Ant("L"+i, room.name, room.left, room.top)
            antObj.addAnt()
            Ants.push(antObj)
        }

        moveAnt(farmData.NbStep, Ants, Rooms)
    })
}

/* Déplacement des fourmis */
function moveAnt(nbStep, ants, rooms){
    async function fetchStep() {
        const response = await fetch('./static/json/step.json')
        if (!response.ok) {
            console.error(`Error fetch step.json : ${response.status}`)
        }
        const stepData = await response.json();
        return stepData
    }

    fetchStep().then(stepData => {
        document.getElementById("NbStep").textContent = stepData.Id

        //console.log(stepData);
        let room = rooms[0]
        let antObj = ants[0]
        const decleration = 15
        let coefX = 0
        let coefY = 0        
        let iStep = 0
        let stepresult = false
        loop()

        function loop(){ // Animation: frame by frame (60Hz)
            let step = stepData[iStep]
            let ligStepResult = ""        

            step.Ants.forEach(ant => {
                ligStepResult = ligStepResult + ant.Name + "-" + ant.NameRoom + " "

                room = Room.getRoomByName(ant.NameRoom, rooms)
                antObj = Ant.getAntByName(ant.Name, ants)

                coefX = (room.left - antObj.left)/decleration
                coefY = (room.top - antObj.top)/decleration
                
                antObj.left = antObj.left + coefX
                antObj.top = antObj.top + coefY
                
                antObj.moveAnt(antObj.left, antObj.top)
            })

            if (!stepresult) {
                document.getElementById("NbStep").textContent = step.Id + " / " + nbStep
                addStepResult(step.Id, ligStepResult)
                stepresult = true
            }

            if(Math.round(antObj.left) == room.left && Math.round(antObj.top) == room.top){
                iStep++
                stepresult = false
            }

            if (iStep < nbStep) {
                requestAnimationFrame(loop)
            }
        }    
    })
}

// Fonctions -----------------------------------------------------------------------

/* Affiche une ligne de liaison entre 2 Rooms */
function drawLine(ctx, xStart, yStart, xEnd, yEnd, color="grey"){
    ctx.beginPath()
    ctx.moveTo(xStart + 55, yStart + 55)
    ctx.lineTo(xEnd + 55, yEnd + 55)
    ctx.strokeStyle = color
    ctx.stroke()
}

/* Affichage d'une ligne de résultat dans l'espace Result */
function addStepResult(step, ligResult){
    const resultHTML = document.getElementById("result")
    const stepHTML = document.createElement('div')
    stepHTML.id = "step" + step
    stepHTML.className = "step"
    stepHTML.textContent = "Step " + step + " : " + ligResult
    resultHTML.append(stepHTML)
}

/* Affichage des chemins choisis dans l'espace Result */
function addPathResult(ipath, AntsByPath, ligResult){
    const resultHTML = document.getElementById("listPath")
    const patHTML = document.createElement('div')
    patHTML.id = "path" + ipath
    patHTML.className = "step"
    patHTML.textContent = "Path" + " " + ipath + " (" + AntsByPath + " ants) : " + ligResult
    resultHTML.append(patHTML)
}
