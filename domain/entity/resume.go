package entity

type Resume struct {
	UserPersonalDetails
	EducationDetails      []UserEducationField `json:"educationDetails"`
	WorkExperienceDetails []UserWorkExperience `json:"workExperienceDetails"`
	Skills                UserSkills           `json:"userSkills"`
	AchievementDetails    []UserAchievements   `json:"achievementDetails"`
}

type UserPersonalDetails struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	LinkedInURL string `json:"linkedInURL"`
	GitHubUrl   string `json:"githubURL"`
}

type UserEducationField struct {
	Degree              string `json:"degree"`
	Major               string `json:"major"`
	College             string `json:"college"`
	City                string `json:"city"`
	EducationStartDate  string `json:"educationStartDate"`
	EducationEndingDate string `json:"educationEndingDate"`
}

type UserWorkExperience struct {
	Organization     string `json:"organization"`
	Designation      string `json:"designation"`
	Location         string `json:"location"`
	Responsibilities string `json:"responsibilities"`
	JobStartDate     string `json:"jobStartDate"`
	JobEndingDate    string `json:"jobEndingDate"`
}

type UserSkills struct {
	ProgrammingLanguages string `json:"programmingLanguages"`
	Database             string `json:"database"`
	SoftwareTools        string `json:"softwareTools"`
	Other                string `json:"other"`
}

type UserAchievements struct {
	Achievement       string `json:"achievement"`
	DateOfAchievement string `json:"dateOfAchievement"`
}
