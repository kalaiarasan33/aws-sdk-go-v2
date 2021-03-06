// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Provides the error of the batch create variable API.
type BatchCreateVariableError struct {
	_ struct{} `type:"structure"`

	// The error code.
	Code *int64 `locationName:"code" type:"integer"`

	// The error message.
	Message *string `locationName:"message" type:"string"`

	// The name.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s BatchCreateVariableError) String() string {
	return awsutil.Prettify(s)
}

// Provides the error of the batch get variable API.
type BatchGetVariableError struct {
	_ struct{} `type:"structure"`

	// The error code.
	Code *int64 `locationName:"code" type:"integer"`

	// The error message.
	Message *string `locationName:"message" type:"string"`

	// The error name.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s BatchGetVariableError) String() string {
	return awsutil.Prettify(s)
}

// The model training validation messages.
type DataValidationMetrics struct {
	_ struct{} `type:"structure"`

	// The field-specific model training validation messages.
	FieldLevelMessages []FieldValidationMessage `locationName:"fieldLevelMessages" type:"list"`

	// The file-specific model training validation messages.
	FileLevelMessages []FileValidationMessage `locationName:"fileLevelMessages" type:"list"`
}

// String returns the string representation
func (s DataValidationMetrics) String() string {
	return awsutil.Prettify(s)
}

// The detector.
type Detector struct {
	_ struct{} `type:"structure"`

	// The detector ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Timestamp of when the detector was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The detector description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The detector ID.
	DetectorId *string `locationName:"detectorId" min:"1" type:"string"`

	// The name of the event type.
	EventTypeName *string `locationName:"eventTypeName" min:"1" type:"string"`

	// Timestamp of when the detector was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`
}

// String returns the string representation
func (s Detector) String() string {
	return awsutil.Prettify(s)
}

// The summary of the detector version.
type DetectorVersionSummary struct {
	_ struct{} `type:"structure"`

	// The detector version description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The detector version ID.
	DetectorVersionId *string `locationName:"detectorVersionId" min:"1" type:"string"`

	// Timestamp of when the detector version was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The detector version status.
	Status DetectorVersionStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DetectorVersionSummary) String() string {
	return awsutil.Prettify(s)
}

// The entity details.
type Entity struct {
	_ struct{} `type:"structure"`

	// The entity ID. If you do not know the entityId, you can pass unknown, which
	// is areserved string literal.
	EntityId *string `locationName:"entityId" min:"1" type:"string"`

	// The entity type.
	EntityType *string `locationName:"entityType" type:"string"`
}

// String returns the string representation
func (s Entity) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Entity) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Entity"}
	if s.EntityId != nil && len(*s.EntityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The entity type details.
type EntityType struct {
	_ struct{} `type:"structure"`

	// The entity type ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Timestamp of when the entity type was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The entity type description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// Timestamp of when the entity type was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The entity type name.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s EntityType) String() string {
	return awsutil.Prettify(s)
}

// The event type details.
type EventType struct {
	_ struct{} `type:"structure"`

	// The entity type ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Timestamp of when the event type was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The event type description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The event type entity types.
	EntityTypes []string `locationName:"entityTypes" min:"1" type:"list"`

	// The event type event variables.
	EventVariables []string `locationName:"eventVariables" type:"list"`

	// The event type labels.
	Labels []string `locationName:"labels" type:"list"`

	// Timestamp of when the event type was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The event type name.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s EventType) String() string {
	return awsutil.Prettify(s)
}

// Details for the external events data used for model version training.
type ExternalEventsDetail struct {
	_ struct{} `type:"structure"`

	// The ARN of the role that provides Amazon Fraud Detector access to the data
	// location.
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `locationName:"dataAccessRoleArn" min:"1" type:"string" required:"true"`

	// The Amazon S3 bucket location for the data.
	//
	// DataLocation is a required field
	DataLocation *string `locationName:"dataLocation" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ExternalEventsDetail) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExternalEventsDetail) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExternalEventsDetail"}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 1))
	}

	if s.DataLocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataLocation"))
	}
	if s.DataLocation != nil && len(*s.DataLocation) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataLocation", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The Amazon SageMaker model.
type ExternalModel struct {
	_ struct{} `type:"structure"`

	// The model ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Timestamp of when the model was last created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The event type names.
	EventTypeName *string `locationName:"eventTypeName" min:"1" type:"string"`

	// The input configuration.
	InputConfiguration *ModelInputConfiguration `locationName:"inputConfiguration" type:"structure"`

	// Timestamp of when the model was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The Amazon SageMaker model endpoints.
	ModelEndpoint *string `locationName:"modelEndpoint" type:"string"`

	// The Amazon Fraud Detector status for the external model endpoint
	ModelEndpointStatus ModelEndpointStatus `locationName:"modelEndpointStatus" type:"string" enum:"true"`

	// The source of the model.
	ModelSource ModelSource `locationName:"modelSource" type:"string" enum:"true"`

	// The output configuration.
	OutputConfiguration *ModelOutputConfiguration `locationName:"outputConfiguration" type:"structure"`

	// The role used to invoke the model.
	Role *Role `locationName:"role" type:"structure"`
}

// String returns the string representation
func (s ExternalModel) String() string {
	return awsutil.Prettify(s)
}

// The message details.
type FieldValidationMessage struct {
	_ struct{} `type:"structure"`

	// The message content.
	Content *string `locationName:"content" type:"string"`

	// The field name.
	FieldName *string `locationName:"fieldName" type:"string"`

	// The message ID.
	Identifier *string `locationName:"identifier" type:"string"`

	// The message title.
	Title *string `locationName:"title" type:"string"`

	// The message type.
	Type *string `locationName:"type" type:"string"`
}

// String returns the string representation
func (s FieldValidationMessage) String() string {
	return awsutil.Prettify(s)
}

// The message details.
type FileValidationMessage struct {
	_ struct{} `type:"structure"`

	// The message content.
	Content *string `locationName:"content" type:"string"`

	// The message title.
	Title *string `locationName:"title" type:"string"`

	// The message type.
	Type *string `locationName:"type" type:"string"`
}

// String returns the string representation
func (s FileValidationMessage) String() string {
	return awsutil.Prettify(s)
}

// The KMS key details.
type KMSKey struct {
	_ struct{} `type:"structure"`

	// The encryption key ARN.
	KmsEncryptionKeyArn *string `locationName:"kmsEncryptionKeyArn" min:"7" type:"string"`
}

// String returns the string representation
func (s KMSKey) String() string {
	return awsutil.Prettify(s)
}

// The label details.
type Label struct {
	_ struct{} `type:"structure"`

	// The label ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Timestamp of when the event type was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The label description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// Timestamp of when the label was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The label name.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s Label) String() string {
	return awsutil.Prettify(s)
}

// The label schema.
type LabelSchema struct {
	_ struct{} `type:"structure"`

	// The label mapper maps the Amazon Fraud Detector supported model classification
	// labels (FRAUD, LEGIT) to the appropriate event type labels. For example,
	// if "FRAUD" and "LEGIT" are Amazon Fraud Detector supported labels, this mapper
	// could be: {"FRAUD" => ["0"], "LEGIT" => ["1"]} or {"FRAUD" => ["false"],
	// "LEGIT" => ["true"]} or {"FRAUD" => ["fraud", "abuse"], "LEGIT" => ["legit",
	// "safe"]}. The value part of the mapper is a list, because you may have multiple
	// label variants from your event type for a single Amazon Fraud Detector label.
	//
	// LabelMapper is a required field
	LabelMapper map[string][]string `locationName:"labelMapper" type:"map" required:"true"`
}

// String returns the string representation
func (s LabelSchema) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LabelSchema) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LabelSchema"}

	if s.LabelMapper == nil {
		invalidParams.Add(aws.NewErrParamRequired("LabelMapper"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Model performance metrics data points.
type MetricDataPoint struct {
	_ struct{} `type:"structure"`

	// The false positive rate. This is the percentage of total legitimate events
	// that are incorrectly predicted as fraud.
	Fpr *float64 `locationName:"fpr" type:"float"`

	// The percentage of fraud events correctly predicted as fraudulent as compared
	// to all events predicted as fraudulent.
	Precision *float64 `locationName:"precision" type:"float"`

	// The model threshold that specifies an acceptable fraud capture rate. For
	// example, a threshold of 500 means any model score 500 or above is labeled
	// as fraud.
	Threshold *float64 `locationName:"threshold" type:"float"`

	// The true positive rate. This is the percentage of total fraud the model detects.
	// Also known as capture rate.
	Tpr *float64 `locationName:"tpr" type:"float"`
}

// String returns the string representation
func (s MetricDataPoint) String() string {
	return awsutil.Prettify(s)
}

// The model.
type Model struct {
	_ struct{} `type:"structure"`

	// The ARN of the model.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// Timestamp of when the model was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The model description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The name of the event type.
	EventTypeName *string `locationName:"eventTypeName" type:"string"`

	// Timestamp of last time the model was updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The model ID.
	ModelId *string `locationName:"modelId" min:"1" type:"string"`

	// The model type.
	ModelType ModelTypeEnum `locationName:"modelType" type:"string" enum:"true"`
}

// String returns the string representation
func (s Model) String() string {
	return awsutil.Prettify(s)
}

// A pre-formed Amazon SageMaker model input you can include if your detector
// version includes an imported Amazon SageMaker model endpoint with pass-through
// input configuration.
type ModelEndpointDataBlob struct {
	_ struct{} `type:"structure"`

	// The byte buffer of the Amazon SageMaker model endpoint input data blob.
	//
	// ByteBuffer is automatically base64 encoded/decoded by the SDK.
	ByteBuffer []byte `locationName:"byteBuffer" type:"blob"`

	// The content type of the Amazon SageMaker model endpoint input data blob.
	ContentType *string `locationName:"contentType" min:"1" type:"string"`
}

// String returns the string representation
func (s ModelEndpointDataBlob) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModelEndpointDataBlob) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModelEndpointDataBlob"}
	if s.ContentType != nil && len(*s.ContentType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContentType", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The Amazon SageMaker model input configuration.
type ModelInputConfiguration struct {
	_ struct{} `type:"structure"`

	// Template for constructing the CSV input-data sent to SageMaker. At event-evaluation,
	// the placeholders for variable-names in the template will be replaced with
	// the variable values before being sent to SageMaker.
	CsvInputTemplate *string `locationName:"csvInputTemplate" type:"string"`

	// The format of the model input configuration. The format differs depending
	// on if it is passed through to SageMaker or constructed by Amazon Fraud Detector.
	Format ModelInputDataFormat `locationName:"format" type:"string" enum:"true"`

	// Template for constructing the JSON input-data sent to SageMaker. At event-evaluation,
	// the placeholders for variable names in the template will be replaced with
	// the variable values before being sent to SageMaker.
	JsonInputTemplate *string `locationName:"jsonInputTemplate" type:"string"`

	// The event variables.
	//
	// UseEventVariables is a required field
	UseEventVariables *bool `locationName:"useEventVariables" type:"boolean" required:"true"`
}

// String returns the string representation
func (s ModelInputConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModelInputConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModelInputConfiguration"}

	if s.UseEventVariables == nil {
		invalidParams.Add(aws.NewErrParamRequired("UseEventVariables"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Provides the Amazon Sagemaker model output configuration.
type ModelOutputConfiguration struct {
	_ struct{} `type:"structure"`

	// A map of CSV index values in the SageMaker response to the Amazon Fraud Detector
	// variables.
	CsvIndexToVariableMap map[string]string `locationName:"csvIndexToVariableMap" type:"map"`

	// The format of the model output configuration.
	//
	// Format is a required field
	Format ModelOutputDataFormat `locationName:"format" type:"string" required:"true" enum:"true"`

	// A map of JSON keys in response from SageMaker to the Amazon Fraud Detector
	// variables.
	JsonKeyToVariableMap map[string]string `locationName:"jsonKeyToVariableMap" type:"map"`
}

// String returns the string representation
func (s ModelOutputConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModelOutputConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModelOutputConfiguration"}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The fraud prediction scores.
type ModelScores struct {
	_ struct{} `type:"structure"`

	// The model version.
	ModelVersion *ModelVersion `locationName:"modelVersion" type:"structure"`

	// The model's fraud prediction scores.
	Scores map[string]float64 `locationName:"scores" type:"map"`
}

// String returns the string representation
func (s ModelScores) String() string {
	return awsutil.Prettify(s)
}

// The model version.
type ModelVersion struct {
	_ struct{} `type:"structure"`

	// The model version ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The model ID.
	//
	// ModelId is a required field
	ModelId *string `locationName:"modelId" min:"1" type:"string" required:"true"`

	// The model type.
	//
	// ModelType is a required field
	ModelType ModelTypeEnum `locationName:"modelType" type:"string" required:"true" enum:"true"`

	// The model version number.
	//
	// ModelVersionNumber is a required field
	ModelVersionNumber *string `locationName:"modelVersionNumber" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ModelVersion) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModelVersion) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModelVersion"}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}

	if s.ModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelId"))
	}
	if s.ModelId != nil && len(*s.ModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ModelId", 1))
	}
	if len(s.ModelType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ModelType"))
	}

	if s.ModelVersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelVersionNumber"))
	}
	if s.ModelVersionNumber != nil && len(*s.ModelVersionNumber) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ModelVersionNumber", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The details of the model version.
type ModelVersionDetail struct {
	_ struct{} `type:"structure"`

	// The model version ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The timestamp when the model was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The event details.
	ExternalEventsDetail *ExternalEventsDetail `locationName:"externalEventsDetail" type:"structure"`

	// The timestamp when the model was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The model ID.
	ModelId *string `locationName:"modelId" min:"1" type:"string"`

	// The model type.
	ModelType ModelTypeEnum `locationName:"modelType" type:"string" enum:"true"`

	// The model version number.
	ModelVersionNumber *string `locationName:"modelVersionNumber" type:"string"`

	// The status of the model version.
	Status *string `locationName:"status" type:"string"`

	// The training data schema.
	TrainingDataSchema *TrainingDataSchema `locationName:"trainingDataSchema" type:"structure"`

	// The model version training data source.
	TrainingDataSource TrainingDataSourceEnum `locationName:"trainingDataSource" type:"string" enum:"true"`

	// The training results.
	TrainingResult *TrainingResult `locationName:"trainingResult" type:"structure"`
}

// String returns the string representation
func (s ModelVersionDetail) String() string {
	return awsutil.Prettify(s)
}

// The outcome.
type Outcome struct {
	_ struct{} `type:"structure"`

	// The outcome ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The timestamp when the outcome was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The outcome description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The timestamp when the outcome was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The outcome name.
	Name *string `locationName:"name" min:"1" type:"string"`
}

// String returns the string representation
func (s Outcome) String() string {
	return awsutil.Prettify(s)
}

// The role used to invoke external model endpoints.
type Role struct {
	_ struct{} `type:"structure"`

	// The role ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" type:"string" required:"true"`

	// The role name.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s Role) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Role) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Role"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A rule.
type Rule struct {
	_ struct{} `type:"structure"`

	// The detector for which the rule is associated.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The rule ID.
	//
	// RuleId is a required field
	RuleId *string `locationName:"ruleId" min:"1" type:"string" required:"true"`

	// The rule version.
	//
	// RuleVersion is a required field
	RuleVersion *string `locationName:"ruleVersion" type:"string" required:"true"`
}

// String returns the string representation
func (s Rule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Rule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Rule"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if s.RuleVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The details of the rule.
type RuleDetail struct {
	_ struct{} `type:"structure"`

	// The rule ARN.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The timestamp of when the rule was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The rule description.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The detector for which the rule is associated.
	DetectorId *string `locationName:"detectorId" min:"1" type:"string"`

	// The rule expression.
	Expression *string `locationName:"expression" min:"1" type:"string" sensitive:"true"`

	// The rule language.
	Language Language `locationName:"language" type:"string" enum:"true"`

	// Timestamp of the last time the rule was updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The rule outcomes.
	Outcomes []string `locationName:"outcomes" min:"1" type:"list"`

	// The rule ID.
	RuleId *string `locationName:"ruleId" min:"1" type:"string"`

	// The rule version.
	RuleVersion *string `locationName:"ruleVersion" type:"string"`
}

// String returns the string representation
func (s RuleDetail) String() string {
	return awsutil.Prettify(s)
}

// The rule results.
type RuleResult struct {
	_ struct{} `type:"structure"`

	// The outcomes of the matched rule, based on the rule execution mode.
	Outcomes []string `locationName:"outcomes" type:"list"`

	// The rule ID that was matched, based on the rule execution mode.
	RuleId *string `locationName:"ruleId" type:"string"`
}

// String returns the string representation
func (s RuleResult) String() string {
	return awsutil.Prettify(s)
}

// A key and value pair.
type Tag struct {
	_ struct{} `type:"structure"`

	// A tag key.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// A value assigned to a tag key.
	//
	// Value is a required field
	Value *string `locationName:"value" type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The training data schema.
type TrainingDataSchema struct {
	_ struct{} `type:"structure"`

	// The label schema.
	//
	// LabelSchema is a required field
	LabelSchema *LabelSchema `locationName:"labelSchema" type:"structure" required:"true"`

	// The training data schema variables.
	//
	// ModelVariables is a required field
	ModelVariables []string `locationName:"modelVariables" type:"list" required:"true"`
}

// String returns the string representation
func (s TrainingDataSchema) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TrainingDataSchema) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TrainingDataSchema"}

	if s.LabelSchema == nil {
		invalidParams.Add(aws.NewErrParamRequired("LabelSchema"))
	}

	if s.ModelVariables == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelVariables"))
	}
	if s.LabelSchema != nil {
		if err := s.LabelSchema.Validate(); err != nil {
			invalidParams.AddNested("LabelSchema", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The training metric details.
type TrainingMetrics struct {
	_ struct{} `type:"structure"`

	// The area under the curve. This summarizes true positive rate (TPR) and false
	// positive rate (FPR) across all possible model score thresholds. A model with
	// no predictive power has an AUC of 0.5, whereas a perfect model has a score
	// of 1.0.
	Auc *float64 `locationName:"auc" type:"float"`

	// The data points details.
	MetricDataPoints []MetricDataPoint `locationName:"metricDataPoints" type:"list"`
}

// String returns the string representation
func (s TrainingMetrics) String() string {
	return awsutil.Prettify(s)
}

// The training result details.
type TrainingResult struct {
	_ struct{} `type:"structure"`

	// The validation metrics.
	DataValidationMetrics *DataValidationMetrics `locationName:"dataValidationMetrics" type:"structure"`

	// The training metric details.
	TrainingMetrics *TrainingMetrics `locationName:"trainingMetrics" type:"structure"`
}

// String returns the string representation
func (s TrainingResult) String() string {
	return awsutil.Prettify(s)
}

// The variable.
type Variable struct {
	_ struct{} `type:"structure"`

	// The ARN of the variable.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time when the variable was created.
	CreatedTime *string `locationName:"createdTime" type:"string"`

	// The data source of the variable.
	DataSource DataSource `locationName:"dataSource" type:"string" enum:"true"`

	// The data type of the variable.
	DataType DataType `locationName:"dataType" type:"string" enum:"true"`

	// The default value of the variable.
	DefaultValue *string `locationName:"defaultValue" type:"string"`

	// The description of the variable.
	Description *string `locationName:"description" type:"string"`

	// The time when variable was last updated.
	LastUpdatedTime *string `locationName:"lastUpdatedTime" type:"string"`

	// The name of the variable.
	Name *string `locationName:"name" type:"string"`

	// The variable type of the variable.
	//
	// Valid Values: AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 |
	// BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE
	// | BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS |
	// FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID
	// | PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1
	// | SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME
	// | SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT | SHIPPING_ZIP
	// | USERAGENT
	VariableType *string `locationName:"variableType" type:"string"`
}

// String returns the string representation
func (s Variable) String() string {
	return awsutil.Prettify(s)
}

// A variable in the list of variables for the batch create variable request.
type VariableEntry struct {
	_ struct{} `type:"structure"`

	// The data source of the variable.
	DataSource *string `locationName:"dataSource" type:"string"`

	// The data type of the variable.
	DataType *string `locationName:"dataType" type:"string"`

	// The default value of the variable.
	DefaultValue *string `locationName:"defaultValue" type:"string"`

	// The description of the variable.
	Description *string `locationName:"description" type:"string"`

	// The name of the variable.
	Name *string `locationName:"name" type:"string"`

	// The type of the variable.
	//
	// Valid Values: AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 |
	// BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE
	// | BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS |
	// FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID
	// | PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1
	// | SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME
	// | SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT | SHIPPING_ZIP
	// | USERAGENT
	VariableType *string `locationName:"variableType" type:"string"`
}

// String returns the string representation
func (s VariableEntry) String() string {
	return awsutil.Prettify(s)
}
