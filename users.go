package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.Users"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []Users
	for rows.Next() {
		var elem Users
		if err := rows.Scan(&elem.IDUser, &elem.Username, &elem.Password); err != nil {
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

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	bID := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.Users WHERE IDUser = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Users
	for rows.Next() {
		var elem Users
		if err := rows.Scan(&elem.IDUser, &elem.Username, &elem.Password); err != nil {
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

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bUsername := mux.Vars(r)["username"]
	selectStringUsername := ("SELECT * FROM medassist_db.Users WHERE Username = \"" + bUsername + "\"")
	rows, err := db.Query(selectStringUsername)
	if err != nil {
		panic(err)
	}
	var elements []Users
	for rows.Next() {
		var elem Users
		if err := rows.Scan(&elem.IDUser, &elem.Username, &elem.Password); err != nil {
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDUser")
	bUsername := parameters.Get("Username")
	bPassword := parameters.Get("Password")
	insertStringUsers := fmt.Sprintf("INSERT INTO medassist_db.Users (`IDUser`, `Username`, `Password`) VALUES ('%s', '%s', '%s')", bID, bUsername, bPassword)
	_, err := db.Exec(insertStringUsers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bUsername, bPassword)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDUser")
	bUsername := parameters.Get("Username")
	bPassword := parameters.Get("Password")
	updateStringUsers := fmt.Sprintf("UPDATE medassist_db.Users SET `IDUser` = '%s', `Username` = '%s', `Password` = '%s' WHERE `IDUser` = '%s'", bID, bUsername, bPassword, bID)
	_, err := db.Exec(updateStringUsers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bUsername, bPassword)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringUsers := fmt.Sprintf("DELETE FROM medassist_db.Users WHERE `IDUser` = '%s'", bID)
	_, err := db.Exec(deleteStringUsers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
