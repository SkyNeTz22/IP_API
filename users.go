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
		if err := rows.Scan(&elem.IDUser, &elem.Username, &elem.Password, &elem.Type); err != nil {
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
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.Users WHERE `IDUser` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []Users
	for rows.Next() {
		var elem Users
		if err := rows.Scan(&elem.IDUser, &elem.Username, &elem.Password, &elem.Type); err != nil {
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
	selectStringUsername := fmt.Sprintf("SELECT * FROM medassist_db.Users WHERE `Username` = '%s'", bUsername)
	rows, err := db.Query(selectStringUsername)
	if err != nil {
		panic(err)
	}
	var elements []Users
	for rows.Next() {
		var elem Users
		if err := rows.Scan(&elem.IDUser, &elem.Username, &elem.Password, &elem.Type); err != nil {
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
	bUsername := r.FormValue("Username")
	bPassword := r.FormValue("Password")
	bType := r.FormValue("Type")
	insertStringUsers := fmt.Sprintf("INSERT INTO medassist_db.Users (`Username`, `Password`, `Type`) VALUES ('%s', '%s', '%s')", bUsername, bPassword, bType)
	_, err := db.Exec(insertStringUsers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bUsername, bPassword)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bUsername := r.FormValue("Username")
	bPassword := r.FormValue("Password")
	bType := r.FormValue("Type")
	updateStringUsers := fmt.Sprintf("UPDATE medassist_db.Users SET `Username` = '%s', `Password` = '%s', `Type` = '%s' WHERE IDUser = '%s'", bUsername, bPassword, bType, bID)
	_, err := db.Exec(updateStringUsers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : '%s' , '%s', '%s' pentru campul cu ID : '%s'", bUsername, bPassword, bType, bID))
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringUsers := fmt.Sprintf("DELETE FROM medassist_db.Users WHERE `IDUser` = '%s'", bID)
	_, err := db.Exec(deleteStringUsers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
