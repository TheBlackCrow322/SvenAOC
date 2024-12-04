package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Definiere eine Pair Struktur
type Pair struct {
	First  int64
	Second int64
}

func main() {
	// Die Datei, die die Liste von Paaren enthält
	fileName := "input1.txt" // Ersetze dies durch den Pfad zu deiner Datei

	// Öffne die Datei
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	var pairs []Pair
	scanner := bufio.NewScanner(file)

	// Lies die Datei Zeile für Zeile
	for scanner.Scan() {
		line := scanner.Text() // Aktuelle Zeile

		// Splitte die Zeile anhand von Leerzeichen
		parts := strings.Fields(line)
		if len(parts) != 2 {
			// Falls die Zeile nicht genau 2 Zahlen enthält, überspringe sie
			continue
		}

		// Konvertiere die Strings in int64
		first, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Println("Fehler beim Parsen der Zahl für First:", err)
			continue
		}

		second, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Println("Fehler beim Parsen der Zahl für Second:", err)
			continue
		}

		// Erstelle das Paar und füge es zur Liste hinzu
		pair := Pair{First: first, Second: second}
		pairs = append(pairs, pair)
	}

	// Überprüfe auf Fehler beim Lesen der Datei
	if err := scanner.Err(); err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return
	}

	// Trenne die Paare in separate Listen für First und Second
	var firstList, secondList []int64
	for _, pair := range pairs {
		firstList = append(firstList, pair.First)
		secondList = append(secondList, pair.Second)
	}

	// Sortiere beide Listen
	sort.Slice(firstList, func(i, j int) bool { return firstList[i] < firstList[j] })
	sort.Slice(secondList, func(i, j int) bool { return secondList[i] < secondList[j] })

	// Ausgabe der sortierten Listen
	fmt.Println("Sortierte First-Liste:", firstList)
	fmt.Println("Sortierte Second-Liste:", secondList)

	// Berechne und gebe die Differenzen für die jeweiligen Paare aus
	fmt.Println("\nDifferenzen (First - Second) der entsprechenden Elemente:")
	for i := 0; i < len(firstList); i++ {
		difference := absoluteDifference(firstList[i], secondList[i])
		fmt.Printf("Position %d: Differenz = |%d - %d| = %d\n", i+1, firstList[i], secondList[i], difference)
	}

	// Berechne die Summe der Differenzen
	totalDifference := calculateSumOfDifferences(firstList, secondList)
	fmt.Printf("\nSumme der Differenzen (absolut): %d\n", totalDifference)

	// Zähle, wie oft jede Zahl der ersten Liste in der rechten Liste vorkommt und multipliziere mit der Häufigkeit
	totalMultiplication := multiplyByFrequencyAndSum(firstList, secondList)
	fmt.Printf("\nGesamtsumme der Produkte der Zahlen mit ihrer Häufigkeit in der rechten Liste: %d\n", totalMultiplication)
}

// Funktion, die die absolute Differenz zwischen zwei Zahlen berechnet
func absoluteDifference(a, b int64) int64 {
	return int64(math.Abs(float64(a - b)))
}

// Funktion, die die Summe der absoluten Differenzen berechnet
func calculateSumOfDifferences(firstList, secondList []int64) int64 {
	var total int64
	for i := 0; i < len(firstList); i++ {
		total += absoluteDifference(firstList[i], secondList[i])
	}
	return total
}

// Funktion, die die Häufigkeit der Zahlen der linken Liste in der rechten Liste zählt, mit der Häufigkeit multipliziert und alles summiert
func multiplyByFrequencyAndSum(firstList, secondList []int64) int64 {
	// Häufigkeitstabelle für die rechte Liste erstellen
	countMap := make(map[int64]int64)
	for _, num := range secondList {
		countMap[num]++
	}

	// Variable für die Gesamtsumme der Produkte
	var totalMultiplication int64

	// Multipliziere jede Zahl aus der linken Liste mit ihrer Häufigkeit in der rechten Liste und addiere sie zur Gesamtsumme
	fmt.Println("\nMultiplikation der Zahlen aus der linken Liste mit ihrer Häufigkeit in der rechten Liste:")
	for _, num := range firstList {
		// Überprüfe, ob die Zahl in der rechten Liste vorkommt
		count, exists := countMap[num]
		if exists && count > 0 {
			result := num * count
			totalMultiplication += result
			fmt.Printf("Zahl %d kommt %d Mal vor, Produkt = %d * %d = %d\n", num, count, num, count, result)
		}
	}

	return totalMultiplication
}
