package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

// конфиг логин и пароль
func main() {
	var err error
	db, err = sql.Open("mysql", "user:123456@tcp(mysql:3306)/api-db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := Router()

	log.Fatal(http.ListenAndServe(":8080", r))

}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/segment", CreateSegment).Methods("POST")
	router.HandleFunc("/segment", DeleteSegment).Methods("DELETE")
	router.HandleFunc("/user/segment/manage", ManageUserSegment).Methods("PUT", "DELETE")
	router.HandleFunc("/user/segment", GetUsersSegments).Methods("GET")

	return router
}

type Segment struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
}

type UserSegment struct {
	UserID      int    `json:"user_id"`
	SegmentID   int    `json:"segment_id"`
	SegmentSlug string `json:"segment_slug"`
}

func CreateSegment(w http.ResponseWriter, r *http.Request) {
	var segment Segment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM segments WHERE slug = ?", segment.Slug).Scan(&count)

	if count != 0 {
		http.Error(w, "Segment is already exists", http.StatusNotFound)
		return
	}

	_, err = db.Exec("INSERT INTO segments (slug) VALUES (?)", segment.Slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Segment created")

	w.WriteHeader(http.StatusCreated)
}

func DeleteSegment(w http.ResponseWriter, r *http.Request) {
	var segment Segment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM segments WHERE slug = ?", segment.Slug).Scan(&count)

	if count == 0 {
		http.Error(w, "No such segment exists", http.StatusNotFound)
		return
	}

	_, err = db.Exec("DELETE FROM segments WHERE slug = ?", segment.Slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Segment removed")
}

func ManageUserSegment(w http.ResponseWriter, r *http.Request) {

	var userSegment []UserSegment
	err := json.NewDecoder(r.Body).Decode(&userSegment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	switch r.Method {
	case "PUT":
		for i := 0; i < len(userSegment); i++ {
			var count int
			err = db.QueryRow("SELECT COUNT(*) FROM segments WHERE slug = ?", userSegment[i].SegmentSlug).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if count == 0 {
				http.Error(w, "Segment not found", http.StatusNotFound)
				return
			}

			err = db.QueryRow("SELECT ID FROM segments WHERE slug = ?", userSegment[i].SegmentSlug).Scan(&userSegment[i].SegmentID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			count = 0
			err = db.QueryRow("SELECT COUNT(*) FROM user_segments WHERE user_id = ? AND segment_id =  ?", userSegment[i].UserID, userSegment[i].SegmentID).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if count > 1 {
				http.Error(w, "User already in this segment", http.StatusNotFound)
				return
			}

			_, err = db.Exec("INSERT INTO user_segments (user_id, segment_id) VALUES (?, ?)", userSegment[i].UserID, userSegment[i].SegmentID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		fmt.Fprint(w, "User/Users added to segment")

	case "DELETE":
		for i := 0; i < len(userSegment); i++ {
			var count int
			err = db.QueryRow("SELECT COUNT(*) FROM segments WHERE slug = ?", userSegment[i].SegmentSlug).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if count == 0 {
				http.Error(w, "Segment not found", http.StatusNotFound)
				return
			}
			count = 0
			err = db.QueryRow("SELECT COUNT(*) FROM user_segments WHERE user_id = ?", userSegment[i].UserID).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if count == 0 {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			err = db.QueryRow("SELECT ID FROM segments WHERE slug = ?", userSegment[i].SegmentSlug).Scan(&userSegment[i].SegmentID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = db.Exec("DELETE FROM user_segments WHERE user_id = ? AND segment_id = ?", userSegment[i].UserID, userSegment[i].SegmentID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		fmt.Fprint(w, "User/Users removed from segment")
	}
}

func GetUsersSegments(w http.ResponseWriter, r *http.Request) {
	var userSegment UserSegment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userSegment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM user_segments WHERE user_id = ?", userSegment.UserID).Scan(&count)
	if count == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	rows, err := db.Query("SELECT segment_id FROM user_segments WHERE user_id = ?", userSegment.UserID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var userSegment UserSegment
		err = rows.Scan(&userSegment.SegmentID)
		err = db.QueryRow("SELECT slug FROM segments WHERE id = ?", userSegment.SegmentID).Scan(&userSegment.SegmentSlug)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(userSegment.SegmentSlug)
	}
}
