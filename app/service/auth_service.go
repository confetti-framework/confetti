package service

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "src/app/entity"
    "strconv"
)

const getPermissionsEndpoint = "http://confetti-cms-auth:80/users/me/permissions"

type AuthService struct {
    user  entity.User
    error error
}

func (g AuthService) InitByRequest(request *http.Request) AuthService {
    accessToken := request.Header.Get("Authorization")
    if accessToken == "" {
        g.error = errors.New("header Authorization not found or empty")
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
        g.error = errors.New("internal server error: h34jkhjk34")
        fmt.Printf("Errored when sending request to the server: h34jkhjk34: %s", err)
        return g
    }
    // Read response
    defer resp.Body.Close()
    response, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        g.error = errors.New("internal server error: o879der")
        fmt.Printf("Errored when reading response: o879der: %s", err)
        return g
    }

    if resp.StatusCode > 299 {
        if resp.StatusCode >= 500 {
            g.error = errors.New("internal server error: 9jeggre")
            fmt.Printf("Error from response: 9jeggre: status: %s, response body %s", resp.Status, string(response))
            return g
        } else {
            g.error = errors.New(
                "error with status: " + strconv.Itoa(resp.StatusCode) +
                    " with request: " + http.MethodGet + " " + getPermissionsEndpoint +
                    " and response: " + string(response),
            )
            return g
        }
    }

    err = json.Unmarshal(response, &g.user)
    if err != nil {
        g.error = errors.New("internal server error: hasd9ure")
        fmt.Printf("Errored when unmarchal response: hasd9ure: %s, with response: %s", err, string(response))
        return g
    }

    return g
}

func (g AuthService) Can(checkPermissions ...string) error {
    if g.error != nil {
        println("Error: " + g.error.Error())
        return g.error
    }
    for _, permission := range checkPermissions {
        if !g.hasPermission(permission) {
            return errors.New("your permission does not have the required privileges. Permission: " + permission)
        }
    }
    return nil
}

func (g AuthService) hasPermission(checkPermission string) bool {

    for _, permission := range g.user.Permissions {
        if checkPermission == permission.Id {
            return true
        }
    }
    return false
}