package schema

// MethodParam holds the specification of an API interface method parameter.
type MethodParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

func NewMethodParam(
	name string,
	ty string,
	optional bool,
	description string,
) *MethodParam {
	return &MethodParam{
		Name:        name,
		Type:        ty,
		Optional:    optional,
		Description: description,
	}
}

// ValidateValue validates the given value against the parameter specification.
//
// Returns an error of type `*RequiredParameterError` if the parameter is required and the value is
// empty.
func (p *MethodParam) ValidateValue(value string) error {
	if !p.Optional && value == "" {
		return &RequiredParameterError{Param: p}
	}

	return nil
}
