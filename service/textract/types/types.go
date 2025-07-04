// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An adapter selected for use when analyzing documents. Contains an adapter ID
// and a version number. Contains information on pages selected for analysis when
// analyzing documents asychronously.
type Adapter struct {

	// A unique identifier for the adapter resource.
	//
	// This member is required.
	AdapterId *string

	// A string that identifies the version of the adapter.
	//
	// This member is required.
	Version *string

	// Pages is a parameter that the user inputs to specify which pages to apply an
	// adapter to. The following is a list of rules for using this parameter.
	//
	//   - If a page is not specified, it is set to ["1"] by default.
	//
	//   - The following characters are allowed in the parameter's string: 0 1 2 3 4 5
	//   6 7 8 9 - * . No whitespace is allowed.
	//
	//   - When using * to indicate all pages, it must be the only element in the list.
	//
	//   - You can use page intervals, such as ["1-3", "1-1", "4-*"] . Where *
	//   indicates last page of document.
	//
	//   - Specified pages must be greater than 0 and less than or equal to the number
	//   of pages in the document.
	Pages []string

	noSmithyDocumentSerde
}

// Contains information on the adapter, including the adapter ID, Name, Creation
// time, and feature types.
type AdapterOverview struct {

	// A unique identifier for the adapter resource.
	AdapterId *string

	// A string naming the adapter resource.
	AdapterName *string

	// The date and time that the adapter was created.
	CreationTime *time.Time

	// The feature types that the adapter is operating on.
	FeatureTypes []FeatureType

	noSmithyDocumentSerde
}

// Contains information about adapters used when analyzing a document, with each
// adapter specified using an AdapterId and version
type AdaptersConfig struct {

	// A list of adapters to be used when analyzing the specified document.
	//
	// This member is required.
	Adapters []Adapter

	noSmithyDocumentSerde
}

// The dataset configuration options for a given version of an adapter. Can
// include an Amazon S3 bucket if specified.
type AdapterVersionDatasetConfig struct {

	// The S3 bucket name and file name that identifies the document.
	//
	// The AWS Region for the S3 bucket that contains the document must match the
	// Region that you use for Amazon Textract operations.
	//
	// For Amazon Textract to process a file in an S3 bucket, the user must have
	// permission to access the S3 bucket and file.
	ManifestS3Object *S3Object

	noSmithyDocumentSerde
}

// Contains information on the metrics used to evalute the peformance of a given
// adapter version. Includes data for baseline model performance and individual
// adapter version perfromance.
type AdapterVersionEvaluationMetric struct {

	// The F1 score, precision, and recall metrics for the baseline model.
	AdapterVersion *EvaluationMetric

	// The F1 score, precision, and recall metrics for the baseline model.
	Baseline *EvaluationMetric

	// Indicates the feature type being analyzed by a given adapter version.
	FeatureType FeatureType

	noSmithyDocumentSerde
}

// Summary info for an adapter version. Contains information on the AdapterId,
// AdapterVersion, CreationTime, FeatureTypes, and Status.
type AdapterVersionOverview struct {

	// A unique identifier for the adapter associated with a given adapter version.
	AdapterId *string

	// An identified for a given adapter version.
	AdapterVersion *string

	// The date and time that a given adapter version was created.
	CreationTime *time.Time

	// The feature types that the adapter version is operating on.
	FeatureTypes []FeatureType

	// Contains information on the status of a given adapter version.
	Status AdapterVersionStatus

	// A message explaining the status of a given adapter vesion.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Used to contain the information detected by an AnalyzeID operation.
type AnalyzeIDDetections struct {

	// Text of either the normalized field or value associated with it.
	//
	// This member is required.
	Text *string

	// The confidence score of the detected text.
	Confidence *float32

	// Only returned for dates, returns the type of value detected and the date
	// written in a more machine readable way.
	NormalizedValue *NormalizedValue

	noSmithyDocumentSerde
}

// A Block represents items that are recognized in a document within a group of
// pixels close to each other. The information returned in a Block object depends
// on the type of operation. In text detection for documents (for example DetectDocumentText), you
// get information about the detected words and lines of text. In text analysis
// (for example AnalyzeDocument), you can also get information about the fields, tables, and
// selection elements that are detected in the document.
//
// An array of Block objects is returned by both synchronous and asynchronous
// operations. In synchronous operations, such as DetectDocumentText, the array of Block objects is
// the entire set of results. In asynchronous operations, such as GetDocumentAnalysis, the array is
// returned over one or more responses.
//
// For more information, see [How Amazon Textract Works].
//
// [How Amazon Textract Works]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works.html
type Block struct {

	// The type of text item that's recognized. In operations for text detection, the
	// following types are returned:
	//
	//   - PAGE - Contains a list of the LINE Block objects that are detected on a
	//   document page.
	//
	//   - WORD - A word detected on a document page. A word is one or more ISO basic
	//   Latin script characters that aren't separated by spaces.
	//
	//   - LINE - A string of space-delimited, contiguous words that are detected on a
	//   document page.
	//
	// In text analysis operations, the following types are returned:
	//
	//   - PAGE - Contains a list of child Block objects that are detected on a
	//   document page.
	//
	//   - KEY_VALUE_SET - Stores the KEY and VALUE Block objects for linked text
	//   that's detected on a document page. Use the EntityType field to determine if a
	//   KEY_VALUE_SET object is a KEY Block object or a VALUE Block object.
	//
	//   - WORD - A word that's detected on a document page. A word is one or more ISO
	//   basic Latin script characters that aren't separated by spaces.
	//
	//   - LINE - A string of tab-delimited, contiguous words that are detected on a
	//   document page.
	//
	//   - TABLE - A table that's detected on a document page. A table is grid-based
	//   information with two or more rows or columns, with a cell span of one row and
	//   one column each.
	//
	//   - TABLE_TITLE - The title of a table. A title is typically a line of text
	//   above or below a table, or embedded as the first row of a table.
	//
	//   - TABLE_FOOTER - The footer associated with a table. A footer is typically a
	//   line or lines of text below a table or embedded as the last row of a table.
	//
	//   - CELL - A cell within a detected table. The cell is the parent of the block
	//   that contains the text in the cell.
	//
	//   - MERGED_CELL - A cell in a table whose content spans more than one row or
	//   column. The Relationships array for this cell contain data from individual
	//   cells.
	//
	//   - SELECTION_ELEMENT - A selection element such as an option button (radio
	//   button) or a check box that's detected on a document page. Use the value of
	//   SelectionStatus to determine the status of the selection element.
	//
	//   - SIGNATURE - The location and confidence score of a signature detected on a
	//   document page. Can be returned as part of a Key-Value pair or a detected cell.
	//
	//   - QUERY - A question asked during the call of AnalyzeDocument. Contains an
	//   alias and an ID that attaches it to its answer.
	//
	//   - QUERY_RESULT - A response to a question asked during the call of analyze
	//   document. Comes with an alias and ID for ease of locating in a response. Also
	//   contains location and confidence score.
	//
	// The following BlockTypes are only returned for Amazon Textract Layout.
	//
	//   - LAYOUT_TITLE - The main title of the document.
	//
	//   - LAYOUT_HEADER - Text located in the top margin of the document.
	//
	//   - LAYOUT_FOOTER - Text located in the bottom margin of the document.
	//
	//   - LAYOUT_SECTION_HEADER - The titles of sections within a document.
	//
	//   - LAYOUT_PAGE_NUMBER - The page number of the documents.
	//
	//   - LAYOUT_LIST - Any information grouped together in list form.
	//
	//   - LAYOUT_FIGURE - Indicates the location of an image in a document.
	//
	//   - LAYOUT_TABLE - Indicates the location of a table in the document.
	//
	//   - LAYOUT_KEY_VALUE - Indicates the location of form key-values in a document.
	//
	//   - LAYOUT_TEXT - Text that is present typically as a part of paragraphs in
	//   documents.
	BlockType BlockType

	// The column in which a table cell appears. The first column position is 1.
	// ColumnIndex isn't returned by DetectDocumentText and GetDocumentTextDetection .
	ColumnIndex *int32

	// The number of columns that a table cell spans. ColumnSpan isn't returned by
	// DetectDocumentText and GetDocumentTextDetection .
	ColumnSpan *int32

	// The confidence score that Amazon Textract has in the accuracy of the recognized
	// text and the accuracy of the geometry points around the recognized text.
	Confidence *float32

	// The type of entity.
	//
	// The following entity types can be returned by FORMS analysis:
	//
	//   - KEY - An identifier for a field on the document.
	//
	//   - VALUE - The field text.
	//
	// The following entity types can be returned by TABLES analysis:
	//
	//   - COLUMN_HEADER - Identifies a cell that is a header of a column.
	//
	//   - TABLE_TITLE - Identifies a cell that is a title within the table.
	//
	//   - TABLE_SECTION_TITLE - Identifies a cell that is a title of a section within
	//   a table. A section title is a cell that typically spans an entire row above a
	//   section.
	//
	//   - TABLE_FOOTER - Identifies a cell that is a footer of a table.
	//
	//   - TABLE_SUMMARY - Identifies a summary cell of a table. A summary cell can be
	//   a row of a table or an additional, smaller table that contains summary
	//   information for another table.
	//
	//   - STRUCTURED_TABLE - Identifies a table with column headers where the content
	//   of each row corresponds to the headers.
	//
	//   - SEMI_STRUCTURED_TABLE - Identifies a non-structured table.
	//
	// EntityTypes isn't returned by DetectDocumentText and GetDocumentTextDetection .
	EntityTypes []EntityType

	// The location of the recognized text on the image. It includes an axis-aligned,
	// coarse bounding box that surrounds the text, and a finer-grain polygon for more
	// accurate spatial information.
	Geometry *Geometry

	// The identifier for the recognized text. The identifier is only unique for a
	// single operation.
	Id *string

	// The page on which a block was detected. Page is returned by synchronous and
	// asynchronous operations. Page values greater than 1 are only returned for
	// multipage documents that are in PDF or TIFF format. A scanned image (JPEG/PNG)
	// provided to an asynchronous operation, even if it contains multiple document
	// pages, is considered a single-page document. This means that for scanned images
	// the value of Page is always 1.
	Page *int32

	//
	Query *Query

	// A list of relationship objects that describe how blocks are related to each
	// other. For example, a LINE block object contains a CHILD relationship type with
	// the WORD blocks that make up the line of text. There aren't Relationship objects
	// in the list for relationships that don't exist, such as when the current block
	// has no child blocks.
	Relationships []Relationship

	// The row in which a table cell is located. The first row position is 1. RowIndex
	// isn't returned by DetectDocumentText and GetDocumentTextDetection .
	RowIndex *int32

	// The number of rows that a table cell spans. RowSpan isn't returned by
	// DetectDocumentText and GetDocumentTextDetection .
	RowSpan *int32

	// The selection status of a selection element, such as an option button or check
	// box.
	SelectionStatus SelectionStatus

	// The word or line of text that's recognized by Amazon Textract.
	Text *string

	// The kind of text that Amazon Textract has detected. Can check for handwritten
	// text and printed text.
	TextType TextType

	noSmithyDocumentSerde
}

// The bounding box around the detected page, text, key-value pair, table, table
// cell, or selection element on a document page. The left (x-coordinate) and top
// (y-coordinate) are coordinates that represent the top and left sides of the
// bounding box. Note that the upper-left corner of the image is the origin (0,0).
//
// The top and left values returned are ratios of the overall document page size.
// For example, if the input image is 700 x 200 pixels, and the top-left coordinate
// of the bounding box is 350 x 50 pixels, the API returns a left value of 0.5
// (350/700) and a top value of 0.25 (50/200).
//
// The width and height values represent the dimensions of the bounding box as a
// ratio of the overall document page dimension. For example, if the document page
// size is 700 x 200 pixels, and the bounding box width is 70 pixels, the width
// returned is 0.1.
type BoundingBox struct {

	// The height of the bounding box as a ratio of the overall document page height.
	Height float32

	// The left coordinate of the bounding box as a ratio of overall document page
	// width.
	Left float32

	// The top coordinate of the bounding box as a ratio of overall document page
	// height.
	Top float32

	// The width of the bounding box as a ratio of the overall document page width.
	Width float32

	noSmithyDocumentSerde
}

// A structure that holds information regarding a detected signature on a page.
type DetectedSignature struct {

	// The page a detected signature was found on.
	Page *int32

	noSmithyDocumentSerde
}

// The input document, either as bytes or as an S3 object.
//
// You pass image bytes to an Amazon Textract API operation by using the Bytes
// property. For example, you would use the Bytes property to pass a document
// loaded from a local file system. Image bytes passed by using the Bytes property
// must be base64 encoded. Your code might not need to encode document file bytes
// if you're using an AWS SDK to call Amazon Textract API operations.
//
// You pass images stored in an S3 bucket to an Amazon Textract API operation by
// using the S3Object property. Documents stored in an S3 bucket don't need to be
// base64 encoded.
//
// The AWS Region for the S3 bucket that contains the S3 object must match the AWS
// Region that you use for Amazon Textract operations.
//
// If you use the AWS CLI to call Amazon Textract operations, passing image bytes
// using the Bytes property isn't supported. You must first upload the document to
// an Amazon S3 bucket, and then call the operation using the S3Object property.
//
// For Amazon Textract to process an S3 object, the user must have permission to
// access the S3 object.
type Document struct {

	// A blob of base64-encoded document bytes. The maximum size of a document that's
	// provided in a blob of bytes is 5 MB. The document bytes must be in PNG or JPEG
	// format.
	//
	// If you're using an AWS SDK to call Amazon Textract, you might not need to
	// base64-encode image bytes passed using the Bytes field.
	Bytes []byte

	// Identifies an S3 object as the document source. The maximum size of a document
	// that's stored in an S3 bucket is 5 MB.
	S3Object *S3Object

	noSmithyDocumentSerde
}

// Summary information about documents grouped by the same document type.
type DocumentGroup struct {

	// A list of the detected signatures found in a document group.
	DetectedSignatures []DetectedSignature

	// An array that contains information about the pages of a document, defined by
	// logical boundary.
	SplitDocuments []SplitDocument

	// The type of document that Amazon Textract has detected. See [Analyze Lending Response Objects] for a list of all
	// types returned by Textract.
	//
	// [Analyze Lending Response Objects]: https://docs.aws.amazon.com/textract/latest/dg/lending-response-objects.html
	Type *string

	// A list of any expected signatures not found in a document group.
	UndetectedSignatures []UndetectedSignature

	noSmithyDocumentSerde
}

// The Amazon S3 bucket that contains the document to be processed. It's used by
// asynchronous operations.
//
// The input document can be an image file in JPEG or PNG format. It can also be a
// file in PDF format.
type DocumentLocation struct {

	// The Amazon S3 bucket that contains the input document.
	S3Object *S3Object

	noSmithyDocumentSerde
}

// Information about the input document.
type DocumentMetadata struct {

	// The number of pages that are detected in the document.
	Pages *int32

	noSmithyDocumentSerde
}

// The evaluation metrics (F1 score, Precision, and Recall) for an adapter version.
type EvaluationMetric struct {

	// The F1 score for an adapter version.
	F1Score float32

	// The Precision score for an adapter version.
	Precision float32

	// The Recall score for an adapter version.
	Recall float32

	noSmithyDocumentSerde
}

// Returns the kind of currency detected.
type ExpenseCurrency struct {

	// Currency code for detected currency. the current supported codes are:
	//
	//   - USD
	//
	//   - EUR
	//
	//   - GBP
	//
	//   - CAD
	//
	//   - INR
	//
	//   - JPY
	//
	//   - CHF
	//
	//   - AUD
	//
	//   - CNY
	//
	//   - BZR
	//
	//   - SEK
	//
	//   - HKD
	Code *string

	// Percentage confideence in the detected currency.
	Confidence *float32

	noSmithyDocumentSerde
}

// An object used to store information about the Value or Label detected by Amazon
// Textract.
type ExpenseDetection struct {

	// The confidence in detection, as a percentage
	Confidence *float32

	// Information about where the following items are located on a document page:
	// detected page, text, key-value pairs, tables, table cells, and selection
	// elements.
	Geometry *Geometry

	// The word or line of text recognized by Amazon Textract
	Text *string

	noSmithyDocumentSerde
}

// The structure holding all the information returned by AnalyzeExpense
type ExpenseDocument struct {

	// This is a block object, the same as reported when DetectDocumentText is run on
	// a document. It provides word level recognition of text.
	Blocks []Block

	// Denotes which invoice or receipt in the document the information is coming
	// from. First document will be 1, the second 2, and so on.
	ExpenseIndex *int32

	// Information detected on each table of a document, seperated into LineItems .
	LineItemGroups []LineItemGroup

	// Any information found outside of a table by Amazon Textract.
	SummaryFields []ExpenseField

	noSmithyDocumentSerde
}

// Breakdown of detected information, seperated into the catagories Type,
// LabelDetection, and ValueDetection
type ExpenseField struct {

	// Shows the kind of currency, both the code and confidence associated with any
	// monatary value detected.
	Currency *ExpenseCurrency

	// Shows which group a response object belongs to, such as whether an address line
	// belongs to the vendor's address or the recipent's address.
	GroupProperties []ExpenseGroupProperty

	// The explicitly stated label of a detected element.
	LabelDetection *ExpenseDetection

	// The page number the value was detected on.
	PageNumber *int32

	// The implied label of a detected element. Present alongside LabelDetection for
	// explicit elements.
	Type *ExpenseType

	// The value of a detected element. Present in explicit and implicit elements.
	ValueDetection *ExpenseDetection

	noSmithyDocumentSerde
}

// Shows the group that a certain key belongs to. This helps differentiate between
// names and addresses for different organizations, that can be hard to determine
// via JSON response.
type ExpenseGroupProperty struct {

	// Provides a group Id number, which will be the same for each in the group.
	Id *string

	// Informs you on whether the expense group is a name or an address.
	Types []string

	noSmithyDocumentSerde
}

// An object used to store information about the Type detected by Amazon Textract.
type ExpenseType struct {

	// The confidence of accuracy, as a percentage.
	Confidence *float32

	// The word or line of text detected by Amazon Textract.
	Text *string

	noSmithyDocumentSerde
}

// Contains information extracted by an analysis operation after using
// StartLendingAnalysis.
type Extraction struct {

	// The structure holding all the information returned by AnalyzeExpense
	ExpenseDocument *ExpenseDocument

	// The structure that lists each document processed in an AnalyzeID operation.
	IdentityDocument *IdentityDocument

	// Holds the structured data returned by AnalyzeDocument for lending documents.
	LendingDocument *LendingDocument

	noSmithyDocumentSerde
}

// Information about where the following items are located on a document page:
// detected page, text, key-value pairs, tables, table cells, and selection
// elements.
type Geometry struct {

	// An axis-aligned coarse representation of the location of the recognized item on
	// the document page.
	BoundingBox *BoundingBox

	// Within the bounding box, a fine-grained polygon around the recognized item.
	Polygon []Point

	// Provides a numerical value corresponding to the rotation of the text.
	RotationAngle *float32

	noSmithyDocumentSerde
}

// Shows the results of the human in the loop evaluation. If there is no
// HumanLoopArn, the input did not trigger human review.
type HumanLoopActivationOutput struct {

	// Shows the result of condition evaluations, including those conditions which
	// activated a human review.
	//
	// This value conforms to the media type: application/json
	HumanLoopActivationConditionsEvaluationResults *string

	// Shows if and why human review was needed.
	HumanLoopActivationReasons []string

	// The Amazon Resource Name (ARN) of the HumanLoop created.
	HumanLoopArn *string

	noSmithyDocumentSerde
}

// Sets up the human review workflow the document will be sent to if one of the
// conditions is met. You can also set certain attributes of the image before
// review.
type HumanLoopConfig struct {

	// The Amazon Resource Name (ARN) of the flow definition.
	//
	// This member is required.
	FlowDefinitionArn *string

	// The name of the human workflow used for this image. This should be kept unique
	// within a region.
	//
	// This member is required.
	HumanLoopName *string

	// Sets attributes of the input data.
	DataAttributes *HumanLoopDataAttributes

	noSmithyDocumentSerde
}

// Allows you to set attributes of the image. Currently, you can declare an image
// as free of personally identifiable information and adult content.
type HumanLoopDataAttributes struct {

	// Sets whether the input image is free of personally identifiable information or
	// adult content.
	ContentClassifiers []ContentClassifier

	noSmithyDocumentSerde
}

// The structure that lists each document processed in an AnalyzeID operation.
type IdentityDocument struct {

	// Individual word recognition, as returned by document detection.
	Blocks []Block

	// Denotes the placement of a document in the IdentityDocument list. The first
	// document is marked 1, the second 2 and so on.
	DocumentIndex *int32

	// The structure used to record information extracted from identity documents.
	// Contains both normalized field and value of the extracted text.
	IdentityDocumentFields []IdentityDocumentField

	noSmithyDocumentSerde
}

// Structure containing both the normalized type of the extracted information and
// the text associated with it. These are extracted as Type and Value respectively.
type IdentityDocumentField struct {

	// Used to contain the information detected by an AnalyzeID operation.
	Type *AnalyzeIDDetections

	// Used to contain the information detected by an AnalyzeID operation.
	ValueDetection *AnalyzeIDDetections

	noSmithyDocumentSerde
}

// The results extracted for a lending document.
type LendingDetection struct {

	// The confidence level for the text of a detected value in a lending document.
	Confidence *float32

	// Information about where the following items are located on a document page:
	// detected page, text, key-value pairs, tables, table cells, and selection
	// elements.
	Geometry *Geometry

	// The selection status of a selection element, such as an option button or check
	// box.
	SelectionStatus SelectionStatus

	// The text extracted for a detected value in a lending document.
	Text *string

	noSmithyDocumentSerde
}

// Holds the structured data returned by AnalyzeDocument for lending documents.
type LendingDocument struct {

	// An array of LendingField objects.
	LendingFields []LendingField

	// A list of signatures detected in a lending document.
	SignatureDetections []SignatureDetection

	noSmithyDocumentSerde
}

// Holds the normalized key-value pairs returned by AnalyzeDocument, including the
// document type, detected text, and geometry.
type LendingField struct {

	// The results extracted for a lending document.
	KeyDetection *LendingDetection

	// The type of the lending document.
	Type *string

	// An array of LendingDetection objects.
	ValueDetections []LendingDetection

	noSmithyDocumentSerde
}

// Contains the detections for each page analyzed through the Analyze Lending API.
type LendingResult struct {

	// An array of Extraction to hold structured data. e.g. normalized key value pairs
	// instead of raw OCR detections .
	Extractions []Extraction

	// The page number for a page, with regard to whole submission.
	Page *int32

	// The classifier result for a given page.
	PageClassification *PageClassification

	noSmithyDocumentSerde
}

// Contains information regarding DocumentGroups and UndetectedDocumentTypes.
type LendingSummary struct {

	// Contains an array of all DocumentGroup objects.
	DocumentGroups []DocumentGroup

	// UndetectedDocumentTypes.
	UndetectedDocumentTypes []string

	noSmithyDocumentSerde
}

// A structure that holds information about the different lines found in a
// document's tables.
type LineItemFields struct {

	// ExpenseFields used to show information from detected lines on a table.
	LineItemExpenseFields []ExpenseField

	noSmithyDocumentSerde
}

// A grouping of tables which contain LineItems, with each table identified by the
// table's LineItemGroupIndex .
type LineItemGroup struct {

	// The number used to identify a specific table in a document. The first table
	// encountered will have a LineItemGroupIndex of 1, the second 2, etc.
	LineItemGroupIndex *int32

	// The breakdown of information on a particular line of a table.
	LineItems []LineItemFields

	noSmithyDocumentSerde
}

// Contains information relating to dates in a document, including the type of
// value, and the value.
type NormalizedValue struct {

	// The value of the date, written as Year-Month-DayTHour:Minute:Second.
	Value *string

	// The normalized type of the value detected. In this case, DATE.
	ValueType ValueType

	noSmithyDocumentSerde
}

// The Amazon Simple Notification Service (Amazon SNS) topic to which Amazon
// Textract publishes the completion status of an asynchronous document operation.
type NotificationChannel struct {

	// The Amazon Resource Name (ARN) of an IAM role that gives Amazon Textract
	// publishing permissions to the Amazon SNS topic.
	//
	// This member is required.
	RoleArn *string

	// The Amazon SNS topic that Amazon Textract posts the completion status to.
	//
	// This member is required.
	SNSTopicArn *string

	noSmithyDocumentSerde
}

// Sets whether or not your output will go to a user created bucket. Used to set
// the name of the bucket, and the prefix on the output file.
//
// OutputConfig is an optional parameter which lets you adjust where your output
// will be placed. By default, Amazon Textract will store the results internally
// and can only be accessed by the Get API operations. With OutputConfig enabled,
// you can set the name of the bucket the output will be sent to the file prefix of
// the results where you can download your results. Additionally, you can set the
// KMSKeyID parameter to a customer master key (CMK) to encrypt your output.
// Without this parameter set Amazon Textract will encrypt server-side using the
// AWS managed CMK for Amazon S3.
//
// Decryption of Customer Content is necessary for processing of the documents by
// Amazon Textract. If your account is opted out under an AI services opt out
// policy then all unencrypted Customer Content is immediately and permanently
// deleted after the Customer Content has been processed by the service. No copy of
// of the output is retained by Amazon Textract. For information about how to opt
// out, see [Managing AI services opt-out policy.]
//
// For more information on data privacy, see the [Data Privacy FAQ].
//
// [Managing AI services opt-out policy.]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html
// [Data Privacy FAQ]: https://aws.amazon.com/compliance/data-privacy-faq/
type OutputConfig struct {

	// The name of the bucket your output will go to.
	//
	// This member is required.
	S3Bucket *string

	// The prefix of the object key that the output will be saved to. When not
	// enabled, the prefix will be “textract_output".
	S3Prefix *string

	noSmithyDocumentSerde
}

// The class assigned to a Page object detected in an input document. Contains
// information regarding the predicted type/class of a document's page and the page
// number that the Page object was detected on.
type PageClassification struct {

	//  The page number the value was detected on, relative to Amazon Textract's
	// starting position.
	//
	// This member is required.
	PageNumber []Prediction

	// The class, or document type, assigned to a detected Page object. The class, or
	// document type, assigned to a detected Page object.
	//
	// This member is required.
	PageType []Prediction

	noSmithyDocumentSerde
}

// The X and Y coordinates of a point on a document page. The X and Y values that
// are returned are ratios of the overall document page size. For example, if the
// input document is 700 x 200 and the operation returns X=0.5 and Y=0.25, then the
// point is at the (350,50) pixel coordinate on the document page.
//
// An array of Point objects, Polygon , is returned by DetectDocumentText. Polygon represents a
// fine-grained polygon around detected text. For more information, see Geometry in
// the Amazon Textract Developer Guide.
type Point struct {

	// The value of the X coordinate for a point on a Polygon .
	X float32

	// The value of the Y coordinate for a point on a Polygon .
	Y float32

	noSmithyDocumentSerde
}

// Contains information regarding predicted values returned by Amazon Textract
// operations, including the predicted value and the confidence in the predicted
// value.
type Prediction struct {

	// Amazon Textract's confidence in its predicted value.
	Confidence *float32

	// The predicted value of a detected object.
	Value *string

	noSmithyDocumentSerde
}

type QueriesConfig struct {

	//
	//
	// This member is required.
	Queries []Query

	noSmithyDocumentSerde
}

// Each query contains the question you want to ask in the Text and the alias you
// want to associate.
type Query struct {

	// Question that Amazon Textract will apply to the document. An example would be
	// "What is the customer's SSN?"
	//
	// This member is required.
	Text *string

	// Alias attached to the query, for ease of location.
	Alias *string

	// Pages is a parameter that the user inputs to specify which pages to apply a
	// query to. The following is a list of rules for using this parameter.
	//
	//   - If a page is not specified, it is set to ["1"] by default.
	//
	//   - The following characters are allowed in the parameter's string: 0 1 2 3 4 5
	//   6 7 8 9 - * . No whitespace is allowed.
	//
	//   - When using * to indicate all pages, it must be the only element in the list.
	//
	//   - You can use page intervals, such as [“1-3”, “1-1”, “4-*”] . Where *
	//   indicates last page of document.
	//
	//   - Specified pages must be greater than 0 and less than or equal to the number
	//   of pages in the document.
	Pages []string

	noSmithyDocumentSerde
}

// Information about how blocks are related to each other. A Block object contains
// 0 or more Relation objects in a list, Relationships . For more information, see Block
// .
//
// The Type element provides the type of the relationship for all blocks in the IDs
// array.
type Relationship struct {

	// An array of IDs for related blocks. You can get the type of the relationship
	// from the Type element.
	Ids []string

	// The type of relationship between the blocks in the IDs array and the current
	// block. The following list describes the relationship types that can be returned.
	//
	//   - VALUE - A list that contains the ID of the VALUE block that's associated
	//   with the KEY of a key-value pair.
	//
	//   - CHILD - A list of IDs that identify blocks found within the current block
	//   object. For example, WORD blocks have a CHILD relationship to the LINE block
	//   type.
	//
	//   - MERGED_CELL - A list of IDs that identify each of the MERGED_CELL block
	//   types in a table.
	//
	//   - ANSWER - A list that contains the ID of the QUERY_RESULT block that’s
	//   associated with the corresponding QUERY block.
	//
	//   - TABLE - A list of IDs that identify associated TABLE block types.
	//
	//   - TABLE_TITLE - A list that contains the ID for the TABLE_TITLE block type in
	//   a table.
	//
	//   - TABLE_FOOTER - A list of IDs that identify the TABLE_FOOTER block types in
	//   a table.
	Type RelationshipType

	noSmithyDocumentSerde
}

// The S3 bucket name and file name that identifies the document.
//
// The AWS Region for the S3 bucket that contains the document must match the
// Region that you use for Amazon Textract operations.
//
// For Amazon Textract to process a file in an S3 bucket, the user must have
// permission to access the S3 bucket and file.
type S3Object struct {

	// The name of the S3 bucket. Note that the # character is not valid in the file
	// name.
	Bucket *string

	// The file name of the input document. Image files may be in PDF, TIFF, JPEG, or
	// PNG format.
	Name *string

	// If the bucket has versioning enabled, you can specify the object version.
	Version *string

	noSmithyDocumentSerde
}

// Information regarding a detected signature on a page.
type SignatureDetection struct {

	// The confidence, from 0 to 100, in the predicted values for a detected signature.
	Confidence *float32

	// Information about where the following items are located on a document page:
	// detected page, text, key-value pairs, tables, table cells, and selection
	// elements.
	Geometry *Geometry

	noSmithyDocumentSerde
}

// Contains information about the pages of a document, defined by logical boundary.
type SplitDocument struct {

	// The index for a given document in a DocumentGroup of a specific Type.
	Index *int32

	// An array of page numbers for a for a given document, ordered by logical
	// boundary.
	Pages []int32

	noSmithyDocumentSerde
}

// A structure containing information about an undetected signature on a page
// where it was expected but not found.
type UndetectedSignature struct {

	// The page where a signature was expected but not found.
	Page *int32

	noSmithyDocumentSerde
}

// A warning about an issue that occurred during asynchronous text analysis (StartDocumentAnalysis ) or
// asynchronous document text detection (StartDocumentTextDetection ).
type Warning struct {

	// The error code for the warning.
	ErrorCode *string

	// A list of the pages that the warning applies to.
	Pages []int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
