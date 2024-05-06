package NeoWs

type NeoWsResponse struct {
	Links            links                        `json:"links"`
	ElementCount     int                          `json:"element_count"`
	NearEarthObjects map[string][]NearEarthObject `json:"near_earth_objects"`
}

type links struct {
	Next string
	Prev string
	Self string
}

type NearEarthObject struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID                             string              `json:"id"`
	NeoReferenceID                 string              `json:"neo_reference_id"`
	Name                           string              `json:"name"`
	NasaJPLURL                     string              `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH             float64             `json:"absolute_magnitude_h"`
	EstimatedDiameter              estimatedDiameter   `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []closeApproachData `json:"close_approach_data"`
	IsSentryObject                 bool                `json:"is_sentry_object"`
	SentryData                     string              `json:"sentry_data,omitempty"`
}

type estimatedDiameter struct {
	Kilometers struct {
		EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
		EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
	} `json:"kilometers"`
	Meters struct {
		EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
		EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
	} `json:"meters"`
	Miles struct {
		EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
		EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
	} `json:"miles"`
	Feet struct {
		EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
		EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
	} `json:"feet"`
}

type closeApproachData struct {
	CloseApproachDate      string           `json:"close_approach_date"`
	CloseApproachDateFull  string           `json:"close_approach_date_full"`
	EpochDateCloseApproach int64            `json:"epoch_date_close_approach"`
	RelativeVelocity       relativeVelocity `json:"relative_velocity"`
	MissDistance           missDistance     `json:"miss_distance"`
	OrbitingBody           string           `json:"orbiting_body"`
}

type relativeVelocity struct {
	KilometersPerSecond string `json:"kilometers_per_second"`
	KilometersPerHour   string `json:"kilometers_per_hour"`
	MilesPerHour        string `json:"miles_per_hour"`
}

type missDistance struct {
	Astronomical string `json:"astronomical"`
	Lunar        string `json:"lunar"`
	Kilometers   string `json:"kilometers"`
	Miles        string `json:"miles"`
}
