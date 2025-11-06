package vehicle

type Vehicle struct {
	ID              int     `json:"ModelloID"`
	Model           string  `json:"Modello"`
	BasePrice       float32 `json:"PrezzoBase"`
	BackgroundImage string  `json:"FileImageSfondo"`
}
