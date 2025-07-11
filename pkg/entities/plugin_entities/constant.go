package plugin_entities

const (
	SECRET_INPUT   = "secret-input"
	TEXT_INPUT     = "text-input"
	SELECT         = "select"
	STRING         = "string"
	NUMBER         = "number"
	FILE           = "file"
	FILES          = "files"
	BOOLEAN        = "boolean"
	APP_SELECTOR   = "app-selector"
	MODEL_SELECTOR = "model-selector"
	// TOOL_SELECTOR  = "tool-selector"
	TOOLS_SELECTOR = "array[tools]"
	ANY            = "any"
	// DynamicSelect
	DYNAMIC_SELECT = "dynamic-select"
)

type ParameterOption struct {
	Value string     `json:"value" yaml:"value" validate:"required"`
	Label I18nObject `json:"label" yaml:"label" validate:"required"`
	Icon  string     `json:"icon" yaml:"icon" validate:"omitempty"`
}
