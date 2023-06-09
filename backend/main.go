package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"

        "github.com/gorilla/websocket"
	"github.com/sgade/randomorg"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
  WriteBufferSize: 1024,

  // We'll need to check the origin of our connection
  // this will allow us to make requests from our React
  // development server to here.
  // For now, we'll do no checking and just allow any connection
  CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)

  // upgrade this connection to a WebSocket
  // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
  }
  // listen indefinitely for new messages coming
  // through on our WebSocket connection
    reader(ws)
}

func setupRoutes() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
  })
  // mape our `/ws` endpoint to the `serveWs` function
    http.HandleFunc("/ws", serveWs)
}

func makeClient() *randomorg.Random {
	apiKey := os.Getenv("RANDOMORG_API_KEY")
	client := randomorg.NewRandom(apiKey)
	return client
}

func rollD6() int {
	random := makeClient()
	value, err := random.GenerateIntegers(1, 1, 6)
	//fmt.Printf("Rolled a %v\n", value)
	if err != nil {
		panic(err)
	}
	return int(value[0])
}

func rollAbilityScores() [][]int {
	allRolls := make([][]int, 6)
	for i := 0; i < 6; i++ {
		allRolls[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			allRolls[i][j] = rollD6()
		}
	}
	return allRolls
}

func generateAbilityScores() []int {
	abilityRolls := rollAbilityScores()
	abilityScores := make([]int, 6)
	for i := 0; i < 6; i++ {
		sort.Ints(abilityRolls[i])
		for j := 1; j < 4; j++ {
			abilityScores[i] += abilityRolls[i][j]
		}
	}
	sort.Ints(abilityScores)
	return abilityScores
}

func getCharInfo() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your character's name? ")
	charName, _ := reader.ReadString('\n')
	charName = strings.TrimSuffix(charName, "\n")
	charName = strings.Title(charName)

	fmt.Print("What is your character's race? ")
	charRace, _ := reader.ReadString('\n')
	charRace = strings.TrimSuffix(charRace, "\n")
	charRace = strings.Title(charRace)

	fmt.Print("What is your character's subrace? ")
	charSubRace, _ := reader.ReadString('\n')
	charSubRace = strings.TrimSuffix(charSubRace, "\n")
	charSubRace = strings.Title(charSubRace)

	fmt.Print("What is your character's class? ")
	charClass, _ := reader.ReadString('\n')
	charClass = strings.TrimSuffix(charClass, "\n")
	charClass = strings.Title(charClass)

	test := "Bard"
	fmt.Println(reflect.TypeOf(test))
	fmt.Println(reflect.TypeOf(charClass))
	fmt.Println(charClass == "Bard")

	return charName, charRace, charSubRace, charClass
}

func finalizeScores() []int {
	charName, charRace, charSubRace, charClass := getCharInfo()

	fmt.Printf("Name: %v\n", charName)
	fmt.Printf("Race: %v\n", charRace)
	fmt.Printf("Subrace: %v\n", charSubRace)
	fmt.Printf("Class: %v\n", charClass)

	rolledScores := generateAbilityScores()

	if charClass == "Barbarian" {
		strength := rolledScores[len(rolledScores)-1]
		constitution := rolledScores[len(rolledScores)-2]
		wisdom := rolledScores[len(rolledScores)-3]
		dexterity := rolledScores[len(rolledScores)-4]
		charisma := rolledScores[len(rolledScores)-5]
		intelligence := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Bard" {
		charisma := rolledScores[len(rolledScores)-1]
		dexterity := rolledScores[len(rolledScores)-2]
		constitution := rolledScores[len(rolledScores)-3]
		intelligence := rolledScores[len(rolledScores)-4]
		wisdom := rolledScores[len(rolledScores)-5]
		strength := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Cleric" {
		wisdom := rolledScores[len(rolledScores)-1]
		strength := rolledScores[len(rolledScores)-2]
		constitution := rolledScores[len(rolledScores)-3]
		dexterity := rolledScores[len(rolledScores)-4]
		intelligence := rolledScores[len(rolledScores)-5]
		charisma := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Druid" {
		wisdom := rolledScores[len(rolledScores)-1]
		constitution := rolledScores[len(rolledScores)-2]
		dexterity := rolledScores[len(rolledScores)-3]
		charisma := rolledScores[len(rolledScores)-4]
		intelligence := rolledScores[len(rolledScores)-5]
		strength := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Fighter" {
		strength := rolledScores[len(rolledScores)-1]
		constitution := rolledScores[len(rolledScores)-2]
		dexterity := rolledScores[len(rolledScores)-3]
		wisdom := rolledScores[len(rolledScores)-4]
		charisma := rolledScores[len(rolledScores)-5]
		intelligence := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Monk" {
		dexterity := rolledScores[len(rolledScores)-1]
		wisdom := rolledScores[len(rolledScores)-2]
		constitution := rolledScores[len(rolledScores)-3]
		intelligence := rolledScores[len(rolledScores)-4]
		strength := rolledScores[len(rolledScores)-5]
		charisma := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Paladin" {
		strength := rolledScores[len(rolledScores)-1]
		charisma := rolledScores[len(rolledScores)-2]
		constitution := rolledScores[len(rolledScores)-3]
		wisdom := rolledScores[len(rolledScores)-4]
		intelligence := rolledScores[len(rolledScores)-5]
		dexterity := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Paladin" {
		strength := rolledScores[len(rolledScores)-1]
		charisma := rolledScores[len(rolledScores)-2]
		constitution := rolledScores[len(rolledScores)-3]
		wisdom := rolledScores[len(rolledScores)-4]
		intelligence := rolledScores[len(rolledScores)-5]
		dexterity := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Ranger" {
		dexterity := rolledScores[len(rolledScores)-1]
		wisdom := rolledScores[len(rolledScores)-2]
		strength := rolledScores[len(rolledScores)-3]
		intelligence := rolledScores[len(rolledScores)-4]
		constitution := rolledScores[len(rolledScores)-5]
		charisma := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Rogue" {
		dexterity := rolledScores[len(rolledScores)-1]
		charisma := rolledScores[len(rolledScores)-2]
		intelligence := rolledScores[len(rolledScores)-3]
		wisdom := rolledScores[len(rolledScores)-4]
		constitution := rolledScores[len(rolledScores)-5]
		strength := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Sorcerer" {
		charisma := rolledScores[len(rolledScores)-1]
		constitution := rolledScores[len(rolledScores)-2]
		intelligence := rolledScores[len(rolledScores)-3]
		dexterity := rolledScores[len(rolledScores)-4]
		wisdom := rolledScores[len(rolledScores)-5]
		strength := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Warlock" {
		charisma := rolledScores[len(rolledScores)-1]
		constitution := rolledScores[len(rolledScores)-2]
		dexterity := rolledScores[len(rolledScores)-3]
		strength := rolledScores[len(rolledScores)-4]
		intelligence := rolledScores[len(rolledScores)-5]
		wisdom := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else if charClass == "Wizard" {
		intelligence := rolledScores[len(rolledScores)-1]
		dexterity := rolledScores[len(rolledScores)-2]
		charisma := rolledScores[len(rolledScores)-3]
		constitution := rolledScores[len(rolledScores)-4]
		wisdom := rolledScores[len(rolledScores)-5]
		strength := rolledScores[len(rolledScores)-6]
		abilityScores := []int{strength, dexterity, constitution, intelligence, wisdom, charisma}
		return abilityScores

	} else {
		return []int{0, 0, 0, 0, 0, 0}
	}
}


func main() {
	/*
	scores := finalizeScores()
	fmt.Printf("Strength: %v\n", scores[0])
	fmt.Printf("Dexterity: %v\n", scores[1])
	fmt.Printf("Constitution: %v\n", scores[2])
	fmt.Printf("Intelligence: %v\n", scores[3])
	fmt.Printf("Wisdom: %v\n", scores[4])
	fmt.Printf("Charisma: %v\n", scores[5])
	*/
	fmt.Println("Chat App v0.01")
        setupRoutes()
        http.ListenAndServe(":8080", nil)
}
