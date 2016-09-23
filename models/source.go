package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Source source

swagger:model Source
*/
type Source struct {

	/* Unique identifier representing a specific data source.

	Read Only: true
	*/
	ID string `json:"id,omitempty"`

	/* links
	 */
	Links *SourceLinks `json:"links,omitempty"`

	/* User facing name of data source

	Required: true
	*/
	Name *string `json:"name"`

	/* Password in cleartext!
	 */
	Password string `json:"password,omitempty"`

	/* Format of the data source

	Required: true
	*/
	Type *string `json:"type"`

	/* Username for authentication to data source
	 */
	Username string `json:"username,omitempty"`
}

// Validate validates this source
func (m *Source) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Source) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Source) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var sourceTypeTypePropEnum []interface{}

// prop value enum
func (m *Source) validateTypeEnum(path, location string, value string) error {
	if sourceTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["influx","influx-enterprise"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			sourceTypeTypePropEnum = append(sourceTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, sourceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Source) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

/*SourceLinks source links

swagger:model SourceLinks
*/
type SourceLinks struct {

	/* URL location of the monitored services endpoint for this source
	 */
	Monitored string `json:"monitored,omitempty"`

	/* URL location of the permissions endpoint for this source
	 */
	Permissions string `json:"permissions,omitempty"`

	/* URL location of proxy endpoint for this source
	 */
	Proxy string `json:"proxy,omitempty"`

	/* URL location of the roles endpoint for this source
	 */
	Roles string `json:"roles,omitempty"`

	/* Self link mapping to this resource
	 */
	Self string `json:"self,omitempty"`

	/* URL location of users endpoint for this source
	 */
	Users string `json:"users,omitempty"`
}

// Validate validates this source links
func (m *SourceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
