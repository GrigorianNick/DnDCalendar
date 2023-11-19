package calendar

type Calendar struct {
	Years    map[int]*Year
	Campaign string
}

type Year struct {
	Months   map[int]*Month
	YearName int
}

type Month struct {
	Days map[int]*Day
	Name string
}

type Day struct {
	Events []*Event
	Name   string
}

type Event struct {
	Description string
}
