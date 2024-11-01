package igdb

import "time"

type Token struct {
	AccessToken string
	ExpiresAt   time.Time
	TokenType   string
}

type AuthenticateResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Game struct {
	ID                    int     `json:"id"`
	AgeRatings            []int   `json:"age_ratings"`
	AggregatedRating      float64 `json:"aggregated_rating"`
	AggregatedRatingCount int     `json:"aggregated_rating_count"`
	AlternativeNames      []int   `json:"alternative_names"`
	Artworks              []int   `json:"artworks"`
	Bundles               []int   `json:"bundles"`
	Category              int     `json:"category"`
	Cover                 Cover   `json:"cover"`
	CreatedAt             int     `json:"created_at"`
	Dlcs                  []int   `json:"dlcs"`
	ExternalGames         []int   `json:"external_games"`
	FirstReleaseDate      int     `json:"first_release_date"`
	Follows               int     `json:"follows"`
	GameEngines           []int   `json:"game_engines"`
	GameModes             []int   `json:"game_modes"`
	Genres                []int   `json:"genres"`
	Hypes                 int     `json:"hypes"`
	InvolvedCompanies     []int   `json:"involved_companies"`
	Keywords              []int   `json:"keywords"`
	MultiplayerModes      []int   `json:"multiplayer_modes"`
	Name                  string  `json:"name"`
	Platforms             []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"platforms"`
	PlayerPerspectives []int   `json:"player_perspectives"`
	Rating             float64 `json:"rating"`
	RatingCount        int     `json:"rating_count"`
	ReleaseDates       []int   `json:"release_dates"`
	Screenshots        []int   `json:"screenshots"`
	SimilarGames       []int   `json:"similar_games"`
	Slug               string  `json:"slug"`
	Summary            string  `json:"summary"`
	Tags               []int   `json:"tags"`
	Themes             []int   `json:"themes"`
	TotalRating        float64 `json:"total_rating"`
	TotalRatingCount   int     `json:"total_rating_count"`
	UpdatedAt          int     `json:"updated_at"`
	Url                string  `json:"url"`
	Videos             []int   `json:"videos"`
	Websites           []int   `json:"websites"`
	Checksum           string  `json:"checksum"`
	LanguageSupports   []int   `json:"language_supports"`
	GameLocalizations  []int   `json:"game_localizations"`
}

type Cover struct {
	Url string `json:"url"`
}
