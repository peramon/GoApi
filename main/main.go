package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type dataStudent struct {
	ID                           int    `json:ID`
	University                   string `json:University`
	Course                       string `json:Course`
	Student                      string `json:Student`
	Period                       string `json:Period`
	PreferredProgrammingLanguage string `json:PreferredProgrammingLanguage`
	PostGraduationAspiration     string `json:PostGraduationAspiration`
}

type allDataStudent []dataStudent

var dataS = allDataStudent{
	{
		ID:                           1,
		University:                   "UTPL",
		Course:                       "SOFTWARE ENGINEERING PROCESSES",
		Student:                      "PAUL ENRIQUE RAMON SUQUILANDA",
		Period:                       "April/August - 2021",
		PreferredProgrammingLanguage: "Java",
		PostGraduationAspiration:     "Continue learning especially in development issues and improving as a human being.",
	},
}

/*func indexRoute(w http.ResponseWriter, r *http.Request) {
 */

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", getDataStudent)

	log.Fatal(http.ListenAndServe(":3000", router))

}

func getDataStudent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dataS)
}
