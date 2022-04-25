// Copyright(c) 2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: profile/local_profile.proto

package profile

import (
	info "github.com/lf-edge/eve/api/go/info"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppCommand_Command int32

const (
	AppCommand_COMMAND_UNSPECIFIED AppCommand_Command = 0
	// Application instance, which is either running or transitioning to a running state,
	// will be stopped and subsequently started again, preserving the mutated run time state.
	AppCommand_COMMAND_RESTART AppCommand_Command = 1
	// Application instance, which is either running or transitioning to a running state,
	// will be stopped and the mutated run time state of the app is deleted.
	// A subsequent action to start the app will start it with a pristine runtime state.
	// This command will purge ALL volumes used by the application.
	AppCommand_COMMAND_PURGE AppCommand_Command = 2
)

// Enum value maps for AppCommand_Command.
var (
	AppCommand_Command_name = map[int32]string{
		0: "COMMAND_UNSPECIFIED",
		1: "COMMAND_RESTART",
		2: "COMMAND_PURGE",
	}
	AppCommand_Command_value = map[string]int32{
		"COMMAND_UNSPECIFIED": 0,
		"COMMAND_RESTART":     1,
		"COMMAND_PURGE":       2,
	}
)

func (x AppCommand_Command) Enum() *AppCommand_Command {
	p := new(AppCommand_Command)
	*p = x
	return p
}

func (x AppCommand_Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppCommand_Command) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_local_profile_proto_enumTypes[0].Descriptor()
}

func (AppCommand_Command) Type() protoreflect.EnumType {
	return &file_profile_local_profile_proto_enumTypes[0]
}

func (x AppCommand_Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppCommand_Command.Descriptor instead.
func (AppCommand_Command) EnumDescriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{7, 0}
}

// LocalProfile message is sent in response to a GET to
// the api/v1/local_profile API
type LocalProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalProfile string `protobuf:"bytes,1,opt,name=local_profile,json=localProfile,proto3" json:"local_profile,omitempty"`
	ServerToken  string `protobuf:"bytes,2,opt,name=server_token,json=serverToken,proto3" json:"server_token,omitempty"`
}

func (x *LocalProfile) Reset() {
	*x = LocalProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalProfile) ProtoMessage() {}

func (x *LocalProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalProfile.ProtoReflect.Descriptor instead.
func (*LocalProfile) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{0}
}

func (x *LocalProfile) GetLocalProfile() string {
	if x != nil {
		return x.LocalProfile
	}
	return ""
}

func (x *LocalProfile) GetServerToken() string {
	if x != nil {
		return x.ServerToken
	}
	return ""
}

// RadioStatus message is sent in the POST request to the api/v1/radio API.
type RadioStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// radio_silence is True if the Radio-Silence mode is enabled by config AND successfully
	// applied for all radio devices.
	// If the radio silence was requested but this field is still returned as False,
	// expect non-empty config_error attached.
	RadioSilence bool `protobuf:"varint,1,opt,name=radio_silence,json=radioSilence,proto3" json:"radio_silence,omitempty"`
	// If the last radio configuration change failed, error message is reported here.
	// Please note that there is also a per-modem configuration error reported under CellularStatus.
	ConfigError string `protobuf:"bytes,2,opt,name=config_error,json=configError,proto3" json:"config_error,omitempty"`
	// Status of every LTE network.
	CellularStatus []*CellularStatus `protobuf:"bytes,3,rep,name=cellular_status,json=cellularStatus,proto3" json:"cellular_status,omitempty"` // XXX Later we can add status for every WiFi network adapter.
}

func (x *RadioStatus) Reset() {
	*x = RadioStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioStatus) ProtoMessage() {}

func (x *RadioStatus) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioStatus.ProtoReflect.Descriptor instead.
func (*RadioStatus) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{1}
}

func (x *RadioStatus) GetRadioSilence() bool {
	if x != nil {
		return x.RadioSilence
	}
	return false
}

func (x *RadioStatus) GetConfigError() string {
	if x != nil {
		return x.ConfigError
	}
	return ""
}

func (x *RadioStatus) GetCellularStatus() []*CellularStatus {
	if x != nil {
		return x.CellularStatus
	}
	return nil
}

// CellularStatus contains status information for a single LTE network.
type CellularStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Logical label assigned to the physical cellular modem.
	Logicallabel string                    `protobuf:"bytes,1,opt,name=logicallabel,proto3" json:"logicallabel,omitempty"`
	Module       *info.ZCellularModuleInfo `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	SimCards     []*info.ZSimcardInfo      `protobuf:"bytes,3,rep,name=sim_cards,json=simCards,proto3" json:"sim_cards,omitempty"`
	Providers    []*info.ZCellularProvider `protobuf:"bytes,4,rep,name=providers,proto3" json:"providers,omitempty"`
	ConfigError  string                    `protobuf:"bytes,10,opt,name=config_error,json=configError,proto3" json:"config_error,omitempty"`
	ProbeError   string                    `protobuf:"bytes,11,opt,name=probe_error,json=probeError,proto3" json:"probe_error,omitempty"`
}

func (x *CellularStatus) Reset() {
	*x = CellularStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CellularStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CellularStatus) ProtoMessage() {}

func (x *CellularStatus) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CellularStatus.ProtoReflect.Descriptor instead.
func (*CellularStatus) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{2}
}

func (x *CellularStatus) GetLogicallabel() string {
	if x != nil {
		return x.Logicallabel
	}
	return ""
}

func (x *CellularStatus) GetModule() *info.ZCellularModuleInfo {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *CellularStatus) GetSimCards() []*info.ZSimcardInfo {
	if x != nil {
		return x.SimCards
	}
	return nil
}

func (x *CellularStatus) GetProviders() []*info.ZCellularProvider {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *CellularStatus) GetConfigError() string {
	if x != nil {
		return x.ConfigError
	}
	return ""
}

func (x *CellularStatus) GetProbeError() string {
	if x != nil {
		return x.ProbeError
	}
	return ""
}

// RadioConfig message may be returned in the response from a POST request
// sent to the api/v1/radio API.
type RadioConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Security token. EVE will verify that serverToken matches the profile server
	// token received from the controller.
	ServerToken string `protobuf:"bytes,1,opt,name=server_token,json=serverToken,proto3" json:"server_token,omitempty"`
	// If enabled, EVE will disable radio transmission on all wireless devices available
	// to the host (i.e. it does not cover wireless devices directly attached to applications).
	RadioSilence bool `protobuf:"varint,2,opt,name=radio_silence,json=radioSilence,proto3" json:"radio_silence,omitempty"`
}

func (x *RadioConfig) Reset() {
	*x = RadioConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioConfig) ProtoMessage() {}

func (x *RadioConfig) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioConfig.ProtoReflect.Descriptor instead.
func (*RadioConfig) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{3}
}

func (x *RadioConfig) GetServerToken() string {
	if x != nil {
		return x.ServerToken
	}
	return ""
}

func (x *RadioConfig) GetRadioSilence() bool {
	if x != nil {
		return x.RadioSilence
	}
	return false
}

// LocalAppInfoList contains information about all app on EdgeNode
// sent to the api/v1/appinfo
type LocalAppInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppsInfo []*LocalAppInfo `protobuf:"bytes,1,rep,name=apps_info,json=appsInfo,proto3" json:"apps_info,omitempty"`
}

func (x *LocalAppInfoList) Reset() {
	*x = LocalAppInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalAppInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalAppInfoList) ProtoMessage() {}

func (x *LocalAppInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalAppInfoList.ProtoReflect.Descriptor instead.
func (*LocalAppInfoList) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{4}
}

func (x *LocalAppInfoList) GetAppsInfo() []*LocalAppInfo {
	if x != nil {
		return x.AppsInfo
	}
	return nil
}

// LocalAppInfo contains information about app on EdgeNode
type LocalAppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version string          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Name    string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Err     *info.ErrorInfo `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	State   info.ZSwState   `protobuf:"varint,5,opt,name=state,proto3,enum=org.lfedge.eve.info.ZSwState" json:"state,omitempty"`
	// Value of the field `timestamp` from the last `AppCommand` that was
	// requested by the Local profile server, received by EVE and has completed
	// its execution for this application instance.
	LastCmdTimestamp uint64 `protobuf:"varint,6,opt,name=last_cmd_timestamp,json=lastCmdTimestamp,proto3" json:"last_cmd_timestamp,omitempty"`
}

func (x *LocalAppInfo) Reset() {
	*x = LocalAppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalAppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalAppInfo) ProtoMessage() {}

func (x *LocalAppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalAppInfo.ProtoReflect.Descriptor instead.
func (*LocalAppInfo) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{5}
}

func (x *LocalAppInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LocalAppInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LocalAppInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LocalAppInfo) GetErr() *info.ErrorInfo {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *LocalAppInfo) GetState() info.ZSwState {
	if x != nil {
		return x.State
	}
	return info.ZSwState_INVALID
}

func (x *LocalAppInfo) GetLastCmdTimestamp() uint64 {
	if x != nil {
		return x.LastCmdTimestamp
	}
	return 0
}

// LocalAppCmds message may be returned in the response from a POST request
// sent to the api/v1/appinfo API.
type LocalAppCmdList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Security token. EVE will verify that server_token matches the profile server
	// token received from the controller.
	ServerToken string `protobuf:"bytes,1,opt,name=server_token,json=serverToken,proto3" json:"server_token,omitempty"`
	// A list of commands requested to be executed for locally running application instances.
	// A new request created for the same application should overwrite the previous entry
	// with the 'timestamp' field updated. In other words, the list should contain at most
	// one entry for each application instance.
	// It is not required for the Local profile server to persist command requests.
	// Also, it is not required for the Local profile server to stop submitting command
	// requests that have been already processed by EVE. Using the `timestamp` field,
	// EVE is able to determine if a given command request has been already handled or not.
	// To check if the last requested command has completed, compare its timestamp with
	// 'last_cmd_timestamp' from `LocalAppInfo` message, submitted by EVE in the request
	// body of the api/v1/appinfo API.
	AppCommands []*AppCommand `protobuf:"bytes,2,rep,name=app_commands,json=appCommands,proto3" json:"app_commands,omitempty"`
}

func (x *LocalAppCmdList) Reset() {
	*x = LocalAppCmdList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalAppCmdList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalAppCmdList) ProtoMessage() {}

func (x *LocalAppCmdList) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalAppCmdList.ProtoReflect.Descriptor instead.
func (*LocalAppCmdList) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{6}
}

func (x *LocalAppCmdList) GetServerToken() string {
	if x != nil {
		return x.ServerToken
	}
	return ""
}

func (x *LocalAppCmdList) GetAppCommands() []*AppCommand {
	if x != nil {
		return x.AppCommands
	}
	return nil
}

// AppCommand references a running application instance by UUID and/or displayname,
// and describes a command to execute for this instance.
type AppCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference the application instance by its ID (which is an instance of UUID).
	// At least one of the id and displayname should be defined.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Reference the application instance by the user-friendly displayname.
	// At least one of the id and displayname should be defined.
	Displayname string `protobuf:"bytes,2,opt,name=displayname,proto3" json:"displayname,omitempty"`
	// Timestamp to record when the request to run the command was made.
	// The format of the timestamp is not defined. It can be a Unix timestamp
	// or a different time representation. It is not even required for the timestamp
	// to match the real time or to be in-sync with the device clock.
	// What is required, however, is that two successive but distinct requests made
	// for the same application will have different timestamps attached.
	// This requirement applies even between restarts of the Local profile server.
	// A request made after a restart should not have the same timestamp attached
	// as the previous request made for the same application before the restart.
	//
	// EVE guarantees that a newly added command request or a change of the timestamp
	// will result in the command being triggered ASAP. Even if the execution of a command
	// is interrupted by a device reboot/crash, the eventuality of the command completion
	// is still guaranteed. The only exception is if Local Profile Server restarts/crashes
	// shortly after a request is made, in which case it can get lost before EVE is able
	// to receive it. For this scenario to be avoided, a persistence of command requests
	// on the side of the Local Profile server is necessary.
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Command to run.
	Command AppCommand_Command `protobuf:"varint,4,opt,name=command,proto3,enum=org.lfedge.eve.profile.AppCommand_Command" json:"command,omitempty"`
}

func (x *AppCommand) Reset() {
	*x = AppCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_local_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppCommand) ProtoMessage() {}

func (x *AppCommand) ProtoReflect() protoreflect.Message {
	mi := &file_profile_local_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppCommand.ProtoReflect.Descriptor instead.
func (*AppCommand) Descriptor() ([]byte, []int) {
	return file_profile_local_profile_proto_rawDescGZIP(), []int{7}
}

func (x *AppCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppCommand) GetDisplayname() string {
	if x != nil {
		return x.Displayname
	}
	return ""
}

func (x *AppCommand) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AppCommand) GetCommand() AppCommand_Command {
	if x != nil {
		return x.Command
	}
	return AppCommand_COMMAND_UNSPECIFIED
}

var File_profile_local_profile_proto protoreflect.FileDescriptor

var file_profile_local_profile_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f,
	0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa6,
	0x01, 0x0a, 0x0b, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x69, 0x6c, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4f, 0x0a, 0x0f, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc0, 0x02, 0x0a, 0x0e, 0x43, 0x65, 0x6c, 0x6c,
	0x75, 0x6c, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x40,
	0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x5a, 0x43, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x3e, 0x0a, 0x09, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x5a, 0x53, 0x69, 0x6d, 0x63, 0x61,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x73, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x44, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x5a, 0x43, 0x65, 0x6c, 0x6c, 0x75,
	0x6c, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x55, 0x0a, 0x0b, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0x55, 0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x61, 0x70, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x5a,
	0x53, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74,
	0x43, 0x6d, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7b, 0x0a, 0x0f,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x45, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c,
	0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0b, 0x61, 0x70,
	0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x0a, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x4a,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x5f, 0x50, 0x55, 0x52, 0x47, 0x45, 0x10, 0x02, 0x42, 0x3f, 0x0a, 0x16, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_profile_local_profile_proto_rawDescOnce sync.Once
	file_profile_local_profile_proto_rawDescData = file_profile_local_profile_proto_rawDesc
)

func file_profile_local_profile_proto_rawDescGZIP() []byte {
	file_profile_local_profile_proto_rawDescOnce.Do(func() {
		file_profile_local_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_local_profile_proto_rawDescData)
	})
	return file_profile_local_profile_proto_rawDescData
}

var file_profile_local_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profile_local_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_profile_local_profile_proto_goTypes = []interface{}{
	(AppCommand_Command)(0),          // 0: org.lfedge.eve.profile.AppCommand.Command
	(*LocalProfile)(nil),             // 1: org.lfedge.eve.profile.LocalProfile
	(*RadioStatus)(nil),              // 2: org.lfedge.eve.profile.RadioStatus
	(*CellularStatus)(nil),           // 3: org.lfedge.eve.profile.CellularStatus
	(*RadioConfig)(nil),              // 4: org.lfedge.eve.profile.RadioConfig
	(*LocalAppInfoList)(nil),         // 5: org.lfedge.eve.profile.LocalAppInfoList
	(*LocalAppInfo)(nil),             // 6: org.lfedge.eve.profile.LocalAppInfo
	(*LocalAppCmdList)(nil),          // 7: org.lfedge.eve.profile.LocalAppCmdList
	(*AppCommand)(nil),               // 8: org.lfedge.eve.profile.AppCommand
	(*info.ZCellularModuleInfo)(nil), // 9: org.lfedge.eve.info.ZCellularModuleInfo
	(*info.ZSimcardInfo)(nil),        // 10: org.lfedge.eve.info.ZSimcardInfo
	(*info.ZCellularProvider)(nil),   // 11: org.lfedge.eve.info.ZCellularProvider
	(*info.ErrorInfo)(nil),           // 12: org.lfedge.eve.info.ErrorInfo
	(info.ZSwState)(0),               // 13: org.lfedge.eve.info.ZSwState
}
var file_profile_local_profile_proto_depIdxs = []int32{
	3,  // 0: org.lfedge.eve.profile.RadioStatus.cellular_status:type_name -> org.lfedge.eve.profile.CellularStatus
	9,  // 1: org.lfedge.eve.profile.CellularStatus.module:type_name -> org.lfedge.eve.info.ZCellularModuleInfo
	10, // 2: org.lfedge.eve.profile.CellularStatus.sim_cards:type_name -> org.lfedge.eve.info.ZSimcardInfo
	11, // 3: org.lfedge.eve.profile.CellularStatus.providers:type_name -> org.lfedge.eve.info.ZCellularProvider
	6,  // 4: org.lfedge.eve.profile.LocalAppInfoList.apps_info:type_name -> org.lfedge.eve.profile.LocalAppInfo
	12, // 5: org.lfedge.eve.profile.LocalAppInfo.err:type_name -> org.lfedge.eve.info.ErrorInfo
	13, // 6: org.lfedge.eve.profile.LocalAppInfo.state:type_name -> org.lfedge.eve.info.ZSwState
	8,  // 7: org.lfedge.eve.profile.LocalAppCmdList.app_commands:type_name -> org.lfedge.eve.profile.AppCommand
	0,  // 8: org.lfedge.eve.profile.AppCommand.command:type_name -> org.lfedge.eve.profile.AppCommand.Command
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_profile_local_profile_proto_init() }
func file_profile_local_profile_proto_init() {
	if File_profile_local_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_local_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalProfile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CellularStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalAppInfoList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalAppInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalAppCmdList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_profile_local_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppCommand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_local_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_profile_local_profile_proto_goTypes,
		DependencyIndexes: file_profile_local_profile_proto_depIdxs,
		EnumInfos:         file_profile_local_profile_proto_enumTypes,
		MessageInfos:      file_profile_local_profile_proto_msgTypes,
	}.Build()
	File_profile_local_profile_proto = out.File
	file_profile_local_profile_proto_rawDesc = nil
	file_profile_local_profile_proto_goTypes = nil
	file_profile_local_profile_proto_depIdxs = nil
}
