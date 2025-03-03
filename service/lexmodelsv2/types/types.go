// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The location of audio log files collected when conversation logging is enabled
// for a bot.
type AudioLogDestination struct {

	// The Amazon S3 bucket where the audio log files are stored. The IAM role
	// specified in the roleArn parameter of the CreateBot operation must have
	// permission to write to this bucket.
	//
	// This member is required.
	S3Bucket *S3BucketLogDestination
}

// Settings for logging audio of conversations between Amazon Lex and a user. You
// specify whether to log audio and the Amazon S3 bucket where the audio file is
// stored.
type AudioLogSetting struct {

	// The location of audio log files collected when conversation logging is enabled
	// for a bot.
	//
	// This member is required.
	Destination *AudioLogDestination

	// Determines whether audio logging in enabled for the bot.
	//
	// This member is required.
	Enabled bool
}

// Provides a record of an event that affects a bot alias. For example, when the
// version of a bot that the alias points to changes.
type BotAliasHistoryEvent struct {

	// The version of the bot that was used in the event.
	BotVersion *string

	// The date and time that the event ended.
	EndDate *time.Time

	// The date and time that the event started.
	StartDate *time.Time
}

// Specifies settings that are unique to a locale. For example, you can use
// different Lambda function depending on the bot's locale.
type BotAliasLocaleSettings struct {

	// Determines whether the locale is enabled for the bot. If the value is false, the
	// locale isn't available for use.
	//
	// This member is required.
	Enabled bool

	// Specifies the Lambda function that should be used in the locale.
	CodeHookSpecification *CodeHookSpecification
}

// Summary information about bot aliases returned from the ListBotAliases
// operation.
type BotAliasSummary struct {

	// The unique identifier assigned to the bot alias. You can use this ID to get
	// detailed information about the alias using the DescribeBotAlias operation.
	BotAliasId *string

	// The name of the bot alias.
	BotAliasName *string

	// The current state of the bot alias. If the status is Available, the alias is
	// ready for use.
	BotAliasStatus BotAliasStatus

	// The version of the bot that the bot alias references.
	BotVersion *string

	// A timestamp of the date and time that the bot alias was created.
	CreationDateTime *time.Time

	// The description of the bot alias.
	Description *string

	// A timestamp of the date and time that the bot alias was last updated.
	LastUpdatedDateTime *time.Time
}

// Filters the responses returned by the ListBots operation.
type BotFilter struct {

	// The name of the field to filter the list of bots.
	//
	// This member is required.
	Name BotFilterName

	// The operator to use for the filter. Specify EQ when the ListBots operation
	// should return only aliases that equal the specified value. Specify CO when the
	// ListBots operation should return aliases that contain the specified value.
	//
	// This member is required.
	Operator BotFilterOperator

	// The value to use for filtering the list of bots.
	//
	// This member is required.
	Values []string
}

// Filters responses returned by the ListBotLocales operation.
type BotLocaleFilter struct {

	// The name of the field to filter the list of bots.
	//
	// This member is required.
	Name BotLocaleFilterName

	// The operator to use for the filter. Specify EQ when the ListBotLocales operation
	// should return only aliases that equal the specified value. Specify CO when the
	// ListBotLocales operation should return aliases that contain the specified value.
	//
	// This member is required.
	Operator BotLocaleFilterOperator

	// The value to use for filtering the list of bots.
	//
	// This member is required.
	Values []string
}

// Provides information about an event that occurred affecting the bot locale.
type BotLocaleHistoryEvent struct {

	// A description of the event that occurred.
	//
	// This member is required.
	Event *string

	// A timestamp of the date and time that the event occurred.
	//
	// This member is required.
	EventDate *time.Time
}

// Specifies attributes for sorting a list of bot locales.
type BotLocaleSortBy struct {

	// The bot locale attribute to sort by.
	//
	// This member is required.
	Attribute BotLocaleSortAttribute

	// Specifies whether to sort the bot locales in ascending or descending order.
	//
	// This member is required.
	Order SortOrder
}

// Summary information about bot locales returned by the ListBotLocales operation.
type BotLocaleSummary struct {

	// The current status of the bot locale. When the status is Built the locale is
	// ready for use.
	BotLocaleStatus BotLocaleStatus

	// The description of the bot locale.
	Description *string

	// A timestamp of the date and time that the bot locale was last built.
	LastBuildSubmittedDateTime *time.Time

	// A timestamp of the date and time that the bot locale was last updated.
	LastUpdatedDateTime *time.Time

	// The language and locale of the bot locale.
	LocaleId *string

	// The name of the bot locale.
	LocaleName *string
}

// Specifies attributes for sorting a list of bots.
type BotSortBy struct {

	// The attribute to use to sort the list of bots.
	//
	// This member is required.
	Attribute BotSortAttribute

	// The order to sort the list. You can choose ascending or descending.
	//
	// This member is required.
	Order SortOrder
}

// Summary information about a bot returned by the ListBots operation.
type BotSummary struct {

	// The unique identifier assigned to the bot. Use this ID to get detailed
	// information about the bot with the DescribeBot operation.
	BotId *string

	// The name of the bot.
	BotName *string

	// The current status of the bot. When the status is Available the bot is ready for
	// use.
	BotStatus BotStatus

	// The description of the bot.
	Description *string

	// The date and time that the bot was last updated.
	LastUpdatedDateTime *time.Time

	// The latest numerical version in use for the bot.
	LatestBotVersion *string
}

// The version of a bot used for a bot locale.
type BotVersionLocaleDetails struct {

	// The version of a bot used for a bot locale.
	//
	// This member is required.
	SourceBotVersion *string
}

// Specifies attributes for sorting a list of bot versions.
type BotVersionSortBy struct {

	// The attribute to use to sort the list of versions.
	//
	// This member is required.
	Attribute BotVersionSortAttribute

	// The order to sort the list. You can specify ascending or descending order.
	//
	// This member is required.
	Order SortOrder
}

// Summary information about a bot version returned by the ListBotVersions
// operation.
type BotVersionSummary struct {

	// The name of the bot associated with the version.
	BotName *string

	// The status of the bot. When the status is available, the version of the bot is
	// ready for use.
	BotStatus BotStatus

	// The numeric version of the bot, or DRAFT to indicate that this is the version of
	// the bot that can be updated..
	BotVersion *string

	// A timestamp of the date and time that the version was created.
	CreationDateTime *time.Time

	// The description of the version.
	Description *string
}

// Specifies attributes for sorting a list of built-in intents.
type BuiltInIntentSortBy struct {

	// The attribute to use to sort the list of built-in intents.
	//
	// This member is required.
	Attribute BuiltInIntentSortAttribute

	// The order to sort the list. You can specify ascending or descending order.
	//
	// This member is required.
	Order SortOrder
}

// Provides summary information about a built-in intent for the ListBuiltInIntents
// operation.
type BuiltInIntentSummary struct {

	// The description of the intent.
	Description *string

	// The signature of the built-in intent. Use this to specify the parent intent of a
	// derived intent.
	IntentSignature *string
}

// Specifies attributes for sorting a list of built-in slot types.
type BuiltInSlotTypeSortBy struct {

	// The attribute to use to sort the list of built-in intents.
	//
	// This member is required.
	Attribute BuiltInSlotTypeSortAttribute

	// The order to sort the list. You can choose ascending or descending.
	//
	// This member is required.
	Order SortOrder
}

// Provides summary information about a built-in slot type for the
// ListBuiltInSlotTypes operation.
type BuiltInSlotTypeSummary struct {

	// The description of the built-in slot type.
	Description *string

	// The signature of the built-in slot type. Use this to specify the parent slot
	// type of a derived slot type.
	SlotTypeSignature *string
}

// Describes a button to use on a response card used to gather slot values from a
// user.
type Button struct {

	// The text that appears on the button. Use this to tell the user what value is
	// returned when they choose this button.
	//
	// This member is required.
	Text *string

	// The value returned to Amazon Lex when the user chooses this button. This must be
	// one of the slot values configured for the slot.
	//
	// This member is required.
	Value *string
}

// The Amazon CloudWatch Logs log group where the text and metadata logs are
// delivered. The log group must exist before you enable logging.
type CloudWatchLogGroupLogDestination struct {

	// The Amazon Resource Name (ARN) of the log group where text and metadata logs are
	// delivered.
	//
	// This member is required.
	CloudWatchLogGroupArn *string

	// The prefix of the log stream name within the log group that you specified
	//
	// This member is required.
	LogPrefix *string
}

// Contains information about code hooks that Amazon Lex calls during a
// conversation.
type CodeHookSpecification struct {

	// Specifies a Lambda function that verifies requests to a bot or fulfilles the
	// user's request to a bot.
	//
	// This member is required.
	LambdaCodeHook *LambdaCodeHook
}

// Configures conversation logging that saves audio, text, and metadata for the
// conversations with your users.
type ConversationLogSettings struct {

	// The Amazon S3 settings for logging audio to an S3 bucket.
	AudioLogSettings []AudioLogSetting

	// The Amazon CloudWatch Logs settings for logging text and metadata.
	TextLogSettings []TextLogSetting
}

// A custom response string that Amazon Lex sends to your application. You define
// the content and structure the string.
type CustomPayload struct {

	// The string that is sent to your application.
	//
	// This member is required.
	Value *string
}

// By default, data stored by Amazon Lex is encrypted. The DataPrivacy structure
// provides settings that determine how Amazon Lex handles special cases of
// securing the data for your bot.
type DataPrivacy struct {

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to a
	// website, program, or other application that is directed or targeted, in whole or
	// in part, to children under age 13 and subject to COPPA. By specifying false in
	// the childDirected field, you confirm that your use of Amazon Lex is not related
	// to a website, program, or other application that is directed or targeted, in
	// whole or in part, to children under age 13 and subject to COPPA. You may not
	// specify a default value for the childDirected field that does not accurately
	// reflect whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. If your use of Amazon Lex relates to a
	// website, program, or other application that is directed in whole or in part, to
	// children under age 13, you must obtain any required verifiable parental consent
	// under COPPA. For information regarding the use of Amazon Lex in connection with
	// websites, programs, or other applications that are directed or targeted, in
	// whole or in part, to children under age 13, see the Amazon Lex FAQ
	// (https://aws.amazon.com/lex/faqs#data-security).
	//
	// This member is required.
	ChildDirected bool
}

// Settings that determine the Lambda function that Amazon Lex uses for processing
// user responses.
type DialogCodeHookSettings struct {

	// Enables the dialog code hook so that it processes user requests.
	//
	// This member is required.
	Enabled bool
}

// Determines if a Lambda function should be invoked for a specific intent.
type FulfillmentCodeHookSettings struct {

	// Indicates whether a Lambda function should be invoked to fulfill a specific
	// intent.
	//
	// This member is required.
	Enabled bool
}

// A card that is shown to the user by a messaging platform. You define the
// contents of the card, the card is displayed by the platform. When you use a
// response card, the response from the user is constrained to the text associated
// with a button on the card.
type ImageResponseCard struct {

	// The title to display on the response card. The format of the title is determined
	// by the platform displaying the response card.
	//
	// This member is required.
	Title *string

	// A list of buttons that should be displayed on the response card. The arrangement
	// of the buttons is determined by the platform that displays the button.
	Buttons []Button

	// The URL of an image to display on the response card. The image URL must be
	// publicly available so that the platform displaying the response card has access
	// to the image.
	ImageUrl *string

	// The subtitle to display on the response card. The format of the subtitle is
	// determined by the platform displaying the response card.
	Subtitle *string
}

// The name of a context that must be active for an intent to be selected by Amazon
// Lex.
type InputContext struct {

	// The name of the context.
	//
	// This member is required.
	Name *string
}

// Provides a statement the Amazon Lex conveys to the user when the intent is
// successfully fulfilled.
type IntentClosingSetting struct {

	// The response that Amazon Lex sends to the user when the intent is complete.
	//
	// This member is required.
	ClosingResponse *ResponseSpecification
}

// Provides a prompt for making sure that the user is ready for the intent to be
// fulfilled.
type IntentConfirmationSetting struct {

	// When the user answers "no" to the question defined in promptSpecification,
	// Amazon Lex responds with this response to acknowledge that the intent was
	// canceled.
	//
	// This member is required.
	DeclinationResponse *ResponseSpecification

	// Prompts the user to confirm the intent. This question should have a yes or no
	// answer. Amazon Lex uses this prompt to ensure that the user acknowledges that
	// the intent is ready for fulfillment. For example, with the OrderPizza intent,
	// you might want to confirm that the order is correct before placing it. For other
	// intents, such as intents that simply respond to user questions, you might not
	// need to ask the user for confirmation before providing the information.
	//
	// This member is required.
	PromptSpecification *PromptSpecification
}

// Filters the response from the ListIntents operation.
type IntentFilter struct {

	// The name of the field to use for the filter.
	//
	// This member is required.
	Name IntentFilterName

	// The operator to use for the filter. Specify EQ when the ListIntents operation
	// should return only aliases that equal the specified value. Specify CO when the
	// ListIntents operation should return aliases that contain the specified value.
	//
	// This member is required.
	Operator IntentFilterOperator

	// The value to use for the filter.
	//
	// This member is required.
	Values []string
}

// Specifies attributes for sorting a list of intents.
type IntentSortBy struct {

	// The attribute to use to sort the list of intents.
	//
	// This member is required.
	Attribute IntentSortAttribute

	// The order to sort the list. You can choose ascending or descending.
	//
	// This member is required.
	Order SortOrder
}

// Summary information about an intent returned by the ListIntents operation.
type IntentSummary struct {

	// The description of the intent.
	Description *string

	// The input contexts that must be active for this intent to be considered for
	// recognition.
	InputContexts []InputContext

	// The unique identifier assigned to the intent. Use this ID to get detailed
	// information about the intent with the DescribeIntent operation.
	IntentId *string

	// The name of the intent.
	IntentName *string

	// The timestamp of the date and time that the intent was last updated.
	LastUpdatedDateTime *time.Time

	// The output contexts that are activated when this intent is fulfilled.
	OutputContexts []OutputContext

	// If this intent is derived from a built-in intent, the name of the parent intent.
	ParentIntentSignature *string
}

// Provides configuration information for the AMAZON.KendraSearchIntent intent.
// When you use this intent, Amazon Lex searches the specified Amazon Kendra index
// and returns documents from the index that match the user's utterance.
type KendraConfiguration struct {

	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the
	// AMAZON.KendraSearchIntent intent to search. The index must be in the same
	// account and Region as the Amazon Lex bot.
	//
	// This member is required.
	KendraIndex *string

	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response
	// from a query. The filter is in the format defined by Amazon Kendra. For more
	// information, see Filtering queries
	// (https://docs.aws.amazon.com/kendra/latest/dg/filtering.html).
	QueryFilterString *string

	// Determines whether the AMAZON.KendraSearchIntent intent uses a custom query
	// string to query the Amazon Kendra index.
	QueryFilterStringEnabled bool
}

// Specifies a Lambda function that verifies requests to a bot or fulfilles the
// user's request to a bot.
type LambdaCodeHook struct {

	// The version of the request-response that you want Amazon Lex to use to invoke
	// your Lambda function.
	//
	// This member is required.
	CodeHookInterfaceVersion *string

	// The Amazon Resource Name (ARN) of the Lambda function.
	//
	// This member is required.
	LambdaARN *string
}

// The object that provides message text and it's type.
type Message struct {

	// A message in a custom format defined by the client application.
	CustomPayload *CustomPayload

	// A message that defines a response card that the client application can show to
	// the user.
	ImageResponseCard *ImageResponseCard

	// A message in plain text format.
	PlainTextMessage *PlainTextMessage

	// A message in Speech Synthesis Markup Language (SSML).
	SsmlMessage *SSMLMessage
}

// Provides one or more messages that Amazon Lex should send to the user.
type MessageGroup struct {

	// The primary message that Amazon Lex should send to the user.
	//
	// This member is required.
	Message *Message

	// Message variations to send to the user. When variations are defined, Amazon Lex
	// chooses the primary message or one of the variations to send to the user.
	Variations []Message
}

// Determines whether Amazon Lex obscures slot values in conversation logs.
type ObfuscationSetting struct {

	// Value that determines whether Amazon Lex obscures slot values in conversation
	// logs. The default is to obscure the values.
	//
	// This member is required.
	ObfuscationSettingType ObfuscationSettingType
}

// Describes a session context that is activated when an intent is fulfilled.
type OutputContext struct {

	// The name of the output context.
	//
	// This member is required.
	Name *string

	// The amount of time, in seconds, that the output context should remain active.
	// The time is figured from the first time the context is sent to the user.
	//
	// This member is required.
	TimeToLiveInSeconds *int32

	// The number of conversation turns that the output context should remain active.
	// The number of turns is counted from the first time that the context is sent to
	// the user.
	//
	// This member is required.
	TurnsToLive *int32
}

// Defines an ASCII text message to send to the user.
type PlainTextMessage struct {

	// The message to send to the user.
	//
	// This member is required.
	Value *string
}

// Specifies a list of message groups that Amazon Lex sends to a user to elicit a
// response.
type PromptSpecification struct {

	// The maximum number of times the bot tries to elicit a resonse from the user
	// using this prompt.
	//
	// This member is required.
	MaxRetries *int32

	// A collection of messages that Amazon Lex can send to the user. Amazon Lex
	// chooses the actual message to send at runtime.
	//
	// This member is required.
	MessageGroups []MessageGroup

	// Indicates whether the user can interrupt a speech prompt from the bot.
	AllowInterrupt *bool
}

// Specifies a list of message groups that Amazon Lex uses to respond the user
// input.
type ResponseSpecification struct {

	// A collection of responses that Amazon Lex can send to the user. Amazon Lex
	// chooses the actual response to send at runtime.
	//
	// This member is required.
	MessageGroups []MessageGroup

	// Indicates whether the user can interrupt a speech response from Amazon Lex.
	AllowInterrupt *bool
}

// Specifies an Amazon S3 bucket for logging audio conversations
type S3BucketLogDestination struct {

	// The S3 prefix to assign to audio log files.
	//
	// This member is required.
	LogPrefix *string

	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are
	// stored.
	//
	// This member is required.
	S3BucketArn *string

	// The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for
	// encrypting audio log files stored in an S3 bucket.
	KmsKeyArn *string
}

// A sample utterance that invokes an intent or respond to a slot elicitation
// prompt.
type SampleUtterance struct {

	// The sample utterance that Amazon Lex uses to build its machine-learning model to
	// recognize intents.
	//
	// This member is required.
	Utterance *string
}

// Defines one of the values for a slot type.
type SampleValue struct {

	// The value that can be used for a slot type.
	//
	// This member is required.
	Value *string
}

// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment
// of user utterances.
type SentimentAnalysisSettings struct {

	// Sets whether Amazon Lex uses Amazon Comprehend to detect the sentiment of user
	// utterances.
	//
	// This member is required.
	DetectSentiment bool
}

// Specifies the default value to use when a user doesn't provide a value for a
// slot.
type SlotDefaultValue struct {

	// The default value to use when a user doesn't provide a value for a slot.
	//
	// This member is required.
	DefaultValue *string
}

// Defines a list of values that Amazon Lex should use as the default value for a
// slot.
type SlotDefaultValueSpecification struct {

	// A list of default values. Amazon Lex chooses the default value to use in the
	// order that they are presented in the list.
	//
	// This member is required.
	DefaultValueList []SlotDefaultValue
}

// Filters the response from the ListSlots operation.
type SlotFilter struct {

	// The name of the field to use for filtering.
	//
	// This member is required.
	Name SlotFilterName

	// The operator to use for the filter. Specify EQ when the ListSlots operation
	// should return only aliases that equal the specified value. Specify CO when the
	// ListSlots operation should return aliases that contain the specified value.
	//
	// This member is required.
	Operator SlotFilterOperator

	// The value to use to filter the response.
	//
	// This member is required.
	Values []string
}

// Sets the priority that Amazon Lex should use when eliciting slot values from a
// user.
type SlotPriority struct {

	// The priority that a slot should be elicited.
	//
	// This member is required.
	Priority *int32

	// The unique identifier of the slot.
	//
	// This member is required.
	SlotId *string
}

// Specifies attributes for sorting a list of bots.
type SlotSortBy struct {

	// The attribute to use to sort the list.
	//
	// This member is required.
	Attribute SlotSortAttribute

	// The order to sort the list. You can choose ascending or descending.
	//
	// This member is required.
	Order SortOrder
}

// Summary information about a slot, a value that the bot elicits from the user.
type SlotSummary struct {

	// The description of the slot.
	Description *string

	// The timestamp of the last date and time that the slot was updated.
	LastUpdatedDateTime *time.Time

	// Whether the slot is required or optional. An intent is complete when all
	// required slots are filled.
	SlotConstraint SlotConstraint

	// The unique identifier of the slot.
	SlotId *string

	// The name given to the slot.
	SlotName *string

	// The unique identifier for the slot type that defines the values for the slot.
	SlotTypeId *string

	// Prompts that are sent to the user to elicit a value for the slot.
	ValueElicitationPromptSpecification *PromptSpecification
}

// Filters the response from the ListSlotTypes operation.
type SlotTypeFilter struct {

	// The name of the field to use for filtering.
	//
	// This member is required.
	Name SlotTypeFilterName

	// The operator to use for the filter. Specify EQ when the ListSlotTypes operation
	// should return only aliases that equal the specified value. Specify CO when the
	// ListSlotTypes operation should return aliases that contain the specified value.
	//
	// This member is required.
	Operator SlotTypeFilterOperator

	// The value to use to filter the response.
	//
	// This member is required.
	Values []string
}

// Specifies attributes for sorting a list of slot types.
type SlotTypeSortBy struct {

	// The attribute to use to sort the list of slot types.
	//
	// This member is required.
	Attribute SlotTypeSortAttribute

	// The order to sort the list. You can say ascending or descending.
	//
	// This member is required.
	Order SortOrder
}

// Provides summary information about a slot type.
type SlotTypeSummary struct {

	// The description of the slot type.
	Description *string

	// A timestamp of the date and time that the slot type was last updated.
	LastUpdatedDateTime *time.Time

	// If the slot type is derived from a built-on slot type, the name of the parent
	// slot type.
	ParentSlotTypeSignature *string

	// The unique identifier assigned to the slot type.
	SlotTypeId *string

	// The name of the slot type.
	SlotTypeName *string
}

// Each slot type can have a set of values. Each SlotTypeValue represents a value
// that the slot type can take.
type SlotTypeValue struct {

	// The value of the slot type entry.
	SampleValue *SampleValue

	// Additional values releated to the slot type entry.
	Synonyms []SampleValue
}

// Settings that you can use for eliciting a slot value.
type SlotValueElicitationSetting struct {

	// Specifies whether the slot is required or optional.
	//
	// This member is required.
	SlotConstraint SlotConstraint

	// A list of default values for a slot. Default values are used when Amazon Lex
	// hasn't determined a value for a slot. You can specify default values from
	// context variables, sesion attributes, and defined values.
	DefaultValueSpecification *SlotDefaultValueSpecification

	// The prompt that Amazon Lex uses to elicit the slot value from the user.
	PromptSpecification *PromptSpecification

	// If you know a specific pattern that users might respond to an Amazon Lex request
	// for a slot value, you can provide those utterances to improve accuracy. This is
	// optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances []SampleUtterance

	// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer
	// input.
	WaitAndContinueSpecification *WaitAndContinueSpecification
}

// Provides a regular expression used to validate the value of a slot.
type SlotValueRegexFilter struct {

	// A regular expression used to validate the value of a slot. Use a standard
	// regular expression. Amazon Lex supports the following characters in the regular
	// expression:
	//
	// * A-Z, a-z
	//
	// * 0-9
	//
	// * Unicode characters ("\ u")
	//
	// Represent Unicode
	// characters with four digits, for example "\u0041" or "\u005A". The following
	// regular expression operators are not supported:
	//
	// * Infinite repeaters: *, +, or
	// {x,} with no upper bound.
	//
	// * Wild card (.)
	//
	// This member is required.
	Pattern *string
}

// Contains settings used by Amazon Lex to select a slot value.
type SlotValueSelectionSetting struct {

	// Determines the slot resolution strategy that Amazon Lex uses to return slot type
	// values. The field can be set to one of the following values:
	//
	// * OriginalValue -
	// Returns the value entered by the user, if the user value is similar to the slot
	// value.
	//
	// * TopResolution - If there is a resolution list for the slot, return the
	// first value in the resolution list as the slot type value. If there is no
	// resolution list, null is returned.
	//
	// If you don't specify the
	// valueSelectionStrategy, the default is OriginalValue.
	//
	// This member is required.
	ResolutionStrategy SlotValueResolutionStrategy

	// A regular expression used to validate the value of a slot.
	RegexFilter *SlotValueRegexFilter
}

// Defines a Speech Synthesis Markup Language (SSML) prompt.
type SSMLMessage struct {

	// The SSML text that defines the prompt.
	//
	// This member is required.
	Value *string
}

// Defines the messages that Amazon Lex sends to a user to remind them that the bot
// is waiting for a response.
type StillWaitingResponseSpecification struct {

	// How often a message should be sent to the user. Minimum of 1 second, maximum of
	// 5 minutes.
	//
	// This member is required.
	FrequencyInSeconds *int32

	// One or more message groups, each containing one or more messages, that define
	// the prompts that Amazon Lex sends to the user.
	//
	// This member is required.
	MessageGroups []MessageGroup

	// If Amazon Lex waits longer than this length of time for a response, it will stop
	// sending messages.
	//
	// This member is required.
	TimeoutInSeconds *int32

	// Indicates that the user can interrupt the response by speaking while the message
	// is being played.
	AllowInterrupt *bool
}

// Defines the Amazon CloudWatch Logs destination log group for conversation text
// logs.
type TextLogDestination struct {

	// Defines the Amazon CloudWatch Logs log group where text and metadata logs are
	// delivered.
	//
	// This member is required.
	CloudWatch *CloudWatchLogGroupLogDestination
}

// Defines settings to enable text conversation logs.
type TextLogSetting struct {

	// Defines the Amazon CloudWatch Logs destination log group for conversation text
	// logs.
	//
	// This member is required.
	Destination *TextLogDestination

	// Determines whether conversation logs should be stored for an alias.
	//
	// This member is required.
	Enabled bool
}

// Defines settings for using an Amazon Polly voice to communicate with a user.
type VoiceSettings struct {

	// The identifier of the Amazon Polly voice to use.
	//
	// This member is required.
	VoiceId *string
}

// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer
// input.
type WaitAndContinueSpecification struct {

	// The response that Amazon Lex sends to indicate that the bot is ready to continue
	// the conversation.
	//
	// This member is required.
	ContinueResponse *ResponseSpecification

	// The response that Amazon Lex sends to indicate that the bot is waiting for the
	// conversation to continue.
	//
	// This member is required.
	WaitingResponse *ResponseSpecification

	// A response that Amazon Lex sends periodically to the user to indicate that the
	// bot is still waiting for input from the user.
	StillWaitingResponse *StillWaitingResponseSpecification
}
