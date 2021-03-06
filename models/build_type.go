// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuildType Represents a build configuration.
//
// swagger:model buildType
type BuildType struct {

	// agent requirements
	AgentRequirements *AgentRequirements `json:"agent-requirements,omitempty"`

	// artifact dependencies
	ArtifactDependencies *ArtifactDependencies `json:"artifact-dependencies,omitempty"`

	// branches
	Branches *Branches `json:"branches,omitempty"`

	// builds
	Builds *Builds `json:"builds,omitempty"`

	// compatible agents
	CompatibleAgents *Agents `json:"compatibleAgents,omitempty"`

	// description
	Description string `json:"description,omitempty" xml:"description,attr,omitempty"`

	// external status allowed
	ExternalStatusAllowed bool `json:"externalStatusAllowed,omitempty" xml:"externalStatusAllowed,attr,omitempty"`

	// features
	Features *Features `json:"features,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href,attr,omitempty"`

	// id
	ID string `json:"id,omitempty" xml:"id,attr,omitempty"`

	// inherited
	Inherited bool `json:"inherited,omitempty" xml:"inherited,attr,omitempty"`

	// internal Id
	InternalID string `json:"internalId,omitempty" xml:"internalId,attr,omitempty"`

	// investigations
	Investigations *Investigations `json:"investigations,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator,attr,omitempty"`

	// name
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`

	// parameters
	Parameters *Properties `json:"parameters,omitempty"`

	// paused
	Paused bool `json:"paused,omitempty" xml:"paused,attr,omitempty"`

	// project
	Project *Project `json:"project,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty" xml:"projectId,attr,omitempty"`

	// project internal Id
	ProjectInternalID string `json:"projectInternalId,omitempty" xml:"projectInternalId,attr,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty" xml:"projectName,attr,omitempty"`

	// settings
	Settings *Properties `json:"settings,omitempty"`

	// snapshot dependencies
	SnapshotDependencies *SnapshotDependencies `json:"snapshot-dependencies,omitempty"`

	// steps
	Steps *Steps `json:"steps,omitempty"`

	// template
	Template *BuildType `json:"template,omitempty"`

	// template flag
	TemplateFlag bool `json:"templateFlag,omitempty" xml:"templateFlag,attr,omitempty"`

	// templates
	Templates *BuildTypes `json:"templates,omitempty"`

	// triggers
	Triggers *Triggers `json:"triggers,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type,attr,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty" xml:"uuid,attr,omitempty"`

	// vcs root entries
	VcsRootEntries *VcsRootEntries `json:"vcs-root-entries,omitempty"`

	// vcs root instances
	VcsRootInstances *VcsRootInstances `json:"vcsRootInstances,omitempty"`

	// web Url
	WebURL string `json:"webUrl,omitempty" xml:"webUrl,attr,omitempty"`
}

// Validate validates this build type
func (m *BuildType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifactDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompatibleAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvestigations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsRootEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsRootInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildType) validateAgentRequirements(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentRequirements) { // not required
		return nil
	}

	if m.AgentRequirements != nil {
		if err := m.AgentRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent-requirements")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent-requirements")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateArtifactDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.ArtifactDependencies) { // not required
		return nil
	}

	if m.ArtifactDependencies != nil {
		if err := m.ArtifactDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact-dependencies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact-dependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateBranches(formats strfmt.Registry) error {
	if swag.IsZero(m.Branches) { // not required
		return nil
	}

	if m.Branches != nil {
		if err := m.Branches.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branches")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branches")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateBuilds(formats strfmt.Registry) error {
	if swag.IsZero(m.Builds) { // not required
		return nil
	}

	if m.Builds != nil {
		if err := m.Builds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("builds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("builds")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateCompatibleAgents(formats strfmt.Registry) error {
	if swag.IsZero(m.CompatibleAgents) { // not required
		return nil
	}

	if m.CompatibleAgents != nil {
		if err := m.CompatibleAgents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compatibleAgents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compatibleAgents")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateInvestigations(formats strfmt.Registry) error {
	if swag.IsZero(m.Investigations) { // not required
		return nil
	}

	if m.Investigations != nil {
		if err := m.Investigations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("investigations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("investigations")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateSnapshotDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotDependencies) { // not required
		return nil
	}

	if m.SnapshotDependencies != nil {
		if err := m.SnapshotDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot-dependencies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot-dependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	if m.Steps != nil {
		if err := m.Steps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if m.Templates != nil {
		if err := m.Templates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateTriggers(formats strfmt.Registry) error {
	if swag.IsZero(m.Triggers) { // not required
		return nil
	}

	if m.Triggers != nil {
		if err := m.Triggers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("triggers")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateVcsRootEntries(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsRootEntries) { // not required
		return nil
	}

	if m.VcsRootEntries != nil {
		if err := m.VcsRootEntries.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs-root-entries")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcs-root-entries")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateVcsRootInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsRootInstances) { // not required
		return nil
	}

	if m.VcsRootInstances != nil {
		if err := m.VcsRootInstances.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsRootInstances")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcsRootInstances")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this build type based on the context it is used
func (m *BuildType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgentRequirements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArtifactDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBranches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuilds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompatibleAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvestigations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTriggers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVcsRootEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVcsRootInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildType) contextValidateAgentRequirements(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentRequirements != nil {
		if err := m.AgentRequirements.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent-requirements")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent-requirements")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateArtifactDependencies(ctx context.Context, formats strfmt.Registry) error {

	if m.ArtifactDependencies != nil {
		if err := m.ArtifactDependencies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact-dependencies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifact-dependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateBranches(ctx context.Context, formats strfmt.Registry) error {

	if m.Branches != nil {
		if err := m.Branches.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branches")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("branches")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateBuilds(ctx context.Context, formats strfmt.Registry) error {

	if m.Builds != nil {
		if err := m.Builds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("builds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("builds")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateCompatibleAgents(ctx context.Context, formats strfmt.Registry) error {

	if m.CompatibleAgents != nil {
		if err := m.CompatibleAgents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compatibleAgents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compatibleAgents")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	if m.Features != nil {
		if err := m.Features.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateInvestigations(ctx context.Context, formats strfmt.Registry) error {

	if m.Investigations != nil {
		if err := m.Investigations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("investigations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("investigations")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {
		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {
		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateSnapshotDependencies(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotDependencies != nil {
		if err := m.SnapshotDependencies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot-dependencies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot-dependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	if m.Steps != nil {
		if err := m.Steps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	if m.Templates != nil {
		if err := m.Templates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateTriggers(ctx context.Context, formats strfmt.Registry) error {

	if m.Triggers != nil {
		if err := m.Triggers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("triggers")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateVcsRootEntries(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsRootEntries != nil {
		if err := m.VcsRootEntries.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs-root-entries")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcs-root-entries")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) contextValidateVcsRootInstances(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsRootInstances != nil {
		if err := m.VcsRootInstances.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsRootInstances")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcsRootInstances")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildType) UnmarshalBinary(b []byte) error {
	var res BuildType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
