# Projet
Lem-in: Projet Zone01.

  Toutes les fourmis partent d'un point de départ et doivent arriver au point d'arrivé avec effectuant le moins d'étapes possible afin
d'optimiser le temps de parcours de l'ensemble de la colonnie.
L'algorytme cherche dans un premier temps tous les chemins possibles, puis ne garde que ceux qui ne se croisent pas. Ensuite un calcul de pondération permet d'affecter le bon nombres de fourmis sur chaque chemins choisis. Enfin les fourmis sont déplacées étape par étape à travers les chambres de la fourmiliaire jusqu'à la dernière.

Le programme est fait en Go. L'interface de visualisation est une appli web en JS, html, le css est en Sass.

Lancement:
  go run main.go example01.txt
  essayez tous les exemples contenu dans le dossier "examples"
  un resultat conforme au éxigences de l'exercice s'affiche dans le terminal. 
  l'interface web optionnelle est accessible sur localhost 8000.