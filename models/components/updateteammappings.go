// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Mappings struct {
	Group   *string  `json:"group,omitempty"`
	TeamIds []string `json:"team_ids,omitempty"`
}

func (o *Mappings) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *Mappings) GetTeamIds() []string {
	if o == nil {
		return nil
	}
	return o.TeamIds
}

type UpdateTeamMappings struct {
	// The mappings object.
	Mappings []Mappings `json:"mappings,omitempty"`
}

func (o *UpdateTeamMappings) GetMappings() []Mappings {
	if o == nil {
		return nil
	}
	return o.Mappings
}
