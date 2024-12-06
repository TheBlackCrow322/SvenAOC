package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func main() {
	// Pfad zur Datei
	filePath := "input3.txt" // Ersetze dies durch den Pfad zu deiner Datei

	// Datei einlesen
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err) // Bei Fehler wird das Programm gestoppt
	}

	// Inhalt als String behandeln
	text := string(content)

	// Regular Expression für "mul(x,y)" (z.B. mul(2,3))
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Alle Treffer finden
	matches := re.FindAllStringSubmatch(text, -1)

	// Überprüfen, ob Treffer gefunden wurden
	if len(matches) == 0 {
		fmt.Println("Keine Multiplikationen gefunden.")
		return
	}

	// Variable für das Gesamtergebnis
	var totalSum int

	// Alle gefundenen Multiplikationen durchgehen
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

		// Multiplikation berechnen
		result := x * y

		// Ergebnis ausgeben
		fmt.Printf("Multiplikation gefunden: mul(%d, %d) = %d\n", x, y, result)

		// Addiere das Ergebnis zur Gesamtsumme
		totalSum += result
	}

	// Gesamtergebnis ausgeben
	fmt.Printf("\nGesamtergebnis der Multiplikationen: %d\n", totalSum)
}
