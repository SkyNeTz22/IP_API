package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetAlert(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
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
	parameters := r.URL.Query()
	bID := parameters.Get("IDAlerta")
	bGravitate := parameters.Get("Gravitate")
	bMesaj := parameters.Get("Mesaj")
	bIDS := parameters.Get("IDS")
	insertStringAlerte := fmt.Sprintf("INSERT INTO medassist_db.Alerte (`IDAlerta`, `Gravitate`,`Mesaj`, `IDSender`) VALUES ('%s', '%s', '%s', '%s')", bID, bGravitate, bMesaj, bIDS)
	_, err := db.Exec(insertStringAlerte)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bGravitate, bMesaj, bIDS)
	}
}

func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDAlerta")
	bGravitate := parameters.Get("Gravitate")
	bMesaj := parameters.Get("Mesaj")
	bIDS := parameters.Get("IDS")
	updateStringAlerte := fmt.Sprintf("UPDATE medassist_db.Alerte SET `IDAlerta` = '%s', `Gravitate` = '%s', `Mesaj` = '%s', `IDSender` = '%s' WHERE `IDAlerta` = '%s'", bID, bGravitate, bMesaj, bIDS, bID)
	_, err := db.Exec(updateStringAlerte)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bGravitate, bMesaj, bIDS)
	}
}

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringAlerte := fmt.Sprintf("DELETE FROM medassist_db.Alerte WHERE `IDAlerta` = '%s'", bID)
	_, err := db.Exec(deleteStringAlerte)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
