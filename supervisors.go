package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSupervisors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Supraveghetori"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Supraveghetori
	for rows.Next() {
		var elem Supraveghetori
		if err := rows.Scan(&elem.IDSupraveghetor, &elem.Nume, &elem.Prenume); err != nil {
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

func GetSupervisorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Supraveghetori WHERE `IDSupraveghetor` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Supraveghetori
	for rows.Next() {
		var elem Supraveghetori
		if err := rows.Scan(&elem.IDSupraveghetor, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
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

func GetSupervisorByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := mux.Vars(r)["username"]
	selectStringUsername := fmt.Sprintf("SELECT * FROM medassist_db.Medici WHERE `Username` = '%s'", bUsername)
	rows, err := db.Query(selectStringUsername)
	if err != nil {
		panic(err)
	}
	var elements []Supraveghetori
	for rows.Next() {
		var elem Supraveghetori
		if err := rows.Scan(&elem.IDSupraveghetor, &elem.Nume, &elem.Prenume, &elem.Username); err != nil {
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

func CreateSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bUsername := r.FormValue("Username")
	insertStringSupraveghetori := fmt.Sprintf("INSERT INTO medassist_db.Supraveghetori (`Nume`,`Prenume`, `Username`) VALUES ('%s', '%s', '%s')", bNume, bPrenume, bUsername)
	_, err := db.Exec(insertStringSupraveghetori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bNume, bPrenume, bUsername)
	}
}

func UpdateSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bUsername := r.FormValue("Username")
	updateStringSupraveghetori := fmt.Sprintf("UPDATE medassist_db.Supraveghetori SET `Nume` = '%s', `Prenume` = '%s' WHERE `IDSupraveghetor` = '%s'", bNume, bPrenume, bUsername, bID)
	_, err := db.Exec(updateStringSupraveghetori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : '%s' , '%s', '%s' pentru campul cu ID : '%s'", bNume, bPrenume, bUsername, bID))
	}
}

func DeleteSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringSupraveghetori := fmt.Sprintf("DELETE FROM medassist_db.Supraveghetori WHERE `IDSupraveghetor` = '%s'", bID)
	_, err := db.Exec(deleteStringSupraveghetori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
