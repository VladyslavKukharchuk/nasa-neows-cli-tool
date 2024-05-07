package NeoWs

type NeoWsResponse struct {
	Links            Links                        `json:"links"`
	ElementCount     int                          `json:"element_count"`
	NearEarthObjects map[string][]NearEarthObject `json:"near_earth_objects"`
}

type Links struct {
	Next string
	Prev string
	Self string
}

type NearEarthObject struct {
	Links                          NearEarthObjectLinks `json:"links"`
	ID                             string               `json:"id"`
	NeoReferenceID                 string               `json:"neo_reference_id"`
	Name                           string               `json:"name"`
	NasaJPLURL                     string               `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH             float64              `json:"absolute_magnitude_h"`
	EstimatedDiameter              EstimatedDiameter    `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                 `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []CloseApproachData  `json:"close_approach_data"`
	IsSentryObject                 bool                 `json:"is_sentry_object"`
	SentryData                     string               `json:"sentry_data,omitempty"`
}

type NearEarthObjectLinks struct {
	Self string `json:"self"`
}

type EstimatedDiameter struct {
	Kilometers Kilometers `json:"kilometers"`
	Meters     Meters     `json:"meters"`
	Miles      Miles      `json:"miles"`
	Feet       Feet       `json:"feet"`
}

type Kilometers struct {
	EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
}

type Meters struct {
	EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
}

type Miles struct {
	EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
}

type Feet struct {
	EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
}

type CloseApproachData struct {
	CloseApproachDate      string           `json:"close_approach_date"`
	CloseApproachDateFull  string           `json:"close_approach_date_full"`
	EpochDateCloseApproach int64            `json:"epoch_date_close_approach"`
	RelativeVelocity       RelativeVelocity `json:"relative_velocity"`
	MissDistance           MissDistance     `json:"miss_distance"`
	OrbitingBody           string           `json:"orbiting_body"`
}

type RelativeVelocity struct {
	KilometersPerSecond string `json:"kilometers_per_second"`
	KilometersPerHour   string `json:"kilometers_per_hour"`
	MilesPerHour        string `json:"miles_per_hour"`
}

type MissDistance struct {
	Astronomical string `json:"astronomical"`
	Lunar        string `json:"lunar"`
	Kilometers   string `json:"kilometers"`
	Miles        string `json:"miles"`
}
