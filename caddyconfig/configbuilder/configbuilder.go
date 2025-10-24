package configbuilder

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig"
)

type Builder struct {
	c *caddy.Config

	warnings *[]caddyconfig.Warning
}

func New() *Builder {
	return &Builder{
		c: &caddy.Config{
			AppsRaw: make(caddy.ModuleMap),
		},
		warnings: &[]caddyconfig.Warning{},
	}
}

func From(cfg *caddy.Config) *Builder {
	return &Builder{
		c:        cfg,
		warnings: &[]caddyconfig.Warning{},
	}
}

// Build returns the built config and a slice of warnings.
func (b *Builder) Build() (*caddy.Config, *[]caddyconfig.Warning) {
	return b.c, b.warnings
}

// Config returns the built config.
func (b *Builder) Config() *caddy.Config {
	return b.c
}

// Warnings returns the current warnings
func (b *Builder) Warnings() *[]caddyconfig.Warning {
	return b.warnings
}

func (b *Builder) AddRawApp(name string, app json.RawMessage) {
	b.c.AppsRaw[name] = app
}

func (b *Builder) AddJsonApp(name string, val any) {
	b.c.AppsRaw[name] = caddyconfig.JSON(val, b.warnings)
}

func (b *Builder) SetStorage(
	storage caddy.StorageConverter,
) {
	b.c.StorageRaw = caddyconfig.JSONModuleObject(
		storage,
		"module",
		storage.(caddy.Module).CaddyModule().ID.Name(),
		b.warnings,
	)
}

func (b *Builder) Mutate(fn func(cfg *caddy.Config, warnings *[]caddyconfig.Warning)) {
	fn(b.c, b.warnings)
}
