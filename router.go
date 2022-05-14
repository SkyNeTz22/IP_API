package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	alertHook := r.PathPrefix("/api/alerts").Subrouter()
	alertHook.HandleFunc("/", GetAlerts).Methods("GET")
	alertHook.HandleFunc("/id/{id}/", GetAlertByID).Methods("GET")
	alertHook.HandleFunc("/create/", CreateAlert).Methods("POST")
	alertHook.HandleFunc("/id/{id}/", UpdateAlert).Methods("PUT")
	alertHook.HandleFunc("/id/{id}/", DeleteAlert).Methods("DELETE")

	// r.HandleFunc("/api/alerts", GetAlerts).Methods("GET")
	// r.HandleFunc("/api/alerts/{id}", GetAlertByID).Methods("GET")
	// r.HandleFunc("/api/alerts", CreateAlert).Methods("POST")
	// r.HandleFunc("/api/alerts", UpdateAlert).Methods("PUT")
	// r.HandleFunc("/api/alerts/{id}", DeleteAlert).Methods("DELETE")

	androidHook := r.PathPrefix("/api/android").Subrouter()
	androidHook.HandleFunc("/", GetAllMobileData).Methods("GET")
	androidHook.HandleFunc("/id/{id}/", GetMobileDataByID).Methods("GET")
	androidHook.HandleFunc("/create/", InsertMobileData).Methods("POST")
	androidHook.HandleFunc("/id/{id}/", UpdateMobileData).Methods("PUT")
	androidHook.HandleFunc("/id/{id}/", DeleteMobileData).Methods("DELETE")

	// r.HandleFunc("/api/android", GetAllMobileData).Methods("GET")
	// r.HandleFunc("/api/android/{id}", GetMobileData).Methods("GET")
	// r.HandleFunc("/api/android", InsertMobileData).Methods("POST")
	// r.HandleFunc("/api/android", UpdateMobileData).Methods("PUT")
	// r.HandleFunc("/api/android/{id}", DeleteMobileData).Methods("DELETE")

	caretakerHook := r.PathPrefix("/api/caretakers").Subrouter()
	caretakerHook.HandleFunc("/", GetCaretakers).Methods("GET")
	caretakerHook.HandleFunc("/id/{id}/", GetCaretakerByID).Methods("GET")
	caretakerHook.HandleFunc("/username/{username}/", GetCaretakerByUsername).Methods("GET")
	caretakerHook.HandleFunc("/create/", CreateCaretaker).Methods("POST")
	caretakerHook.HandleFunc("/id/{id}/", UpdateCaretaker).Methods("PUT")
	caretakerHook.HandleFunc("/id/{id}/", DeleteCaretaker).Methods("DELETE")

	// r.HandleFunc("/api/caretakers", GetCaretakers).Methods("GET")
	// r.HandleFunc("/api/caretakers/{id}", GetCaretaker).Methods("GET")
	// r.HandleFunc("/api/caretakers", CreateCaretaker).Methods("POST")
	// r.HandleFunc("/api/caretakers", UpdateCaretaker).Methods("PUT")
	// r.HandleFunc("/api/caretakers/{id}", DeleteCaretaker).Methods("DELETE")

	medicalFileHook := r.PathPrefix("/api//medicalfile").Subrouter()
	medicalFileHook.HandleFunc("/", GetMedicalFiles).Methods("GET")
	medicalFileHook.HandleFunc("/id/{id}/", GetMedicalFileByID).Methods("GET")
	medicalFileHook.HandleFunc("/create/", CreateMedicalFile).Methods("POST")
	medicalFileHook.HandleFunc("/id/{id}/", UpdateMedicalFile).Methods("PUT")
	medicalFileHook.HandleFunc("/id/{id}/", DeleteMedicalFile).Methods("DELETE")

	// r.HandleFunc("/api/medicalfile", GetMedicalFiles).Methods("GET")
	// r.HandleFunc("/api/medicalfile/{id}", GetMedicalFile).Methods("GET")
	// r.HandleFunc("/api/medicalfile", CreateMedicalFile).Methods("POST")
	// r.HandleFunc("/api/medicalfile", UpdateMedicalFile).Methods("PUT")
	// r.HandleFunc("/api/medicalfile/{id}", DeleteMedicalFile).Methods("DELETE")

	medicHook := r.PathPrefix("/api/medics").Subrouter()
	medicHook.HandleFunc("/", GetMedics).Methods("GET")
	medicHook.HandleFunc("/id/{id}/", GetMedicByID).Methods("GET")
	medicHook.HandleFunc("/username/{username}/", GetMedicByUsername).Methods("GET")
	medicHook.HandleFunc("/create/", CreateMedic).Methods("POST")
	medicHook.HandleFunc("/id/{id}/", UpdateMedic).Methods("PUT")
	medicHook.HandleFunc("/id/{id}/", DeleteMedic).Methods("DELETE")

	// r.HandleFunc("/api/medics", GetMedics).Methods("GET")
	// r.HandleFunc("/api/medics/id/{id}", GetMedicById).Methods("GET")
	// r.HandleFunc("/api/medics/username", GetMedicByUsername).Methods("GET")
	// r.HandleFunc("/api/medics", CreateMedic).Methods("POST")
	// r.HandleFunc("/api/medics", UpdateMedic).Methods("PUT")
	// r.HandleFunc("/api/medics/{id}", DeleteMedic).Methods("DELETE")

	patientHook := r.PathPrefix("/api/patients").Subrouter()
	patientHook.HandleFunc("/", GetPatients).Methods("GET")
	patientHook.HandleFunc("/id/{id}/", GetPatientByID).Methods("GET")
	patientHook.HandleFunc("/username/{username}/", GetPatientByUsername).Methods("GET")
	patientHook.HandleFunc("/create/", CreatePatient).Methods("POST")
	patientHook.HandleFunc("/id/{id}/", UpdatePatient).Methods("PUT")
	patientHook.HandleFunc("/id/{id}/", DeletePatient).Methods("DELETE")

	// r.HandleFunc("/api/patients", GetPatients).Methods("GET")
	// r.HandleFunc("/api/patients/{id}", GetPatient).Methods("GET")
	// r.HandleFunc("/api/patients", CreatePatient).Methods("POST")
	// r.HandleFunc("/api/patients", UpdatePatient).Methods("PUT")
	// r.HandleFunc("/api/patients/{id}", DeletePatient).Methods("DELETE")

	sensorHook := r.PathPrefix("/api/sensors").Subrouter()
	sensorHook.HandleFunc("/", GetAllSensorData).Methods("GET")
	sensorHook.HandleFunc("/id/{id}/", GetSensorDataByID).Methods("GET")
	sensorHook.HandleFunc("/create/", InsertSensorData).Methods("POST")
	sensorHook.HandleFunc("/id/{id}/", UpdateSensorData).Methods("PUT")
	sensorHook.HandleFunc("/id/{id}/", DeleteSensorData).Methods("DELETE")

	// r.HandleFunc("/api/sensors", GetAllSensorData).Methods("GET")
	// r.HandleFunc("/api/sensors/{id}", GetSensorData).Methods("GET")
	// r.HandleFunc("/api/sensors", InsertSensorData).Methods("POST")
	// r.HandleFunc("/api/sensors", UpdateSensorData).Methods("PUT")
	// r.HandleFunc("/api/sensors/{id}", DeleteSensorData).Methods("DELETE")

	userHook := r.PathPrefix("/api/users").Subrouter()
	userHook.HandleFunc("/", GetUsers).Methods("GET")
	userHook.HandleFunc("/id/{id}/", GetUserByID).Methods("GET")
	userHook.HandleFunc("/username/{username}/", GetUserByUsername).Methods("GET")
	userHook.HandleFunc("/create/", CreateUser).Methods("POST")
	userHook.HandleFunc("/id/{id}/", UpdateUser).Methods("PUT")
	userHook.HandleFunc("/id/{id}/", DeleteUser).Methods("DELETE")

	// r.HandleFunc("/api/users", GetUsers).Methods("GET")
	// r.HandleFunc("/api/users/id", GetUserByID).Methods("GET")
	// r.HandleFunc("/api/users/username", GetUserByUsername).Methods("GET")
	// r.HandleFunc("/api/users", CreateUser).Methods("POST")
	// r.HandleFunc("/api/users", UpdateUser).Methods("PUT")
	// r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	supervisorHook := r.PathPrefix("/api/supervisors").Subrouter()
	supervisorHook.HandleFunc("/", GetSupervisors).Methods("GET")
	supervisorHook.HandleFunc("/id/{id}/", GetSupervisorByID).Methods("GET")
	supervisorHook.HandleFunc("/username/{username}/", GetSupervisorByUsername).Methods("GET")
	supervisorHook.HandleFunc("/create/", CreateSupervisor).Methods("POST")
	supervisorHook.HandleFunc("/id/{id}/", UpdateSupervisor).Methods("PUT")
	supervisorHook.HandleFunc("/id/{id}/", DeleteSupervisor).Methods("DELETE")

	// r.HandleFunc("/api/supervisors", GetSupervisors).Methods("GET")
	// r.HandleFunc("/api/supervisors/{id}", GetSupervisor).Methods("GET")
	// r.HandleFunc("/api/supervisors", CreateSupervisor).Methods("POST")
	// r.HandleFunc("/api/supervisors", UpdateSupervisor).Methods("PUT")
	// r.HandleFunc("/api/supervisors/{id}", DeleteSupervisor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))

}
