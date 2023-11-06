package model

type TopInfo struct {
	Pet             Pet `json:"pet"`
	DosageSchedules struct {
		Today Task `json:"today"`
		Next  Task `json:"next"`
	} `json:"dosage_schedules"`
	PhysicalCondition struct {
		Food   int `json:"food"`
		Sweat  int `json:"sweat"`
		Toilet int `json:"toilet"`
	} `json:"physical_condition"`
	Memo      Task       `json:"memo"`
	Schedules []Schedule `json:"schedules"`
}
