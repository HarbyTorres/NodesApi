package models

type Node struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Data     map[string]interface{} `json:"data"`
	Class    string                 `json:"class"`
	HTML     string                 `json:"html"`
	Typenode string                 `json:"typenode"`
	Inputs   struct {
		Input1 struct {
			Connections []struct {
				Node   string `json:"node"`
				Output string `json:"output"`
			} `json:"connections"`
		} `json:"input_1"`
		Input2 struct {
			Connections []struct {
				Node   string `json:"node"`
				Output string `json:"output"`
			} `json:"connections"`
		} `json:"input_2"`
	} `json:"inputs"`
	Outputs struct {
		Output1 struct {
			Connections []struct {
				Node   string `json:"node"`
				Output string `json:"output"`
			} `json:"connections"`
		} `json:"output_1"`
		Output2 struct {
			Connections []struct {
				Node   string `json:"node"`
				Output string `json:"output"`
			} `json:"connections"`
		} `json:"output_2"`
	} `json:"outputs"`
	PosX float64 `json:"pos_x"`
	PosY float64 `json:"pos_y"`
}
