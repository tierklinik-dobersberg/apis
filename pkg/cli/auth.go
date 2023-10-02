package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
)

type TokenFile struct {
	AccessToken   string `json:"accessToken"`
	RefreshCookie string `json:"refreshCookie"`
}

func (root *Root) Login(username string, password string, totp string) error {
	res, err := root.Auth().Login(root.Context(), connect.NewRequest(&idmv1.LoginRequest{
		AuthType: idmv1.AuthType_AUTH_TYPE_PASSWORD,
		Auth: &idmv1.LoginRequest_Password{
			Password: &idmv1.PasswordAuth{
				Username: username,
				Password: password,
			},
		},
	}))

	if err != nil {
		return err
	}

	switch res.Msg.Response.(type) {
	case *idmv1.LoginResponse_MfaRequired:
		if totp == "" {
			return fmt.Errorf("TOTP code required")
		}

		res, err = root.Auth().Login(root.Context(), connect.NewRequest(&idmv1.LoginRequest{
			AuthType: idmv1.AuthType_AUTH_TYPE_TOTP,
			Auth: &idmv1.LoginRequest_Totp{
				Totp: &idmv1.TotpAuth{
					Code: totp,
				},
			},
		}))

		if err != nil {
			return err
		}
	}

	token, ok := res.Msg.Response.(*idmv1.LoginResponse_AccessToken)
	if !ok {
		return fmt.Errorf("expected an access token but got %T", res.Msg.Response)
	}

	dummy := &http.Request{Header: http.Header{
		"Cookie": res.Header()["Set-Cookie"],
	}}

	refreshCookie, err := dummy.Cookie("cis_idm_refresh")
	if err != nil {
		return fmt.Errorf("failed to parse cis_idm_refresh cookie: %w", err)
	}

	root.tokens = TokenFile{
		AccessToken:   token.AccessToken.Token,
		RefreshCookie: refreshCookie.String(),
	}

	return root.writeTokenFile()
}

func (root *Root) refreshToken() error {
	if root.tokens.RefreshCookie == "" {
		return fmt.Errorf("failed to refresh: no refresh cookie available")
	}

	req := connect.NewRequest(&idmv1.RefreshTokenRequest{})
	req.Header().Add("Cookie", root.tokens.RefreshCookie)

	res, err := root.Auth().RefreshToken(root.Context(), req)
	if err == nil {
		root.tokens.AccessToken = res.Msg.GetAccessToken().GetToken()

		if err := root.writeTokenFile(); err != nil {
			return err
		}

	} else {
		return fmt.Errorf("failed to refresh: %w", err)
	}

	return nil
}

func (root *Root) writeTokenFile() error {
	blob, err := json.Marshal(root.tokens)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := os.WriteFile(root.TokenPath, blob, 0600); err != nil {
		return fmt.Errorf("write to %q: %w", root.TokenPath, err)
	}

	return nil
}

func (root *Root) readTokenFile() error {
	content, err := os.ReadFile(root.TokenPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", root.TokenPath, err)
	}

	if err := json.Unmarshal(content, &root.tokens); err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	return nil
}
