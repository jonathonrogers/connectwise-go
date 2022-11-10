package connectwise

import (
	"encoding/json"
	"fmt"
	"time"
)

// Tickets holds an array of Ticket structs
type Tickets []Ticket

// Ticket encapsulates a Tickets object returned from the Manage API
type Ticket struct {
	ID                         int             `json:"id"`
	Summary                    string          `json:"summary"`
	RecordType                 string          `json:"recordType"`
	Board                      Board           `json:"board"`
	Status                     Status          `json:"status"`
	WorkRole                   WorkRole        `json:"workRole"`
	WorkType                   WorkType        `json:"workType"`
	Company                    Company         `json:"company"`
	Site                       Site            `json:"site"`
	SiteName                   string          `json:"siteName"`
	AddressLine1               string          `json:"addressLine1"`
	City                       string          `json:"city"`
	StateIdentifier            string          `json:"stateIdentifier"`
	Zip                        string          `json:"zip"`
	Country                    Country         `json:"country"`
	Contact                    Contact         `json:"contact"`
	ContactName                string          `json:"contactName"`
	ContactPhoneNumber         string          `json:"contactPhoneNumber"`
	ContactEmailAddress        string          `json:"contactEmailAddress"`
	Type                       Type            `json:"type"`
	SubType                    SubType         `json:"subType"`
	Team                       Team            `json:"team"`
	Owner                      Owner           `json:"owner"`
	Priority                   Priority        `json:"priority"`
	ServiceLocation            ServiceLocation `json:"serviceLocation"`
	Source                     Source          `json:"source"`
	Agreement                  Agreement       `json:"agreement"`
	Severity                   string          `json:"severity"`
	Impact                     string          `json:"impact"`
	AllowAllClientsPortalView  bool            `json:"allowAllClientsPortalView"`
	CustomerUpdatedFlag        bool            `json:"customerUpdatedFlag"`
	AutomaticEmailContactFlag  bool            `json:"automaticEmailContactFlag"`
	AutomaticEmailResourceFlag bool            `json:"automaticEmailResourceFlag"`
	AutomaticEmailCcFlag       bool            `json:"automaticEmailCcFlag"`
	ClosedFlag                 bool            `json:"closedFlag"`
	ActualHours                float64         `json:"actualHours"`
	Approved                   bool            `json:"approved"`
	EstimatedExpenseCost       float64         `json:"estimatedExpenseCost"`
	EstimatedExpenseRevenue    float64         `json:"estimatedExpenseRevenue"`
	EstimatedProductCost       float64         `json:"estimatedProductCost"`
	EstimatedProductRevenue    float64         `json:"estimatedProductRevenue"`
	EstimatedTimeCost          float64         `json:"estimatedTimeCost"`
	EstimatedTimeRevenue       float64         `json:"estimatedTimeRevenue"`
	BillingMethod              string          `json:"billingMethod"`
	SubBillingMethod           string          `json:"subBillingMethod"`
	DateResponded              time.Time       `json:"dateResponded"`
	ResolveMinutes             int             `json:"resolveMinutes"`
	ResPlanMinutes             int             `json:"resPlanMinutes"`
	RespondMinutes             int             `json:"respondMinutes"`
	IsInSLA                    bool            `json:"isInSla"`
	Resources                  string          `json:"resources"`
	HasChildTicket             bool            `json:"hasChildTicket"`
	HasMergedChildTicketFlag   bool            `json:"hasMergedChildTicketFlag"`
	BillTime                   string          `json:"billTime"`
	BillExpenses               string          `json:"billExpenses"`
	BillProducts               string          `json:"billProducts"`
	Location                   Location        `json:"location"`
	Department                 Department      `json:"department"`
	MobileGUID                 string          `json:"mobileGuid"`
	SLA                        SLA             `json:"sla"`
	SLAStatus                  string          `json:"slaStatus"`
	RequestForChangeFlag       bool            `json:"requestForChangeFlag"`
	Currency                   Currency        `json:"currency"`
	Info                       TicketInfo      `json:"_info"`
	EscalationStartDateUTC     time.Time       `json:"escalationStartDateUTC"`
	EscalationLevel            int             `json:"escalationLevel"`
	MinutesBeforeWaiting       int             `json:"minutesBeforeWaiting"`
	RespondedSkippedMinutes    int             `json:"respondedSkippedMinutes"`
	ResplanSkippedMinutes      int             `json:"resplanSkippedMinutes"`
	ClosedDate                 time.Time       `json:"closedDate,omitempty"`
	ClosedBy                   string          `json:"closedBy,omitempty"`
	DateResolved               time.Time       `json:"dateResolved,omitempty"`
	DateResplan                time.Time       `json:"dateResplan,omitempty"`
}
type BoardInfo struct {
	BoardHref string `json:"board_href"`
}
type Board struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Info BoardInfo `json:"_info"`
}
type StatusInfo struct {
	StatusHref string `json:"status_href"`
}
type Status struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Info StatusInfo `json:"_info"`
}
type WorkRoleInfo struct {
	WorkRoleHref string `json:"workRole_href"`
}
type WorkRole struct {
	ID   int          `json:"id"`
	Name string       `json:"name"`
	Info WorkRoleInfo `json:"_info"`
}
type WorkTypeInfo struct {
	WorkTypeHref string `json:"workType_href"`
}
type WorkType struct {
	ID   int          `json:"id"`
	Name string       `json:"name"`
	Info WorkTypeInfo `json:"_info"`
}
type CompanyInfo struct {
	CompanyHref string `json:"company_href"`
	MobileGUID  string `json:"mobileGuid"`
}
type Company struct {
	ID         int         `json:"id"`
	Identifier string      `json:"identifier"`
	Name       string      `json:"name"`
	Info       CompanyInfo `json:"_info"`
}
type SiteInfo struct {
	SiteHref   string `json:"site_href"`
	MobileGUID string `json:"mobileGuid"`
}
type Site struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Info SiteInfo `json:"_info"`
}
type CountryInfo struct {
	CountryHref string `json:"country_href"`
}
type Country struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Info CountryInfo `json:"_info"`
}
type ContactInfo struct {
	MobileGUID  string `json:"mobileGuid"`
	ContactHref string `json:"contact_href"`
}
type Contact struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Info ContactInfo `json:"_info"`
}
type TypeInfo struct {
	TypeHref string `json:"type_href"`
}
type Type struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Info TypeInfo `json:"_info"`
}
type SubTypeInfo struct {
	SubTypeHref string `json:"subType_href"`
}
type SubType struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Info SubTypeInfo `json:"_info"`
}
type TeamInfo struct {
	TeamHref string `json:"team_href"`
}
type Team struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Info TeamInfo `json:"_info"`
}
type OwnerInfo struct {
	MemberHref string `json:"member_href"`
	ImageHref  string `json:"image_href"`
}
type Owner struct {
	ID         int       `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
	Info       OwnerInfo `json:"_info"`
}
type PriorityInfo struct {
	PriorityHref string `json:"priority_href"`
	ImageHref    string `json:"image_href"`
}
type Priority struct {
	ID   int          `json:"id"`
	Name string       `json:"name"`
	Sort int          `json:"sort"`
	Info PriorityInfo `json:"_info"`
}
type ServiceLocationInfo struct {
	LocationHref string `json:"location_href"`
}
type ServiceLocation struct {
	ID   int                 `json:"id"`
	Name string              `json:"name"`
	Info ServiceLocationInfo `json:"_info"`
}
type SourceInfo struct {
	SourceHref string `json:"source_href"`
}
type Source struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Info SourceInfo `json:"_info"`
}
type AgreementInfo struct {
	AgreementHref string `json:"agreement_href"`
}
type Agreement struct {
	ID   int           `json:"id"`
	Name string        `json:"name"`
	Info AgreementInfo `json:"_info"`
}
type LocationInfo struct {
	LocationHref string `json:"location_href"`
}
type Location struct {
	ID   int          `json:"id"`
	Name string       `json:"name"`
	Info LocationInfo `json:"_info"`
}
type DepartmentInfo struct {
	DepartmentHref string `json:"department_href"`
}
type Department struct {
	ID         int            `json:"id"`
	Identifier string         `json:"identifier"`
	Name       string         `json:"name"`
	Info       DepartmentInfo `json:"_info"`
}
type SLAInfo struct {
	SLAHref string `json:"sla_href"`
}
type SLA struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Info SLAInfo `json:"_info"`
}
type CurrencyInfo struct {
	CurrencyHref string `json:"currency_href"`
}
type Currency struct {
	ID                      int          `json:"id"`
	Symbol                  string       `json:"symbol"`
	CurrencyCode            string       `json:"currencyCode"`
	DecimalSeparator        string       `json:"decimalSeparator"`
	NumberOfDecimals        int          `json:"numberOfDecimals"`
	ThousandsSeparator      string       `json:"thousandsSeparator"`
	NegativeParenthesesFlag bool         `json:"negativeParenthesesFlag"`
	DisplaySymbolFlag       bool         `json:"displaySymbolFlag"`
	CurrencyIdentifier      string       `json:"currencyIdentifier"`
	DisplayIDFlag           bool         `json:"displayIdFlag"`
	RightAlign              bool         `json:"rightAlign"`
	Name                    string       `json:"name"`
	Info                    CurrencyInfo `json:"_info"`
}
type TicketInfo struct {
	LastUpdated         time.Time `json:"lastUpdated"`
	UpdatedBy           string    `json:"updatedBy"`
	DateEntered         time.Time `json:"dateEntered"`
	EnteredBy           string    `json:"enteredBy"`
	ActivitiesHref      string    `json:"activities_href"`
	ScheduleentriesHref string    `json:"scheduleentries_href"`
	DocumentsHref       string    `json:"documents_href"`
	ConfigurationsHref  string    `json:"configurations_href"`
	TasksHref           string    `json:"tasks_href"`
	NotesHref           string    `json:"notes_href"`
	ProductsHref        string    `json:"products_href"`
	TimeentriesHref     string    `json:"timeentries_href"`
	ExpenseEntriesHref  string    `json:"expenseEntries_href"`
}

func (tickets Tickets) String() string {
	strTickets := ""
	for _, t := range tickets {
		strTickets += t.String() + "\n"
	}
	return strTickets
}

func (ticket Ticket) String() string {
	strTicket := ""
	if ticket.ID > 0 {
		strTicket += fmt.Sprintf("%d - %s", ticket.ID, ticket.Summary)
	}
	return strTicket
}

const path = "/service/tickets"

func GetTicketByID(client *CwClient, id int) (t Ticket, err error) {
	data, err := client.Get(
		path,
		CwOption{Key: "conditions", Value: fmt.Sprintf("id=%d", id)},
	)
	tickets := decode(data)
	return tickets[0], err
}

func decode(data []byte) Tickets {
	var tickets Tickets
	json.Unmarshal(data, &tickets)
	return tickets
}
