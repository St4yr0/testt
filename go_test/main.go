package main

import (
	"errors"
	"fmt"
	"go_test/interface_test"
)

// Private visibility
var test1 = "private variable"

// Public visibility
var Test2 = "public variable"

// Private visiblity
func greet() string {
	return "hello"
}

// Public visibility
func Greet2() string {
	return "hello world"
}

// basic func
func deloiu_8_def(nirs string, practika string, leviev string, ocenka float64) {
	fmt.Printf("Delo iu8 ebanoe govno:\n"+
		" Nirs - %s,\n"+
		" Practika - %s,\n a Leviev - %s\n Ocenka etoi xuine - %.5f", nirs, practika, leviev, ocenka)
}

func nirs_xueta(mnenie string) bool {
	return true
}

func ebanati_natriya(ebanati *[2]string, kolvo int) {
	// infinite loop
	for {
		fmt.Println((*ebanati)[kolvo-2] + " - gandon")
		fmt.Println((*ebanati)[kolvo-1] + " - ebanatka")
		break
	}
	for i := range *ebanati {
		if i == 0 {
			fmt.Println((*ebanati)[i] + " - gandon")
			fmt.Println((*ebanati)[i+1] + " - ebanatka")
		} else {
			fmt.Println((*ebanati)[i-1] + " - gandon")
			fmt.Println((*ebanati)[i] + " - ebanatka")
		}
	}
}

func komu_bolnee(pesiki *map[string]int) {
	for pesik, pipiska := range *pesiki {
		fmt.Println(pesik, pipiska)
	}

	(*pesiki)["Serega"] = 12
	fmt.Println(*pesiki)

	delete(*pesiki, "Serega")

}

// pointer finc
func who_is(name *string, whois string) (string, error) {
	if whois == " - Pes" {
		*name = *name + whois
		return *name, nil
	} else {
		return *name, errors.New("Artem vse taki pes")
	}
}

func main() {
	// basic
	message := "ebani chlenosos"
	var nirs = "xueta"
	var practika string
	practika = "also xueta"
	ocenka := 0.00001
	deloiu_8_def(nirs, practika, message, ocenka)
	fmt.Println("\n-------------------------------------------------------------------")
	// pointers
	name := "Artem"
	pName := &name
	pNameChange := &pName
	fmt.Println(name, pName, *pName, "\n", pNameChange, *pNameChange)
	// Artem 0xc000014280 Artem
	// 0xc000012030 0xc000014280
	_, err := who_is(&name, "Pes")
	if err == nil {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Println(name, pName, *pName, "\n", pNameChange, *pNameChange)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println("\n-------------------------------------------------------------------")
	// slices and arrays
	gandoni := [2]string{"Leviev", "Zaitseva"} // array
	kolvo_gandonov := len(gandoni)
	norm_cheli := []string{} // slice
	fmt.Println(norm_cheli)
	ebanati_natriya(&gandoni, kolvo_gandonov)
	// map
	pipiski := map[string]int{"Serega": 10, "Artem": 12}
	// m1 := make(map[string]float64)      // Empty map of string-float64 pairs
	// m2 := make(map[string]float64, 100) // Preallocate room for 100 entries
	// 	_, found := m["pi"]        // found == true
	// _, found = m["pie"]        // found == false
	// if x, found := m["pi"]; found {
	// 	fmt.Println(x)
	// }
	komu_bolnee(&pipiski)
	fmt.Println("\n-------------------------------------------------------------------")
	// structs and interfaces
	// p := interface_test.Ne_pidarasi{"Serega", true}
	ebeni := []interface_test.Eben{interface_test.Pidarasi{"Serega", true, "Ebet Artema"}, interface_test.Ne_pidarasi{"Artem", true}}
	for _, eb := range ebeni {
		fmt.Println(eb.Eben_or_not())
	}

}
