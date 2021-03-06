// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MsgVpnLinks msg vpn links
// swagger:model MsgVpnLinks
type MsgVpnLinks struct {

	// The URI of this MsgVpn's aclProfiles collection.
	ACLProfilesURI string `json:"aclProfilesUri,omitempty"`

	// The URI of this MsgVpn's authorizationGroups collection.
	AuthorizationGroupsURI string `json:"authorizationGroupsUri,omitempty"`

	// The URI of this MsgVpn's bridges collection.
	BridgesURI string `json:"bridgesUri,omitempty"`

	// The URI of this MsgVpn's clientProfiles collection.
	ClientProfilesURI string `json:"clientProfilesUri,omitempty"`

	// The URI of this MsgVpn's clientUsernames collection.
	ClientUsernamesURI string `json:"clientUsernamesUri,omitempty"`

	// The URI of this MsgVpn's jndiConnectionFactories collection. Available since 2.2.
	JndiConnectionFactoriesURI string `json:"jndiConnectionFactoriesUri,omitempty"`

	// The URI of this MsgVpn's jndiQueues collection. Available since 2.2.
	JndiQueuesURI string `json:"jndiQueuesUri,omitempty"`

	// The URI of this MsgVpn's jndiTopics collection. Available since 2.2.
	JndiTopicsURI string `json:"jndiTopicsUri,omitempty"`

	// The URI of this MsgVpn's mqttSessions collection. Available since 2.1.
	MqttSessionsURI string `json:"mqttSessionsUri,omitempty"`

	// The URI of this MsgVpn's queues collection.
	QueuesURI string `json:"queuesUri,omitempty"`

	// The URI of this MsgVpn's replayLogs collection. Available since 2.10.
	ReplayLogsURI string `json:"replayLogsUri,omitempty"`

	// The URI of this MsgVpn's replicatedTopics collection. Available since 2.9.
	ReplicatedTopicsURI string `json:"replicatedTopicsUri,omitempty"`

	// The URI of this MsgVpn's restDeliveryPoints collection.
	RestDeliveryPointsURI string `json:"restDeliveryPointsUri,omitempty"`

	// The URI of this MsgVpn's sequencedTopics collection.
	SequencedTopicsURI string `json:"sequencedTopicsUri,omitempty"`

	// The URI of this MsgVpn's topicEndpoints collection. Available since 2.1.
	TopicEndpointsURI string `json:"topicEndpointsUri,omitempty"`

	// The URI of this MsgVpn object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn links
func (m *MsgVpnLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
