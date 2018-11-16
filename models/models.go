package models

// APIError represents an api error
type APIError struct {
	Code    int
	Message string
}

// User represents the Keycloak User Structure
type User struct {
	ID                         string        `json:"id,omitempty"`
	CreatedTimestamp           int64         `json:"createdTimestamp,omitempty"`
	Username                   string        `json:"username,omitempty"`
	Enabled                    bool          `json:"enabled,omitempty"`
	Totp                       bool          `json:"totp,omitempty"`
	EmailVerified              bool          `json:"emailVerified,omitempty"`
	FirstName                  string        `json:"firstName,omitempty"`
	LastName                   string        `json:"lastName,omitempty"`
	Email                      string        `json:"email,omitempty"`
	FederationLink             string        `json:"federationLink,omitempty"`
	Attributes                 Attributes    `json:"attributes,omitempty"`
	DisableableCredentialTypes []interface{} `json:"disableableCredentialTypes,omitempty"`
	RequiredActions            []interface{} `json:"requiredActions,omitempty"`
	Access                     Access        `json:"access,omitempty"`
}

// Attributes holds Attributes
type Attributes struct {
	LDAPENTRYDN []string `json:"LDAP_ENTRY_DN"`
	LDAPID      []string `json:"LDAP_ID"`
}

// Access represents access
type Access struct {
	ManageGroupMembership bool `json:"manageGroupMembership"`
	View                  bool `json:"view"`
	MapRoles              bool `json:"mapRoles"`
	Impersonate           bool `json:"impersonate"`
	Manage                bool `json:"manage"`
}

// UserGroup is a UserGroup
type UserGroup struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

// Group is a Group
type Group struct {
	ID        string        `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	Path      string        `json:"path,omitempty"`
	SubGroups []interface{} `json:"subGroups,omitempty"`
}

// Role is a role
type Role struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	ScopeParamRequired bool   `json:"scopeParamRequired,omitempty"`
	Composite          bool   `json:"composite,omitempty"`
	ClientRole         bool   `json:"clientRole,omitempty"`
	ContainerID        string `json:"containerId,omitempty"`
	Description        string `json:"description,omitempty,omitempty"`
}

// RoleMapping is a role mapping
type RoleMapping struct {
	ID       string                  `json:"id"`
	Client   string                  `json:"client"`
	Mappings []ClientRoleMappingRole `json:"mappings"`
}

// ClientRoleMappingRole is a client role mapping role
type ClientRoleMappingRole struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	ScopeParamRequired bool   `json:"scopeParamRequired"`
	Composite          bool   `json:"composite"`
	ClientRole         bool   `json:"clientRole"`
	ContainerID        string `json:"containerId"`
}

// Client is a Client
type Client struct {
	ID       string `json:"id"`
	ClientID string `json:"clientId"`
}
