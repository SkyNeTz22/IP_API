package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMedicalFiles(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	selectString := "SELECT * FROM medassist_db.FisaMed"
	rows, err := db.Query(selectString)
	if err != nil {
		panic(err)
	}
	var elements []FisaMed
	for rows.Next() {
		var elem FisaMed
		if err := rows.Scan(&elem.IDFisa, &elem.Istoric_Medical, &elem.Lista_Alergii, &elem.Recomandari, &elem.Schema_Medicatie); err != nil {
			fmt.Println(err)
		}
		elements = append(elements, elem)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	encoded, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	w.Write(encoded)
}

func GetMedicalFileByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	selectStringId := fmt.Sprintf("SELECT * FROM medassist_db.FisaMed WHERE `IDFisa` = '%s'", bID)
	rows, err := db.Query(selectStringId)
	if err != nil {
		panic(err)
	}
	var elements []FisaMed
	for rows.Next() {
		var elem FisaMed
		if err := rows.Scan(&elem.IDFisa, &elem.Istoric_Medical, &elem.Lista_Alergii, &elem.Recomandari, &elem.Schema_Medicatie); err != nil {
			fmt.Println(err)
		}
		elements = append(elements, elem)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	encoded, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	w.Write(encoded)
}

func CreateMedicalFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bIstoricMedical := r.FormValue("Istoric_Medical")
	bListaAlergii := r.FormValue("Lista_Alergii")
	bRecomandari := r.FormValue("Recomandari")
	bSchemaMedicatie := r.FormValue("Schema_Medicatie")
	insertStringFisaMed := fmt.Sprintf("INSERT INTO medassist_db.FisaMed (`Istoric_Medical`, `Lista_Alergii`, `Recomandari`, `Schema_Medicatie`) VALUES ('%s', '%s', '%s', '%s')", bIstoricMedical, bListaAlergii, bRecomandari, bSchemaMedicatie)
	_, err := db.Exec(insertStringFisaMed)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", bIstoricMedical, bListaAlergii, bRecomandari, bSchemaMedicatie)
	}
}

func UpdateMedicalFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	bIstoricMedical := r.FormValue("Istoric_Medical")
	bListaAlergii := r.FormValue("Lista_Alergii")
	bRecomandari := r.FormValue("Recomandari")
	bSchemaMedicatie := r.FormValue("Schema_Medicatie")
	updateStringFisaMed := fmt.Sprintf("UPDATE medassist_db.FisaMed SET `Istoric_Medical` = '%s', `Lista_Alergii` = '%s', `Recomandari` = '%s', `Schema_Medicatie` = '%s' WHERE `IDFisa` = '%s'", bIstoricMedical, bListaAlergii, bRecomandari, bSchemaMedicatie, bID)
	_, err := db.Exec(updateStringFisaMed)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("Actualizarea s-a efectuat cu succes! Urmatoarele date au fost actualizate : '%s' , '%s', '%s' pentru campul cu ID : '%s'", bIstoricMedical, bListaAlergii, bRecomandari, bSchemaMedicatie, bID))
	}
}

func DeleteMedicalFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bID := mux.Vars(r)["id"]
	deleteStringFisaMed := fmt.Sprintf("DELETE FROM medassist_db.FisaMed WHERE `IDFisa` = '%s'", bID)
	_, err := db.Exec(deleteStringFisaMed)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Stergerea s-a efectuat cu succes! Randul CU ID-ul " + bID + " a fost sters.")
	}
}
