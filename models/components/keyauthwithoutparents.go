// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type KeyAuthWithoutParents struct {
	Key  *string  `json:"key,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func (o *KeyAuthWithoutParents) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *KeyAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
