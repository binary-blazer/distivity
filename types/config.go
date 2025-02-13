package types

import "github.com/gofiber/fiber/v2"

type Config struct {
	Webserver Webserver `json:"webserver"`
	App       App       `json:"app"`
	Routes    []Route   `json:"routes"`
}

type WebserverPaths struct {
	Root   string `json:"root"`
	Static string `json:"static"`
}

type Webserver struct {
	Port  int            `json:"port"`
	Paths WebserverPaths `json:"paths"`
	Fiber FiberSettings  `json:"fiber"`
}

type App struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Environment string   `json:"environment"`
	Authors     []Author `json:"authors"`
}

type Author struct {
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Email    string `json:"email"`
}

type FiberSettings struct {
	Prefork       bool `json:"prefork"`
	CaseSensitive bool `json:"case_sensitive"`
	StrictRouting bool `json:"strict_routing"`
}

type Route struct {
	Path    string        `json:"path"`
	Handler fiber.Handler `json:"handler"`
}
