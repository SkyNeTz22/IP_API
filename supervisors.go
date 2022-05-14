package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetSupervisor(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.Supraveghetori WHERE IDSupraveghetor = " + bID)
	rows, err := db.Query(selectStringId)
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

func CreateSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDSupraveghetor")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	insertStringSupraveghetori := fmt.Sprintf("INSERT INTO medassist_db.Supraveghetori (`IDSupraveghetor`, `Nume`,`Prenume`) VALUES ('%s', '%s', '%s')", bID, bNume, bPrenume)
	_, err := db.Exec(insertStringSupraveghetori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bNume, bPrenume)
	}
}

func UpdateSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDSupraveghetor")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	updateStringSupraveghetori := fmt.Sprintf("UPDATE medassist_db.Supraveghetori SET `IDSupraveghetor` = '%s', `Nume` = '%s', `Prenume` = '%s' WHERE `IDSupraveghetor` = '%s'", bID, bNume, bPrenume, bID)
	_, err := db.Exec(updateStringSupraveghetori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bNume, bPrenume)
	}
}

func DeleteSupervisor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringSupraveghetori := fmt.Sprintf("DELETE FROM medassist_db.Supraveghetori WHERE `IDSupraveghetor` = '%s'", bID)
	_, err := db.Exec(deleteStringSupraveghetori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
