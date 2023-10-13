/*
06/01/2023
@author: 
    Tony Quedeville (tquedevi)
Zone01
Projet Lem-in

lien utile : https://tech.mozfr.org/post/2015/08/12/ES6-en-details-:-les-sous-classes-et-l-heritage
*/
/*------------------------------------------------------------------------------------------*/

class Room {
    constructor(name, type, x, y){
        this.name = name
        this.type = type
        this.offsetx = 60
        this.offsety = 120
        this.left = x * 80 
        this.top = y * 80
    }

    addRoom(){
        const farmHTML = document.getElementById("farm")
        this.roomHTML = document.createElement('div')
        this.roomHTML.id = this.name
        this.roomHTML.classList = "room type-" + this.type
        this.roomHTML.textContent = this.name
        this.roomHTML.style.left = this.left + this.offsetx + 'px'
        this.roomHTML.style.top = this.top + this.offsety + 'px'
        farmHTML.append(this.roomHTML)
    }

    static getRoomByName(name, rooms){
        var result = null
    
        rooms.forEach(room => {
            if (room.name == name) {
                result = room
            }
        })
        return result
    }

    static getRoomByType(type, rooms){
        var result = null
    
        rooms.forEach(room => {
            if (room.type == type) {
                result = room
            }
        })
        return result
    }
}

class Ant {
    constructor(name, nameRoom, x, y){
        this.name = name
        this.nameRoom = nameRoom
        this.offsetx = 60
        this.offsety = 120
        this.left = x 
        this.top = y 
    }

    static getAntByName(name, ants){
        var result = null

        ants.forEach(ant => {
            if (ant.name == name){
                result = ant
            }
        })

        return result
    }

    addAnt(){
        const farmHTML = document.getElementById("farm")
        this.antHTML = document.createElement('div')
        this.antHTML.id = this.name
        this.antHTML.classList = "ant"
        this.antHTML.textContent = this.name
        this.antHTML.style.left = this.left + this.offsetx + 'px'
        this.antHTML.style.top = this.top + this.offsety + 'px'
        farmHTML.append(this.antHTML)
    }

    static set left(l){
        this._left = l
    }
    static set top(l){
        this._top = l
    }
    
    moveAnt(x,y){
        this.antHTML.style.left = x + this.offsetx + 'px'
        this.antHTML.style.top = y + this.offsety + 'px'
    }
}

