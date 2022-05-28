package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		if err := rows.Scan(&elem.IDSenzor, &elem.Puls, &elem.Lumina, &elem.Alerta_Proximitate, &elem.Temperatura_Amb, &elem.Hum_Alert, &elem.Gas_Alert, &elem.IDPacient, &elem.Data_Introducerii); err != nil {
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

func GetSensorDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.DateSenzori WHERE IDSenzor = '%s' ", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []DateSenzori
	for rows.Next() {
		var elem DateSenzori
		if err := rows.Scan(&elem.IDSenzor, &elem.Puls, &elem.Lumina, &elem.Alerta_Proximitate, &elem.Temperatura_Amb, &elem.Hum_Alert, &elem.Gas_Alert, &elem.IDPacient, &elem.Data_Introducerii); err != nil {
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

func GetSensorDataByPatientID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.DateSenzori WHERE `IDPacient` = '%s' ORDER BY IDSenzor DESC LIMIT 1", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []DateSenzori
	for rows.Next() {
		var elem DateSenzori
		if err := rows.Scan(&elem.IDSenzor, &elem.Puls, &elem.Lumina, &elem.Alerta_Proximitate, &elem.Temperatura_Amb, &elem.Hum_Alert, &elem.Gas_Alert, &elem.IDPacient, &elem.Data_Introducerii); err != nil {
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
	bPuls := r.FormValue("Puls")
	bLumina := r.FormValue("Lumina")
	bAlertaProximitate := r.FormValue("Alerta_Proximitate")
	bTemperaturaAmb := r.FormValue("Temperatura_Amb")
	bHumAlert := r.FormValue("Hum_Alert")
	bGasAlert := r.FormValue("Gas_Alert")
	bIDP := r.FormValue("IDPacient")
	bDataIntroducerii := r.FormValue("Data_Introducerii")
	insertStringDateSenzori := fmt.Sprintf("INSERT INTO medassist_db.DateSenzori (`Puls`, `Lumina`, `Alerta_Proximitate`, `Temperatura_Amb`, `Hum_Alert`, `Gas_Alert`, `IDPacient`, `Data_Introducerii`) VALUES ( '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP, bDataIntroducerii)
	_, err := db.Exec(insertStringDateSenzori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP, bDataIntroducerii)
	}
}

func UpdateSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bPuls := r.FormValue("Puls")
	bLumina := r.FormValue("Lumina")
	bAlertaProximitate := r.FormValue("Alerta_Proximitate")
	bTemperaturaAmb := r.FormValue("Temperatura_Amb")
	bHumAlert := r.FormValue("Hum_Alert")
	bGasAlert := r.FormValue("Gas_Alert")
	bIDP := r.FormValue("IDPacient")
	bDataIntroducerii := r.FormValue("Data_Introducerii")
	updateStringDateSenzori := fmt.Sprintf("UPDATE medassist_db.DateSenzori SET `Puls` = '%s', `Lumina` = '%s', `Alerta_Proximitate` = '%s', `Temperatura_Amb` = '%s', `Hum_Alert` = '%s', `Gas_Alert` = '%s', `IDPacient` = '%s', `Data_Introducerii` = '%s' WHERE `IDSenzor` = '%s'", bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP, bDataIntroducerii, bID)
	_, err := db.Exec(updateStringDateSenzori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : ", bPuls, bLumina, bAlertaProximitate, bTemperaturaAmb, bHumAlert, bGasAlert, bIDP, bDataIntroducerii)
	}
}

func DeleteSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringDateSenzori := fmt.Sprintf("DELETE FROM medassist_db.DateSenzori WHERE `IDSenzor` = '%s'", bID)
	_, err := db.Exec(deleteStringDateSenzori)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
