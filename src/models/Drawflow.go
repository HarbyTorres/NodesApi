package models

import (
	"time"
)

type Drawflow struct {
	Id         uint64
	Body       DrawflowMap
	CreateDate time.Time
}

type CreateDrawflow struct {
	Body DrawflowMap `json:"body"`
}

type DrawflowMap struct {
	Drawflow struct {
		Home struct {
			Data struct {
				Num1 struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Data struct {
						Num string `json:"num"`
					} `json:"data"`
					Class    string `json:"class"`
					HTML     string `json:"html"`
					Typenode string `json:"typenode"`
					Inputs   struct {
					} `json:"inputs"`
					Outputs struct {
						Output1 struct {
							Connections []struct {
								Node   string `json:"node"`
								Output string `json:"output"`
							} `json:"connections"`
						} `json:"output_1"`
					} `json:"outputs"`
					PosX int `json:"pos_x"`
					PosY int `json:"pos_y"`
				} `json:"1"`
				Num2 struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Data struct {
						Num string `json:"num"`
					} `json:"data"`
					Class    string `json:"class"`
					HTML     string `json:"html"`
					Typenode string `json:"typenode"`
					Inputs   struct {
					} `json:"inputs"`
					Outputs struct {
						Output1 struct {
							Connections []struct {
								Node   string `json:"node"`
								Output string `json:"output"`
							} `json:"connections"`
						} `json:"output_1"`
					} `json:"outputs"`
					PosX int `json:"pos_x"`
					PosY int `json:"pos_y"`
				} `json:"2"`
				Num3 struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Data struct {
						Num1 string `json:"num1"`
						Num2 string `json:"num2"`
					} `json:"data"`
					Class    string `json:"class"`
					HTML     string `json:"html"`
					Typenode string `json:"typenode"`
					Inputs   struct {
						Input1 struct {
							Connections []struct {
								Node  string `json:"node"`
								Input string `json:"input"`
							} `json:"connections"`
						} `json:"input_1"`
						Input2 struct {
							Connections []struct {
								Node  string `json:"node"`
								Input string `json:"input"`
							} `json:"connections"`
						} `json:"input_2"`
					} `json:"inputs"`
					Outputs struct {
						Output1 struct {
							Connections []interface{} `json:"connections"`
						} `json:"output_1"`
					} `json:"outputs"`
					PosX int `json:"pos_x"`
					PosY int `json:"pos_y"`
				} `json:"3"`
				Num4 struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Data struct {
						Num string `json:"num"`
					} `json:"data"`
					Class    string `json:"class"`
					HTML     string `json:"html"`
					Typenode string `json:"typenode"`
					Inputs   struct {
					} `json:"inputs"`
					Outputs struct {
						Output1 struct {
							Connections []interface{} `json:"connections"`
						} `json:"output_1"`
					} `json:"outputs"`
					PosX int `json:"pos_x"`
					PosY int `json:"pos_y"`
				} `json:"4"`
			} `json:"data"`
		} `json:"Home"`
	} `json:"drawflow"`
}
