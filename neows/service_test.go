package neows

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

type mockClient struct {
	mock.Mock
}

func (m *mockClient) GetNeoWsByTimePeriod(startDate, endDate string) (*NeoWsResponse, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).(*NeoWsResponse), args.Error(1)
}

func TestGetNEOsByDates(t *testing.T) {
	type args struct {
		URL    string
		apiKey string
		dates  []string
	}

	tests := []struct {
		name                        string
		args                        args
		mockGetNeoWsByTimePeriod    []*NeoWsResponse
		mockGetNeoWsByTimePeriodErr []error
		expected                    NeoWs
		expectedErr                 error
	}{
		{
			name: "success for single record",
			args: args{
				URL:    "https://api.nasa.gov/neo/rest/v1/",
				apiKey: os.Getenv("API_KEY"),
				dates:  []string{"2024-05-06"},
			},
			mockGetNeoWsByTimePeriod: []*NeoWsResponse{&NeoWsResponse{
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
			}},
			mockGetNeoWsByTimePeriodErr: []error{nil},
			expected: NeoWs{
				Total: 2,
				NearEarthObjects: []NearEarthObjects{
					{
						Date:                           "2024-05-06",
						ID:                             "2474425",
						Name:                           "474425 (2002 YF4)",
						IsPotentiallyHazardousAsteroid: false,
					},
					{
						Date:                           "2024-05-06",
						ID:                             "3277400",
						Name:                           "(2005 HN3)",
						IsPotentiallyHazardousAsteroid: true,
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "success for multiple records",
			args: args{
				URL:    "https://api.nasa.gov/neo/rest/v1/",
				apiKey: os.Getenv("API_KEY"),
				dates:  []string{"2024-05-05", "2024-05-06"},
			},
			mockGetNeoWsByTimePeriod: []*NeoWsResponse{
				{
					Links: Links{
						Next: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-06&end_date=2024-05-06&detailed=false&api_key=DEMO_KEY",
						Prev: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-04&end_date=2024-05-04&detailed=false&api_key=DEMO_KEY",
						Self: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-05&end_date=2024-05-05&detailed=false&api_key=DEMO_KEY",
					},
					ElementCount: 1,
					NearEarthObjects: map[string][]NearEarthObject{
						"2024-05-05": {
							{
								Links: NearEarthObjectLinks{
									Self: "http://api.nasa.gov/neo/rest/v1/neo/2474425?api_key=DEMO_KEY",
								},
								ID:                 "3374350",
								NeoReferenceID:     "3374350",
								Name:               "(2007 HW3)",
								NasaJPLURL:         "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=3374350",
								AbsoluteMagnitudeH: 20.54,
								EstimatedDiameter: EstimatedDiameter{
									Kilometers: Kilometers{
										EstimatedDiameterMin: 0.2072788434,
										EstimatedDiameterMax: 0.4634895841,
									},
									Meters: Meters{
										EstimatedDiameterMin: 207.2788433771,
										EstimatedDiameterMax: 463.4895840887,
									},
									Miles: Miles{
										EstimatedDiameterMin: 0.1287970622,
										EstimatedDiameterMax: 0.2879989864,
									},
									Feet: Feet{
										EstimatedDiameterMin: 680.0487205053,
										EstimatedDiameterMax: 1520.6351670615,
									},
								},
								IsPotentiallyHazardousAsteroid: false,
								CloseApproachData: []CloseApproachData{
									{
										CloseApproachDate:      "2024-05-05",
										CloseApproachDateFull:  "2024-May-05 23:06",
										EpochDateCloseApproach: 1714950360000,
										RelativeVelocity: RelativeVelocity{
											KilometersPerSecond: "21.7704360991",
											KilometersPerHour:   "78373.5699568493",
											MilesPerHour:        "48698.2885101323",
										},
										MissDistance: MissDistance{
											Astronomical: "0.3397126895",
											Lunar:        "132.1482362155",
											Kilometers:   "50820294.761171365",
											Miles:        "31578266.860900237",
										},
										OrbitingBody: "Earth",
									},
								},
								IsSentryObject: false,
							},
						},
					},
				},
				{
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
			},
			mockGetNeoWsByTimePeriodErr: []error{nil, nil},
			expected: NeoWs{
				Total: 3,
				NearEarthObjects: []NearEarthObjects{
					{
						Date:                           "2024-05-06",
						ID:                             "2474425",
						Name:                           "474425 (2002 YF4)",
						IsPotentiallyHazardousAsteroid: false,
					},
					{
						Date:                           "2024-05-06",
						ID:                             "3277400",
						Name:                           "(2005 HN3)",
						IsPotentiallyHazardousAsteroid: true,
					},
					{
						Date:                           "2024-05-05",
						ID:                             "3374350",
						Name:                           "(2007 HW3)",
						IsPotentiallyHazardousAsteroid: false,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockClient := new(mockClient)

			service := NewService(mockClient)

			for i, date := range test.args.dates {
				mockClient.On("GetNeoWsByTimePeriod", date, date).Return(test.mockGetNeoWsByTimePeriod[i], test.mockGetNeoWsByTimePeriodErr[i])
			}

			result, err := service.GetNEOsByDates(test.args.dates)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
