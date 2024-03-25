package resources

import (
	"fmt"
	"golang-demo/utils"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Map() {
	utils.LineBreak()
	fmt.Println("Create an Empty Map")
	fmt.Println("")
	var mapA = make(map[string]string)
	var mapB map[string]string
	fmt.Printf("mapA value\t\t: %v\n", mapA)
	fmt.Printf("mapA isnill\t\t: %t\n", mapA == nil)
	fmt.Printf("mapB value\t\t: %v\n", mapB)
	fmt.Printf("mapB isnill\t\t: %t\n", mapB == nil)
	fmt.Println("")
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Create Map Using var Keyword")
	fmt.Println("")
	var mapC = map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	fmt.Printf("mapC value\t\t: %v\n", mapC)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Create Map Using := Notation")
	fmt.Println("")
	mapD := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	fmt.Printf("mapD value\t\t: %v\n", mapD)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Create Maps Using make()Function:")
	fmt.Println("")
	mapE := make(map[string]string)
	fmt.Printf("mapE value\t\t: %v\n", mapE)
	mapE["name"] = "Rangga Wiraguna"
	mapE["role"] = "Mobile Developer"
	fmt.Printf("mapE value\t\t: %v\n", mapE)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Access Map Element")
	fmt.Println("")
	mapF := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	fmt.Printf("Name\t\t\t: %v\n", mapF["name"])
	fmt.Printf("Role\t\t\t: %v\n", mapF["role"])

	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Update and Add Map Element")
	fmt.Println("")
	mapG := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	fmt.Printf("mapG Name\t\t: %v\n", mapG["name"])
	fmt.Printf("mapG Role\t\t: %v\n", mapG["role"])
	mapG["name"] = "Bujang Enam"
	mapG["role"] = "DevOps Engineer"
	fmt.Printf("mapG Name Updated\t: %v\n", mapG["name"])
	fmt.Printf("mapG Role Updated\t: %v\n", mapG["role"])
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Remove Element from Map")
	fmt.Println("")
	mapH := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	fmt.Printf("mapH value\t\t: %v\n", mapH)
	delete(mapH, "role")
	fmt.Printf("mapH value Deleted\t: %v\n", mapH)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Check For Specific Elements in a Map")
	fmt.Println("")
	mapI := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	name, nameAvailable := mapI["name"]    // Checking for existing key and its value
	email, emailAvailable := mapI["email"] // Checking for non-existing key and its value
	_, roleAvailable := mapI["role"]       // Only checking for existing key and not its value
	fmt.Printf("Name Value\t\t: %s (%t)\n", name, nameAvailable)
	fmt.Printf("Email Value\t\t: %s (%t)\n", email, emailAvailable)
	fmt.Printf("Role Available\t\t: (%t)\n", roleAvailable)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Map Characteristic")
	fmt.Println("")
	mapJ := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer"}
	mapK := mapJ
	fmt.Printf("mapJ value\t\t: %v\n", mapJ)
	fmt.Printf("mapK value\t\t: %v\n", mapK)
	mapJ["role"] = "Fullstack Developer"
	mapK["name"] = "Bujang Enam"
	delete(mapK, "role")
	fmt.Printf("mapJ value\t\t: %v\n", mapJ)
	fmt.Printf("mapK value\t\t: %v\n", mapK)
	fmt.Println("Note\t\t\t: The values of mapJ and mapK use the same reference")
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Iterate Over Map")
	fmt.Println("")
	mapL := map[string]string{"gender": "Male", "name": "Rangga Wiraguna", "birth": "20 Juni 2000", "role": "Mobile Developer"}
	for key, value := range mapL {
		fmt.Printf("%v\t\t: %v\n", cases.Title(language.Und, cases.NoLower).String(key), value)
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Iterate Over Map in a Specific Order")
	fmt.Println("")
	mapM := map[string]string{"name": "Rangga Wiraguna", "role": "Mobile Developer", "gender": "Male", "birth": "20 Juni 2000"}
	var sliceN []string
	sliceN = append(sliceN, "name", "role", "gender", "birth")
	for _, key := range sliceN {
		fmt.Printf("%v\t\t: %v\n", cases.Title(language.Und, cases.NoLower).String(key), mapM[key])
	}
	utils.LineBreak()
}
