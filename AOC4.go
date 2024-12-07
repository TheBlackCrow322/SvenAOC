package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Funktion zum Suchen von "XMAS" in alle Richtungen
func countXMAS(grid []string) (int, [][]bool) {
	word := "XMAS"
	directions := [][2]int{
		{0, 1},   // Rechts
		{1, 0},   // Nach unten
		{1, 1},   // Diagonal nach unten rechts
		{1, -1},  // Diagonal nach unten links
		{0, -1},  // Links
		{-1, 0},  // Nach oben
		{-1, -1}, // Diagonal nach oben links
		{-1, 1},  // Diagonal nach oben rechts
	}

	// Ein 2D-Array, das anzeigt, wo "XMAS" gefunden wurde
	mask := make([][]bool, len(grid))
	for i := range mask {
		mask[i] = make([]bool, len(grid[i]))
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Über alle Positionen im Gitter iterieren
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Über alle Richtungen iterieren
			for _, dir := range directions {
				x, y := i, j
				found := true
				// Überprüfen, ob wir XMAS in dieser Richtung finden können
				for k := 0; k < len(word); k++ {
					if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] != word[k] {
						found = false
						break
					}
					x += dir[0]
					y += dir[1]
				}
				// Wenn wir "XMAS" gefunden haben, erhöhen wir den Zähler und setzen die Maske
				if found {
					count++
					x, y = i, j
					for k := 0; k < len(word); k++ {
						mask[x][y] = true
						x += dir[0]
						y += dir[1]
					}
				}
			}
		}
	}
	return count, mask
}

// Funktion zum Einlesen des Gitters aus einer Datei
func readGridFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Zeile einlesen und als String in das Gitter-Array einfügen
		grid = append(grid, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

// Funktion zum Erstellen des maskierten Gitters
func createMaskedGrid(grid []string, mask [][]bool) []string {
	maskedGrid := make([]string, len(grid))
	for i := 0; i < len(grid); i++ {
		var maskedLine []rune
		for j := 0; j < len(grid[i]); j++ {
			if mask[i][j] {
				maskedLine = append(maskedLine, rune(grid[i][j]))
			} else {
				maskedLine = append(maskedLine, '.')
			}
		}
		maskedGrid[i] = string(maskedLine)
	}
	return maskedGrid
}

func main() {
	// Datei einlesen
	filename := "input4.txt" // Hier den Namen der Textdatei eintragen
	grid, err := readGridFromFile(filename)
	if err != nil {
		fmt.Println("Fehler beim Einlesen der Datei:", err)
		return
	}

	// Zählen der Vorkommen von XMAS und das Masken-Gitter erstellen
	count, mask := countXMAS(grid)

	// Maskiertes Gitter erstellen
	maskedGrid := createMaskedGrid(grid, mask)

	// Ausgabe
	fmt.Println("Original Gitter:")
	for _, line := range grid {
		fmt.Println(line)
	}

	fmt.Println("\nMaskiertes Gitter:")
	for _, line := range maskedGrid {
		fmt.Println(line)
	}

	fmt.Printf("\nDas Wort XMAS erscheint %d mal.\n", count)
}
