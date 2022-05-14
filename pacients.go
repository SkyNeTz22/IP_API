package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPatients(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Pacienti"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetPatient(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.Pacienti WHERE IDPacient = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDPacient")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	bVarsta := parameters.Get("Varsta")
	bAdresa := parameters.Get("Adresa")
	bNrTel := parameters.Get("NrTel")
	bMail := parameters.Get("Mail")
	bProfesia := parameters.Get("Profesia")
	bLocDeMunca := parameters.Get("LocDeMunca")
	bIDMedic := parameters.Get("IDMedic")
	bIDIngrijitor := parameters.Get("IDIngrijitor")
	bIDSupraveghetor := parameters.Get("IDSupraveghetor")
	bIDFisa := parameters.Get("IDFisa")
	insertStringPacienti := fmt.Sprintf("INSERT INTO medassist_db.Pacienti (`IDPacient`, `Nume`, `Prenume`, `Varsta`, `Adresa`, `NrTel`, `Mail`, `Profesia`, `LocDeMunca`, `IDMedic`, `IDIngrijitor`, `IDSupraveghetor`, `IDFisa`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", bID, bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa)
	_, err := db.Exec(insertStringPacienti)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa)
	}
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDPacient")
	bNume := parameters.Get("Nume")
	bPrenume := parameters.Get("Prenume")
	bVarsta := parameters.Get("Varsta")
	bAdresa := parameters.Get("Adresa")
	bNrTel := parameters.Get("NrTel")
	bMail := parameters.Get("Mail")
	bProfesia := parameters.Get("Profesia")
	bLocDeMunca := parameters.Get("LocDeMunca")
	bIDMedic := parameters.Get("IDMedic")
	bIDIngrijitor := parameters.Get("IDIngrijitor")
	bIDSupraveghetor := parameters.Get("IDSupraveghetor")
	bIDFisa := parameters.Get("IDFisa")
	updateStringPacienti := fmt.Sprintf("UPDATE medassist_db.Pacienti SET `IDPacient` = '%s', `Nume` = '%s', `Prenume` = '%s', `Varsta` = '%s', `Adresa` = '%s', `NrTel` = '%s', `Mail` = '%s', `Profesia` = '%s', `LocDeMunca` = '%s', `IDMedic` = '%s', `IDIngrijitor` = '%s', `IDSupraveghetor` = '%s', `IDFisa` = '%s' WHERE `IDPacient` = '%s'", bID, bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa, bID)
	_, err := db.Exec(updateStringPacienti)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa)
	}
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringPacienti := fmt.Sprintf("DELETE FROM medassist_db.Pacienti WHERE `IDPacient` = '%s'", bID)
	_, err := db.Exec(deleteStringPacienti)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
