package query

type Field string
type Value string

const (
	FieldName                = "name"
	FieldUID                 = "uid"
	FieldCreationTimeStamp   = "creationTimestamp"
	FieldLastUpdateTimestamp = "lastUpdateTimestamp"
	FieldLabel               = "label"
	FieldAnnotation          = "annotation"
	FieldClusterName         = "clusterName"
	FieldNamespace           = "namespace"
	FieldStatus              = "status"
	FieldOwnerReference      = "ownerReference"
	FieldOwnerKind           = "ownerKind"
)

var SortableFields = []Field{
	FieldCreationTimeStamp,
	FieldLastUpdateTimestamp,
	FieldName,
}

// Field contains all the query field that can be compared
var ComparableFields = []Field{
	FieldName,
	FieldUID,
	FieldLabel,
	FieldAnnotation,
	FieldClusterName,
	FieldNamespace,
	FieldStatus,
	FieldOwnerReference,
	FieldOwnerKind,
}
