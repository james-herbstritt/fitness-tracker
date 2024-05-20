package main

import ( 
	// "image"
	"time"
	_ "image/png"
	_ "image/jpeg"
	"net/http"
	"io"
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
)

type ProfileTime struct {
	time.Time
}

func (t *ProfileTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

// TODO there are security concerns for decoding images 
// https://pkg.go.dev/image#hdr-Security_Considerations
type Profile struct {
	User struct {
		AboutMe string `json:"aboutMe,omitempty"`
		Age int64 `json:"age,omitempty"`
		Ambassador bool `json:"ambassador,omitempty"`
		AutoStrideEnabled bool `json:"autoStrideEnabled,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Avatar150 string `json:"avatar150,omitempty"`
		Avatar640 string `json:"avatar640,omitempty"`
		AverageDailySteps float64 `json:"averageDailySteps,omitempty"`
		ChallengesBeta bool `json:"challengesBeta,omitempty"`
		City string `json:"city,omitempty"`
		ClockTimeDisplayFormat string `json:"clockTimeDisplayFormat,omitempty"`
		Country string `json:"country,omitempty"`
		Corporate bool `json:"corporate,omitempty"`
		CorporateAdmin bool `json:"corporateAdmin,omitempty"`
		DateOfBirth ProfileTime `json:"dateOfBirth,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
		DisplayNameSetting string `json:"displayNameSetting,omitempty"`
		DistanceUnit string `json:"distanceUnit,omitempty"`
		EncodedId string `json:"endcodedId,omitempty"`
		Features struct {
			ExerciseGoal bool `json:"exerciseGoal,omitempty"`
		} `json:"features,omitempty"`
		FirstName string `json:"firstName,omitempty"`
		FoodsLocale string `json:"foodsLocale,omitempty"`
		FullName string `json:"fullName,omitempty"`
		Gender string `json:"gender,omitempty"`
		GlucoseUnit string `json:"glucoseUnit,omitempty"`
		Height float64 `json:"height,omitempty"`
		HeightUnit string `json:"heightUnit,omitempty"`
		IsBugReportEnabled bool `json:"isBugReportEnabled,omitempty"`
		IsChild bool `json:"isChild,omitempty"`
		IsCoach bool `json:"isCoach,omitempty"`
		LanguageLocale string `json:"languageLocale,omitempty"`
		LastName string `json:"lastName,omitempty"`
		LegalTermsAcceptRequired bool `json:"legalTermsAcceptRequired,omitempty"`
		Locale string `json:"locale,omitempty"`
		MemberSince ProfileTime `json:"memberSince,omitempty"`
		MfaEnabled bool `json:"mfaEnabled,omitempty"`
		OffsetFromUTCMillis int64 `json:"offsetFromUTCMillis,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
		SdkDeveloper bool `json:"sdkDeveloper,omitempty"`
		SleepTracking string `json:"sleepTracking,omitempty"`
		StartDayOfWeek string `json:"startDayOfWeek,omitempty"`
		State string `json:"state,omitempty"`
		StrideLengthRunning float64 `json:"strideLengthRunning,omitempty"`
		StrideLengthRunningType string `json:"strideLengthRunningType,omitempty"`
		StrideLengthWalking float64 `json:"strideLengthWalking,omitempty"`
		StrideLengthWalkingType string `json:"strideLengthWalkingType,omitempty"`
		SwimUnit string `json:"swimUnit,omitempty"`
		TemperatureUnit string `json:"temperatureUnit,omitempty"`
		Timezone string `json:"timezone,omitempty"`
		TopBadges []*Badge `json:"topBadges,omitempty"`
		WaterUnit string `json:"waterUnit,omitempty"`
		WaterUnitName string `json:"waterUnitName,omitempty"`
		Weight float64 `json:"weight,omitempty"`
		WeightUnit string `json:"weightUnit,omitempty"`
	} `json:"user,omitempty"`
}

func GetProfile(c *gin.Context, userId string, accessToken string) *Profile {
	userUrl := fmt.Sprintf("https://api.fitbit.com/1/user/%s/profile.json", userId)
	bearerToken := fmt.Sprintf("Bearer %s", accessToken)
	client := http.Client{}
	req, err := http.NewRequestWithContext(c, "GET", userUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"accept":        {"application/json"},
		"accept-langauge": {"en_US"},
		"accept-locale": {"en_US"},
		"Authorization": {bearerToken},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var profile Profile
	err = json.Unmarshal(b, &profile)
	if err != nil {
		panic(err)
	}

	return &profile
}

// TODO:: Write this func
// Should this take in a profile? I Think maybe it should just take in some query params
func UpdateProfile(p *Profile) *Profile {
	return p
}