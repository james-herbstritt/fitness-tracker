package fitbit

type ActivityGoals struct {
	Goals struct {
		ActiveMinutes     int64   `json:"activeMinutes,omitempty"`
		ActiveZoneMinutes int64   `json:"activeZoneMinutes,omitempty"`
		CaloriesOut       int64   `json:"caloriesOut,omitempty"`
		Distance          float64 `json:"distance,omitempty"`
		Floors            int64   `json:"floors,omitempty"`
		Steps             int64   `json:"steps,omitempty"`
	} `json:"goals,omitempty"`
}
