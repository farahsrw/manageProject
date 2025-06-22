package models

import "github.com/golang-jwt/jwt/v5"

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Project struct {
	XataID          string   `json:"xata_id,omitempty"`
	UserID          string   `json:"user,omitempty"`
	Name            string   `json:"name,omitempty"`
	RepoURL         string   `json:"repo_url,omitempty"`
	SiteURL         string   `json:"site_url,omitempty"`
	Description     string   `json:"description,omitempty"`
	Dependencies    []string `json:"dependencies,omitempty"`
	DevDependencies []string `json:"dev_dependencies,omitempty"`
	Status          string   `json:"status,omitempty"`
}

type RouteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

type UserResponse struct {
	XataID   string `json:"xata_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Claims struct {
	Username string `json:"username"`
	XataID   string `json:"xata_id"`
	jwt.RegisteredClaims
}
