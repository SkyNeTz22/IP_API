package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCaretakers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Ingrijitori"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Ingrijitori
	for rows.Next() {
		var elem Ingrijitori
		if err := rows.Scan(&elem.IDIngrijitor, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
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

func GetCaretakerByID(w http.ResponseWriter, r *http.Request) {
	bID := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Ingrijitori WHERE IDIngrijitor = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Ingrijitori
	for rows.Next() {
		var elem Ingrijitori
		if err := rows.Scan(&elem.IDIngrijitor, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
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

func GetCaretakerByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := mux.Vars(r)["username"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Ingrijitori WHERE Username = '%s'", bUsername)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Ingrijitori
	for rows.Next() {
		var elem Ingrijitori
		if err := rows.Scan(&elem.IDIngrijitor, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
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

func CreateCaretaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bUsername := r.FormValue("Username")
	insertStringIngrijitori := fmt.Sprintf("INSERT INTO medassist_db.Ingrijitori (`Nume`, `Prenume`, `Username`) VALUES ('%s', '%s', '%s')", bNume, bPrenume, bUsername)
	_, err := db.Exec(insertStringIngrijitori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bNume, bPrenume, bUsername)
	}
}

func UpdateCaretaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bUsername := r.FormValue("Username")
	updateStringIngrijitori := fmt.Sprintf("UPDATE medassist_db.Ingrijitori SET `Nume` = '%s', `Prenume` = '%s', `Username` = '%s' WHERE `IDIngrijitor` = '%s'", bNume, bPrenume, bUsername, bID)
	_, err := db.Exec(updateStringIngrijitori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : '%s' , '%s', '%s' pentru campul cu ID : '%s'", bNume, bPrenume, bUsername, bID))
	}
}

func DeleteCaretaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringIngrijitori := fmt.Sprintf("DELETE FROM medassist_db.Ingrijitori WHERE `IDIngrijitor` = '%s'", bID)
	_, err := db.Exec(deleteStringIngrijitori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
