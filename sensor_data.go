package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllSensorData(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.DateSenzori"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []DateSenzori
	for rows.Next() {
		var elem DateSenzori
		if err := rows.Scan(&elem.IDSenzor, &elem.Puls, &elem.Lumina, &elem.Alerta_Proximitate, &elem.Temperatura_Amb, &elem.Hum_Alert, &elem.Gas_Alert, &elem.IDPacient); err != nil {
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

func GetSensorData(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	w.Header().Set("Content-Type", "application/json")
	selectStringId := ("SELECT * FROM medassist_db.DateSenzori WHERE IDSenzor = " + bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []DateSenzori
	for rows.Next() {
		var elem DateSenzori
		if err := rows.Scan(&elem.IDSenzor, &elem.Puls, &elem.Lumina, &elem.Alerta_Proximitate, &elem.Temperatura_Amb, &elem.Hum_Alert, &elem.Gas_Alert, &elem.IDPacient); err != nil {
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

func InsertSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDSenzor")
	bPuls := parameters.Get("Puls")
	bLumina := parameters.Get("Lumina")
	bAlertaProximitate := parameters.Get("Alerta_Proximitate")
	bTemperaturaAmb := parameters.Get("Temperatura_Amb")
	bHumAlert := parameters.Get("Hum_Alert")
	bGasAlert := parameters.Get("Gas_Alert")
	bIDP := parameters.Get("IDPacient")
	insertStringDateSenzori := fmt.Sprintf("INSERT INTO medassist_db.DateSenzori (`IDSenzor`, `Puls`, `Lumina`, `Alerta_Proximitate`, `Temperatura_Amb`, `Hum_Alert`, `Gas_Alert`, `IDPacient`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", bID, bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP)
	_, err := db.Exec(insertStringDateSenzori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bID, bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP)
	}
}

func UpdateSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := r.URL.Query()
	bID := parameters.Get("IDSenzor")
	bPuls := parameters.Get("Puls")
	bLumina := parameters.Get("Lumina")
	bAlertaProximitate := parameters.Get("Alerta_Proximitate")
	bTemperaturaAmb := parameters.Get("Temperatura_Amb")
	bHumAlert := parameters.Get("Hum_Alert")
	bGasAlert := parameters.Get("Gas_Alert")
	bIDP := parameters.Get("IDPacient")
	updateStringDateSenzori := fmt.Sprintf("UPDATE medassist_db.DateSenzori SET `IDSenzor` = '%s', `Puls` = '%s', `Lumina` = '%s', `Alerta_Proximitate` = '%s', `Temperatura_Amb` = '%s', `Hum_Alert` = '%s', `Gas_Alert` = '%s', `IDPacient` = '%s' WHERE `IDSenzor` = '%s'", bID, bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP, bID)
	_, err := db.Exec(updateStringDateSenzori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bID, bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP)
	}
}

func DeleteSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlString := r.URL.String()
	bID := urlString[len(urlString)-1:]
	deleteStringDateSenzori := fmt.Sprintf("DELETE FROM medassist_db.DateSenzori WHERE `IDSenzor` = '%s'", bID)
	_, err := db.Exec(deleteStringDateSenzori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
