package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/shelojara/collection-api/devops"
	"github.com/shelojara/collection-api/port"
	"github.com/shelojara/collection-api/shared/igdb"
)

var _ port.IGDBService = (*IGDBService)(nil)

var igdbToken *igdb.Token
var igdbTokenMutex sync.Mutex

type IGDBService struct {
	clientID     string
	clientSecret string

	ctx context.Context
}

func NewIGDBService(ctx context.Context, config *devops.Config) *IGDBService {
	return &IGDBService{clientID: config.IGDB.ClientID, clientSecret: config.IGDB.ClientSecret, ctx: ctx}
}

func (s *IGDBService) Authenticate() error {
	tokenURL, err := url.Parse("https://id.twitch.tv/oauth2/token")
	if err != nil {
		return err
	}

	query := tokenURL.Query()
	query.Set("client_id", s.clientID)
	query.Set("client_secret", s.clientSecret)
	query.Set("grant_type", igdb.GrantType)
	tokenURL.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(s.ctx, "POST", tokenURL.String(), nil)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	payload := &igdb.AuthenticateResponse{}
	if err := json.NewDecoder(response.Body).Decode(payload); err != nil {
		return err
	}

	igdbToken = &igdb.Token{
		AccessToken: payload.AccessToken,
		ExpiresAt:   time.Now().Add(time.Duration(int(float64(payload.ExpiresIn)*0.9)) * time.Second),
		TokenType:   payload.TokenType,
	}

	return nil
}

func (s *IGDBService) Reauthenticate() error {
	igdbTokenMutex.Lock()
	defer igdbTokenMutex.Unlock()

	if igdbToken == nil || igdbToken.ExpiresAt.Before(time.Now()) {
		return s.Authenticate()
	}

	return nil
}

func (s *IGDBService) Search(title string) ([]*igdb.Game, error) {
	if err := s.Reauthenticate(); err != nil {
		return nil, err
	}

	body := bytes.NewBufferString(`
		search "` + title + `"; 
		where category = (0, 1, 2, 8, 9, 10, 11, 12); 
		fields *, platforms.name, cover.url;
	`)
	request, err := http.NewRequestWithContext(s.ctx, "POST", "https://api.igdb.com/v4/games/", body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Client-ID", s.clientID)
	request.Header.Set("Authorization", "Bearer "+igdbToken.AccessToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	games := make([]*igdb.Game, 0)
	if err := json.NewDecoder(response.Body).Decode(&games); err != nil {
		return nil, err
	}

	return games, nil
}

func (s *IGDBService) Get(id string) (*igdb.Game, error) {
	if err := s.Reauthenticate(); err != nil {
		return nil, err
	}

	body := bytes.NewBufferString(`where id = "` + id + `"; fields *;`)
	request, err := http.NewRequestWithContext(s.ctx, "POST", "https://api.igdb.com/v4/games/", body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Client-ID", s.clientID)
	request.Header.Set("Authorization", "Bearer "+igdbToken.AccessToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	games := make([]*igdb.Game, 0)
	if err := json.NewDecoder(response.Body).Decode(&games); err != nil {
		return nil, err
	}

	return games[0], nil
}
