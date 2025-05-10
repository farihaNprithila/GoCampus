package fourthclass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
		listEventHandler(w, r)
	case http.MethodPost:
		createEvent(w, r)
	case http.MethodDelete:
		deleteEventHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if err := deleteEventValidation(name, eventList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, event := range eventList {
		if event.Name != nil && *event.Name == name {
			eventList = append(eventList[:i], eventList[i+1:]...)
			fmt.Printf("Deleted event with name: %s\n", name)

			resp := map[string]string{"message": "Event deleted successfully"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
}

func deleteEventValidation(name string, eventList []Event) error {
	if name == "" {
		return fmt.Errorf("'name' field is required for deletion")
	}

	for _, event := range eventList {
		if event.Name != nil && *event.Name == name {
			return nil // Valid, event found
		}
	}

	return fmt.Errorf("event with name '%s' not found", name)
}

type Event struct {
	Name *string `json:"name"`
	Type *string `json:"type"`
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

	//input validation
	err := validateEventInput(event, eventList)
	if err != nil {
		return
	}

	eventList = append(eventList, event)

	resp := map[string]interface{}{"Message": "Successfully Created",
		"data": event}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func validateEventInput(event Event, existingEvents []Event) error {
	if event.Name == nil || strings.TrimSpace(*event.Name) == "" {
		return fmt.Errorf("'name' field is required and cannot be null or empty")
	}
	if event.Type == nil || strings.TrimSpace(*event.Type) == "" {
		return fmt.Errorf("'type' field is required and cannot be null or empty")
	}

	for _, e := range existingEvents {
		if e.Name == event.Name {
			return fmt.Errorf("event with name '%s' already exists", event.Name)
		}
	}

	return nil
}

func listEventHandler(w http.ResponseWriter, r *http.Request) {
	// Default values for pagination
	page := 1
	size := 10

	// Get the 'page' and 'size' query parameters
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Parse the 'page' parameter
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			http.Error(w, "Invalid 'page' parameter", http.StatusBadRequest)
			return
		}
	}

	// Parse the 'size' parameter
	if sizeStr != "" {
		var err error
		size, err = strconv.Atoi(sizeStr)
		if err != nil || size <= 0 {
			http.Error(w, "Invalid 'size' parameter", http.StatusBadRequest)
			return
		}
	}

	// Get the paginated events
	paginatedEvents, err := getPaginatedEvents(page, size, eventList)
	if err != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	// Set the response header and encode the paginated events to JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paginatedEvents); err != nil {
		http.Error(w, "Failed to encode events", http.StatusInternalServerError)
		return
	}
}

func getPaginatedEvents(page, size int, eventList []Event) ([]Event, error) {
	start := (page - 1) * size
	end := start + size

	if start > len(eventList) {
		start = len(eventList)
	}
	if end > len(eventList) {
		end = len(eventList)
	}

	return eventList[start:end], nil
}
