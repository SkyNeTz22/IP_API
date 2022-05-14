package main

// var db *sql.DB
// var config Config
// var citireReusita bool

func main() {

	InitialMigration()
	initializeRouter()

	// file, err := os.ReadFile("config.json")

	// if err != nil {
	// 	panic(err)
	// }
	// json.Unmarshal(file, &config)

	// cfg := mysql.Config{
	// 	User:   config.MySql.Username,
	// 	Passwd: config.MySql.Password,
	// 	Net:    config.MySql.Network,
	// 	Addr:   config.MySql.HostName,
	// 	DBName: config.MySql.DbName,
	// }

	// jsonFile, err1 := os.Open("medici.json")
	// if err1 != nil {
	// 	fmt.Println("No bueno! || {}", err1)
	// 	citireReusita = false
	// } else {
	// 	citireReusita = true
	// }
	// fmt.Println("Am deschis JSON-urile cu succes!")
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// var objMedici Medici
	// json.Unmarshal(byteValue, &objMedici)
	// fmt.Println("ID Medic : " + strconv.Itoa(objMedici.IDMedic))
	// fmt.Println("Numele medicului : " + objMedici.Nume)
	// fmt.Println("Prenumele medicului : " + objMedici.Prenume)
	// fmt.Println("Citirea a fost realizata cu succes!")
	// if citireReusita {
	// 	fmt.Println("Se realizeaza conexiunea catre baza de date...")
	// 	db, err = sql.Open("mysql", cfg.FormatDSN())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pingErr := db.Ping()
	// 	if pingErr != nil {
	// 		log.Fatal(pingErr)
	// 	}
	// 	fmt.Println("Conectat!")
	// 	fmt.Println("Se incearca inserarea...")
	// 	insertString := fmt.Sprintf("INSERT INTO Medici (`IDMedic`, `Nume`, `Prenume`) VALUES ('%d', '%s', '%s')", objMedici.IDMedic, objMedici.Nume, objMedici.Prenume)
	// 	_, err = db.Exec(insertString)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("Inserarea s-a efectuat cu succes! Urmatoarele variabile au fost inserate : ", objMedici.IDMedic, objMedici.Nume, objMedici.Prenume)
	// 	}

	// 	selectString := "SELECT * FROM medassist_db.Medici"

	// 	var medicID int
	// 	var medicNume string
	// 	var medicPrenume string

	// 	fmt.Println("Se incearca selectarea...")
	// 	row := db.QueryRow(selectString)

	// 	switch err := row.Scan(&medicID, &medicNume, &medicPrenume); err {
	// 	case sql.ErrNoRows:
	// 		fmt.Println("Ne-au furat ucrainienii randurile crezand ca sunt tancuri rusesti...")
	// 	case nil:
	// 		fmt.Println("Selectarea s-a efectuat cu succes!")
	// 		fmt.Println("| " + fmt.Sprint(medicID) + " | " + medicNume + " | " + medicPrenume + " |")
	// 	default:
	// 		panic(err)
	// 	}
	// }

}

/*
func selectMedicDetails(db *sql.DB, medicID int) (int, error) {
    log.Printf("Trying to fetch data about medics ...\n")
    query := `SELECT * FROM Medici`
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()

    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when preparing SQL statement", err)
        return 0, err
    }
    defer stmt.Close()

    var medicID int
    row := stmt.QueryRowContext(ctx, medicID)
    if err := row.Scan(&medicID); err != nil {
        return 0, err
    }
    return medicID, nil
}

func main(){
	//	********		INCEPUT CITIRE JSON			********  //
	jsonFile1, err1 := os.Open("medici.json")

	if err1 != nil {
		fmt.Println("No bueno! || {}", err1)
	}
	fmt.Println("Am deschis JSON-urile cu succes!\n")
	defer jsonFile1.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile1)
	var objMedici Medici
	json.Unmarshal(byteValue, &objMedici)
	fmt.Println("ID Medic : " 				+ strconv.Itoa(objMedici.IDMedic))
	fmt.Println("Numele medicului : " 		+ objMedici.Nume)
	fmt.Println("Prenumele medicului : " 	+ objMedici.Prenume)
	fmt.Println("\n")
	//	********		SFARSIT CITIRE JSON			********  //

	db, err := sql.Open("mysql", dsn(""))
    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        return
    }
    defer db.Close()

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS " + dbName)
    if err != nil {
        log.Printf("Error %s when creating DB\n", err)
        return
    }
    no, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when fetching rows", err)
        return
    }
    log.Printf("rows affected %d\n", no)

    db.Close()
    db, err = sql.Open("mysql", dsn(dbName))
    if err != nil {
        log.Printf("Error %s when opening DB", err)
        return
    }
    defer db.Close()

    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(time.Minute * 5)

    ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return
    }
    log.Printf("Connected to DB %s successfully\n", dbName)

}
*/

/*
func main(){
	cfg := mysql.Config{
        User:   "dbuser",
        Passwd: "parolasql2",
        Net:    "tcp",
        Addr:   "185.207.251.196:3306",
        DBName: "medassist_db",
    }

	jsonFile1, err1 := os.Open("medici.json")

	if err1 != nil {
		fmt.Println("No bueno! || {}", err1)
	}
	fmt.Println("Am deschis JSON-urile cu succes!\n")
	defer jsonFile1.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile1)
	var objMedici Medici
	json.Unmarshal(byteValue, &objMedici)
	fmt.Println("ID Medic : " 				+ strconv.Itoa(objMedici.IDMedic))
	fmt.Println("Numele medicului : " 		+ objMedici.Nume)
	fmt.Println("Prenumele medicului : " 	+ objMedici.Prenume)
	fmt.Println("\n")

	fmt.Println("Connecting to database ....\n")

	var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	query := "SELECT * FROM medassist_db.Medici"

	rows, selectErr := db.Query("SELECT * FROM medassist_db.Medici")
	if selectErr != nil {
		log.Fatal(selectErr)
	}

	for rows.Next() {
		selectErr := rows.Scan(&medicID, &medicNume, &medicPrenume)
		if selectErr != nil {
			log.Fatal(selectErr)
		}
	}
	defer rows.Close()
	fmt.Println(medicID, medicNume, medicPrenume)

	if medicID != objMedici.IDMedic {
		fmt.Println("Attempting insert...")

		result, insertErr := db.Exec("INSERT INTO Medici VALUES (? ? ?)", objMedici.IDMedic, objMedici.Nume, objMedici.Prenume)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println("Insert successful!")
		fmt.Println(result)
	} else {
		fmt.Println("The medic is already in the DB!")
	}

	selectErr = rows.Err()
	if selectErr != nil {
		log.Fatal(selectErr)
	}
	db.Close()
}
*/

/*

type MainBody struct{
	Widget struct {
		Debug	string	`json:"debug"`

		Window struct {
			Title 			string	`json:"title"`
			NameWindow		string	`json:"name"`
			Width			int		`json:"width"`
			Height			int		`json:"height"`
		}	`json:"window"`

		Image struct {
			Src				string	`json:"src"`
			NameImg			string	`json:"name"`
			HOffsetImage	int		`json:"hOffset"`
			VOffsetImage	int		`json:"vOffset"`
			Allignment		string	`json:"alignment"`
		}	`json:"image"`

		Text struct {
			Label			string	`json:"data"`
			Size			int		`json:"size"`
			Style			string	`json:"style"`
			NameText		string	`json:"name"`
			HOffsetText		int		`json:"hOffset"`
			VOffsetText		int		`json:"vOffset"`
			Allignment		string	`json:"alignment"`
			OnMouseUp			string	`json:"onMouseUp"`
		}	`json:"text"`
	}	`json:"widget"`
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name	string 	`json:"name"`
	Type	string 	`json:"type"`
	Age		int		`json:"age"`
	Social	Social	`'json:"social"`
}

type Social struct{
	Facebook	string	`json:"facebook"`
	Twitter		string	`json:"twitter"`
}

func main() {
	// se deschide primul fisier
	jsonFile1, err1 := os.Open("users.json")

	if err1 != nil {
		fmt.Println("No bueno! || {}", err1)
	}

	// se deschide al doilea fisier
	jsonFile2, err2 := os.Open("widgets.json")

	if err2 != nil {
		fmt.Println("No bueno! || {}", err2)
	}

	fmt.Println("Am deschis JSON-urile cu succes!\n")

	// se inchid fisierele, deoarece deja sunt salvate in memorie
	defer jsonFile1.Close()
	defer jsonFile2.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile1)
	var users Users
	json.Unmarshal(byteValue, &users)

	for index := 0; index < len(users.Users); index++ {
		fmt.Println("Tipul de utilizator : " 	+ users.Users[index].Type)
		fmt.Println("Varsta utilizatorului : " 	+ strconv.Itoa(users.Users[index].Age))
		fmt.Println("Numele utilizatorului : " 	+ users.Users[index].Name)
		fmt.Println("URL FB : " 				+ users.Users[index].Social.Facebook)
		fmt.Println("\n")
	}

	byteValue, _ = ioutil.ReadAll(jsonFile2)
	var valuesBody MainBody
	json.Unmarshal(byteValue, &valuesBody)

	fmt.Println("Numele ferestrei : " 	+ valuesBody.Widget.Window.Title) //.Window.NameWindow)
	fmt.Println("Sursa imaginii : " + valuesBody.Widget.Image.Src)
}
*/
