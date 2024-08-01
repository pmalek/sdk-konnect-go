// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type KeySet struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *KeySet) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KeySet) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeySet) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *KeySet) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KeySet) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
