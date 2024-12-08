package main

import (
	"bufio"
	"fmt"
	"os"
)

// Richtungstyp für Bewegungen
type direction struct {
	dx, dy int
}

// Definition der Bewegungsrichtungen (rechts, unten, links, oben)
var directions = []direction{
	{dx: 1, dy: 0},  // rechts
	{dx: 0, dy: 1},  // unten
	{dx: -1, dy: 0}, // links
	{dx: 0, dy: -1}, // oben
}

// Funktion zum Einlesen der Datei
func readFileIntoGrid(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Öffnen der Datei: %v", err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen der Datei: %v", err)
	}

	return grid, nil
}

// Funktion zum Bearbeiten des Gitters
func processGrid(grid [][]rune) ([][]rune, int) {
	rows := len(grid)
	if rows == 0 {
		return grid, 0
	}
	cols := len(grid[0])

	xCount := 0 // Zähler für `X`

	// Starte die Suche an allen Positionen mit '^'
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell == '^' {
				// Ersetze '^' durch 'X'
				grid[rowIndex][colIndex] = 'X'
				xCount++

				// Beginne die Suche ab dieser Position
				x, y := colIndex, rowIndex
				directionIndex := 3 // Starte mit der Richtung nach oben

				for {
					// Bewege in der aktuellen Richtung
					nextX := x + directions[directionIndex].dx
					nextY := y + directions[directionIndex].dy

					// Prüfe, ob der Rand des Gitters erreicht wurde
					if nextX < 0 || nextX >= cols || nextY < 0 || nextY >= rows {
						break
					}

					// Wenn ein `#` gefunden wird, ändere die Richtung vor dem `#`
					if grid[nextY][nextX] == '#' {
						directionIndex = (directionIndex + 1) % 4 // Nächste Richtung im Uhrzeigersinn
						continue
					}

					// Ersetze `.` oder `X` durch `X` und bewege weiter
					if grid[nextY][nextX] == '.' || grid[nextY][nextX] == 'X' {
						if grid[nextY][nextX] == '.' { // Nur zählen, wenn es vorher ein `.`
							xCount++
						}
						grid[nextY][nextX] = 'X'
						x, y = nextX, nextY
					} else {
						// Stoppe, falls ein anderes Zeichen außer `.` oder `X` gefunden wird
						break
					}
				}
			}
		}
	}

	return grid, xCount
}

// Funktion zur Ausgabe des Gitters
func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	// Dateipfad angeben
	filePath := "input6.txt"

	// Datei in ein Gitter einlesen
	grid, err := readFileIntoGrid(filePath)
	if err != nil {
		fmt.Printf("Fehler: %v\n", err)
		return
	}

	// Originalgitter ausgeben
	fmt.Println("Originales Gitter:")
	printGrid(grid)

	// Gitter verarbeiten und Anzahl der `X` zählen
	newGrid, xCount := processGrid(grid)

	// Neues Gitter ausgeben
	fmt.Println("\nNeues Gitter:")
	printGrid(newGrid)

	// Anzahl der `X` ausgeben
	fmt.Printf("\nAnzahl der ersetzten 'X': %d\n", xCount)
}
