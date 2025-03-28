package fitbit

type BadgeList struct {
	Badges []Badge `json:"badges,omitempty"`
}

type Badge struct {
	BadgeGradientEndColor   string      `json:"badgeGradientEndColor,omitempty"`
	BadgeGradientStartColor string      `json:"badgeGradientStartColor,omitempty"`
	BadgeType               string      `json:"badgeType,omitempty"`
	Category                string      `json:"category,omitempty"`
	Cheers                  []string    `json:"cheers,omitempty"`
	DateTime                ProfileTime `json:"dateTime,omitempty"`
	Description             string      `json:"description,omitempty"`
	EarnedMessage           string      `json:"earnedMessage,omitempty"`
	EncodedId               string      `json:"encodedId,omitempty"`
	Image100px              string      `json:"image100px,omitempty"`
	Image125px              string      `json:"image125px,omitempty"`
	Image300px              string      `json:"image300px,omitempty"`
	Image50px               string      `json:"image50px,omitempty"`
	Image75px               string      `json:"image75px,omitempty"`
	MarketingDescription    string      `json:"marketingDescription,omitempty"`
	MobileDescription       string      `json:"mobileDescription,omitempty"`
	Name                    string      `json:"name,omitempty"`
	ShareImage640px         string      `json:"shareImage640px,omitempty"`
	ShareText               string      `json:"shareText,omitempty"`
	ShortDescription        string      `json:"shortDescription,omitempty"`
	ShortName               string      `json:"shortName,omitempty"`
	TimesAchieved           int64       `json:"timesAchieved,omitempty"`
	Value                   int64       `json:"value,omitempty"`
}
