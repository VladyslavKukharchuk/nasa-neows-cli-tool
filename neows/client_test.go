package neows

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetNeoWsByTimePeriod(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
	}

	createServerMock := func(statusCode int, jsonResponse string) *httptest.Server {
		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(statusCode)
					w.Write([]byte(jsonResponse))
				},
			),
		)

		return mockServer
	}

	tests := []struct {
		name           string
		args           args
		mockApiKey     string
		mockResponse   string
		mockStatusCode int
		expected       *NeoWsResponse
		expectedErr    error
	}{
		{
			name: "success",
			args: args{
				startDate: "2024-05-06",
				endDate:   "2024-05-06",
			},
			mockApiKey: "DEMO_KEY",
			mockResponse: `{
				"links": {
					"next": "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-07&end_date=2024-05-07&detailed=false&api_key=DEMO_KEY",
					"prev": "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-05&end_date=2024-05-05&detailed=false&api_key=DEMO_KEY",
					"self": "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-06&end_date=2024-05-06&detailed=false&api_key=DEMO_KEY"
				},
				"element_count": 2,
				"near_earth_objects": {
					"2024-05-06": [
						{
							"links": {
								"self": "http://api.nasa.gov/neo/rest/v1/neo/2474425?api_key=DEMO_KEY"
							},
							"id": "2474425",
							"neo_reference_id": "2474425",
							"name": "474425 (2002 YF4)",
							"nasa_jpl_url": "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=2474425",
							"absolute_magnitude_h": 18.81,
							"estimated_diameter": {
								"kilometers": {
									"estimated_diameter_min": 0.4597851883,
									"estimated_diameter_max": 1.028110936
								},
								"meters": {
									"estimated_diameter_min": 459.7851882794,
									"estimated_diameter_max": 1028.1109360402
								},
								"miles": {
									"estimated_diameter_min": 0.2856971822,
									"estimated_diameter_max": 0.6388383204
								},
								"feet": {
									"estimated_diameter_min": 1508.4816371145,
									"estimated_diameter_max": 3373.0674833982
								}
							},
							"is_potentially_hazardous_asteroid": false,
							"close_approach_data": [
								{
									"close_approach_date": "2024-05-06",
									"close_approach_date_full": "2024-May-06 14:12",
									"epoch_date_close_approach": 1715004720000,
									"relative_velocity": {
										"kilometers_per_second": "16.8705110197",
										"kilometers_per_hour": "60733.8396709463",
										"miles_per_hour": "37737.6460999833"
									},
									"miss_distance": {
										"astronomical": "0.4369597061",
										"lunar": "169.9773256729",
										"kilometers": "65368241.308386007",
										"miles": "40617941.6700486166"
									},
									"orbiting_body": "Earth"
								}
							],
							"is_sentry_object": false
						},
						{
							"links": {
								"self": "http://api.nasa.gov/neo/rest/v1/neo/3277400?api_key=DEMO_KEY"
							},
							"id": "3277400",
							"neo_reference_id": "3277400",
							"name": "(2005 HN3)",
							"nasa_jpl_url": "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=3277400",
							"absolute_magnitude_h": 18.81,
							"estimated_diameter": {
								"kilometers": {
									"estimated_diameter_min": 0.4597851883,
									"estimated_diameter_max": 1.028110936
								},
								"meters": {
									"estimated_diameter_min": 459.7851882794,
									"estimated_diameter_max": 1028.1109360402
								},
								"miles": {
									"estimated_diameter_min": 0.2856971822,
									"estimated_diameter_max": 0.6388383204
								},
								"feet": {
									"estimated_diameter_min": 1508.4816371145,
									"estimated_diameter_max": 3373.0674833982
								}
							},
							"is_potentially_hazardous_asteroid": true,
							"close_approach_data": [
								{
									"close_approach_date": "2024-05-06",
									"close_approach_date_full": "2024-May-06 14:12",
									"epoch_date_close_approach": 1715004720000,
									"relative_velocity": {
										"kilometers_per_second": "16.8705110197",
										"kilometers_per_hour": "60733.8396709463",
										"miles_per_hour": "37737.6460999833"
									},
									"miss_distance": {
										"astronomical": "0.4369597061",
										"lunar": "169.9773256729",
										"kilometers": "65368241.308386007",
										"miles": "40617941.6700486166"
									},
									"orbiting_body": "Earth"
								}
							],
							"is_sentry_object": false
						}
					]
				}
			}`,
			mockStatusCode: http.StatusOK,
			expected: &NeoWsResponse{
				Links: Links{
					Next: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-07&end_date=2024-05-07&detailed=false&api_key=DEMO_KEY",
					Prev: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-05&end_date=2024-05-05&detailed=false&api_key=DEMO_KEY",
					Self: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-06&end_date=2024-05-06&detailed=false&api_key=DEMO_KEY",
				},
				ElementCount: 2,
				NearEarthObjects: map[string][]NearEarthObject{
					"2024-05-06": {
						{
							Links: NearEarthObjectLinks{
								Self: "http://api.nasa.gov/neo/rest/v1/neo/2474425?api_key=DEMO_KEY",
							},
							ID:                 "2474425",
							NeoReferenceID:     "2474425",
							Name:               "474425 (2002 YF4)",
							NasaJPLURL:         "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=2474425",
							AbsoluteMagnitudeH: 18.81,
							EstimatedDiameter: EstimatedDiameter{
								Kilometers: Kilometers{
									EstimatedDiameterMin: 0.4597851883,
									EstimatedDiameterMax: 1.028110936,
								},
								Meters: Meters{
									EstimatedDiameterMin: 459.7851882794,
									EstimatedDiameterMax: 1028.1109360402,
								},
								Miles: Miles{
									EstimatedDiameterMin: 0.2856971822,
									EstimatedDiameterMax: 0.6388383204,
								},
								Feet: Feet{
									EstimatedDiameterMin: 1508.4816371145,
									EstimatedDiameterMax: 3373.0674833982,
								},
							},
							IsPotentiallyHazardousAsteroid: false,
							CloseApproachData: []CloseApproachData{
								{
									CloseApproachDate:      "2024-05-06",
									CloseApproachDateFull:  "2024-May-06 14:12",
									EpochDateCloseApproach: 1715004720000,
									RelativeVelocity: RelativeVelocity{
										KilometersPerSecond: "16.8705110197",
										KilometersPerHour:   "60733.8396709463",
										MilesPerHour:        "37737.6460999833",
									},
									MissDistance: MissDistance{
										Astronomical: "0.4369597061",
										Lunar:        "169.9773256729",
										Kilometers:   "65368241.308386007",
										Miles:        "40617941.6700486166",
									},
									OrbitingBody: "Earth",
								},
							},
							IsSentryObject: false,
						},
						{
							Links: NearEarthObjectLinks{
								Self: "http://api.nasa.gov/neo/rest/v1/neo/3277400?api_key=DEMO_KEY",
							},
							ID:                 "3277400",
							NeoReferenceID:     "3277400",
							Name:               "(2005 HN3)",
							NasaJPLURL:         "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=3277400",
							AbsoluteMagnitudeH: 18.81,
							EstimatedDiameter: EstimatedDiameter{
								Kilometers: Kilometers{
									EstimatedDiameterMin: 0.4597851883,
									EstimatedDiameterMax: 1.028110936,
								},
								Meters: Meters{
									EstimatedDiameterMin: 459.7851882794,
									EstimatedDiameterMax: 1028.1109360402,
								},
								Miles: Miles{
									EstimatedDiameterMin: 0.2856971822,
									EstimatedDiameterMax: 0.6388383204,
								},
								Feet: Feet{
									EstimatedDiameterMin: 1508.4816371145,
									EstimatedDiameterMax: 3373.0674833982,
								},
							},
							IsPotentiallyHazardousAsteroid: true,
							CloseApproachData: []CloseApproachData{
								{
									CloseApproachDate:      "2024-05-06",
									CloseApproachDateFull:  "2024-May-06 14:12",
									EpochDateCloseApproach: 1715004720000,
									RelativeVelocity: RelativeVelocity{
										KilometersPerSecond: "16.8705110197",
										KilometersPerHour:   "60733.8396709463",
										MilesPerHour:        "37737.6460999833",
									},
									MissDistance: MissDistance{
										Astronomical: "0.4369597061",
										Lunar:        "169.9773256729",
										Kilometers:   "65368241.308386007",
										Miles:        "40617941.6700486166",
									},
									OrbitingBody: "Earth",
								},
							},
							IsSentryObject: false,
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "fail, request without API_KEY",
			args: args{
				startDate: "2024-05-06",
				endDate:   "2024-05-06",
			},
			mockApiKey: "",
			mockResponse: `{
    			"error": {
        			"code": "API_KEY_MISSING",
        			"message": "No api_key was supplied. Get one at https://api.nasa.gov:443"
    			}
			}`,
			mockStatusCode: http.StatusForbidden,
			expected:       nil,
			expectedErr:    errors.New("No api_key was supplied. Get one at https://api.nasa.gov:443"),
		},
		{
			name: "fail, request with invalid API_KEY",
			args: args{
				startDate: "2024-05-06",
				endDate:   "2024-05-06",
			},
			mockApiKey: "123456",
			mockResponse: `{
    			"error": {
        			"code": "API_KEY_INVALID",
        			"message": "An invalid api_key was supplied. Get one at https://api.nasa.gov:443"
    			}
			}`,
			mockStatusCode: http.StatusForbidden,
			expected:       nil,
			expectedErr:    errors.New("An invalid api_key was supplied. Get one at https://api.nasa.gov:443"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockServer := createServerMock(test.mockStatusCode, test.mockResponse)

			client := NewClient(mockServer.URL, test.mockApiKey)

			result, err := client.GetNeoWsByTimePeriod(test.args.startDate, test.args.endDate)

			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedErr, err)

			defer mockServer.Close()
		})
	}
}
