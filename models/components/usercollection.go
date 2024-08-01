// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UserCollection struct {
	// returns the pagination information
	Meta *PaginatedMeta `json:"meta,omitempty"`
	Data []User         `json:"data,omitempty"`
}

func (o *UserCollection) GetMeta() *PaginatedMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *UserCollection) GetData() []User {
	if o == nil {
		return nil
	}
	return o.Data
}
