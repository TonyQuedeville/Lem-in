/*
06/01/2023
@author: 
    Tony Quedeville (tquedevi)
Zone01
Projet Lem-in

lien utile : https://tech.mozfr.org/post/2015/08/12/ES6-en-details-:-les-sous-classes-et-l-heritage
*/
/*------------------------------------------------------------------------------------------*/

// Connexion TCP
let socket = new WebSocket("ws://127.0.0.1:8080/ws");
console.log("Tentative de connexion ...")

socket.onopen = () => {
    console.log("Connexion OK !")
    //socket.send("Connexion client OK !")
    dataValidate()
}

socket.onclose = event => {
    console.log("Connexion terminée !", event)
    //socket.send("Connexion client terminée !")
}

socket.onmessage = (msg) => {
    //console.log(msg)
}

socket.onerror = error => {
    console.error("Erreur Connexion !", error)
}

/*------------------------------------------------------------------------------------------*/

function dataValidate(){
    var url = document.location.href // url de la page courante 
    //console.log(url);

    if(url == "http://localhost:8080/"){ 
        indexApp()
    }
}

/*------------------------------------------------------------------------------------------*/
