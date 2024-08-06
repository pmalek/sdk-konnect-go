// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ConsumerGroup struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *ConsumerGroup) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConsumerGroup) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConsumerGroup) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ConsumerGroup) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
