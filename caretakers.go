package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		if err := rows.Scan(&elem.IDIngrijitor, &elem.Nume, &elem.Prenume); err != nil {
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

func GetCaretaker(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.Ingrijitori WHERE IDIngrijitor = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Ingrijitori
	for rows.Next() {
		var elem Ingrijitori
		if err := rows.Scan(&elem.IDIngrijitor, &elem.Nume, &elem.Prenume); err != nil {
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
	parameters := r.URL.Query()
	bID := parameters.Get("IDIngrijitor")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	insertStringIngrijitori := fmt.Sprintf("INSERT INTO medassist_db.Ingrijitori (`IDIngrijitor`, `Nume`,`Prenume`) VALUES ('%s', '%s', '%s')", bID, bNume, bPrenume)
	_, err := db.Exec(insertStringIngrijitori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bNume, bPrenume)
	}
}

func UpdateCaretaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDIngrijitor")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	updateStringIngrijitori := fmt.Sprintf("UPDATE medassist_db.Ingrijitori SET `IDIngrijitor` = '%s', `Nume` = '%s', `Prenume` = '%s' WHERE `IDIngrijitor` = '%s'", bID, bNume, bPrenume, bID)
	_, err := db.Exec(updateStringIngrijitori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bNume, bPrenume)
	}
}

func DeleteCaretaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringIngrijitori := fmt.Sprintf("DELETE FROM medassist_db.Ingrijitori WHERE `IDIngrijitor` = '%s'", bID)
	_, err := db.Exec(deleteStringIngrijitori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
