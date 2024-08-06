// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GroupConflict - The Group Conflict object contains information about a conflict in a control plane group.
type GroupConflict struct {
	// The ID of a control plane member of a control plane group.
	ClusterID *string `json:"cluster_id,omitempty"`
	// The description of the conflict.
	Description *string `json:"description,omitempty"`
	// A resource causing a conflict in a control plane group.
	Resource *GroupConflictResource `json:"resource,omitempty"`
}

func (o *GroupConflict) GetClusterID() *string {
	if o == nil {
		return nil
	}
	return o.ClusterID
}

func (o *GroupConflict) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GroupConflict) GetResource() *GroupConflictResource {
	if o == nil {
		return nil
	}
	return o.Resource
}
