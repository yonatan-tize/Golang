package models

type Result struct{
	ID string `json:"id"`
	RunnerID string `json:"runner_id"`
	RaceResult string `json:"race_result"`
	Location string `json:"location"`
	Position string `json:"position,omitempty"`
	Year string `json:"year"`
}