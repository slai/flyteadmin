package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flyteorg/flyteadmin/auth/config"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestCookieManager_SetTokenCookies(t *testing.T) {
	ctx := context.Background()
	// These were generated for unit testing only.
	hashKeyEncoded := "wG4pE1ccdw/pHZ2ml8wrD5VJkOtLPmBpWbKHmezWXktGaFbRoAhXidWs8OpbA3y7N8vyZhz1B1E37+tShWC7gA" //nolint:goconst
	blockKeyEncoded := "afyABVgGOvWJFxVyOvCWCupoTn6BkNl4SOHmahho16Q"                                           //nolint:goconst
	cookieSetting := config.CookieSettings{
		SameSitePolicy:    config.SameSiteDefaultMode,
		DomainMatchPolicy: config.DomainMatchSubdomains,
	}
	manager, err := NewCookieManager(ctx, hashKeyEncoded, blockKeyEncoded, cookieSetting)
	assert.NoError(t, err)

	token := &oauth2.Token{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}

	token = token.WithExtra(map[string]interface{}{
		"id_token": "id token",
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/projects", nil)
	assert.NoError(t, err)
	err = manager.SetTokenCookies(ctx, req, w, token)
	assert.NoError(t, err)
	fmt.Println(w.Header().Get("Set-Cookie"))
	c := w.Result().Cookies()
	assert.Equal(t, "flyte_at", c[0].Name)
	assert.Equal(t, "flyte_idt", c[1].Name)
	assert.Equal(t, "flyte_rt", c[2].Name)
}

func TestCookieManager_RetrieveTokenValues(t *testing.T) {
	ctx := context.Background()
	// These were generated for unit testing only.
	hashKeyEncoded := "wG4pE1ccdw/pHZ2ml8wrD5VJkOtLPmBpWbKHmezWXktGaFbRoAhXidWs8OpbA3y7N8vyZhz1B1E37+tShWC7gA" //nolint:goconst
	blockKeyEncoded := "afyABVgGOvWJFxVyOvCWCupoTn6BkNl4SOHmahho16Q"                                           //nolint:goconst

	cookieSetting := config.CookieSettings{
		SameSitePolicy:    config.SameSiteDefaultMode,
		DomainMatchPolicy: config.DomainMatchSubdomains,
	}

	manager, err := NewCookieManager(ctx, hashKeyEncoded, blockKeyEncoded, cookieSetting)
	assert.NoError(t, err)

	token := &oauth2.Token{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}

	token = token.WithExtra(map[string]interface{}{
		"id_token": "id token",
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/projects", nil)
	assert.NoError(t, err)
	err = manager.SetTokenCookies(ctx, req, w, token)
	assert.NoError(t, err)

	cookies := w.Result().Cookies()
	req, err = http.NewRequest("GET", "/api/v1/projects", nil)
	assert.NoError(t, err)
	for _, c := range cookies {
		req.AddCookie(c)
	}

	idToken, access, refresh, err := manager.RetrieveTokenValues(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, "id token", idToken)
	assert.Equal(t, "access", access)
	assert.Equal(t, "refresh", refresh)
}

func TestGetLogoutAccessCookie(t *testing.T) {
	cookie := getLogoutAccessCookie()
	assert.True(t, time.Now().After(cookie.Expires))
}

func TestGetLogoutRefreshCookie(t *testing.T) {
	cookie := getLogoutRefreshCookie()
	assert.True(t, time.Now().After(cookie.Expires))
}

func TestCookieManager_DeleteCookies(t *testing.T) {
	ctx := context.Background()

	// These were generated for unit testing only.
	hashKeyEncoded := "wG4pE1ccdw/pHZ2ml8wrD5VJkOtLPmBpWbKHmezWXktGaFbRoAhXidWs8OpbA3y7N8vyZhz1B1E37+tShWC7gA" //nolint:goconst
	blockKeyEncoded := "afyABVgGOvWJFxVyOvCWCupoTn6BkNl4SOHmahho16Q"                                           //nolint:goconst
	cookieSetting := config.CookieSettings{
		SameSitePolicy:    config.SameSiteDefaultMode,
		DomainMatchPolicy: config.DomainMatchSubdomains,
	}

	manager, err := NewCookieManager(ctx, hashKeyEncoded, blockKeyEncoded, cookieSetting)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	manager.DeleteCookies(ctx, w)
	cookies := w.Result().Cookies()
	assert.Equal(t, 2, len(cookies))
	assert.True(t, time.Now().After(cookies[0].Expires))
	assert.True(t, time.Now().After(cookies[1].Expires))
}

func TestGetHTTPSameSitePolicy(t *testing.T) {
	ctx := context.Background()

	// These were generated for unit testing only.
	hashKeyEncoded := "wG4pE1ccdw/pHZ2ml8wrD5VJkOtLPmBpWbKHmezWXktGaFbRoAhXidWs8OpbA3y7N8vyZhz1B1E37+tShWC7gA" //nolint:goconst
	blockKeyEncoded := "afyABVgGOvWJFxVyOvCWCupoTn6BkNl4SOHmahho16Q"                                           //nolint:goconst
	cookieSetting := config.CookieSettings{
		SameSitePolicy:    config.SameSiteDefaultMode,
		DomainMatchPolicy: config.DomainMatchSubdomains,
	}

	manager, err := NewCookieManager(ctx, hashKeyEncoded, blockKeyEncoded, cookieSetting)
	assert.NoError(t, err)
	assert.Equal(t, http.SameSiteDefaultMode, manager.getHTTPSameSitePolicy())

	manager.sameSitePolicy = config.SameSiteLaxMode
	assert.Equal(t, http.SameSiteLaxMode, manager.getHTTPSameSitePolicy())

	manager.sameSitePolicy = config.SameSiteStrictMode
	assert.Equal(t, http.SameSiteStrictMode, manager.getHTTPSameSitePolicy())

	manager.sameSitePolicy = config.SameSiteNoneMode
	assert.Equal(t, http.SameSiteNoneMode, manager.getHTTPSameSitePolicy())
}

func TestGetCookieDomain(t *testing.T) {
	ctx := context.Background()

	// These were generated for unit testing only.
	hashKeyEncoded := "wG4pE1ccdw/pHZ2ml8wrD5VJkOtLPmBpWbKHmezWXktGaFbRoAhXidWs8OpbA3y7N8vyZhz1B1E37+tShWC7gA" //nolint:goconst
	blockKeyEncoded := "afyABVgGOvWJFxVyOvCWCupoTn6BkNl4SOHmahho16Q"                                           //nolint:goconst
	cookieSetting := config.CookieSettings{
		SameSitePolicy:    config.SameSiteDefaultMode,
		DomainMatchPolicy: config.DomainMatchSubdomains,
	}

	req, err := http.NewRequest("GET", "http://localhost/api/v1/projects", nil)
	assert.NoError(t, err)

	manager, err := NewCookieManager(ctx, hashKeyEncoded, blockKeyEncoded, cookieSetting)
	assert.NoError(t, err)
	assert.Equal(t, ".localhost", manager.getCookieDomain(req))

	manager.domainMatchPolicy = config.DomainMatchExact
	assert.Equal(t, "", manager.getCookieDomain(req))
}
