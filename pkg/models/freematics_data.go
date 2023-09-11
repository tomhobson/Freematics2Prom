package models

type FreematicsFile struct {
	Data []FreematicsData
}

type FreematicsData struct {
	Name  string
	Value interface{}
}
