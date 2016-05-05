package coverart_test

import (
	"fmt"
	"github.com/piraveen/go-coverart"
	"os"
	"testing"
)

func TestItunes(t *testing.T) {
	itunes := coverart.Itunes()

	results, err := itunes.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover http://is3.mzstatic.com/image/thumb/Music4/v4/38/42/2b/38422b5a-d597-c4ac-5287-be05cd05dc9e/source/100x100bb.jpg
	}
}

func TestLastFm(t *testing.T) {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	apiKey := os.Getenv("LASTFM_APIKEY")
	lastfm, err := coverart.LastFm(apiKey)

	if err != nil {
		// Abort action
		fmt.Printf("No API Key or incorrectly set\n")
		// Output: No API Key or incorrectly set
		return
	}

	results, err := lastfm.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: ArtistCover http://img2-ak.lst.fm/i/u/arQ/eb410194931c9427e2240023426be62b.png
	}
}

func TestSpotify(t *testing.T) {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	// c := os.Getenv("SPOTIFY_CLIENTID")
	// s := os.Getenv("SPOTIFY_CLIENTSECRET")

	// Note: Providing Spotify Client Id and Client Secret is optional, but it
	// would help you increase the Spotify rate limit for requests
	// if err := spotifyart.Configure(c, s); err != nil {
	// 	// Abort action
	// 	return
	// }

	spotify := coverart.Spotify()
	results, err := spotify.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover https://i.scdn.co/image/c649d891ee6e0b86bf460bca264bd66715bd87f4
	}
}

func ExampleItunes() {
	itunes := coverart.Itunes()

	results, err := itunes.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
	}
}

func ExampleLastFm() {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	// apiKey := os.Getenv("LASTFM_APIKEY")
	lastfm, err := coverart.LastFm("LASTFM_APIKEY")

	if err != nil {
		// You probably didn't set the API Key
		return
	}

	results, err := lastfm.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
	}
}

func ExampleSpotify() {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	// c := os.Getenv("SPOTIFY_CLIENTID")
	// s := os.Getenv("SPOTIFY_CLIENTSECRET")

	// Note: Providing Spotify Client Id and Client Secret is optional, but it
	// would help you increase the Spotify rate limit for requests
	// if err := spotifyart.Configure(c, s); err != nil {
	// 	// Abort action
	// 	return
	// }
	//

	spotify := coverart.Spotify()
	results, err := spotify.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
	}
}
