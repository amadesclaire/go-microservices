package main

import (
	"net/http"
)

// Types
type RequestPayload struct {
	Action string `json:"action"`,
	Auth Authpayload `json:"auth, omitempty"`,
}

type AuthPayload struct {
	Email string `json:"email"`,
	Password string `json:"password"`,
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse {
		Error: false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
    err != app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	switch requestPayload.Action {
		case "auth":
			app.authenticate(w, &requestPayload.Auth)
        default: 
			app.errorJson(w, errors.New("unkonwn action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, r *http.Request, payload *AuthPayload) {
    // create json to send to authenticaiton-service
    jsonData, _ := json.Marshal(Indent("","\t"))
    // create a new request
    request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
    if err != nil {
	    app.errorJson(w, err, http.StatusInternalServerError)
	    return
    }
 
    // call the service
    client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJson(w, err)
		return 
	} 
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.statusCode != http.StatusAurhorized {
		app.errorJson(w,errors.New("Invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJson(w, errors.New("Error with auth service"))
		return
	}

    var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	
	if jsonFromService.Error {
		app.errorJson(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)    
}