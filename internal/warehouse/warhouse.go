package warehouse

type Contact struct {
	Phones []string `json:"phones"`
	Emails []string `json:"emails"`
}

type Warehouse struct {
	ID      string  `json:"id"`
	LLC     string  `json:"llc"`
	Name    string  `json:"name"`
	Contact Contact `json:"contact"`
	Address string  `json:"address"`
	State   string  `json:"state"`
	City    string  `json:"city"`
	ZipCode string  `json:"zip_code"`
}
