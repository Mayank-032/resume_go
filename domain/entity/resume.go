package entity

type Resume struct {
	UserPersonalDetails
	UserEducationField
	UserWorkExperience
	UserAchievements
}

type UserPersonalDetails struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	LinkedInURL string `json:"linkedInURL"`
	GitHubUrl   string `json:"githubURL"`
}

type UserEducationField struct {
	Degree              string `json:"degree"`
	Major               string `json:"major"`
	EducationStartDate  string `json:"educationStartDate"`
	EducationEndingDate string `json:"educationEndingDate"`
}

type UserWorkExperience struct {
	Organization  string `json:"organization"`
	Duration      int    `json:"duration"`
	JobStartDate  string `json:"jobStartDate"`
	JobEndingDate string `json:"jobEndingDate"`
}

type UserAchievements struct {
	Achievement       string `json:"achievement"`
	DateOfAchievement string `json:"dateOfAchievement"`
}
