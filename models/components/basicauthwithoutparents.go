// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type BasicAuthWithoutParents struct {
	Password *string  `json:"password,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Username *string  `json:"username,omitempty"`
}

func (o *BasicAuthWithoutParents) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *BasicAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BasicAuthWithoutParents) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
