// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

type TargetWithoutParents struct {
	// An optional set of strings associated with the Target for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The target address (ip or hostname) and port. If the hostname resolves to an SRV record, the `port` value will be overridden by the value from the DNS record.
	Target *string `json:"target,omitempty"`
	// The weight this target gets within the upstream loadbalancer (`0`-`65535`). If the hostname resolves to an SRV record, the `weight` value will be overridden by the value from the DNS record.
	Weight *int64 `default:"100" json:"weight"`
}

func (t TargetWithoutParents) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TargetWithoutParents) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TargetWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TargetWithoutParents) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *TargetWithoutParents) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}
