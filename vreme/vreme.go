// Package vreme omogoča shranjevanje vremenskih postaj, dodajanje meritev
// ter izpis meritev in statističnih poročil.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom, krajem in meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo postajo v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev doda temperaturno meritev postaji, če je v dovoljenem območju.
// Če temperaturne meje niso podane, uporabi območje od -50 do 50 stopinj
// Celzija.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, temperaturneMeje ...float64) {
	minTemp, maxTemp := -50.0, 50.0
	if len(temperaturneMeje) == 2 {
		minTemp, maxTemp = temperaturneMeje[0], temperaturneMeje[1]
	}

	if temperatura >= minTemp && temperatura <= maxTemp {
		if postaja, ok := postaje[oznaka]; ok {
			postaja.meritve = append(postaja.meritve, temperatura)
			postaje[oznaka] = postaja
		} else {
			fmt.Printf("Postaja z oznako %s ne obstaja.\n", oznaka)
		}
	} else {
		fmt.Printf("Temperatura %.1f ni v dovoljenem območju (%.1f do %.1f).\n", temperatura, minTemp, maxTemp)
	}
}

// statistika izračuna najnižjo, najvišjo in povprečno temperaturo postaje.
// Uspeh vrne le, če postaja obstaja in ima dovolj meritev.
func statistika(postaje map[string]Postaja, oznaka string, najmanjMeritev ...int) (najnizja, najvisja, povprecna float64, ok bool) {
	stMeritev := 3
	if len(najmanjMeritev) == 1 {
		stMeritev = najmanjMeritev[0]
	}

	if postaja, ok := postaje[oznaka]; ok {
		if len(postaja.meritve) >= stMeritev {
			najnizja := postaja.meritve[0]
			najvisja := postaja.meritve[0]
			vsota := postaja.meritve[0]
			for _, meritev := range postaja.meritve[1:] {
				najnizja = min(najnizja, meritev)
				najvisja = max(najvisja, meritev)
				vsota += meritev
			}
			povprecna := vsota / float64(len(postaja.meritve))
			return najnizja, najvisja, povprecna, true
		} else {
			fmt.Printf("Postaja z oznako %s nima dovolj meritev (vsaj %d so potrebne).\n", oznaka, stMeritev)
			return 0, 0, 0, false
		}
	} else {
		fmt.Printf("Postaja z oznako %s ne obstaja.\n", oznaka)
		return 0, 0, 0, false
	}
}

// IzpisMeritev izpiše vse postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// IzpisPorocila izpiše statistično poročilo za vse postaje.
// Če najmanjše število meritev ni podano, so potrebne tri meritve.
func IzpisPorocila(postaje map[string]Postaja, najmanjMeritev ...int) {
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, najmanjMeritev...)
		if !ok {
			fmt.Printf("%s: Premalo meritev\n", postaja.ime)
			return
		}
		opis := ""
		switch {
		case povprecna >= 25:
			opis = "Vroče"
		case povprecna >= 15 && povprecna < 25:
			opis = "Toplo"
		case povprecna >= 5 && povprecna < 15:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}
		fmt.Printf("%s: min %.1f°C, max %.1f°C, povprečje %.1f°C -> %s\n", postaja.ime, najnizja, najvisja, povprecna, opis)
	}
}
