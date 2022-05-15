package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSenders(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Senders"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Senders
	for rows.Next() {
		var elem Senders
		if err := rows.Scan(&elem.IDSender, &elem.Username, &elem.Type); err != nil {
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

func GetSenderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Senders WHERE `IDSender` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Senders
	for rows.Next() {
		var elem Senders
		if err := rows.Scan(&elem.IDSender, &elem.Username, &elem.Type); err != nil {
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

func GetSenderByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := mux.Vars(r)["username"]
	selectStringUsername := fmt.Sprintf("SELECT * FROM medassist_db.Senders WHERE `Username` = '%s'", bUsername)
	rows, err := db.Query(selectStringUsername)
	if err != nil {
		panic(err)
	}
	var elements []Senders
	for rows.Next() {
		var elem Senders
		if err := rows.Scan(&elem.IDSender, &elem.Username, &elem.Type); err != nil {
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

func CreateSender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := r.FormValue("Username")
	bType := r.FormValue("Type")
	insertStringSenders := fmt.Sprintf("INSERT INTO medassist_db.Senders (`Username`, `Type`) VALUES ('%s', '%s')", bUsername, bType)
	_, err := db.Exec(insertStringSenders)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bUsername, bType)
	}
}

func UpdateSender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bUsername := r.FormValue("Username")
	bType := r.FormValue("Password")
	updateStringSenders := fmt.Sprintf("UPDATE medassist_db.Senders SET `Username` = '%s', `Type` = '%s' WHERE `IDSender` = '%s'", bUsername, bType, bID)
	_, err := db.Exec(updateStringSenders)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bUsername, bType)
	}
}

func DeleteSender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringSenders := fmt.Sprintf("DELETE FROM medassist_db.Senders WHERE `IDUser` = '%s'", bID)
	_, err := db.Exec(deleteStringSenders)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
