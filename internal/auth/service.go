package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"src/internal/pkg/handler"
	"strconv"
	"strings"
)

const getPermissionsEndpoint = "http://confetti-cms-cms__auth:80/users/me/permissions"

type Permission struct {
	Id string `json:"id"`
}

type User struct {
	Id          string       `json:"id"`
	UserName    string       `json:"username"`
	Name        string       `json:"name"`
	Picture     string       `json:"picture_url"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type Service struct {
	user  User
	error error
}

func (g Service) InitByRequest(request *http.Request) Service {
	accessToken := request.Header.Get("Authorization")
	if accessToken == "" {
		g.error = handler.NewUserError("header `Authorization` not found or empty", http.StatusUnsupportedMediaType)
		return g
	}
	// Create request
	req, err := http.NewRequest(
		http.MethodGet,
		getPermissionsEndpoint,
		nil,
	)
	if err != nil {
		g.error = errors.New("internal server error: sy8oagfd")
		fmt.Printf("sy8oagfd: %s\n", err)
		return g
	}
	h := req.Header
	h.Add("Content-Type", "application/json")
	h.Add("Authorization", accessToken)
	req.Header = h
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.error = handler.NewSystemError(fmt.Errorf("errored when sending request to the server: %w", err), "h34jkhjk34")
		return g
	}
	// Read response
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		g.error = handler.NewSystemError(fmt.Errorf("errored when when reading response: %w", err), "o879der")
		return g
	}

	if resp.StatusCode > 299 {
		if resp.StatusCode >= 500 {
			g.error = handler.NewSystemError(fmt.Errorf("error from response: status: %s, response body %s", resp.Status, string(response)), "9jeggre")
			return g
		} else {
			g.error = handler.NewUserError(
				"error with status: "+strconv.Itoa(resp.StatusCode)+
					" with request: "+http.MethodGet+" "+getPermissionsEndpoint+
					" and response: "+string(response), resp.StatusCode)
			return g
		}
	}

	err = json.Unmarshal(response, &g.user)
	if err != nil {
		g.error = handler.NewSystemError(fmt.Errorf("errored when unmarchal response: hasd9ure: %s, with response: %s", err, string(response)), "hasd9ure")
		return g
	}

	return g
}

func (g Service) Can(checkPermissions ...string) error {
	if g.error != nil {
		println("Error: " + g.error.Error())
		return g.error
	}
	for _, permission := range checkPermissions {
		if !g.hasPermission(permission) {
			return handler.NewUserError("You do not have the required privileges for permission: "+permission, http.StatusForbidden)
		}
	}
	return nil
}

func (g Service) hasPermission(checkPermission string) bool {

	for _, permission := range g.user.Permissions {
		// if exact match
		if checkPermission == permission.Id {
			return true
		}
		// if given permission e.g. `/image/store` starts with `/image/`
		if strings.HasPrefix(checkPermission, permission.Id+"/") {
			return true
		}
	}
	return false
}
