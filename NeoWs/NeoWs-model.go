package NeoWs

type NeoWs struct {
	Total            int                `json:"total"`
	NearEarthObjects []NearEarthObjects `json:"near_earth_objects"`
}

type NearEarthObjects struct {
	Date                           string `json:"date"`
	ID                             string `json:"id"`
	Name                           string `json:"name"`
	IsPotentiallyHazardousAsteroid bool   `json:"is_potentially_hazardous_asteroid"`
}
