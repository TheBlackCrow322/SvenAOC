package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Funktion, die prüft, ob die Differenz zwischen benachbarten Zahlen in einem Bereich von 1 bis 3 liegt
func checkStepDifference(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		diff := abs(numbers[i+1] - numbers[i])
		if diff < 1 || diff > 3 {
			return false // Wenn eine Differenz nicht im Bereich 1 bis 3 liegt, wird false zurückgegeben
		}
	}
	return true // Alle Differenzen waren im Bereich 1 bis 3
}

// Hilfsfunktion, um den absoluten Wert der Differenz zu berechnen
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Funktion, um zu prüfen, ob eine Liste sortiert ist (aufsteigend oder absteigend)
func isSorted(numbers []int) bool {
	// Aufsteigend sortiert
	ascending := true
	// Absteigend sortiert
	descending := true

	// Überprüfen, ob die Liste aufsteigend oder absteigend ist
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			descending = false
		}
		if numbers[i] > numbers[i+1] {
			ascending = false
		}
	}

	// Die Liste ist entweder aufsteigend oder absteigend
	return ascending || descending
}

// Funktion zur Berechnung und Ausgabe der Differenzen
func printDifferences(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		fmt.Printf("Differenz zwischen %d und %d: %d\n", numbers[i], numbers[i+1], diff)
	}
}

func main() {
	// Datei öffnen
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	// Zeilenweise Einlesen der Datei
	var listOfNumbers [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Jede Zeile lesen und durch Leerzeichen getrennte Zahlen extrahieren
		line := scanner.Text()
		parts := strings.Fields(line) // Trennt die Zahlen in der Zeile

		var numbers []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Fehler beim Parsen der Zahl:", err)
				return
			}
			numbers = append(numbers, num)
		}

		// Die Liste von Zahlen zur übergeordneten Liste hinzufügen
		listOfNumbers = append(listOfNumbers, numbers)
	}

	// Fehlerprüfung des Scanners
	if err := scanner.Err(); err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return
	}

	// Zähler für Listen, die den Anstieg/Abstieg mit Differenz 1 bis 3 erfüllen
	countStepDifference := 0
	// Zähler für Listen, die sortiert sind
	countSorted := 0

	// Verarbeitung der Zeilen
	fmt.Println("Überprüfung der Differenzen in den Zeilen:")

	for _, numbers := range listOfNumbers {
		// Prüfen, ob die Zeile sortiert ist
		if isSorted(numbers) {
			countSorted++
			fmt.Printf("\nZeile: %v ist sortiert.\n", numbers)
			printDifferences(numbers)

			// Prüfen, ob die Zeile einen Anstieg/Abstieg mit Differenz 1 bis 3 hat
			if checkStepDifference(numbers) {
				countStepDifference++
				fmt.Printf("Zeile: %v erfüllt den Anstieg/Abstieg mit Differenz zwischen 1 und 3.\n", numbers)
			} else {
				fmt.Printf("Zeile: %v erfüllt NICHT den Anstieg/Abstieg mit Differenz zwischen 1 und 3.\n", numbers)
			}

		} else {
			// Wenn die Zeile nicht sortiert ist, wird sie ignoriert
			fmt.Printf("\nZeile: %v ist nicht sortiert und wird ignoriert.\n", numbers)
		}
	}

	// Ausgabe der Anzahl der Listen, die die Anstiegs-/Abstiegsbedingung erfüllen
	fmt.Printf("\nAnzahl der Listen, die einen Anstieg/Abstieg mit Differenz zwischen 1 und 3 haben: %d\n", countStepDifference)

	// Ausgabe der Anzahl der sortierten Listen
	fmt.Printf("\nAnzahl der sortierten Listen: %d\n", countSorted)
}
