package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllMobileData(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.DateMobile"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []DateMobile
	for rows.Next() {
		var elem DateMobile
		if err := rows.Scan(&elem.IDDate, &elem.Data, &elem.Greutate, &elem.Glicemie, &elem.Tensiune, &elem.Temperatura, &elem.IDPacient); err != nil {
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

func GetMobileDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.DateMobile WHERE `IDDate` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []DateMobile
	for rows.Next() {
		var elem DateMobile
		if err := rows.Scan(&elem.IDDate, &elem.Data, &elem.Greutate, &elem.Glicemie, &elem.Tensiune, &elem.Temperatura, &elem.IDPacient); err != nil {
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

func GetMobileDataByPatientID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.DateMobile WHERE `IDPacient` = '%s' ORDER BY `IDDate` DESC LIMIT 1", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []DateMobile
	for rows.Next() {
		var elem DateMobile
		if err := rows.Scan(&elem.IDDate, &elem.Data, &elem.Greutate, &elem.Glicemie, &elem.Tensiune, &elem.Temperatura, &elem.IDPacient); err != nil {
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

func InsertMobileData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bData := r.FormValue("Data")
	bGreutate := r.FormValue("Greutate")
	bGlicemie := r.FormValue("Glicemie")
	bTensiune := r.FormValue("Tensiune")
	bTemperatura := r.FormValue("Temperatura")
	bIDP := r.FormValue("IDPacient")
	insertStringDateMobile := fmt.Sprintf("INSERT INTO medassist_db.DateMobile (`Data`, `Greutate`, `Glicemie`, `Tensiune`, `Temperatura`, `IDPacient`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')", bData, bGreutate, bGlicemie, bTensiune, bTemperatura, bIDP)
	_, err := db.Exec(insertStringDateMobile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bData, bGreutate, bGlicemie, bTensiune, bTemperatura, bIDP)
	}
}

func UpdateMobileData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bData := r.FormValue("Data")
	bGreutate := r.FormValue("Greutate")
	bGlicemie := r.FormValue("Glicemie")
	bTensiune := r.FormValue("Tensiune")
	bTemperatura := r.FormValue("Temperatura")
	bIDP := r.FormValue("IDPacient")
	updateStringDateMobile := fmt.Sprintf("UPDATE medassist_db.DateMobile SET `Data` = '%s', `Greutate` = '%s', `Glicemie` = '%s', `Tensiune` = '%s', `Temperatura` = '%s', `IDPacient` = '%s' WHERE `IDDate` = '%s'", bData, bGreutate, bGlicemie, bTensiune, bTemperatura, bIDP, bID)
	_, err := db.Exec(updateStringDateMobile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bData, bGreutate, bGlicemie, bTensiune, bTemperatura, bIDP)
	}
}

func DeleteMobileData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringDateMobile := fmt.Sprintf("DELETE FROM medassist_db.DateMobile WHERE `IDDate` = '%s'", bID)
	_, err := db.Exec(deleteStringDateMobile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
