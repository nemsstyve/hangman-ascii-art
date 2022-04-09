package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var tabmots []string           // là où on va stocker les mots de words.txt
	f, err := os.Open("words.txt") // on va lire words.txt
	if err != nil {                // si erreur (words.txt vide ou autre)
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f) // on va scanner words.txt
	scanner.Split(bufio.ScanLines) // on va stocker en lisant ligne par ligne (1 ligne = une place)
	for scanner.Scan() {
		tabmots = append(tabmots, (scanner.Text())) // on ajoute la ligne à  tabmots tant qu'il y'a des lignes dans words.txt
	}

	max := len(tabmots) // ici on définit les limites du tableau qui contient les mots pour en choisir un aléatoirement dedans
	min := 0
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(max - min)
	tresor := tabmots[rand]
	q := Initial(tresor)
	fmt.Print("\n")
	fmt.Println("Good Luck, you have 10 Attempts.")
	fmt.Print("\n")
	sop := Majuscule(q)
	fmt.Print("\n\n")
	StrAsci(sop)
	Attempts := 0
	debut_jeu := spacetahlesfous(tresor)
	for Attempts != 10 || q == debut_jeu {
		scan := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
		fmt.Print("\n")
		fmt.Println("Choose:")
		scan.Scan()                      // lancement du scanner
		entreeUtilisateur := scan.Text() // stockage du résultat du scanner dans une variable
		if Is_This_In_TheWord(entreeUtilisateur, tresor) == true {
			fmt.Print("\n")
			ch := Change(debut_jeu, q, entreeUtilisateur)
			w := Majuscule(ch)
			q = w
			StrAsci(w)
			a := Majuscule(debut_jeu)
			b := Majuscule(q)
			if All_Letters_are_in_Theword(a, b) == true {
				fmt.Println("CONGRATS! You won, but don't claim victory too quickly. Go to the next level")
				break
			} else {
				continue
			}
		} else {
			Attempts++
			zh := Josélecondamné(Attempts)
			z := Change(debut_jeu, q, entreeUtilisateur)
			p := Majuscule(z)
			StrAsci(p)
			fmt.Println(zh)
		}
	}
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

func Initial(s string) string {
	var mot []string      // EX : salut
	var motcaché []string // EX : _ _ _ _ _
	for _, x := range s {
		mot = append(mot, string(x))
		motcaché = append(motcaché, "_")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (len(mot))/2-1; i++ {
		pos := rand.Intn(len(mot))
		motcaché[pos] = mot[pos]
	}
	res := ""
	for i := 0; i < len(motcaché); i++ {
		res = res + motcaché[i] + " "
	}
	return res
}

func Is_This_In_TheWord(letter string, word string) bool { // ci la lettre est presente ds le mot
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}

func Change(mot string, motcaché string, lettreproposésamère string) string { // mettre la lettre prposé a la place du '_'
	var tsl []rune
	radd := []rune(lettreproposésamère)
	str := ""
	for _, r := range mot {
		tsl = append(tsl, r) // je met ds le slice tsl les runes de mots
	}
	for i, r1 := range motcaché {
		if r1 == '_' && tsl[i] == radd[0] {
			r1 = tsl[i]
		}
		str += string(r1)
	}
	return str
}

func spacetahlesfous(s string) string { // rajouter l'espace
	haziz := ""
	for t, i := range s {
		if t != len(s) {
			haziz += string(i) + " "
		} else {
			haziz += string(i)
		}
	}
	return haziz
}

func All_Letters_are_in_Theword(moooot string, mmoootcaché string) bool {
	for _, char1 := range mmoootcaché {
		write := 0
		for _, char2 := range moooot {
			if string(char1) == string(char2) {
				write = 1
			}
		}
		if write == 0 {
			return false
		}
	}
	return true
}

func Majuscule(s string) string {
	relou := []rune(s)
	for x := range relou {
		if relou[x] >= 'a' && relou[x] <= 'z' {
			relou[x] = relou[x] - 32
		}
	}
	return string(relou)
}

func Josélecondamné(try int) string {
	fmt.Print("\n")
	if try == 1 {
		fmt.Printf("you suck de ouf, try again my bro, 9 try \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("         \n")
		fmt.Printf("=========\n")
	}
	if try == 2 {
		fmt.Printf("you suck de ouf, try again my bro, 8 try \n")
		fmt.Printf("       \n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 3 {
		fmt.Printf("you suck de ouf, try again my bro, 7 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 4 {
		fmt.Printf("you suck de ouf, try again my bro, 6 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 5 {
		fmt.Printf("you suck de ouf, try again my bro, 5 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 6 {
		fmt.Printf("you suck de ouf, try again my bro, 4 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 7 {
		fmt.Printf("you suck de ouf, try again my bro, 3 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 8 {
		fmt.Printf("you suck de ouf, try again my bro, 2 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 9 {
		fmt.Printf("you suck de ouf, try again my bro, 1 try \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	if try == 10 {
		fmt.Printf("You were hanged, go back to fortnite please \n")
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("=========\n")
	}
	return string(try)
}
