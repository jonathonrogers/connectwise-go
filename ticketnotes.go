package connectwise

// TicketNotes encapsulates a TicketNotes object
type TicketNotes []struct {
	ID                    int    `json:"id"`
	TicketID              int    `json:"ticketId"`
	Text                  string `json:"text"`
	DetailDescriptionFlag bool   `json:"detailDescriptionFlag"`
	InternalAnalysisFlag  bool   `json:"internalAnalysisFlag"`
	ResolutionFlag        bool   `json:"resolutionFlag"`
	IssueFlag             bool   `json:"issueFlag"`
	Contact               struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			ContactHref string `json:"contact_href"`
		} `json:"_info"`
	} `json:"contact,omitempty"`
	DateCreated  time.Time `json:"dateCreated"`
	CreatedBy    string    `json:"createdBy"`
	InternalFlag bool      `json:"internalFlag"`
	ExternalFlag bool      `json:"externalFlag"`
	Info         struct {
		LastUpdated time.Time `json:"lastUpdated"`
		UpdatedBy   string    `json:"updatedBy"`
	} `json:"_info"`
	Member struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string    `json:"member_href"`
			ImageHref  time.Time `json:"image_href"`
		} `json:"_info"`
	} `json:"member,omitempty"`
}
