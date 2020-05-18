package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

func zendesk() {
	log.Println("Request Token.....")

	requestData := url.Values{}
	requestData.Set("grant_type", "password")
	requestData.Set("client_id", clientID)
	requestData.Set("client_secret", secretKey)
	requestData.Set("username", "USERNAME")
	requestData.Set("password", "PASSWORD")
	requestData.Set("scope", "read write")

	response, err := http.Post(subDomain+"/oauth/tokens", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(requestData.Encode())))
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer response.Body.Close()
	log.Println(response.StatusCode)

	if response.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(string(b))
	}

	var zenDeskResponse ZenDeskResponse
	if err := json.NewDecoder(response.Body).Decode(&zenDeskResponse); err != nil {
		return
	}

	log.Println(zenDeskResponse)

	// List Tickets ...

	// List Ticket Fields ...

	request, err := http.NewRequest("GET", subDomain+"/api/v2/ticket_fields.json", nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	request.Header.Set("Authorization", zenDeskResponse.TokenType+" "+zenDeskResponse.AccessToken)
	request.Header.Set("Content-Type", "application/json")

	var client http.Client

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// log.Println(string(b))

	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println(err.Error())
	}
	data["ticket_fields"].([]interface{})[7].(map[string]interface{})["case_id"] = ""
	b, err := json.MarshalIndent(data["ticket_fields"].([]interface{})[7], "", "\t")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(string(b))
	// log.Println(data["ticket_fields"].([]interface{}))

	for _, v := range data["ticket_fields"].([]interface{}) {
		log.Println(v.(map[string]interface{})["title"])

		switch v.(map[string]interface{})["title"] {
		case "Subject":
			subjectURL := v.(map[string]interface{})["url"]
			log.Println(subjectURL)
		case "Description":
			descriptionURL := v.(map[string]interface{})["url"]
			log.Println(descriptionURL)
		case "Status":
			statusURL := v.(map[string]interface{})["url"]
			log.Println(statusURL)
		case "Type":
			typeURL := v.(map[string]interface{})["url"]
			log.Println(typeURL)
		case "Priority":
			priorityURL := v.(map[string]interface{})["url"]
			log.Println(priorityURL)
		case "Group":
			groupURL := v.(map[string]interface{})["url"]
			log.Println(groupURL)
		case "Assignee":
			assigneeURL := v.(map[string]interface{})["url"]
			log.Println(assigneeURL)
		case "Case Id":
			log.Println(v)
			caseIDURL := v.(map[string]interface{})["url"].(string)
			log.Println(caseIDURL)
			log.Println(reflect.TypeOf(v.(map[string]interface{})["case_id"]))

			var responseData ResponseData
			responseData.CreateTicket.Subject = "Updating case Id Field"
			responseData.CreateTicket.Comment.Body = "Creating a new ticket"
			data, err := json.Marshal(responseData)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(string(data))

			request, err := http.NewRequest("POST", subDomain+"/api/v2/tickets.json", bytes.NewBuffer(data))
			if err != nil {
				log.Println(err.Error())
			}

			request.Header.Set("Authorization", zenDeskResponse.TokenType+" "+zenDeskResponse.AccessToken)
			request.Header.Set("Content-Type", "application/json")

			var client http.Client
			response, err := client.Do(request)
			if err != nil {
				log.Println(err.Error())
				return
			}

			defer response.Body.Close()

			log.Println(response.StatusCode)
			if response.StatusCode != http.StatusCreated {
				b, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(string(b))
			}

			b, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println(err.Error())
			}

			log.Println(string(b))

			// Updating Custom Fields for the Created Ticket

			var customFields CustomFields
			customFields.ID = 360023867013
			customFields.Value = "14952173"

			var responseDataUpdateField ResponseDataUpdateField

			responseDataUpdateField.UpdateCustomField.Status = "solved"
			responseDataUpdateField.UpdateCustomField.Comment.Body = "Update Case Id field with Case Number"
			responseDataUpdateField.UpdateCustomField.CustomFields = append(responseDataUpdateField.UpdateCustomField.CustomFields, customFields)

			customData, err := json.Marshal(responseDataUpdateField)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(string(customData))
			req, err := http.NewRequest("PUT", "https://celebal.zendesk.com/api/v2/tickets/8.json", bytes.NewBuffer(customData))
			if err != nil {
				log.Println(err.Error())
				return
			}

			req.Header.Set("Authorization", zenDeskResponse.TokenType+" "+zenDeskResponse.AccessToken)
			req.Header.Set("Content-Type", "application/json")

			var cli http.Client
			resp, err := cli.Do(req)
			if err != nil {
				log.Println(err.Error())
				return
			}

			defer resp.Body.Close()

			log.Println(resp.StatusCode)
			if resp.StatusCode != http.StatusOK {
				z, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(string(z))
			}

			z, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err.Error())
			}

			log.Println(string(z))

		default:
			break
		}
	}
}
