package connectwise

type Tickets []struct {
	ID         int    `json:"id"`
	Summary    string `json:"summary"`
	RecordType string `json:"recordType"`
	Board      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			BoardHref string `json:"board_href"`
		} `json:"_info"`
	} `json:"board"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			StatusHref string `json:"status_href"`
		} `json:"_info"`
	} `json:"status"`
	WorkRole struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkRoleHref string `json:"workRole_href"`
		} `json:"_info"`
	} `json:"workRole"`
	WorkType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			WorkTypeHref string `json:"workType_href"`
		} `json:"_info"`
	} `json:"workType"`
	Company struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			CompanyHref string `json:"company_href"`
			MobileGUID  string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"company"`
	Site struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SiteHref   string `json:"site_href"`
			MobileGUID string `json:"mobileGuid"`
		} `json:"_info"`
	} `json:"site"`
	SiteName        string `json:"siteName"`
	AddressLine1    string `json:"addressLine1"`
	City            string `json:"city"`
	StateIdentifier string `json:"stateIdentifier"`
	Zip             string `json:"zip"`
	Country         struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			CountryHref string `json:"country_href"`
		} `json:"_info"`
	} `json:"country"`
	Contact struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			MobileGUID  string `json:"mobileGuid"`
			ContactHref string `json:"contact_href"`
		} `json:"_info"`
	} `json:"contact"`
	ContactName         string `json:"contactName"`
	ContactPhoneNumber  string `json:"contactPhoneNumber"`
	ContactEmailAddress string `json:"contactEmailAddress"`
	Type                struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TypeHref string `json:"type_href"`
		} `json:"_info"`
	} `json:"type"`
	SubType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SubTypeHref string `json:"subType_href"`
		} `json:"_info"`
	} `json:"subType"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			TeamHref string `json:"team_href"`
		} `json:"_info"`
	} `json:"team"`
	Owner struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			MemberHref string    `json:"member_href"`
			ImageHref  time.Time `json:"image_href"`
		} `json:"_info"`
	} `json:"owner"`
	Priority struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Sort int    `json:"sort"`
		Info struct {
			PriorityHref string    `json:"priority_href"`
			ImageHref    time.Time `json:"image_href"`
		} `json:"_info"`
	} `json:"priority"`
	ServiceLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"serviceLocation"`
	Source struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SourceHref string `json:"source_href"`
		} `json:"_info"`
	} `json:"source"`
	Agreement struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			AgreementHref string `json:"agreement_href"`
		} `json:"_info"`
	} `json:"agreement"`
	Severity                   string    `json:"severity"`
	Impact                     string    `json:"impact"`
	AllowAllClientsPortalView  bool      `json:"allowAllClientsPortalView"`
	CustomerUpdatedFlag        bool      `json:"customerUpdatedFlag"`
	AutomaticEmailContactFlag  bool      `json:"automaticEmailContactFlag"`
	AutomaticEmailResourceFlag bool      `json:"automaticEmailResourceFlag"`
	AutomaticEmailCcFlag       bool      `json:"automaticEmailCcFlag"`
	ClosedFlag                 bool      `json:"closedFlag"`
	ActualHours                float64   `json:"actualHours"`
	Approved                   bool      `json:"approved"`
	EstimatedExpenseCost       float64   `json:"estimatedExpenseCost"`
	EstimatedExpenseRevenue    float64   `json:"estimatedExpenseRevenue"`
	EstimatedProductCost       float64   `json:"estimatedProductCost"`
	EstimatedProductRevenue    float64   `json:"estimatedProductRevenue"`
	EstimatedTimeCost          float64   `json:"estimatedTimeCost"`
	EstimatedTimeRevenue       float64   `json:"estimatedTimeRevenue"`
	BillingMethod              string    `json:"billingMethod"`
	SubBillingMethod           string    `json:"subBillingMethod"`
	DateResponded              time.Time `json:"dateResponded"`
	ResolveMinutes             int       `json:"resolveMinutes"`
	ResPlanMinutes             int       `json:"resPlanMinutes"`
	RespondMinutes             int       `json:"respondMinutes"`
	IsInSLA                    bool      `json:"isInSla"`
	Resources                  string    `json:"resources"`
	HasChildTicket             bool      `json:"hasChildTicket"`
	HasMergedChildTicketFlag   bool      `json:"hasMergedChildTicketFlag"`
	BillTime                   string    `json:"billTime"`
	BillExpenses               string    `json:"billExpenses"`
	BillProducts               string    `json:"billProducts"`
	Location                   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			LocationHref string `json:"location_href"`
		} `json:"_info"`
	} `json:"location"`
	Department struct {
		ID         int    `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		Info       struct {
			DepartmentHref string `json:"department_href"`
		} `json:"_info"`
	} `json:"department"`
	MobileGUID string `json:"mobileGuid"`
	SLA        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Info struct {
			SLAHref string `json:"sla_href"`
		} `json:"_info"`
	} `json:"sla"`
	SLAStatus            string `json:"slaStatus"`
	RequestForChangeFlag bool   `json:"requestForChangeFlag"`
	Currency             struct {
		ID                      int    `json:"id"`
		Symbol                  string `json:"symbol"`
		CurrencyCode            string `json:"currencyCode"`
		DecimalSeparator        string `json:"decimalSeparator"`
		NumberOfDecimals        int    `json:"numberOfDecimals"`
		ThousandsSeparator      string `json:"thousandsSeparator"`
		NegativeParenthesesFlag bool   `json:"negativeParenthesesFlag"`
		DisplaySymbolFlag       bool   `json:"displaySymbolFlag"`
		CurrencyIdentifier      string `json:"currencyIdentifier"`
		DisplayIDFlag           bool   `json:"displayIdFlag"`
		RightAlign              bool   `json:"rightAlign"`
		Name                    string `json:"name"`
		Info                    struct {
			CurrencyHref string `json:"currency_href"`
		} `json:"_info"`
	} `json:"currency"`
	Info struct {
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
	} `json:"_info"`
	EscalationStartDateUTC  time.Time `json:"escalationStartDateUTC"`
	EscalationLevel         int       `json:"escalationLevel"`
	MinutesBeforeWaiting    int       `json:"minutesBeforeWaiting"`
	RespondedSkippedMinutes int       `json:"respondedSkippedMinutes"`
	ResplanSkippedMinutes   int       `json:"resplanSkippedMinutes"`
}
