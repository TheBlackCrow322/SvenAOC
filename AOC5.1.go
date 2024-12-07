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
func isTopologicallySorted(list []int, pairs []Pair) bool {
	position := make(map[int]int)
	for i, value := range list {
		position[value] = i
	}

	for _, pair := range pairs {
		firstPos, firstExists := position[pair.First]
		secondPos, secondExists := position[pair.Second]

		if firstExists && secondExists {
			if firstPos > secondPos {
				return false
			}
		}
	}
	return true
}

// topologicalSort korrigiert die Liste basierend auf den Paaren
func topologicalSort(numbers []int, pairs []Pair) []int {
	// Set für schnelle Überprüfung, welche Zahlen in der Liste vorkommen
	numberSet := make(map[int]bool)
	for _, num := range numbers {
		numberSet[num] = true
	}

	// Graph und In-Degree aufbauen
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Initialisieren
	for _, num := range numbers {
		graph[num] = []int{}
		inDegree[num] = 0
	}

	// Kanten hinzufügen (nur gültige Paare berücksichtigen)
	for _, pair := range pairs {
		if numberSet[pair.First] && numberSet[pair.Second] {
			graph[pair.First] = append(graph[pair.First], pair.Second)
			inDegree[pair.Second]++
		}
	}

	// Knoten mit In-Degree 0 in die Queue
	var queue []int
	for num, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}

	// Topologische Sortierung
	var sorted []int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Rückgabe der sortierten Liste (nur falls vollständig sortiert)
	if len(sorted) == len(numbers) {
		return sorted
	}

	// Fehlerfall: Rückgabe einer leeren Liste
	return []int{}
}

// Funktion zur Berechnung der mittleren Zahl
func calculateMiddleValue(list []int) int {
	if len(list) == 0 {
		fmt.Println("Fehler: Die Liste ist leer und hat keine mittlere Zahl.")
		return 0
	}
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

	// Überprüfung und mittlere Zahlen korrigierter Listen summieren
	totalMiddleValue := 0
	fmt.Println("\nÜberprüfung der Listen:")
	for i, list := range lists {
		fmt.Printf("\nListe %d: %v\n", i+1, list)
		if isTopologicallySorted(list, pairs) {
			fmt.Println("✔ Die Liste ist bereits korrekt sortiert. Ignorieren...")
		} else {
			fmt.Println("❌ Die Liste ist nicht korrekt. Korrigiere...")
			correctedList := topologicalSort(list, pairs)
			if len(correctedList) == 0 {
				fmt.Println("Die korrigierte Liste konnte nicht erstellt werden. Überspringe...")
				continue
			}
			fmt.Printf("Korrigierte Liste: %v\n", correctedList)
			middleValue := calculateMiddleValue(correctedList)
			fmt.Printf("Mittlere Zahl der korrigierten Liste: %d\n", middleValue)
			totalMiddleValue += middleValue
		}
	}

	// Gesamtergebnis ausgeben
	fmt.Printf("\nGesamtsumme der mittleren Zahlen korrigierter Listen: %d\n", totalMiddleValue)
}
