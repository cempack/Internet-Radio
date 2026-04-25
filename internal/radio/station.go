package radio

type Station struct {
	ID      string
	Name    string
	URL     string
	Country string
	Language string
	Tags    []string
	Codec   string
	Bitrate int
	Votes   int
}
