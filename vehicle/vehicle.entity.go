package vehicle

type Vehicle struct {
	ID              int     `json:"ModelloID"`
	Model           string  `json:"Modello"`
	BasePrice       float32 `json:"PrezzoBase"`
	BackgroundImage string  `json:"FileImageSfondo"`
}

type Optional struct {
	ModelID   int     `json:"ModelloID"`
	Name      string  `json:"Optional"`
	Model     string  `json:"Modello"`
	Price     float32 `json:"Prezzo"`
	FileImage string  `json:"FileImage"`
	Default   bool    `json:"Predefinito"`
}
