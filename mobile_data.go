package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		if err := rows.Scan(&elem.IDDate, &elem.Data, &elem.Greutate, &elem.Glicemie, &elem.Tensiune_Mica, &elem.Tensiune_Mare, &elem.Temperatura, &elem.IDPacient); err != nil {
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

func GetMobileData(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.DateMobile WHERE IDSenzor = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []DateMobile
	for rows.Next() {
		var elem DateMobile
		if err := rows.Scan(&elem.IDDate, &elem.Data, &elem.Greutate, &elem.Glicemie, &elem.Tensiune_Mica, &elem.Tensiune_Mare, &elem.Temperatura, &elem.IDPacient); err != nil {
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
	parameters := r.URL.Query()
	bID := parameters.Get("IDDate")
	bData := parameters.Get("Data")
	bGreutate := parameters.Get("Greutate")
	bGlicemie := parameters.Get("Glicemie")
	bTensiuneMica := parameters.Get("Tensiune_Mica")
	bTensiuneMare := parameters.Get("Tensiune_Mare")
	bTemperatura := parameters.Get("Temperatura")
	bIDP := parameters.Get("IDPacient")
	insertStringDateMobile := fmt.Sprintf("INSERT INTO medassist_db.DateMobile (`IDDate`, `Data`, `Greutate`, `Glicemie`, `Tensiune_Mica`, `Tensiune_Mare`, `Temperatura`, `IDPacient`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", bID, bData, bGreutate, bGlicemie, bTensiuneMica, bTensiuneMare, bTemperatura, bIDP)
	_, err := db.Exec(insertStringDateMobile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bData, bGreutate, bGlicemie, bTensiuneMica, bTensiuneMare, bTemperatura, bIDP)
	}
}

func UpdateMobileData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDDate")
	bData := parameters.Get("Data")
	bGreutate := parameters.Get("Greutate")
	bGlicemie := parameters.Get("Glicemie")
	bTensiuneMica := parameters.Get("Tensiune_Mica")
	bTensiuneMare := parameters.Get("Tensiune_Mare")
	bTemperatura := parameters.Get("Temperatura")
	bIDP := parameters.Get("IDPacient")
	updateStringDateMobile := fmt.Sprintf("UPDATE medassist_db.DateMobile SET `IDDate` = '%s', `Data` = '%s', `Greutate` = '%s', `Glicemie` = '%s', `Tensiune_Mica` = '%s', `Tensiune_Mare` = '%s', `Temperatura` = '%s', `IDPacient` = '%s' WHERE `IDDate` = '%s'", bID, bData, bGreutate, bGlicemie, bTensiuneMica, bTensiuneMare, bTemperatura, bIDP, bID)
	_, err := db.Exec(updateStringDateMobile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bData, bGreutate, bGlicemie, bTensiuneMica, bTensiuneMare, bTemperatura, bIDP)
	}
}

func DeleteMobileData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringDateMobile := fmt.Sprintf("DELETE FROM medassist_db.DateMobile WHERE `IDDate` = '%s'", bID)
	_, err := db.Exec(deleteStringDateMobile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
