package main

// ZenDeskResponse ...
type ZenDeskResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	Scope       string `json:"scope"`
}

// CustomFields ...
type CustomFields struct {
	// TicketField []interface{} `json:"ticket,omitempty"`
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// Comment ...
type Comment struct {
	Body string `json:"body"`
}

// ReferenceStruct ...
type ReferenceStruct struct {
	// Subject string  `json:"subject"`
	Comment      Comment        `json:"comment"`
	Status       string         `json:"status"`
	CustomFields []CustomFields `json:"custom_fields"`
}

// CreateTicket ...
type CreateTicket struct {
	Subject string  `json:"subject"`
	Comment Comment `json:"comment"`
}

// UpdateCustomField ...
type UpdateCustomField struct {
	Status       string         `json:"status"`
	Comment      Comment        `json:"comment"`
	CustomFields []CustomFields `json:"custom_fields"`
}

// ResponseDataUpdateField ...
type ResponseDataUpdateField struct {
	UpdateCustomField UpdateCustomField `json:"ticket"`
}

// ResponseData ...
type ResponseData struct {
	CreateTicket CreateTicket `json:"ticket"`
}
