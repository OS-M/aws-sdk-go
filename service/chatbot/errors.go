// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chatbot

import (
	"github.com/OS-M/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// There was an issue processing your request.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeCreateChimeWebhookConfigurationException for service response error code
	// "CreateChimeWebhookConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeCreateChimeWebhookConfigurationException = "CreateChimeWebhookConfigurationException"

	// ErrCodeCreateSlackChannelConfigurationException for service response error code
	// "CreateSlackChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeCreateSlackChannelConfigurationException = "CreateSlackChannelConfigurationException"

	// ErrCodeCreateTeamsChannelConfigurationException for service response error code
	// "CreateTeamsChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeCreateTeamsChannelConfigurationException = "CreateTeamsChannelConfigurationException"

	// ErrCodeDeleteChimeWebhookConfigurationException for service response error code
	// "DeleteChimeWebhookConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDeleteChimeWebhookConfigurationException = "DeleteChimeWebhookConfigurationException"

	// ErrCodeDeleteMicrosoftTeamsUserIdentityException for service response error code
	// "DeleteMicrosoftTeamsUserIdentityException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDeleteMicrosoftTeamsUserIdentityException = "DeleteMicrosoftTeamsUserIdentityException"

	// ErrCodeDeleteSlackChannelConfigurationException for service response error code
	// "DeleteSlackChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDeleteSlackChannelConfigurationException = "DeleteSlackChannelConfigurationException"

	// ErrCodeDeleteSlackUserIdentityException for service response error code
	// "DeleteSlackUserIdentityException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDeleteSlackUserIdentityException = "DeleteSlackUserIdentityException"

	// ErrCodeDeleteSlackWorkspaceAuthorizationFault for service response error code
	// "DeleteSlackWorkspaceAuthorizationFault".
	//
	// There was an issue deleting your Slack workspace.
	ErrCodeDeleteSlackWorkspaceAuthorizationFault = "DeleteSlackWorkspaceAuthorizationFault"

	// ErrCodeDeleteTeamsChannelConfigurationException for service response error code
	// "DeleteTeamsChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDeleteTeamsChannelConfigurationException = "DeleteTeamsChannelConfigurationException"

	// ErrCodeDeleteTeamsConfiguredTeamException for service response error code
	// "DeleteTeamsConfiguredTeamException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDeleteTeamsConfiguredTeamException = "DeleteTeamsConfiguredTeamException"

	// ErrCodeDescribeChimeWebhookConfigurationsException for service response error code
	// "DescribeChimeWebhookConfigurationsException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDescribeChimeWebhookConfigurationsException = "DescribeChimeWebhookConfigurationsException"

	// ErrCodeDescribeSlackChannelConfigurationsException for service response error code
	// "DescribeSlackChannelConfigurationsException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDescribeSlackChannelConfigurationsException = "DescribeSlackChannelConfigurationsException"

	// ErrCodeDescribeSlackUserIdentitiesException for service response error code
	// "DescribeSlackUserIdentitiesException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDescribeSlackUserIdentitiesException = "DescribeSlackUserIdentitiesException"

	// ErrCodeDescribeSlackWorkspacesException for service response error code
	// "DescribeSlackWorkspacesException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeDescribeSlackWorkspacesException = "DescribeSlackWorkspacesException"

	// ErrCodeGetAccountPreferencesException for service response error code
	// "GetAccountPreferencesException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeGetAccountPreferencesException = "GetAccountPreferencesException"

	// ErrCodeGetTeamsChannelConfigurationException for service response error code
	// "GetTeamsChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeGetTeamsChannelConfigurationException = "GetTeamsChannelConfigurationException"

	// ErrCodeInternalServiceError for service response error code
	// "InternalServiceError".
	//
	// Customer/consumer-facing internal service exception. https://w.amazon.com/index.php/AWS/API_Standards/Exceptions#InternalServiceError
	ErrCodeInternalServiceError = "InternalServiceError"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// Your request input doesn't meet the constraints that AWS Chatbot requires.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// Your request input doesn't meet the constraints that AWS Chatbot requires.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// You have exceeded a service limit for AWS Chatbot.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeListMicrosoftTeamsConfiguredTeamsException for service response error code
	// "ListMicrosoftTeamsConfiguredTeamsException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeListMicrosoftTeamsConfiguredTeamsException = "ListMicrosoftTeamsConfiguredTeamsException"

	// ErrCodeListMicrosoftTeamsUserIdentitiesException for service response error code
	// "ListMicrosoftTeamsUserIdentitiesException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeListMicrosoftTeamsUserIdentitiesException = "ListMicrosoftTeamsUserIdentitiesException"

	// ErrCodeListTeamsChannelConfigurationsException for service response error code
	// "ListTeamsChannelConfigurationsException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeListTeamsChannelConfigurationsException = "ListTeamsChannelConfigurationsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// We were not able to find the resource for your request.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The supplied list of tags contains too many tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeUpdateAccountPreferencesException for service response error code
	// "UpdateAccountPreferencesException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeUpdateAccountPreferencesException = "UpdateAccountPreferencesException"

	// ErrCodeUpdateChimeWebhookConfigurationException for service response error code
	// "UpdateChimeWebhookConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeUpdateChimeWebhookConfigurationException = "UpdateChimeWebhookConfigurationException"

	// ErrCodeUpdateSlackChannelConfigurationException for service response error code
	// "UpdateSlackChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeUpdateSlackChannelConfigurationException = "UpdateSlackChannelConfigurationException"

	// ErrCodeUpdateTeamsChannelConfigurationException for service response error code
	// "UpdateTeamsChannelConfigurationException".
	//
	// We can’t process your request right now because of a server issue. Try
	// again later.
	ErrCodeUpdateTeamsChannelConfigurationException = "UpdateTeamsChannelConfigurationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":                           newErrorConflictException,
	"CreateChimeWebhookConfigurationException":    newErrorCreateChimeWebhookConfigurationException,
	"CreateSlackChannelConfigurationException":    newErrorCreateSlackChannelConfigurationException,
	"CreateTeamsChannelConfigurationException":    newErrorCreateTeamsChannelConfigurationException,
	"DeleteChimeWebhookConfigurationException":    newErrorDeleteChimeWebhookConfigurationException,
	"DeleteMicrosoftTeamsUserIdentityException":   newErrorDeleteMicrosoftTeamsUserIdentityException,
	"DeleteSlackChannelConfigurationException":    newErrorDeleteSlackChannelConfigurationException,
	"DeleteSlackUserIdentityException":            newErrorDeleteSlackUserIdentityException,
	"DeleteSlackWorkspaceAuthorizationFault":      newErrorDeleteSlackWorkspaceAuthorizationFault,
	"DeleteTeamsChannelConfigurationException":    newErrorDeleteTeamsChannelConfigurationException,
	"DeleteTeamsConfiguredTeamException":          newErrorDeleteTeamsConfiguredTeamException,
	"DescribeChimeWebhookConfigurationsException": newErrorDescribeChimeWebhookConfigurationsException,
	"DescribeSlackChannelConfigurationsException": newErrorDescribeSlackChannelConfigurationsException,
	"DescribeSlackUserIdentitiesException":        newErrorDescribeSlackUserIdentitiesException,
	"DescribeSlackWorkspacesException":            newErrorDescribeSlackWorkspacesException,
	"GetAccountPreferencesException":              newErrorGetAccountPreferencesException,
	"GetTeamsChannelConfigurationException":       newErrorGetTeamsChannelConfigurationException,
	"InternalServiceError":                        newErrorInternalServiceError,
	"InvalidParameterException":                   newErrorInvalidParameterException,
	"InvalidRequestException":                     newErrorInvalidRequestException,
	"LimitExceededException":                      newErrorLimitExceededException,
	"ListMicrosoftTeamsConfiguredTeamsException":  newErrorListMicrosoftTeamsConfiguredTeamsException,
	"ListMicrosoftTeamsUserIdentitiesException":   newErrorListMicrosoftTeamsUserIdentitiesException,
	"ListTeamsChannelConfigurationsException":     newErrorListTeamsChannelConfigurationsException,
	"ResourceNotFoundException":                   newErrorResourceNotFoundException,
	"ServiceUnavailableException":                 newErrorServiceUnavailableException,
	"TooManyTagsException":                        newErrorTooManyTagsException,
	"UpdateAccountPreferencesException":           newErrorUpdateAccountPreferencesException,
	"UpdateChimeWebhookConfigurationException":    newErrorUpdateChimeWebhookConfigurationException,
	"UpdateSlackChannelConfigurationException":    newErrorUpdateSlackChannelConfigurationException,
	"UpdateTeamsChannelConfigurationException":    newErrorUpdateTeamsChannelConfigurationException,
}
