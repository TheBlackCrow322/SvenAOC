package main

import (
	"bufio"
	"fmt"
	"os"
)

// Funktion zum Einlesen der Textdatei und Rückgabe der Zeilen als 2D-Array
func readFile(filename string) ([][]rune, error) {
	fmt.Printf("Versuche, die Datei '%s' zu öffnen...\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	// Zeilenweise Einlesen der Datei und Umwandeln in ein 2D-Raster (Array von runes)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Zeile eingelesen: %s\n", line) // Debugging-Ausgabe
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Datei erfolgreich eingelesen. Rastergröße: %dx%d\n", len(grid), len(grid[0]))
	return grid, nil
}

// Funktion zum Überprüfen, ob mindestens zwei "MAS"-Muster in den Diagonalen existieren
func checkTwoMASPatterns(grid [][]rune, r, c int) bool {
	fmt.Printf("Überprüfe A an Position (%d, %d)...\n", r, c) // Debugging-Ausgabe

	// Überprüfen der Diagonalen auf das "MAS"-Muster
	diagonal1 := grid[r-1][c-1] == 'M' && grid[r][c] == 'A' && grid[r+1][c+1] == 'S'
	diagonal2 := grid[r-1][c+1] == 'M' && grid[r][c] == 'A' && grid[r+1][c-1] == 'S'
	diagonal3 := grid[r+1][c-1] == 'M' && grid[r][c] == 'A' && grid[r-1][c+1] == 'S'
	diagonal4 := grid[r+1][c+1] == 'M' && grid[r][c] == 'A' && grid[r-1][c-1] == 'S'

	// Debugging-Ausgaben für jede Diagonale
	fmt.Printf("Diagonal1 (oben links -> unten rechts): %v\n", diagonal1)
	fmt.Printf("Diagonal2 (oben rechts -> unten links): %v\n", diagonal2)
	fmt.Printf("Diagonal3 (unten links -> oben rechts): %v\n", diagonal3)
	fmt.Printf("Diagonal4 (unten rechts -> oben links): %v\n", diagonal4)

	// Überprüfen, ob mindestens zwei der vier Diagonalen das "MAS"-Muster enthalten
	result := (diagonal1 && diagonal2) || (diagonal1 && diagonal3) || (diagonal1 && diagonal4) ||
		(diagonal2 && diagonal3) || (diagonal2 && diagonal4) || (diagonal3 && diagonal4)

	fmt.Printf("Ergebnis für Position (%d, %d): %v\n", r, c, result) // Debugging-Ausgabe
	return result
}

func main() {
	// Den Dateinamen angeben
	filename := "input4.txt" // Hier den Pfad zur Datei angeben

	// Einlesen der Datei
	grid, err := readFile(filename)
	if err != nil {
		fmt.Println("Fehler beim Einlesen der Datei:", err)
		return
	}

	count := 0
	var positions [][]int // Speichern der Positionen der gültigen A-Buchstaben

	// Durch das Gitter iterieren und nach A suchen
	for r := 1; r < len(grid)-1; r++ { // Start bei 1, um Ausreißer zu vermeiden
		for c := 1; c < len(grid[r])-1; c++ { // Start bei 1, um Ausreißer zu vermeiden
			// Wenn wir ein A finden, prüfen wir die Diagonalen
			if grid[r][c] == 'A' {
				fmt.Printf("Gefundenes A an Position (%d, %d).\n", r, c) // Debugging-Ausgabe
				if checkTwoMASPatterns(grid, r, c) {
					count++                                    // Wenn das "MAS"-Muster gefunden wurde, zählen
					positions = append(positions, []int{r, c}) // Speichern der Position des A
				}
			}
		}
	}

	// Ausgabe der Anzahl der validen As
	fmt.Printf("Anzahl der berücksichtigten A: %d\n", count)

	// Ausgabe der Positionen der validen As
	fmt.Println("Positionen der berücksichtigten A:")
	for _, pos := range positions {
		fmt.Printf("A-Position: (%d, %d)\n", pos[0], pos[1])
	}
}
