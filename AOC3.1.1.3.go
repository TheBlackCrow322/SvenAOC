package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Struct, um eine Zeile und ihre Multiplikationen zu speichern
type ProcessedLine struct {
	OriginalLine    string
	Multiplications []string
}

func main() {
	// Pfad zur Textdatei angeben
	filePath := "input3.txt" // Ersetze dies mit dem Pfad zu deiner Datei

	// Datei einlesen
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err) // Fehlerbehandlung, falls die Datei nicht gelesen werden kann
	}

	// Inhalt als String behandeln
	text := string(content)

	// Wir fügen nach "do()" und "don't()" Zeilenumbrüche ein, um die Verarbeitung zu erleichtern
	// Nach jedem "do()" und "don't()" einen Zeilenumbruch einfügen
	text = strings.ReplaceAll(text, "do()", "\ndo()")
	text = strings.ReplaceAll(text, "don't()", "\ndon't()")

	// Text in einzelne Zeilen aufteilen
	lines := strings.Split(text, "\n")

	// Regular Expression für "mul(X,Y)" (z.B. mul(2,3))
	re := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)

	// Flag, um zu steuern, ob Multiplikationen berücksichtigt werden
	processMultiplications := true // Zu Beginn sind Multiplikationen aktiviert

	// Slice, um alle verarbeiteten Zeilen und Multiplikationen zu speichern
	var processedLines []ProcessedLine

	// Regular Expressions für "do()" und "don't()"
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	// Zeilenweise durchgehen
	for _, line := range lines {
		// Entfernen von Leerzeichen und unnötigen Zeilenumbrüchen
		line = strings.TrimSpace(line)

		// Wenn die Zeile leer ist, überspringen
		if line == "" {
			continue
		}

		// Erstellt eine Struktur, die die Originalzeile und die gefundenen Multiplikationen speichert
		var currentProcessedLine ProcessedLine
		currentProcessedLine.OriginalLine = line

		// Überprüfen, ob die Zeile "do()" oder "don't()" enthält und den Status entsprechend ändern
		if doRe.MatchString(line) {
			processMultiplications = true
			currentProcessedLine.Multiplications = append(currentProcessedLine.Multiplications, "Multiplikationen aktiviert.")
		}
		if dontRe.MatchString(line) {
			processMultiplications = false
			currentProcessedLine.Multiplications = append(currentProcessedLine.Multiplications, "Multiplikationen deaktiviert.")
		}

		// Suchen und Verarbeiten der "mul(x,y)"-Ausdrücke in der Zeile
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// Extrahierte Zahlen aus der Regex
			xStr := match[1]
			yStr := match[2]

			// Umwandeln der Strings in Integers
			x, err := strconv.Atoi(xStr)
			if err != nil {
				log.Fatal("Fehler beim Umwandeln von x:", err)
			}
			y, err := strconv.Atoi(yStr)
			if err != nil {
				log.Fatal("Fehler beim Umwandeln von y:", err)
			}

			// Wenn die Multiplikation berücksichtigt wird, ausführen
			if processMultiplications {
				result := x * y
				currentProcessedLine.Multiplications = append(currentProcessedLine.Multiplications, fmt.Sprintf("mul(%d, %d) = %d", x, y, result))
			} else {
				// Multiplikation wird ignoriert
				currentProcessedLine.Multiplications = append(currentProcessedLine.Multiplications, fmt.Sprintf("mul(%d, %d) (ignoriert)", x, y))
			}
		}

		// Die aktuelle Zeile mit den Multiplikationen in das Slice aufnehmen
		processedLines = append(processedLines, currentProcessedLine)
	}

	// Gesamtergebnis der Multiplikationen ausgeben
	var totalSum int
	for _, processedLine := range processedLines {
		for _, multiplication := range processedLine.Multiplications {
			// Versuchen, das Ergebnis zu extrahieren und zu summieren
			// Beispiel: "mul(2, 3) = 6"
			if strings.Contains(multiplication, "=") {
				parts := strings.Split(multiplication, "=")
				resultStr := strings.TrimSpace(parts[1])
				result, err := strconv.Atoi(resultStr)
				if err == nil {
					totalSum += result
				}
			}
		}
	}

	// Ausgabe der verarbeiteten Zeilen
	fmt.Println("\nVerarbeitete Zeilen:")
	for _, processedLine := range processedLines {
		fmt.Printf("Zeile: %s\n", processedLine.OriginalLine)
		for _, multiplication := range processedLine.Multiplications {
			fmt.Printf("  %s\n", multiplication)
		}
	}

	// Gesamtergebnis ausgeben
	fmt.Printf("\nGesamtergebnis der Multiplikationen: %d\n", totalSum)
}
