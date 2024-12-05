package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Funktion, die prüft, ob die Differenz zwischen benachbarten Zahlen im Bereich von 1 bis 3 liegt
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

// Funktion, die prüft, ob man durch das Entfernen einer Zahl die Anforderungen erfüllen kann
func canBeFixedByRemovingOne(numbers []int) bool {
	// Versuche, jede Zahl in der Liste zu entfernen und prüfe, ob die verbleibende Liste die Anforderungen erfüllt
	for i := 0; i < len(numbers); i++ {
		// Erstelle eine neue Liste ohne die i-te Zahl
		newList := append([]int{}, numbers[:i]...)
		newList = append(newList, numbers[i+1:]...)

		// Überprüfe, ob die neue Liste sortiert ist und die Differenzen zwischen den Zahlen im Bereich von 1 bis 3 liegen
		if isSorted(newList) && checkStepDifference(newList) {
			return true
		}
	}
	return false
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

	// Zähler für Listen, die die Anforderungen erfüllen (mit oder ohne eine Zahl zu entfernen)
	countFixed := 0
	// Zähler für Listen, die sortiert sind und die Differenzenbedingungen erfüllen
	countSorted := 0

	// Verarbeitung der Zeilen
	fmt.Println("Überprüfung der Differenzen in den Zeilen:")

	for _, numbers := range listOfNumbers {
		// Prüfen, ob die Zeile sortiert ist und die Differenzen den Anforderungen entsprechen
		if isSorted(numbers) && checkStepDifference(numbers) {
			countSorted++
			// Diese Zeile ist bereits sortiert und erfüllt die Differenzanforderungen
			fmt.Printf("\nZeile: %v ist bereits sortiert und erfüllt den Anstieg/Abstieg mit Differenz zwischen 1 und 3.\n", numbers)
		} else {
			// Wenn die Zeile nicht direkt gültig ist, versuchen wir, sie durch Entfernen einer Zahl zu reparieren
			fmt.Printf("\nZeile: %v ist nicht direkt gültig.\n", numbers)
			if canBeFixedByRemovingOne(numbers) {
				countFixed++
				fmt.Printf("Zeile: %v kann durch das Entfernen einer Zahl in eine sortierte Liste mit gültigen Differenzen umgewandelt werden.\n", numbers)
			} else {
				fmt.Printf("Zeile: %v kann durch das Entfernen einer Zahl NICHT in eine gültige Liste umgewandelt werden.\n", numbers)
			}
		}
	}

	// Ausgabe der Anzahl der Listen, die die Anforderungen erfüllen (mit oder ohne eine Zahl zu entfernen)
	fmt.Printf("\nAnzahl der Listen, die durch Entfernen einer Zahl die Anforderungen erfüllen: %d\n", countFixed)

	// Ausgabe der Anzahl der sortierten Listen, die bereits gültig sind
	fmt.Printf("\nAnzahl der sortierten Listen, die die Anforderungen erfüllen: %d\n", countSorted)
}
