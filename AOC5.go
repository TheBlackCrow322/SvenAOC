package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Pair repräsentiert ein Zahlenpaar (a vor b)
type Pair struct {
	First  int
	Second int
}

// Funktion, die die Datei liest und Paare und Listen extrahiert
func readFile(filePath string) ([]Pair, [][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Fehler beim Öffnen der Datei: %v", err)
	}
	defer file.Close()

	var pairs []Pair
	var lists [][]int

	scanner := bufio.NewScanner(file)
	isPairSection := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Leere Zeile signalisiert den Übergang von Paaren zu Listen
		if line == "" {
			isPairSection = false
			continue
		}

		if isPairSection {
			// Paare verarbeiten
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("Ungültiges Paar: %s", line)
			}

			first, err1 := strconv.Atoi(parts[0])
			second, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				return nil, nil, fmt.Errorf("Fehler beim Konvertieren: %s", line)
			}

			pairs = append(pairs, Pair{First: first, Second: second})
		} else {
			// Zahlenlisten verarbeiten
			parts := strings.Split(line, ",")
			var numbers []int
			for _, part := range parts {
				number, err := strconv.Atoi(part)
				if err != nil {
					return nil, nil, fmt.Errorf("Fehler beim Konvertieren der Zahl: %s", part)
				}
				numbers = append(numbers, number)
			}
			lists = append(lists, numbers)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Fehler beim Lesen der Datei: %v", err)
	}

	return pairs, lists, nil
}

// isTopologicallySorted überprüft, ob die Liste topologisch korrekt ist.
func isTopologicallySorted(list []int, pairs []Pair) (bool, string) {
	position := make(map[int]int)
	for i, value := range list {
		position[value] = i
	}

	var explanation []string

	for _, pair := range pairs {
		firstPos, firstExists := position[pair.First]
		secondPos, secondExists := position[pair.Second]

		if firstExists && secondExists {
			if firstPos > secondPos {
				return false, fmt.Sprintf("Fehler: %d (Index %d) muss vor %d (Index %d) stehen.", pair.First, firstPos, pair.Second, secondPos)
			}
			explanation = append(explanation, fmt.Sprintf("%d ist korrekt vor %d, weil %d|%d.", pair.First, pair.Second, pair.First, pair.Second))
		}
	}

	return true, fmt.Sprintf("Die Liste ist korrekt sortiert.\nDetails:\n%s", joinLines(explanation))
}

// Hilfsfunktion zum Verbinden von Zeilen
func joinLines(lines []string) string {
	result := ""
	for _, line := range lines {
		result += line + "\n"
	}
	return result
}

// Funktion zur Berechnung der mittleren Zahl
func calculateMiddleValue(list []int) int {
	middle := len(list) / 2
	return list[middle]
}

func main() {
	// Dateipfad
	filePath := "input5.txt"

	// Datei lesen
	pairs, lists, err := readFile(filePath)
	if err != nil {
		fmt.Printf("Fehler: %v\n", err)
		return
	}

	// Paare ausgeben
	fmt.Println("Paar-Definitionen:")
	for _, pair := range pairs {
		fmt.Printf("%d -> %d\n", pair.First, pair.Second)
	}

	// Überprüfung und mittlere Zahlen summieren
	totalMiddleValue := 0
	fmt.Println("\nÜberprüfung der Listen:")
	for i, list := range lists {
		fmt.Printf("\nListe %d: %v\n", i+1, list)
		valid, explanation := isTopologicallySorted(list, pairs)
		if valid {
			fmt.Println("✔ Die Liste ist topologisch korrekt.")
			// Mittlere Zahl hinzufügen
			middleValue := calculateMiddleValue(list)
			fmt.Printf("Mittlere Zahl: %d\n", middleValue)
			totalMiddleValue += middleValue
		} else {
			fmt.Println("❌ Die Liste ist nicht korrekt.")
		}
		fmt.Println(explanation)
	}

	// Gesamtergebnis ausgeben
	fmt.Printf("\nGesamtsumme der mittleren Zahlen korrekt sortierter Listen: %d\n", totalMiddleValue)
}
