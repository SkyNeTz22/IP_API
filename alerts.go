package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAlerts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Alerte"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Alerte
	for rows.Next() {
		var elem Alerte
		if err := rows.Scan(&elem.IDAlerta, &elem.Gravitate, &elem.Mesaj, &elem.IDSender); err != nil {
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

func GetAlertByID(w http.ResponseWriter, r *http.Request) {
	bID := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.Alerte WHERE IDAlerta = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Alerte
	for rows.Next() {
		var elem Alerte
		if err := rows.Scan(&elem.IDAlerta, &elem.Gravitate, &elem.Mesaj, &elem.IDSender); err != nil {
			w.Write([]byte("Mare eroare la scanare, boule."))
		}
		elements = append(elements, elem)
	}
	if err = rows.Err(); err != nil {
		w.Write([]byte("Mare eroare la scanare, boule."))
	}

	encoded, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	w.Write(encoded)
}

func CreateAlert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bGravitate := r.FormValue("Gravitate")
	bMesaj := r.FormValue("Mesaj")
	bIDS := r.FormValue("IDSender")
	insertStringAlerte := fmt.Sprintf("INSERT INTO medassist_db.Alerte (`Gravitate`,`Mesaj`, `IDSender`) VALUES ('%s', '%s', '%s')", bGravitate, bMesaj, bIDS)
	_, err := db.Exec(insertStringAlerte)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bGravitate, bMesaj, bIDS)
	}
}

func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bGravitate := r.FormValue("Gravitate")
	bMesaj := r.FormValue("Mesaj")
	bIDS := r.FormValue("IDSender")
	updateStringAlerte := fmt.Sprintf("UPDATE medassist_db.Alerte SET `Gravitate` = '%s', `Mesaj` = '%s', `IDSender` = '%s' WHERE `IDAlerta` = '%s'", bGravitate, bMesaj, bIDS, bID)
	_, err := db.Exec(updateStringAlerte)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bGravitate, bMesaj, bIDS)
	}
}

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringAlerte := fmt.Sprintf("DELETE FROM medassist_db.Alerte WHERE `IDAlerta` = '%s'", bID)
	_, err := db.Exec(deleteStringAlerte)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
