package radio

func MockStations() []Station {
	return []Station{
		{
			ID:      "1",
			Name:    "Lofi Girl",
			URL:     "http://localhost:8080/dummy",
			Country: "France",
			Language: "English",
			Tags:    []string{"lofi", "chill", "study"},
			Codec:   "MP3",
			Bitrate: 128,
			Votes:   9999,
		},
		{
			ID:      "2",
			Name:    "SomaFM: Groove Salad",
			URL:     "http://localhost:8080/dummy2",
			Country: "USA",
			Language: "English",
			Tags:    []string{"ambient", "chillout", "electronic"},
			Codec:   "AAC",
			Bitrate: 128,
			Votes:   5000,
		},
		{
			ID:      "3",
			Name:    "NTS Radio 1",
			URL:     "http://localhost:8080/dummy3",
			Country: "UK",
			Language: "English",
			Tags:    []string{"underground", "electronic", "eclectic"},
			Codec:   "MP3",
			Bitrate: 192,
			Votes:   3000,
		},
		{
			ID:      "4",
			Name:    "KEXP 90.3",
			URL:     "http://localhost:8080/dummy4",
			Country: "USA",
			Language: "English",
			Tags:    []string{"indie", "alternative", "rock"},
			Codec:   "MP3",
			Bitrate: 256,
			Votes:   8000,
		},
	}
}
