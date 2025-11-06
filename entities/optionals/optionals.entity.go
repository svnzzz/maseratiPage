package optionals

type Optional struct {
	ModelID   int     `json:"ModelloID"`
	Name      string  `json:"Optional"`
	Model     string  `json:"Modello"`
	Price     float32 `json:"Prezzo"`
	FileImage string  `json:"FileImage"`
	Default   bool    `json:"Predefinito"`
}
