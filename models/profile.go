package models

type ProfileAPIResponse struct {
	Name       string                  `json:"name"`
	Email      string                  `json:"email"`
	Phone      string                  `json:"phone"`
	Skills     []string                `json:"skills"`
	Education  []EducationAPIResponse  `json:"education"`
	Experience []ExperienceAPIResponse `json:"experience"`
}

type EducationAPIResponse struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

type ExperienceAPIResponse struct {
	Name  string   `json:"name"`
	Dates []string `json:"dates"`
	URL   string   `json:"url,omitempty"`
}
