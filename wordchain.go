package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Dictionary map[string]interface{}

// readDictionary is the function used to read the english.json file present in the local location and store all the words and meanings in the map structure
func readDictionary() {
	// Open our jsonFile
	jsonFile, err := os.Open("english.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened english.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// // read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// a map container to decode the JSON structure into
	Dictionary = make(map[string]interface{})

	// unmarschal JSON
	err = json.Unmarshal(byteValue, &Dictionary)

	// Print on error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("DIctionary extraction completed")

}

//formatstring corrects the captitalization of the word, makes CAT -> Cat
func formatString(Input string) (formatedstring string) {

	if len(Input) > 1 {
		Headpart := Input[:1]
		Tailpart := Input[1:]

		formatedstring = fmt.Sprintln(strings.ToUpper(Headpart) + strings.ToLower(Tailpart))
		formatedstring = strings.TrimSpace(formatedstring)
	}
	return formatedstring
}

//StartChain is the method used to print the chain meaningful words after comparing the key with dictionary
func StartChain(Initialword, Finalword string) {

	//start with printing initial input on screen
	fmt.Println(formatString(Initialword))

	var Words = "abcdefghijklmnopqrstuvwxyz"
	placeOfChange := 2

	for {

		//condition to initialize place of change in the word
		if placeOfChange >= len(Initialword) {
			fmt.Println(Finalword)
			placeOfChange = 1
			return
		}

		for _, letter := range Words {

			//Reordering the string as per the place of change
			letterstojoin := []string{}
			firstpart := Initialword[:placeOfChange-1]
			secondpart := Initialword[placeOfChange:]
			letterstojoin = append(letterstojoin, firstpart, string(letter), secondpart)

			//redered word after changing the letter at place of change
			textout := strings.Join(letterstojoin, "")

			if textout != Initialword {

				//confition to compare the keys of the dictionary
				if _, ok := Dictionary[formatString(textout)]; ok {

					if strings.ToUpper(Initialword[placeOfChange-1:placeOfChange]) != strings.ToUpper(textout[placeOfChange-1:placeOfChange]) {

						//priont the meaningful words, which is present in the dictionary given
						fmt.Println(formatString(textout))

						Initialword = strings.ToUpper(textout)

						//condition to exit after reacing to final word
						if strings.ToUpper(Finalword) == strings.ToUpper(textout) {
							return
						}

						break
					}
				}
			}

		}
		placeOfChange++
	}

}

//startServer starts the whole process, to achive the objective
func startServer() {

	// Create a single reader which can be called multiple times
	reader := bufio.NewReader(os.Stdin)

	// Prompt and read input 1
	fmt.Print("Enter String 1: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	//Read Input text 2
	fmt.Print("Enter string 2: ")
	text2, err := reader.ReadString('\n')

	//Handle error
	if err != nil {
		fmt.Println(err)
	}

	readDictionary()
	StartChain(text, text2)

}

func main() {
	startServer()
}
