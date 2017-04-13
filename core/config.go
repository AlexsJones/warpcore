package core

type Resource struct {
	Name     string
	Type     string
	Provider string
}

type Variable struct {
	Name         string
	DeclaredType string `mapstructure:"type"`
	Default      interface{}
	Description  string
}

type ProviderConfig struct {
	Name  string
	Alias string
}

type Config struct {
	Resources       []*Resource
	ProviderConfigs []*ProviderConfig
	Variables       []*Variable
}
