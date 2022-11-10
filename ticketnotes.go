package connectwise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	stripmd "github.com/writeas/go-strip-markdown"
)

// TicketNotes holds an array of TicketNote structs
type TicketNotes []TicketNote

// TicketNote encapsulates a TicketNotes object
type TicketNote struct {
	ID                    int             `json:"id"`
	TicketID              int             `json:"ticketId"`
	Text                  string          `json:"text"`
	DetailDescriptionFlag bool            `json:"detailDescriptionFlag"`
	InternalAnalysisFlag  bool            `json:"internalAnalysisFlag"`
	ResolutionFlag        bool            `json:"resolutionFlag"`
	IssueFlag             bool            `json:"issueFlag"`
	Contact               Contact         `json:"contact,omitempty"`
	DateCreated           time.Time       `json:"dateCreated"`
	CreatedBy             string          `json:"createdBy"`
	InternalFlag          bool            `json:"internalFlag"`
	ExternalFlag          bool            `json:"externalFlag"`
	Info                  TicketNotesInfo `json:"_info"`
	Member                Member          `json:"member,omitempty"`
}

// type ContactInfo struct {
// 	ContactHref string `json:"contact_href"`
// }
// type Contact struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Info ContactInfo   `json:"_info"`
// }
type TicketNotesInfo struct {
	LastUpdated time.Time `json:"lastUpdated"`
	UpdatedBy   string    `json:"updatedBy"`
}
type MemberInfo struct {
	MemberHref string `json:"member_href"`
	ImageHref  string `json:"image_href"`
}
type Member struct {
	ID         int        `json:"id"`
	Identifier string     `json:"identifier"`
	Name       string     `json:"name"`
	Info       MemberInfo `json:"_info"`
}

func (notes TicketNotes) String() string {
	strNotes := ""
	for _, n := range notes {
		strNotes += fmt.Sprintf("{\n\ttext: \"%v\",\n\tcreatedby: \"%v\",\n\tdatecreated: \"%v\"\n},\n", stripmd.Strip(n.Text), n.CreatedBy, n.DateCreated)
	}
	return strNotes
}

func (t *Ticket) GetNotes(client *CwClient) (notes TicketNotes, err error) {
	data, err := client.Get(
		fmt.Sprintf("/service/tickets/%d/notes", t.ID),
	)
	fmt.Printf("Getting notes for %d\n", t.ID)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&notes)
	if err != nil {
		panic(err)
	}
	return notes, err
}
