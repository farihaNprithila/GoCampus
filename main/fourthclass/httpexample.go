package fourthclass

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var eventList []Event

func HttpExample() {

	//mux := http.NewServeMux()
	//mux.HandleFunc("/user", customHandler)

	//h := &CustomHandler{}
	//http.HandleFunc("/", helloHandler)
	//http.HandleFunc("/events", helloHandler)

	//fmt.Println("Server is running on port 8080")
	//http.ListenAndServe(":8080", h)

	http.HandleFunc("/events", eventsHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}

func customHandler(writer http.ResponseWriter, request *http.Request) {

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	r.Method = "GET"
	fmt.Fprint(w, "Hello World")

}

type CustomHandler struct{}

func (c CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listEventHandler(w)
	case http.MethodPost:
		createEvent(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

type Event struct {
	Name string
	Type string
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Print the fields separately for clarity
	fmt.Printf("Name: %s\n", event.Name)
	fmt.Printf("Type: %s\n", event.Type)

	// Or you can print the whole event struct using fmt.Printf for more control over formatting
	fmt.Printf("Event: %+v\n", event)

	eventList = append(eventList, event)

	resp := map[string]interface{}{"Message": "Successfully Created",
		"data": event}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func listEventHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(eventList); err != nil {
		http.Error(w, "Failed to encode events", http.StatusInternalServerError)
		return
	}
}
