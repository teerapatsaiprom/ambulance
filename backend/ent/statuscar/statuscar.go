// Code generated by entc, DO NOT EDIT.

package statuscar

const (
	// Label holds the string label denoting the statuscar type in the database.
	Label = "statuscar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatusDetail holds the string denoting the status_detail field in the database.
	FieldStatusDetail = "status_detail"

	// EdgePredicament holds the string denoting the predicament edge name in mutations.
	EdgePredicament = "predicament"

	// Table holds the table name of the statuscar in the database.
	Table = "statuscars"
	// PredicamentTable is the table the holds the predicament relation/edge.
	PredicamentTable = "predicaments"
	// PredicamentInverseTable is the table name for the Predicament entity.
	// It exists in this package in order to avoid circular dependency with the "predicament" package.
	PredicamentInverseTable = "predicaments"
	// PredicamentColumn is the table column denoting the predicament relation/edge.
	PredicamentColumn = "status_id"
)

// Columns holds all SQL columns for statuscar fields.
var Columns = []string{
	FieldID,
	FieldStatusDetail,
}

var (
	// StatusDetailValidator is a validator for the "status_detail" field. It is called by the builders before save.
	StatusDetailValidator func(string) error
)
