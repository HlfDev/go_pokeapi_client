package models

type Pokemon struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Height    int      `json:"height"`
	Weight    int      `json:"weight"`
	Types     []string `json:"types"`
	Abilities []string `json:"abilities"`
	Sprites   Sprites  `json:"sprites"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
}
