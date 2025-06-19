package models


type Dashboard struct {
    UserName           string          `json:"userName"`
    DiplomasInProgress []Diploma       `json:"diplomasInProgress"`
    UpcomingClasses    []UpcomingClass `json:"upcomingClasses"`
    AICareerSuggestion string          `json:"aiCareerSuggestion"`
    AIStyleSuggestion  string          `json:"aiStyleSuggestion"`
    LastCourse         string          `json:"lastCourse"`
    RecommendedPaths   []string        `json:"recommendedPaths"`

}


