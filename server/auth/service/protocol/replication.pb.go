// Code generated by protoc-gen-go.
// source: replication.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	replication.proto

It has these top-level messages:
	ServiceLinkTicket
	ServiceLinkRequest
	ServiceLinkResponse
	AuthGroup
	AuthSecret
	AuthIPWhitelist
	AuthIPWhitelistAssignment
	AuthDB
	AuthDBRevision
	ChangeNotification
	ReplicationPushRequest
	ReplicationPushResponse
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Status codes.
type ServiceLinkResponse_Status int32

const (
	// The service is now linked and primary will be pushing updates to it.
	ServiceLinkResponse_SUCCESS ServiceLinkResponse_Status = 0
	// Primary do not replies.
	ServiceLinkResponse_TRANSPORT_ERROR ServiceLinkResponse_Status = 1
	// Linking ticket is invalid or expired.
	ServiceLinkResponse_BAD_TICKET ServiceLinkResponse_Status = 2
	// Linking ticket was generated for another app, not the calling one.
	ServiceLinkResponse_AUTH_ERROR ServiceLinkResponse_Status = 3
)

var ServiceLinkResponse_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "TRANSPORT_ERROR",
	2: "BAD_TICKET",
	3: "AUTH_ERROR",
}
var ServiceLinkResponse_Status_value = map[string]int32{
	"SUCCESS":         0,
	"TRANSPORT_ERROR": 1,
	"BAD_TICKET":      2,
	"AUTH_ERROR":      3,
}

func (x ServiceLinkResponse_Status) Enum() *ServiceLinkResponse_Status {
	p := new(ServiceLinkResponse_Status)
	*p = x
	return p
}
func (x ServiceLinkResponse_Status) String() string {
	return proto.EnumName(ServiceLinkResponse_Status_name, int32(x))
}
func (x *ServiceLinkResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServiceLinkResponse_Status_value, data, "ServiceLinkResponse_Status")
	if err != nil {
		return err
	}
	*x = ServiceLinkResponse_Status(value)
	return nil
}
func (ServiceLinkResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// Overall status of the operation.
type ReplicationPushResponse_Status int32

const (
	// Replica accepted the push request and updated its copy of auth db.
	ReplicationPushResponse_APPLIED ReplicationPushResponse_Status = 0
	// Replica has a newer version of AuthDB, the push request is skipped.
	ReplicationPushResponse_SKIPPED ReplicationPushResponse_Status = 1
	// Non fatal error happened, the push request may be retried.
	ReplicationPushResponse_TRANSIENT_ERROR ReplicationPushResponse_Status = 2
	// Fatal error happened, the push request must not be retried.
	ReplicationPushResponse_FATAL_ERROR ReplicationPushResponse_Status = 3
)

var ReplicationPushResponse_Status_name = map[int32]string{
	0: "APPLIED",
	1: "SKIPPED",
	2: "TRANSIENT_ERROR",
	3: "FATAL_ERROR",
}
var ReplicationPushResponse_Status_value = map[string]int32{
	"APPLIED":         0,
	"SKIPPED":         1,
	"TRANSIENT_ERROR": 2,
	"FATAL_ERROR":     3,
}

func (x ReplicationPushResponse_Status) Enum() *ReplicationPushResponse_Status {
	p := new(ReplicationPushResponse_Status)
	*p = x
	return p
}
func (x ReplicationPushResponse_Status) String() string {
	return proto.EnumName(ReplicationPushResponse_Status_name, int32(x))
}
func (x *ReplicationPushResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReplicationPushResponse_Status_value, data, "ReplicationPushResponse_Status")
	if err != nil {
		return err
	}
	*x = ReplicationPushResponse_Status(value)
	return nil
}
func (ReplicationPushResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11, 0}
}

// Error codes, for TRANSIENT_ERROR and FATAL_ERROR statuses.
type ReplicationPushResponse_ErrorCode int32

const (
	// Trying to push an update to service that is not a replica.
	ReplicationPushResponse_NOT_A_REPLICA ReplicationPushResponse_ErrorCode = 1
	// Replica doesn't know about the service that pushing the update.
	ReplicationPushResponse_FORBIDDEN ReplicationPushResponse_ErrorCode = 2
	// Signature headers are missing.
	ReplicationPushResponse_MISSING_SIGNATURE ReplicationPushResponse_ErrorCode = 3
	// Signature is not valid.
	ReplicationPushResponse_BAD_SIGNATURE ReplicationPushResponse_ErrorCode = 4
	// Format of the request is not valid.
	ReplicationPushResponse_BAD_REQUEST ReplicationPushResponse_ErrorCode = 5
)

var ReplicationPushResponse_ErrorCode_name = map[int32]string{
	1: "NOT_A_REPLICA",
	2: "FORBIDDEN",
	3: "MISSING_SIGNATURE",
	4: "BAD_SIGNATURE",
	5: "BAD_REQUEST",
}
var ReplicationPushResponse_ErrorCode_value = map[string]int32{
	"NOT_A_REPLICA":     1,
	"FORBIDDEN":         2,
	"MISSING_SIGNATURE": 3,
	"BAD_SIGNATURE":     4,
	"BAD_REQUEST":       5,
}

func (x ReplicationPushResponse_ErrorCode) Enum() *ReplicationPushResponse_ErrorCode {
	p := new(ReplicationPushResponse_ErrorCode)
	*p = x
	return p
}
func (x ReplicationPushResponse_ErrorCode) String() string {
	return proto.EnumName(ReplicationPushResponse_ErrorCode_name, int32(x))
}
func (x *ReplicationPushResponse_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReplicationPushResponse_ErrorCode_value, data, "ReplicationPushResponse_ErrorCode")
	if err != nil {
		return err
	}
	*x = ReplicationPushResponse_ErrorCode(value)
	return nil
}
func (ReplicationPushResponse_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11, 1}
}

// Generated by Primary, passed to Replica to initiate linking process.
type ServiceLinkTicket struct {
	// GAE application ID of Primary that generated this ticket. Replica will send
	// ServiceLinkRequest to this service when it processes the ticket.
	PrimaryId *string `protobuf:"bytes,1,req,name=primary_id,json=primaryId" json:"primary_id,omitempty"`
	// URL to the root page of a primary service, i.e. https://<...>.appspot.com.
	// Useful when testing on dev appserver and on non-default version.
	PrimaryUrl *string `protobuf:"bytes,2,req,name=primary_url,json=primaryUrl" json:"primary_url,omitempty"`
	// Identity of a user that generated this ticket.
	GeneratedBy *string `protobuf:"bytes,3,req,name=generated_by,json=generatedBy" json:"generated_by,omitempty"`
	// Opaque blob passed back to Primary in ServiceLinkRequest. Its exact
	// structure is an implementation detail of Primary. It contains app_id of
	// a replica this ticket is intended for, timestamp and HMAC tag.
	Ticket           []byte `protobuf:"bytes,4,req,name=ticket" json:"ticket,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServiceLinkTicket) Reset()                    { *m = ServiceLinkTicket{} }
func (m *ServiceLinkTicket) String() string            { return proto.CompactTextString(m) }
func (*ServiceLinkTicket) ProtoMessage()               {}
func (*ServiceLinkTicket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceLinkTicket) GetPrimaryId() string {
	if m != nil && m.PrimaryId != nil {
		return *m.PrimaryId
	}
	return ""
}

func (m *ServiceLinkTicket) GetPrimaryUrl() string {
	if m != nil && m.PrimaryUrl != nil {
		return *m.PrimaryUrl
	}
	return ""
}

func (m *ServiceLinkTicket) GetGeneratedBy() string {
	if m != nil && m.GeneratedBy != nil {
		return *m.GeneratedBy
	}
	return ""
}

func (m *ServiceLinkTicket) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

// Sent from Replica to Primary via direct serivce <-> service HTTP call,
// replicas app_id would be available via X-Appengine-Inbound-Appid header.
type ServiceLinkRequest struct {
	// Same ticket that was passed to Replica via ServiceLinkTicket.
	Ticket []byte `protobuf:"bytes,1,req,name=ticket" json:"ticket,omitempty"`
	// URL to use when making requests to Replica from Primary.
	ReplicaUrl *string `protobuf:"bytes,2,req,name=replica_url,json=replicaUrl" json:"replica_url,omitempty"`
	// Identity of a user that accepted the ticket and initiated this request.
	InitiatedBy      *string `protobuf:"bytes,3,req,name=initiated_by,json=initiatedBy" json:"initiated_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServiceLinkRequest) Reset()                    { *m = ServiceLinkRequest{} }
func (m *ServiceLinkRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceLinkRequest) ProtoMessage()               {}
func (*ServiceLinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceLinkRequest) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func (m *ServiceLinkRequest) GetReplicaUrl() string {
	if m != nil && m.ReplicaUrl != nil {
		return *m.ReplicaUrl
	}
	return ""
}

func (m *ServiceLinkRequest) GetInitiatedBy() string {
	if m != nil && m.InitiatedBy != nil {
		return *m.InitiatedBy
	}
	return ""
}

// Primary's response to ServiceLinkRequest. Always returned with HTTP code 200.
type ServiceLinkResponse struct {
	Status           *ServiceLinkResponse_Status `protobuf:"varint,1,req,name=status,enum=protocol.ServiceLinkResponse_Status" json:"status,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ServiceLinkResponse) Reset()                    { *m = ServiceLinkResponse{} }
func (m *ServiceLinkResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceLinkResponse) ProtoMessage()               {}
func (*ServiceLinkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceLinkResponse) GetStatus() ServiceLinkResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ServiceLinkResponse_SUCCESS
}

// Some user group. Corresponds to AuthGroup entity in model.py.
type AuthGroup struct {
	// Name of the group.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// List of members that are explicitly in this group.
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
	// List of identity-glob expressions (like 'user:*@example.com').
	Globs []string `protobuf:"bytes,3,rep,name=globs" json:"globs,omitempty"`
	// List of nested group names.
	Nested []string `protobuf:"bytes,4,rep,name=nested" json:"nested,omitempty"`
	// Human readable description.
	Description *string `protobuf:"bytes,5,req,name=description" json:"description,omitempty"`
	// When the group was created. Microseconds since epoch.
	CreatedTs *int64 `protobuf:"varint,6,req,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	// Who created the group.
	CreatedBy *string `protobuf:"bytes,7,req,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// When the group was modified last time. Microseconds since epoch.
	ModifiedTs *int64 `protobuf:"varint,8,req,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
	// Who modified the group last time.
	ModifiedBy *string `protobuf:"bytes,9,req,name=modified_by,json=modifiedBy" json:"modified_by,omitempty"`
	// A name of the group that can modify or delete this group.
	Owners           *string `protobuf:"bytes,10,opt,name=owners" json:"owners,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthGroup) Reset()                    { *m = AuthGroup{} }
func (m *AuthGroup) String() string            { return proto.CompactTextString(m) }
func (*AuthGroup) ProtoMessage()               {}
func (*AuthGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthGroup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthGroup) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *AuthGroup) GetGlobs() []string {
	if m != nil {
		return m.Globs
	}
	return nil
}

func (m *AuthGroup) GetNested() []string {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *AuthGroup) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *AuthGroup) GetCreatedTs() int64 {
	if m != nil && m.CreatedTs != nil {
		return *m.CreatedTs
	}
	return 0
}

func (m *AuthGroup) GetCreatedBy() string {
	if m != nil && m.CreatedBy != nil {
		return *m.CreatedBy
	}
	return ""
}

func (m *AuthGroup) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

func (m *AuthGroup) GetModifiedBy() string {
	if m != nil && m.ModifiedBy != nil {
		return *m.ModifiedBy
	}
	return ""
}

func (m *AuthGroup) GetOwners() string {
	if m != nil && m.Owners != nil {
		return *m.Owners
	}
	return ""
}

// Some secret blob. Corresponds to AuthSecret entity in model.py.
type AuthSecret struct {
	// Name of the secret.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Last several values of a secret, with current value in front.
	Values [][]byte `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	// When secret was modified last time. Microseconds since epoch.
	ModifiedTs *int64 `protobuf:"varint,3,req,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
	// Who modified the secret last time.
	ModifiedBy       *string `protobuf:"bytes,4,req,name=modified_by,json=modifiedBy" json:"modified_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthSecret) Reset()                    { *m = AuthSecret{} }
func (m *AuthSecret) String() string            { return proto.CompactTextString(m) }
func (*AuthSecret) ProtoMessage()               {}
func (*AuthSecret) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthSecret) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthSecret) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *AuthSecret) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

func (m *AuthSecret) GetModifiedBy() string {
	if m != nil && m.ModifiedBy != nil {
		return *m.ModifiedBy
	}
	return ""
}

// A named set of whitelisted IP addresses. Corresponds to AuthIPWhitelist
// entity in model.py.
type AuthIPWhitelist struct {
	// Name of the IP whitelist.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The list of IP subnets.
	Subnets []string `protobuf:"bytes,2,rep,name=subnets" json:"subnets,omitempty"`
	// Human readable description.
	Description *string `protobuf:"bytes,3,req,name=description" json:"description,omitempty"`
	// When the list was created. Microseconds since epoch.
	CreatedTs *int64 `protobuf:"varint,4,req,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	// Who created the list.
	CreatedBy *string `protobuf:"bytes,5,req,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// When the list was modified. Microseconds since epoch.
	ModifiedTs *int64 `protobuf:"varint,6,req,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
	// Who modified the list the last time.
	ModifiedBy       *string `protobuf:"bytes,7,req,name=modified_by,json=modifiedBy" json:"modified_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthIPWhitelist) Reset()                    { *m = AuthIPWhitelist{} }
func (m *AuthIPWhitelist) String() string            { return proto.CompactTextString(m) }
func (*AuthIPWhitelist) ProtoMessage()               {}
func (*AuthIPWhitelist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AuthIPWhitelist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthIPWhitelist) GetSubnets() []string {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *AuthIPWhitelist) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *AuthIPWhitelist) GetCreatedTs() int64 {
	if m != nil && m.CreatedTs != nil {
		return *m.CreatedTs
	}
	return 0
}

func (m *AuthIPWhitelist) GetCreatedBy() string {
	if m != nil && m.CreatedBy != nil {
		return *m.CreatedBy
	}
	return ""
}

func (m *AuthIPWhitelist) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

func (m *AuthIPWhitelist) GetModifiedBy() string {
	if m != nil && m.ModifiedBy != nil {
		return *m.ModifiedBy
	}
	return ""
}

// A pair (identity, IP whitelist name) plus some metadata. Corresponds to
// AuthIPWhitelistAssignments.Assignment model in model.py.
type AuthIPWhitelistAssignment struct {
	// Identity name to limit by IP whitelist.
	Identity *string `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	// Name of IP whitelist to use (see AuthIPWhitelist).
	IpWhitelist *string `protobuf:"bytes,2,req,name=ip_whitelist,json=ipWhitelist" json:"ip_whitelist,omitempty"`
	// Why the assignment was created.
	Comment *string `protobuf:"bytes,3,req,name=comment" json:"comment,omitempty"`
	// When the assignment was created. Microseconds since epoch.
	CreatedTs *int64 `protobuf:"varint,4,req,name=created_ts,json=createdTs" json:"created_ts,omitempty"`
	// Who created the assignment.
	CreatedBy        *string `protobuf:"bytes,5,req,name=created_by,json=createdBy" json:"created_by,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthIPWhitelistAssignment) Reset()                    { *m = AuthIPWhitelistAssignment{} }
func (m *AuthIPWhitelistAssignment) String() string            { return proto.CompactTextString(m) }
func (*AuthIPWhitelistAssignment) ProtoMessage()               {}
func (*AuthIPWhitelistAssignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AuthIPWhitelistAssignment) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

func (m *AuthIPWhitelistAssignment) GetIpWhitelist() string {
	if m != nil && m.IpWhitelist != nil {
		return *m.IpWhitelist
	}
	return ""
}

func (m *AuthIPWhitelistAssignment) GetComment() string {
	if m != nil && m.Comment != nil {
		return *m.Comment
	}
	return ""
}

func (m *AuthIPWhitelistAssignment) GetCreatedTs() int64 {
	if m != nil && m.CreatedTs != nil {
		return *m.CreatedTs
	}
	return 0
}

func (m *AuthIPWhitelistAssignment) GetCreatedBy() string {
	if m != nil && m.CreatedBy != nil {
		return *m.CreatedBy
	}
	return ""
}

// An entire database of auth configuration that is being replicated.
// Corresponds to AuthGlobalConfig entity in model.py, plus a list of all groups
// and a list of global secrets.
type AuthDB struct {
	// OAuth2 client_id to use to mint new OAuth2 tokens.
	OauthClientId *string `protobuf:"bytes,1,req,name=oauth_client_id,json=oauthClientId" json:"oauth_client_id,omitempty"`
	// OAuth2 client secret. Not so secret really, since it's passed to clients.
	OauthClientSecret *string `protobuf:"bytes,2,req,name=oauth_client_secret,json=oauthClientSecret" json:"oauth_client_secret,omitempty"`
	// Additional OAuth2 client_ids allowed to access the services.
	OauthAdditionalClientIds []string `protobuf:"bytes,3,rep,name=oauth_additional_client_ids,json=oauthAdditionalClientIds" json:"oauth_additional_client_ids,omitempty"`
	// All groups.
	Groups []*AuthGroup `protobuf:"bytes,4,rep,name=groups" json:"groups,omitempty"`
	// Global secrets shared between services.
	Secrets []*AuthSecret `protobuf:"bytes,5,rep,name=secrets" json:"secrets,omitempty"`
	// All IP whitelists.
	IpWhitelists []*AuthIPWhitelist `protobuf:"bytes,6,rep,name=ip_whitelists,json=ipWhitelists" json:"ip_whitelists,omitempty"`
	// Mapping 'account -> IP whitlist to use for that account'.
	IpWhitelistAssignments []*AuthIPWhitelistAssignment `protobuf:"bytes,7,rep,name=ip_whitelist_assignments,json=ipWhitelistAssignments" json:"ip_whitelist_assignments,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *AuthDB) Reset()                    { *m = AuthDB{} }
func (m *AuthDB) String() string            { return proto.CompactTextString(m) }
func (*AuthDB) ProtoMessage()               {}
func (*AuthDB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AuthDB) GetOauthClientId() string {
	if m != nil && m.OauthClientId != nil {
		return *m.OauthClientId
	}
	return ""
}

func (m *AuthDB) GetOauthClientSecret() string {
	if m != nil && m.OauthClientSecret != nil {
		return *m.OauthClientSecret
	}
	return ""
}

func (m *AuthDB) GetOauthAdditionalClientIds() []string {
	if m != nil {
		return m.OauthAdditionalClientIds
	}
	return nil
}

func (m *AuthDB) GetGroups() []*AuthGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *AuthDB) GetSecrets() []*AuthSecret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *AuthDB) GetIpWhitelists() []*AuthIPWhitelist {
	if m != nil {
		return m.IpWhitelists
	}
	return nil
}

func (m *AuthDB) GetIpWhitelistAssignments() []*AuthIPWhitelistAssignment {
	if m != nil {
		return m.IpWhitelistAssignments
	}
	return nil
}

// Information about some particular revision of auth DB.
type AuthDBRevision struct {
	// GAE App ID of a service holding primary copy of Auth DB.
	PrimaryId *string `protobuf:"bytes,1,req,name=primary_id,json=primaryId" json:"primary_id,omitempty"`
	// Revision of Auth DB being pushed.
	AuthDbRev *int64 `protobuf:"varint,2,req,name=auth_db_rev,json=authDbRev" json:"auth_db_rev,omitempty"`
	// Timestamp of that revision by Primary's clock, microseconds since epoch.
	ModifiedTs       *int64 `protobuf:"varint,3,req,name=modified_ts,json=modifiedTs" json:"modified_ts,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AuthDBRevision) Reset()                    { *m = AuthDBRevision{} }
func (m *AuthDBRevision) String() string            { return proto.CompactTextString(m) }
func (*AuthDBRevision) ProtoMessage()               {}
func (*AuthDBRevision) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AuthDBRevision) GetPrimaryId() string {
	if m != nil && m.PrimaryId != nil {
		return *m.PrimaryId
	}
	return ""
}

func (m *AuthDBRevision) GetAuthDbRev() int64 {
	if m != nil && m.AuthDbRev != nil {
		return *m.AuthDbRev
	}
	return 0
}

func (m *AuthDBRevision) GetModifiedTs() int64 {
	if m != nil && m.ModifiedTs != nil {
		return *m.ModifiedTs
	}
	return 0
}

// Published by Primary into 'auth-db-changed' PubSub topic. The body of the
// message is base64 encoded serialized ChangeNotification. Additional
// attributes are:
//  X-AuthDB-SigKey-v1: <id of a public key>
//  X-AuthDB-SigVal-v1: <base64 encoded RSA-SHA256(blob) signature>
type ChangeNotification struct {
	// New revision of the AuthDB.
	Revision         *AuthDBRevision `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ChangeNotification) Reset()                    { *m = ChangeNotification{} }
func (m *ChangeNotification) String() string            { return proto.CompactTextString(m) }
func (*ChangeNotification) ProtoMessage()               {}
func (*ChangeNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ChangeNotification) GetRevision() *AuthDBRevision {
	if m != nil {
		return m.Revision
	}
	return nil
}

// Sent from Primary to Replica to update Replica's AuthDB.
// Primary signs the entire serialized message with its private key and appends
// two headers to HTTP request that carries the blob:
//  X-AuthDB-SigKey-v1: <id of a public key>
//  X-AuthDB-SigVal-v1: <base64 encoded RSA-SHA256(SHA512(blob)) signature>
type ReplicationPushRequest struct {
	// Revision that is being pushed.
	Revision *AuthDBRevision `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	// An entire database of auth configuration for specific revision.
	AuthDb *AuthDB `protobuf:"bytes,2,opt,name=auth_db,json=authDb" json:"auth_db,omitempty"`
	// Version of 'auth' component on Primary, see components/auth/version.py.
	AuthCodeVersion  *string `protobuf:"bytes,3,opt,name=auth_code_version,json=authCodeVersion" json:"auth_code_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReplicationPushRequest) Reset()                    { *m = ReplicationPushRequest{} }
func (m *ReplicationPushRequest) String() string            { return proto.CompactTextString(m) }
func (*ReplicationPushRequest) ProtoMessage()               {}
func (*ReplicationPushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReplicationPushRequest) GetRevision() *AuthDBRevision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *ReplicationPushRequest) GetAuthDb() *AuthDB {
	if m != nil {
		return m.AuthDb
	}
	return nil
}

func (m *ReplicationPushRequest) GetAuthCodeVersion() string {
	if m != nil && m.AuthCodeVersion != nil {
		return *m.AuthCodeVersion
	}
	return ""
}

// Replica's response to ReplicationPushRequest.
type ReplicationPushResponse struct {
	// Overall status of the operation.
	Status *ReplicationPushResponse_Status `protobuf:"varint,1,req,name=status,enum=protocol.ReplicationPushResponse_Status" json:"status,omitempty"`
	// Revision known by Replica (set for APPLIED and SKIPPED statuses).
	CurrentRevision *AuthDBRevision `protobuf:"bytes,2,opt,name=current_revision,json=currentRevision" json:"current_revision,omitempty"`
	// Present for TRANSIENT_ERROR and FATAL_ERROR statuses.
	ErrorCode *ReplicationPushResponse_ErrorCode `protobuf:"varint,3,opt,name=error_code,json=errorCode,enum=protocol.ReplicationPushResponse_ErrorCode" json:"error_code,omitempty"`
	// Version of 'auth' component on Replica, see components/auth/version.py.
	AuthCodeVersion  *string `protobuf:"bytes,4,opt,name=auth_code_version,json=authCodeVersion" json:"auth_code_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReplicationPushResponse) Reset()                    { *m = ReplicationPushResponse{} }
func (m *ReplicationPushResponse) String() string            { return proto.CompactTextString(m) }
func (*ReplicationPushResponse) ProtoMessage()               {}
func (*ReplicationPushResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReplicationPushResponse) GetStatus() ReplicationPushResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ReplicationPushResponse_APPLIED
}

func (m *ReplicationPushResponse) GetCurrentRevision() *AuthDBRevision {
	if m != nil {
		return m.CurrentRevision
	}
	return nil
}

func (m *ReplicationPushResponse) GetErrorCode() ReplicationPushResponse_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ReplicationPushResponse_NOT_A_REPLICA
}

func (m *ReplicationPushResponse) GetAuthCodeVersion() string {
	if m != nil && m.AuthCodeVersion != nil {
		return *m.AuthCodeVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceLinkTicket)(nil), "protocol.ServiceLinkTicket")
	proto.RegisterType((*ServiceLinkRequest)(nil), "protocol.ServiceLinkRequest")
	proto.RegisterType((*ServiceLinkResponse)(nil), "protocol.ServiceLinkResponse")
	proto.RegisterType((*AuthGroup)(nil), "protocol.AuthGroup")
	proto.RegisterType((*AuthSecret)(nil), "protocol.AuthSecret")
	proto.RegisterType((*AuthIPWhitelist)(nil), "protocol.AuthIPWhitelist")
	proto.RegisterType((*AuthIPWhitelistAssignment)(nil), "protocol.AuthIPWhitelistAssignment")
	proto.RegisterType((*AuthDB)(nil), "protocol.AuthDB")
	proto.RegisterType((*AuthDBRevision)(nil), "protocol.AuthDBRevision")
	proto.RegisterType((*ChangeNotification)(nil), "protocol.ChangeNotification")
	proto.RegisterType((*ReplicationPushRequest)(nil), "protocol.ReplicationPushRequest")
	proto.RegisterType((*ReplicationPushResponse)(nil), "protocol.ReplicationPushResponse")
	proto.RegisterEnum("protocol.ServiceLinkResponse_Status", ServiceLinkResponse_Status_name, ServiceLinkResponse_Status_value)
	proto.RegisterEnum("protocol.ReplicationPushResponse_Status", ReplicationPushResponse_Status_name, ReplicationPushResponse_Status_value)
	proto.RegisterEnum("protocol.ReplicationPushResponse_ErrorCode", ReplicationPushResponse_ErrorCode_name, ReplicationPushResponse_ErrorCode_value)
}

var fileDescriptor0 = []byte{
	// 1029 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x5d, 0x72, 0xe3, 0x44,
	0x10, 0x46, 0xb6, 0x63, 0xc7, 0xed, 0x24, 0xb6, 0x27, 0x4b, 0xf0, 0x42, 0x01, 0x41, 0x50, 0xd4,
	0xc2, 0x56, 0xe5, 0x21, 0xc5, 0x23, 0x50, 0x28, 0xb6, 0x36, 0xab, 0xdd, 0xe0, 0x98, 0x91, 0x02,
	0x4f, 0x94, 0x4a, 0xb6, 0x86, 0x64, 0x6a, 0x6d, 0xc9, 0x68, 0xe4, 0x6c, 0x85, 0x43, 0x70, 0x01,
	0x1e, 0x28, 0x1e, 0xb8, 0x00, 0xe7, 0xe0, 0x12, 0xdc, 0x84, 0x9e, 0xd1, 0xe8, 0x27, 0x4e, 0xb2,
	0xae, 0x82, 0x27, 0xab, 0xbf, 0xe9, 0x9e, 0xe9, 0xef, 0xeb, 0x9e, 0x69, 0x43, 0x3f, 0x61, 0xcb,
	0x39, 0x9f, 0x05, 0x29, 0x8f, 0xa3, 0xa3, 0x65, 0x12, 0xa7, 0x31, 0xd9, 0x56, 0x3f, 0xb3, 0x78,
	0x6e, 0xfe, 0x6a, 0x40, 0xdf, 0x65, 0xc9, 0x35, 0x9f, 0xb1, 0x33, 0x1e, 0xbd, 0xf2, 0xf8, 0xec,
	0x15, 0x4b, 0xc9, 0xfb, 0x00, 0xcb, 0x84, 0x2f, 0x82, 0xe4, 0xc6, 0xe7, 0xe1, 0xc0, 0x38, 0xac,
	0x3d, 0x69, 0xd3, 0xb6, 0x46, 0x9c, 0x90, 0x7c, 0x08, 0x9d, 0x7c, 0x79, 0x95, 0xcc, 0x07, 0x35,
	0xb5, 0x9e, 0x47, 0x5c, 0x24, 0x73, 0xf2, 0x11, 0xec, 0x5c, 0xb2, 0x88, 0x25, 0x41, 0xca, 0x42,
	0x7f, 0x7a, 0x33, 0xa8, 0x2b, 0x8f, 0x4e, 0x81, 0x9d, 0xdc, 0x90, 0x03, 0x68, 0xa6, 0xea, 0xb0,
	0x41, 0x03, 0x17, 0x77, 0xa8, 0xb6, 0xcc, 0x25, 0x90, 0x4a, 0x3e, 0x94, 0xfd, 0xbc, 0x62, 0x22,
	0xad, 0x78, 0x1b, 0x55, 0x6f, 0x99, 0x89, 0x66, 0x57, 0xcd, 0x44, 0x43, 0x3a, 0x13, 0x1e, 0xf1,
	0x94, 0xaf, 0x65, 0x52, 0x60, 0x27, 0x37, 0xe6, 0xef, 0x06, 0xec, 0xdf, 0x3a, 0x52, 0x2c, 0xe3,
	0x48, 0x30, 0xf2, 0x25, 0x34, 0x45, 0x1a, 0xa4, 0x2b, 0xa1, 0xce, 0xdc, 0x3b, 0xfe, 0xe4, 0x28,
	0x57, 0xed, 0xe8, 0x1e, 0xf7, 0x23, 0x57, 0xf9, 0x52, 0x1d, 0x63, 0xbe, 0x80, 0x66, 0x86, 0x90,
	0x0e, 0xb4, 0xdc, 0x8b, 0xe1, 0xd0, 0x76, 0xdd, 0xde, 0x5b, 0x64, 0x1f, 0xba, 0x1e, 0xb5, 0xc6,
	0xee, 0xe4, 0x9c, 0x7a, 0xbe, 0x4d, 0xe9, 0x39, 0xed, 0x19, 0x64, 0x0f, 0xe0, 0xc4, 0x1a, 0xf9,
	0x9e, 0x33, 0x7c, 0x69, 0x7b, 0xbd, 0x9a, 0xb4, 0xad, 0x0b, 0xef, 0xb9, 0x5e, 0xaf, 0x9b, 0x7f,
	0xd4, 0xa0, 0x6d, 0xad, 0xd2, 0xab, 0xd3, 0x24, 0x5e, 0x2d, 0x09, 0x81, 0x46, 0x14, 0x2c, 0x98,
	0x2e, 0x8b, 0xfa, 0x26, 0x03, 0x68, 0x2d, 0xd8, 0x62, 0xca, 0x12, 0x81, 0x1a, 0xd4, 0x11, 0xce,
	0x4d, 0xf2, 0x08, 0xb6, 0x2e, 0xe7, 0xf1, 0x54, 0x20, 0x73, 0x89, 0x67, 0x86, 0xd4, 0x33, 0x42,
	0x5d, 0x59, 0x88, 0xea, 0x4b, 0x58, 0x5b, 0xe4, 0x10, 0x3a, 0x21, 0x13, 0xb3, 0x84, 0x2f, 0x65,
	0xb7, 0x0c, 0xb6, 0x32, 0xb5, 0x2a, 0x90, 0x6c, 0x8d, 0x59, 0xc2, 0x94, 0x9c, 0xa9, 0x18, 0x34,
	0xd1, 0xa1, 0x4e, 0xdb, 0x1a, 0xf1, 0x44, 0x75, 0x19, 0xd5, 0x6e, 0x65, 0x9d, 0xa3, 0x11, 0xac,
	0x3a, 0xd6, 0x6b, 0x11, 0x87, 0xfc, 0x27, 0x9e, 0x85, 0x6f, 0xab, 0x70, 0xc8, 0x21, 0x8c, 0xaf,
	0x3a, 0xe0, 0x06, 0xed, 0xac, 0xa0, 0x39, 0x94, 0xf5, 0x4d, 0xfc, 0x3a, 0x92, 0x44, 0xe1, 0xd0,
	0x90, 0x99, 0x67, 0x96, 0xf9, 0x0b, 0x6a, 0x86, 0x12, 0xb9, 0x0c, 0x0f, 0x4b, 0xef, 0xd5, 0x08,
	0x23, 0xaf, 0x83, 0x39, 0xb6, 0x93, 0x92, 0x08, 0x7b, 0x28, 0xb3, 0xd6, 0x73, 0xaa, 0x6f, 0xca,
	0xa9, 0xb1, 0x9e, 0x93, 0xf9, 0x8f, 0x01, 0x5d, 0x79, 0xb8, 0x33, 0xf9, 0xe1, 0x8a, 0xa7, 0x6c,
	0xce, 0x45, 0xfa, 0x50, 0x95, 0xc4, 0x6a, 0x1a, 0xb1, 0xb4, 0xa8, 0x92, 0x36, 0xd7, 0x75, 0xaf,
	0x6f, 0xd2, 0xbd, 0xf1, 0x66, 0xdd, 0xb7, 0x36, 0xe8, 0xde, 0xdc, 0xc4, 0xb1, 0x75, 0x87, 0xe3,
	0x5f, 0x06, 0x3c, 0x5e, 0xe3, 0x68, 0x09, 0xc1, 0x2f, 0xa3, 0x05, 0x8b, 0x52, 0xf2, 0x2e, 0x6c,
	0xf3, 0x10, 0x3f, 0x78, 0x7a, 0xa3, 0x19, 0x17, 0xb6, 0xba, 0x82, 0x4b, 0xff, 0x75, 0x1e, 0xa5,
	0x2f, 0x69, 0x87, 0x2f, 0x4b, 0xb1, 0x50, 0x98, 0x59, 0xbc, 0x90, 0x3b, 0x69, 0xea, 0xb9, 0xf9,
	0xff, 0x68, 0x9b, 0xbf, 0xd5, 0xa1, 0x29, 0x93, 0x1e, 0x9d, 0x90, 0x4f, 0xa1, 0x1b, 0x07, 0xf8,
	0xe9, 0xcf, 0xe6, 0x1c, 0x37, 0x2e, 0xdf, 0xb5, 0x5d, 0x05, 0x0f, 0x15, 0x8a, 0x6f, 0xdb, 0x11,
	0xec, 0xdf, 0xf2, 0x13, 0xaa, 0xa1, 0x74, 0xd2, 0xfd, 0x8a, 0xaf, 0xee, 0xb4, 0xaf, 0xe0, 0xbd,
	0xcc, 0x3f, 0x08, 0x43, 0x2e, 0x2b, 0x15, 0xcc, 0xcb, 0x23, 0xf2, 0x5b, 0x37, 0x50, 0x2e, 0x56,
	0xe1, 0x91, 0x9f, 0x26, 0xc8, 0x53, 0x68, 0x5e, 0xca, 0x5b, 0x2d, 0xd4, 0x45, 0xec, 0x1c, 0xef,
	0x97, 0x8f, 0x4c, 0x71, 0xe3, 0xa9, 0x76, 0xc1, 0xdc, 0x5a, 0x59, 0x3a, 0x02, 0xa9, 0x4a, 0xef,
	0x47, 0xb7, 0xbd, 0xb3, 0x94, 0x68, 0xee, 0x44, 0xbe, 0x86, 0xdd, 0xaa, 0xf2, 0xb2, 0xee, 0x32,
	0xea, 0xf1, 0xed, 0xa8, 0x4a, 0x45, 0xe9, 0x4e, 0xa5, 0x2a, 0x82, 0xfc, 0x08, 0x83, 0x6a, 0xbc,
	0x1f, 0x14, 0x05, 0x17, 0xd8, 0x21, 0x72, 0xab, 0x8f, 0x1f, 0xdc, 0xaa, 0x6c, 0x0e, 0x7a, 0x50,
	0xd9, 0xb4, 0x84, 0x05, 0x3e, 0xf5, 0x7b, 0x59, 0x71, 0x28, 0xbb, 0xe6, 0x42, 0x37, 0xf9, 0x9b,
	0xe6, 0xce, 0x07, 0xd0, 0x51, 0x52, 0x87, 0x53, 0x3f, 0x61, 0xd7, 0xaa, 0x26, 0xd8, 0x0d, 0x12,
	0x1a, 0x4d, 0x71, 0x8f, 0x8d, 0x37, 0x19, 0x1f, 0x65, 0x32, 0xbc, 0x0a, 0xa2, 0x4b, 0x36, 0x8e,
	0x53, 0xc4, 0xb2, 0x99, 0x48, 0xbe, 0x80, 0xed, 0x44, 0x67, 0x80, 0x67, 0x1a, 0x48, 0x6b, 0x70,
	0x9b, 0x56, 0x99, 0x21, 0x2d, 0x3c, 0xcd, 0x3f, 0x0d, 0x38, 0xa0, 0xe5, 0x64, 0x9d, 0xac, 0xc4,
	0x55, 0x3e, 0xad, 0xfe, 0xd3, 0x86, 0xe4, 0x33, 0x68, 0x69, 0x76, 0xc8, 0x4c, 0x06, 0xf5, 0xee,
	0x04, 0x35, 0x33, 0xae, 0xe4, 0x73, 0xe8, 0x67, 0x3d, 0x1a, 0x87, 0xcc, 0xbf, 0xc6, 0xe7, 0x2f,
	0x7b, 0x34, 0xe4, 0x7b, 0xd8, 0x55, 0x1d, 0x8a, 0xf8, 0xf7, 0x19, 0x6c, 0xfe, 0x5d, 0x87, 0x77,
	0xee, 0xe4, 0xa9, 0x47, 0xdc, 0x37, 0x6b, 0x23, 0xee, 0x49, 0x79, 0xe2, 0x03, 0x21, 0x6b, 0x63,
	0x8e, 0x0c, 0xa1, 0x37, 0x5b, 0x25, 0x89, 0x6c, 0xf7, 0x82, 0x72, 0x6d, 0x03, 0xe5, 0xae, 0x8e,
	0x28, 0xca, 0xfe, 0x02, 0x80, 0x25, 0x49, 0x9c, 0x28, 0x3e, 0x8a, 0xc7, 0xde, 0xf1, 0xd3, 0xcd,
	0xa9, 0xd8, 0x32, 0x46, 0x52, 0xa5, 0x6d, 0x96, 0x7f, 0xde, 0x2f, 0x4d, 0xe3, 0x7e, 0x69, 0x9e,
	0x57, 0x67, 0xb4, 0x35, 0x99, 0x9c, 0x39, 0xf6, 0x08, 0x67, 0xb4, 0x1c, 0xd8, 0x2f, 0x9d, 0xc9,
	0x04, 0x0d, 0xa3, 0x18, 0xd8, 0x8e, 0x3d, 0xce, 0x07, 0x76, 0x8d, 0x74, 0xa1, 0xf3, 0xcc, 0xf2,
	0xac, 0xb3, 0x62, 0x42, 0x5f, 0x41, 0xbb, 0xc8, 0x86, 0xf4, 0x61, 0x77, 0x7c, 0xee, 0xf9, 0x96,
	0x4f, 0x6d, 0xdc, 0x72, 0x68, 0xe1, 0x2e, 0xbb, 0xd0, 0x7e, 0x76, 0x4e, 0x4f, 0x9c, 0xd1, 0xc8,
	0x1e, 0x63, 0xfc, 0xdb, 0xd0, 0xff, 0xd6, 0x71, 0x5d, 0x67, 0x7c, 0xea, 0xbb, 0xce, 0xe9, 0xd8,
	0xf2, 0x2e, 0xa8, 0xdd, 0xab, 0xcb, 0x40, 0xf9, 0x3f, 0xa0, 0x84, 0x1a, 0xf2, 0x24, 0x09, 0x51,
	0xfb, 0xbb, 0x0b, 0xdb, 0xf5, 0x7a, 0x5b, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x13, 0x95, 0x33,
	0xfe, 0xce, 0x09, 0x00, 0x00,
}
