package auth

import (
    "context"
    authentik "goauthentik.io/api/v3"
)

type AuthClient struct {
    Ctx    context.Context
    Client *authentik.APIClient
}

func New(baseURL, token string) *AuthClient {
    cfg := authentik.NewConfiguration()
    cfg.Servers = authentik.ServerConfigurations{
        {
            URL: baseURL,
        },
    }

    cfg.AddDefaultHeader("Authorization", "Bearer "+token)

    client := authentik.NewAPIClient(cfg)

    return &AuthClient{
        Ctx:    context.Background(),
        Client: client,
    }
}