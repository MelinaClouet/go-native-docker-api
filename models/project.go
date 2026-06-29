package models

type Service struct {
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Environment map[string]string `json:"environment"`
	Volumes     []string          `json:"volumes"`
}

type Project struct {
	Name     string    `json:"name"`
	Services []Service `json:"services"`
	Engine   Engine    `json:"engine"` // ✅ ICI
}
