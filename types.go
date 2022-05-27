package main

type Alerte struct {
	IDAlerta  int    `json:"IDAlerta"`
	Gravitate int    `json:"Gravitate"`
	Mesaj     string `json:"Mesaj"`
	IDSender  int    `json:"IDSender"`
	IDPacient int    `json:"IDPacient"`
}

type DateMobile struct {
	IDDate              int     `json:"IDDate"`
	Data                string  `json:"Data"`
	Greutate            float32 `json:"Greutate"`
	Glicemie            float32 `json:"Glicemie"`
	Tensiune_Sistolica  int     `json:"Tensiune_Sistolica"`
	Tensiune_Diastolica int     `json:"Tensiune_Diastolica"`
	Temperatura         float32 `json:"Temperatura"`
	IDPacient           int     `json:"IDPacient"`
}

type DateSenzori struct {
	IDSenzor           int     `json:"IDSenzor"`
	Puls               int     `json:"Puls"`
	Lumina             int     `json:"Lumina"`
	Alerta_Proximitate int8    `json:"Alerta_Proximitate"`
	Temperatura_Amb    float32 `json:"Temperatura_Amb"`
	Hum_Alert          int8    `json:"Hum_Alert"`
	Gas_Alert          int8    `json:"Gas_Alert"`
	IDPacient          int     `json:"IDPacient"`
}

type FisaMed struct {
	IDFisa           int    `json:"IDFisa"`
	Istoric_Medical  string `json:"Istoric_Medical"`
	Lista_Alergii    string `json:"Lista_Alergii"`
	Recomandari      string `json:"Recomandari"`
	Schema_Medicatie string `json:"Schema_Medicatie"`
}

type Ingrijitori struct {
	IDIngrijitor int    `json:"IDIngrijitor"`
	Nume         string `json:"Nume"` // var Nume -> Type String -> Value of Key "Nume" ; "Nume": "Cutarescu"
	Prenume      string `json:"Prenume"`
	Username     string `json:"Username"`
}

type Medici struct {
	IDMedic  int    `json:"IDMedic"`
	Nume     string `json:"Nume"` // var Nume -> Type String -> Value of Key "Nume" ; "Nume": "Cutarescu"
	Prenume  string `json:"Prenume"`
	Username string `json:"Username"`
}

type Pacienti struct {
	IDPacient       int    `json:"IDPacient"`
	Nume            string `json:"Nume"`
	Prenume         string `json:"Prenume"`
	Varsta          int    `json:"Varsta"`
	Adresa          string `json:"Adresa"`
	NrTel           int    `json:"NrTel"`
	Mail            string `json:"Mail"`
	Profesia        string `json:"Profesia"`
	LocDeMunca      string `json:"LocDeMunca"`
	Username        string `json:"Username"`
	IDMedic         int    `json:"IDMedic"`
	IDIngrijitor    int    `json:"IDIngrijitor"`
	IDSupraveghetor int    `json:"IDSupraveghetor"`
	IDFisa          int    `json:"IDFisa"`
}

type Users struct {
	IDUser   int    `json:"IDUser"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Type     string `json:"Type"`
}

type Senders struct {
	IDSender int    `json:"IDSender"`
	Username string `json:"Username"`
	Type     string `json:"Type"`
}

type Supraveghetori struct {
	IDSupraveghetor int    `json:"IDSupraveghetor"`
	Nume            string `json:"Nume"` // var Nume -> Type String -> Value of Key "Nume" ; "Nume": "Cutarescu"
	Prenume         string `json:"Prenume"`
	Username        string `json:"Username"`
}
