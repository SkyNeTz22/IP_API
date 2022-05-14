package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMedics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Medici"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Medici
	for rows.Next() {
		var elem Medici
		if err := rows.Scan(&elem.IDMedic, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
			w.Write([]byte("A intrat pe if"))
		}
		elements = append(elements, elem)
	}
	if err = rows.Err(); err != nil {
		w.Write([]byte("A intrat pe if"))
	}

	encoded, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	w.Write(encoded)
}

func GetMedicByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := mux.Vars(r)["username"]
	//selectStringUsername := ("SELECT * FROM medassist_db.Medici WHERE Username = \"" + bUsername + "\"")
	selectStringUsername := fmt.Sprintf("SELECT * FROM medassist_db.Medici WHERE Username = '%s'", bUsername)
	rows, err := db.Query(selectStringUsername)
	if err != nil {
		panic(err)
	}
	var elements []Medici
	for rows.Next() {
		var elem Medici
		if err := rows.Scan(&elem.IDMedic, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
			w.Write([]byte("A intrat pe if"))
		}
		elements = append(elements, elem)
	}
	if err = rows.Err(); err != nil {
		w.Write([]byte("A intrat pe if"))
	}

	encoded, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	w.Write(encoded)
}

func GetMedicById(w http.ResponseWriter, r *http.Request) {
	bID := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.Medici WHERE IDMedic = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Medici
	for rows.Next() {
		var elem Medici
		if err := rows.Scan(&elem.IDMedic, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
			w.Write([]byte("A intrat pe if"))
		}
		elements = append(elements, elem)
	}
	if err = rows.Err(); err != nil {
		w.Write([]byte("A intrat pe if"))
	}

	encoded, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	w.Write(encoded)
}

func CreateMedic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bUsername := r.FormValue("Username")
	insertStringMedici := fmt.Sprintf("INSERT INTO medassist_db.Medici (`Nume`, `Prenume`, `Username`) VALUES ('%s', '%s', '%s')", bNume, bPrenume, bUsername)
	_, err := db.Exec(insertStringMedici)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bNume, bPrenume, bUsername)
	}
}

func UpdateMedic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	updateID := mux.Vars(r)["id"]
	parameters := r.URL.Query()
	bID := parameters.Get("IDMedic")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	bUsername := parameters.Get("Username")
	updateStringMedici := fmt.Sprintf("UPDATE medassist_db.Medici SET `IDMedic` = '%s', `Nume` = '%s', `Prenume` = '%s', `Username` = '%s' WHERE `IDMedic` = '%s'", bID, bNume, bPrenume, bUsername, updateID)
	_, err := db.Exec(updateStringMedici)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bNume, bPrenume, bUsername)
	}
}

func DeleteMedic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringMedici := fmt.Sprintf("DELETE FROM medassist_db.Medici WHERE `IDMedic` = '%s'", bID)
	_, err := db.Exec(deleteStringMedici)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
