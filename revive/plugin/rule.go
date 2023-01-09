package plugin

type rule struct {
	RuleName string   `json:"RuleName"`
	Failure  string   `json:"Failure"`
	Position position `json:"Position"`
}

type position struct {
	Start location `json:"Start"`
	End   location `json:"End"`
}

type location struct {
	FileName string `json:"Filename"`
	Line     uint   `json:"Line"`
}

func (r rule) fileName() string {
	return r.Position.Start.FileName
}

func (r rule) startLine() uint {
	return r.Position.Start.Line
}

func (r rule) endLine() uint {
	return r.Position.End.Line
}

func (r rule) text() string {
	return r.Failure
}
