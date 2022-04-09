package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("\n")
	StrAsci("___L_")

}

func StrAsci(S string) { //le mot que je transforme en ascii
	moi := []rune(S)     //je la transforme un tableau de rune
	var Styve [][]string //ça sera ma liste d'ascii pour afficher le mot
	for i := range moi { //là je mets dans la liste Styve des listes de string
		Styve = append(Styve, lettreASCI(moi[i]))
	}
	for ligne := range Styve[0] { //maintenant j'affiche, il se répètera 10 fois car 10 ligne
		for mot := range Styve { //il se repetera par le nombre de lettre
			fmt.Printf(Styve[mot][ligne]) //j'affiche ligne par ligne
		}
		fmt.Println() //là je retourne à la ligne
	}
}

func lettreASCI(r rune) []string {
	//fonction qui permet de convertir une rune en caractére assci art
	file, err := os.Open("standard.txt")
	if err == nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var tab []string
		counter := 0
		position := int(r - 32)
		for scanner.Scan() {
			if counter >= 9*position+2 && counter < 9*position+9 {
				tab = append(tab, scanner.Text())
			}
			counter++
		}
		return tab
	}
	return nil
}
