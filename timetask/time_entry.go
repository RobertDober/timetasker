package timetask

import "time"

type TimeEntry struct {
	Client      string
	Project     string
	Task        string
	WorkType    string `yaml:"work_type"`
	Date        time.Time
	Billable    bool
	Description string
}

// Parse date string into a real date
func (te *TimeEntry) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var auxiliary struct {
		Client      string
		Project     string
		Task        string
		WorkType    string `yaml:"work_type"`
		Date        string
		Billable    bool
		Description string
	}

	err := unmarshal(&auxiliary)
	if err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", auxiliary.Date)
	if err != nil {
		return err
	}

	te.Client = auxiliary.Client
	te.Project = auxiliary.Project
	te.Task = auxiliary.Task
	te.WorkType = auxiliary.WorkType
	te.Date = date
	te.Billable = auxiliary.Billable
	te.Description = auxiliary.Description

	return nil
}