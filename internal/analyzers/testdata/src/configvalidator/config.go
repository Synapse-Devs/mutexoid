package configvalidator

// Valid config struct with all required validation tags
type ValidConfig struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `koanf:"age" validate:"required,gt=0"`
	IsAdmin bool   `json:"is_admin" validate:"required"`
}

// Invalid config struct missing validation tags
type InvalidConfig struct {
	Name    string `json:"name"`     // want "config field Name should have a validate tag"
	Age     int    `koanf:"age"`     // want "config field Age should have a validate tag"
	IsAdmin bool   `json:"is_admin"` // want "config field IsAdmin should have a validate tag"
}

// Mixed config struct with some fields having validation
type MixedConfig struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `koanf:"age"`     // want "config field Age should have a validate tag"
	IsAdmin bool   `json:"is_admin"` // want "config field IsAdmin should have a validate tag"
}
