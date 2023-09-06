package http

import (
	"dev11/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func createEvent(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		inputDate := r.FormValue("date")
		userId := r.FormValue("user_id")
		id, err := strconv.Atoi(userId)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		date, ok := validateDate(inputDate)
		if !ok {
			err := writeToJson(w, http.StatusBadRequest, false, "not a valid format date")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		input := models.Event{
			Name:   name,
			Date:   date,
			UserId: id,
		}

		if err := service.CreateEvent(input); err != nil {
			err := writeToJson(w, http.StatusInternalServerError, false, err.Error())
			if err != nil {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
			return
		}

		if err := writeToJson(w, http.StatusCreated, true, "created"); err != nil {
			http.Error(w, "failed to return JSON answer", http.StatusInternalServerError)
		}
	}
}

func updateEvent(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		newName := r.FormValue("new_name")
		newDate := r.FormValue("new_date")
		oldName := r.FormValue("old_name")

		date, ok := validateDate(newDate)
		if !ok {
			err := writeToJson(w, http.StatusBadRequest, false, "Not a valid format date")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		input := models.Event{
			UserId: userID,
			Name:   newName,
			Date:   date,
		}

		if err := service.UpdateEvent(input, oldName); err != nil {
			err := writeToJson(w, http.StatusInternalServerError, false, err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err := writeToJson(w, http.StatusOK, true, "updated"); err != nil {
			http.Error(w, "failed to return JSON answer", http.StatusInternalServerError)
		}
	}
}

func deleteEvent(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")

		input := models.Event{
			Name:   name,
			UserId: userID,
		}

		if err := service.DeleteEvent(input); err != nil {
			err := writeToJson(w, http.StatusInternalServerError, false, err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err := writeToJson(w, http.StatusOK, true, "deleted"); err != nil {
			http.Error(w, "failed to return JSON answer", http.StatusInternalServerError)
		}
	}
}

func getEventForDay(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}
		queryParams := r.URL.Query()
		queryDate := queryParams.Get("date")
		userID := queryParams.Get("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		date, ok := validateDate(queryDate)
		if !ok {
			err := writeToJson(w, http.StatusBadRequest, false, "not a valid format date")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		input := models.Event{
			Date:   date,
			UserId: id,
		}

		events, err := service.GetEventForDay(input)
		if err != nil {
			err := writeToJson(w, http.StatusInternalServerError, false, err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err := writeToJson(w, http.StatusOK, true, events); err != nil {
			http.Error(w, "failed to return JSON answer", http.StatusInternalServerError)
		}
	}
}

func getEventForWeek(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "wrong method", http.StatusMethodNotAllowed)
			return
		}
		queryParams := r.URL.Query()
		queryDate := queryParams.Get("date")
		userID := queryParams.Get("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			log.Fatal(err)
			return
		}

		date, ok := validateDate(queryDate)
		if !ok {
			err := writeToJson(w, http.StatusBadRequest, false, "not a valid format date")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		input := models.Event{
			Date:   date,
			UserId: id,
		}

		events, err := service.GetEventForWeek(input)
		if err != nil {
			err := writeToJson(w, http.StatusInternalServerError, false, err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		if err := writeToJson(w, http.StatusOK, true, events); err != nil {
			http.Error(w, "failed to return JSON answer", http.StatusInternalServerError)
		}
	}
}

func getEventForMonth(service service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "wrong method", http.StatusMethodNotAllowed)
			return
		}
		queryParams := r.URL.Query()
		queryDate := queryParams.Get("date")
		userID := queryParams.Get("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			log.Fatal(err)
			return
		}

		date, ok := validateDate(queryDate)
		if !ok {
			err := writeToJson(w, http.StatusBadRequest, false, "not a valid format date")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		input := models.Event{
			Date:   date,
			UserId: id,
		}

		events, err := service.GetEventForMonth(input)
		if err != nil {
			err := writeToJson(w, http.StatusInternalServerError, false, err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		if err := writeToJson(w, http.StatusOK, true, events); err != nil {
			http.Error(w, "failed to return JSON answer", http.StatusInternalServerError)
		}
	}
}

func writeToJson(w http.ResponseWriter, code int, fl bool, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if fl {
		return json.NewEncoder(w).Encode(models.JsonResult{Result: value})
	} else {
		return json.NewEncoder(w).Encode(models.JsonResult{Error: value})
	}
}

func validateDate(date string) (time.Time, bool) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, false
	}
	return d, true
}
