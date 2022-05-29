package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetPatientByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Pacienti WHERE `IDPacient` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetPatientByCaretakerID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Pacienti WHERE `IDIngrijitor` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetPatientByMedicID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Pacienti WHERE `IDMedic` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetLastPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Pacienti ORDER BY `IDPacient` DESC LIMIT 1")
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetPatientBySupervisorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Pacienti WHERE `IDSupraveghetor` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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

func GetPatientByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := mux.Vars(r)["username"]
	selectStringUsername := fmt.Sprintf("SELECT * FROM medassist_db.Pacienti WHERE `Username` = '%s'", bUsername)
	rows, err := db.Query(selectStringUsername)
	if err != nil {
		panic(err)
	}
	var elements []Pacienti
	for rows.Next() {
		var elem Pacienti
		if err := rows.Scan(&elem.IDPacient, &elem.Nume, &elem.Prenume, &elem.Varsta, &elem.Adresa, &elem.NrTel, &elem.Mail, &elem.Profesia, &elem.LocDeMunca, &elem.Username, &elem.IDMedic, &elem.IDIngrijitor, &elem.IDSupraveghetor, &elem.IDFisa); err != nil {
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
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bVarsta := r.FormValue("Varsta")
	bAdresa := r.FormValue("Adresa")
	bNrTel := r.FormValue("NrTel")
	bMail := r.FormValue("Mail")
	bProfesia := r.FormValue("Profesia")
	bLocDeMunca := r.FormValue("LocDeMunca")
	bUsername := r.FormValue("Username")
	bIDMedic := r.FormValue("IDMedic")
	bIDIngrijitor := r.FormValue("IDIngrijitor")
	bIDSupraveghetor := r.FormValue("IDSupraveghetor")
	bIDFisa := r.FormValue("IDFisa")
	insertStringPacienti := fmt.Sprintf("INSERT INTO medassist_db.Pacienti (`Nume`, `Prenume`, `Varsta`, `Adresa`, `NrTel`, `Mail`, `Profesia`, `LocDeMunca`, `Username`, `IDMedic`, `IDIngrijitor`, `IDSupraveghetor`, `IDFisa`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bUsername, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa)
	_, err := db.Exec(insertStringPacienti)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bUsername, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa)
	}
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bNume := r.FormValue("Nume")
	bPrenume := r.FormValue("Prenume")
	bVarsta := r.FormValue("Varsta")
	bAdresa := r.FormValue("Adresa")
	bNrTel := r.FormValue("NrTel")
	bMail := r.FormValue("Mail")
	bProfesia := r.FormValue("Profesia")
	bLocDeMunca := r.FormValue("LocDeMunca")
	bUsername := r.FormValue("Username")
	bIDMedic := r.FormValue("IDMedic")
	bIDIngrijitor := r.FormValue("IDIngrijitor")
	bIDSupraveghetor := r.FormValue("IDSupraveghetor")
	bIDFisa := r.FormValue("IDFisa")
	updateStringPacienti := fmt.Sprintf("UPDATE medassist_db.Pacienti SET `Nume` = '%s', `Prenume` = '%s', `Varsta` = '%s', `Adresa` = '%s', `NrTel` = '%s', `Mail` = '%s', `Profesia` = '%s', `LocDeMunca` = '%s', `Username` = '%s', `IDMedic` = '%s', `IDIngrijitor` = '%s', `IDSupraveghetor` = '%s', `IDFisa` = '%s' WHERE `IDPacient` = '%s'", bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bUsername, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa, bID)
	_, err := db.Exec(updateStringPacienti)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s' pentru campul cu ID : '%s'", bNume, bPrenume, bVarsta, bAdresa, bNrTel, bMail, bProfesia, bLocDeMunca, bUsername, bIDMedic, bIDIngrijitor, bIDSupraveghetor, bIDFisa, bID))
	}
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringPacienti := fmt.Sprintf("DELETE FROM medassist_db.Pacienti WHERE `IDPacient` = '%s'", bID)
	_, err := db.Exec(deleteStringPacienti)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
