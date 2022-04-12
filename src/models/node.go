package models

type Node struct {
	Id       uint64
	name     string
	data     Data
	class    string
	template string
	typeNode string
	inputs   Input
	outputs  Output
}

type Data struct {
	value map[string]int64
}

type Input struct {
	id uint64
}

type Output struct {
}
