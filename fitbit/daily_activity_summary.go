package fitbit

// Not sure if this is actually there
type ActivityLog struct {
	ActivityId         int64  `json:"activityId,omitempty"`
	ActivityParentId   int64  `json:"activityParentId,omitempty"`
	ActivityParentName string `json:"activityParentName,omitempty"`
}

type DailyActivity struct {
	Calories             int64   `json:"calories,omitempty"`
	Description          string  `json:"description,omitempty"`
	DetailsLink          string  `json:"detailsLink,omitempty"`
	Distance             float64 `json:"distance,omitempty"`
	Duration             int64   `json:"duration,omitempty"`
	HasActiveZoneMinutes bool    `json:"hasActiveZoneMinutes,omitempty"`
	HasStartTime         bool    `json:"hasStartTime,omitempty"`
	IsFavorite           bool    `json:"isFavorite,omitempty"`
	LastModifiedTime     string  `json:"lastModifiedTime,omitempty"`
	LogId                int64   `json:"logId,omitempty"`
	Name                 string  `json:"name,omitempty"`
	StartDate            string  `json:"startDate,omitempty"`
	StartTime            string  `json:"startTime,omitempty"`
	Steps                int64   `json:"steps,omitempty"`
}

type DailyGoals struct {
	ActiveMinutes int64   `json:"activeMinutes,omitempty"`
	CaloriesOut   int64   `json:"caloriesOut,omitempty"`
	Distance      float64 `json:"distance,omitempty"`
	Floors        int64   `json:"floors,omitempty"`
	Steps         int64   `json:"steps,omitempty"`
}

type SummaryDistance struct {
	Activity string  `json:"activity,omitempty"`
	Distance float64 `json:"distance,omitempty"` // Distance in meters
}

type SummaryHeartRateZone struct {
	CaloriesOut int64  `json:"caloriesOut,omitempty"`
	Max         int64  `json:"max,omitempty"`
	Min         int64  `json:"min,omitempty"`
	Minutes     int64  `json:"minutes,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Summary struct {
	ActiveScore            int64                  `json:"activeScore,omitempty"` // No longer supported, always returns -1
	ActivityCalories       int64                  `json:"activityCalories,omitempty"`
	CaloriesEstimationMu   int64                  `json:"caloriesEstimationMu,omitempty"` // Calories Measurement Uncertainty
	CaloriesBMR            int64                  `json:"caloriesBMR,omitempty"`          // Basal Metabolic Rate
	CaloriesOut            int64                  `json:"caloriesOut,omitempty"`
	CaloriesOutUnestimated int64                  `json:"caloriesOutUnestimated,omitempty"`
	Distances              []SummaryDistance      `json:"distances,omitempty"`
	Elevation              float64                `json:"elevation,omitempty"`
	FairlyActiveMinutes    int64                  `json:"fairlyActiveMinutes,omitempty"`
	Floors                 int64                  `json:"floors,omitempty"`
	HeartRateZones         []SummaryHeartRateZone `json:"heartRateZones,omitempty"`
	LightlyActiveMinutes   int64                  `json:"lightlyActiveMinutes,omitempty"`
	MarginalCalories       int64                  `json:"marginalCalories,omitempty"`
	RestingHeartRate       int64                  `json:"restingHeartRate,omitempty"`
	SedentaryMinutes       int64                  `json:"sedentaryMinutes,omitempty"`
	Steps                  int64                  `json:"steps,omitempty"`
	UseEstimation          bool                   `json:"useEstimation,omitempty"`
	VeryActiveMinutes      int64                  `json:"veryActiveMinutes,omitempty"`
}

type DailyActivitySummary struct {
	Activities []DailyActivity `json:"activities,omitempty"`
	Goals      DailyGoals      `json:"goals"`
	Summary    Summary         `json:"summary"`
}
