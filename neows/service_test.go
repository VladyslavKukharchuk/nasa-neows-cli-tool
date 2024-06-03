package neows

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestGetNEOsByDates(t *testing.T) {
	t.Run("success for multiple records", func(t *testing.T) {
		const URL = "https://api.nasa.gov/neo/rest/v1/"
		apiKey := os.Getenv("API_KEY")
		dates := []string{"2024-05-06"}

		result, _ := GetNEOsByDates(URL, apiKey, dates)

		expected := NeoWs{
			Total: 23,
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
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "3607741",
					Name:                           "(2012 QG8)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "3728903",
					Name:                           "(2015 TK)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "3745998",
					Name:                           "(2016 EN84)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "3837800",
					Name:                           "(2019 AP9)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54055121",
					Name:                           "(2020 TS3)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54096684",
					Name:                           "(2020 WN5)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54106347",
					Name:                           "(2021 BO1)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54123628",
					Name:                           "(2021 DW1)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "2612970",
					Name:                           "612970 (2005 HN3)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54296744",
					Name:                           "(2022 QA3)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54339760",
					Name:                           "(2023 BW)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54357784",
					Name:                           "(2023 JV)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54357980",
					Name:                           "(2023 JL1)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54402140",
					Name:                           "(2023 VH)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54432968",
					Name:                           "(2024 GE1)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54438668",
					Name:                           "(2024 HE2)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54438672",
					Name:                           "(2024 HL2)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54438675",
					Name:                           "(2024 HM2)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54439316",
					Name:                           "(2024 HY2)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54439322",
					Name:                           "(2024 JF)",
					IsPotentiallyHazardousAsteroid: false,
				},
				{
					Date:                           "2024-05-06",
					ID:                             "54439420",
					Name:                           "(2024 JQ)",
					IsPotentiallyHazardousAsteroid: false,
				},
			},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("unexpected result. Expected: %+v, Got: %+v", expected, result)
		}
	})
}

func TestFormatNearWsResponses(t *testing.T) {
	tests := []struct {
		name     string
		args     []*NeoWsResponse
		expected NeoWs
	}{
		{
			name: "success for single records",
			args: []*NeoWsResponse{
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
		},
		{
			name: "success for multiple records",
			args: []*NeoWsResponse{
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
				{
					Links: Links{
						Next: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-08&end_date=2024-05-08&detailed=false&api_key=DEMO_KEY",
						Prev: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-08&end_date=2024-05-08&detailed=false&api_key=DEMO_KEY",
						Self: "http://api.nasa.gov/neo/rest/v1/feed?start_date=2024-05-07&end_date=2024-05-07&detailed=false&api_key=DEMO_KEY",
					},
					ElementCount: 3,
					NearEarthObjects: map[string][]NearEarthObject{
						"2024-05-07": {
							{
								Links: NearEarthObjectLinks{
									Self: "http://api.nasa.gov/neo/rest/v1/neo/2481457?api_key=DEMO_KEY",
								},
								ID:                 "2481457",
								NeoReferenceID:     "2481457",
								Name:               "481457 (2006 XD2)",
								NasaJPLURL:         "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=2481457",
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
									Self: "http://api.nasa.gov/neo/rest/v1/neo/3458580?api_key=DEMO_KEY",
								},
								ID:                 "3458580",
								NeoReferenceID:     "3458580",
								Name:               "(2009 HV44)",
								NasaJPLURL:         "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=3458580",
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
							{
								Links: NearEarthObjectLinks{
									Self: "http://api.nasa.gov/neo/rest/v1/neo/3711166?api_key=DEMO_KEY",
								},
								ID:                 "3711166",
								NeoReferenceID:     "3711166",
								Name:               "(2015 DY53)",
								NasaJPLURL:         "https://ssd.jpl.nasa.gov/tools/sbdb_lookup.html#/?sstr=3711166",
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
			expected: NeoWs{
				Total: 5,
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
						Date:                           "2024-05-07",
						ID:                             "2481457",
						Name:                           "481457 (2006 XD2)",
						IsPotentiallyHazardousAsteroid: false,
					},
					{
						Date:                           "2024-05-07",
						ID:                             "3458580",
						Name:                           "(2009 HV44)",
						IsPotentiallyHazardousAsteroid: true,
					},
					{
						Date:                           "2024-05-07",
						ID:                             "3711166",
						Name:                           "(2015 DY53)",
						IsPotentiallyHazardousAsteroid: true,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FormatNearWsResponses(test.args)

			assert.Equal(t, test.expected, result)
		})
	}
}
