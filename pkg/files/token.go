package files

type Upload struct {
	Token           string   `json:"token"`
	Valid           bool     `json:"valid"`
	Expires         int      `json:"expires"`
	AllowedMimeType []string `json:"allowableMimetypes"`
}

type Download struct {
	Token string `json:"token"`
}
