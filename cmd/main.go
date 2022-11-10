package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathonrogers/connectwise-go"
)

var client connectwise.CwClient

func main() {
	env := ".local"
	err := godotenv.Load(".env" + env)
	if err != nil {
		log.Fatal("Error loading env file")
	}

	clientid := os.Getenv("CLIENTID")
	baseurl := os.Getenv("BASEURL")
	company := os.Getenv("COMPANY")
	pubkey := os.Getenv("PUBLIC_KEY")
	privkey := os.Getenv("PRIVATE_KEY")

	client, err = connectwise.NewCwClient(baseurl, clientid, company, pubkey, privkey)

	// client, err := connectwise.NewCwClient("cw.peakesupport.com", "d5ea3b84-e53a-4b8f-87be-b114d058ce5d", "peake", "bX0jqBeeS2f8MQwi", "mDjOKCC7XvWj4Toi")

	if err != nil {
		panic(err)
	}

	//connectwise.CwOption{Key: "conditions", Value: "id >= 714526"},
	//fmt.Printf("%s\n", data)
	// d := stringToMap(string(data))
	// fmt.Println(d)
	//decoder.DisallowUnknownFields()
	//json.Unmarshal(data, &tickets)
	//fmt.Printf("%v\n", tickets)

	//getTickets()

	//test, _ := connectwise.GetTicketByID(&client, 600108)
	test, _ := client.TicketsClient.GetTicketByID(600108)
	// what is ticket 601012?
	//fmt.Printf("%v\n\n", test)
	testNotes, _ := test.GetNotes(&client)
	//fmt.Printf("%d notes", len(testNotes))
	fmt.Printf("\n\n%v\n\n%v", test, testNotes)
}

func getTickets() {
	data, err := client.Get("/service/tickets",
		connectwise.CwOption{Key: "orderby", Value: "id desc"},

		connectwise.CwOption{Key: "pagesize", Value: "10"},
	)
	if err != nil {
		panic(err)
	}

	var tickets connectwise.Tickets
	decoder := json.NewDecoder(bytes.NewReader(data))

	err = decoder.Decode(&tickets)
	if err != nil {
		panic(err)
	}

	if len(tickets) > 0 {

		for _, t := range tickets {
			fmt.Println(t.Summary)
			//getNotes(t)
			notes, _ := t.GetNotes(&client)
			fmt.Printf("%d Notes:\n\n%v\n\n", len(notes), notes)
		}
	}

}

/*
func getNotes(t connectwise.Ticket) {
	data, err := client.Get(
		fmt.Sprintf("/service/tickets/%d/notes", t.ID),
	)
	if err != nil {
		panic(err)
	}
	var ticketNotes connectwise.TicketNotes
	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&ticketNotes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d Notes\n\n", len(ticketNotes))
}
*/
