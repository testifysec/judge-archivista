// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/testifysec/archivista/ent/signature"
	"github.com/testifysec/archivista/ent/timestamp"
)

// Timestamp is the model entity for the Timestamp schema.
type Timestamp struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimestampQuery when eager-loading is set.
	Edges                TimestampEdges `json:"edges"`
	signature_timestamps *int
	selectValues         sql.SelectValues
}

// TimestampEdges holds the relations/edges for other nodes in the graph.
type TimestampEdges struct {
	// Signature holds the value of the signature edge.
	Signature *Signature `json:"signature,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// SignatureOrErr returns the Signature value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimestampEdges) SignatureOrErr() (*Signature, error) {
	if e.loadedTypes[0] {
		if e.Signature == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: signature.Label}
		}
		return e.Signature, nil
	}
	return nil, &NotLoadedError{edge: "signature"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Timestamp) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case timestamp.FieldID:
			values[i] = new(sql.NullInt64)
		case timestamp.FieldType:
			values[i] = new(sql.NullString)
		case timestamp.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case timestamp.ForeignKeys[0]: // signature_timestamps
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Timestamp fields.
func (t *Timestamp) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timestamp.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case timestamp.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = value.String
			}
		case timestamp.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				t.Timestamp = value.Time
			}
		case timestamp.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field signature_timestamps", value)
			} else if value.Valid {
				t.signature_timestamps = new(int)
				*t.signature_timestamps = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Timestamp.
// This includes values selected through modifiers, order, etc.
func (t *Timestamp) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QuerySignature queries the "signature" edge of the Timestamp entity.
func (t *Timestamp) QuerySignature() *SignatureQuery {
	return NewTimestampClient(t.config).QuerySignature(t)
}

// Update returns a builder for updating this Timestamp.
// Note that you need to call Timestamp.Unwrap() before calling this method if this Timestamp
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Timestamp) Update() *TimestampUpdateOne {
	return NewTimestampClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Timestamp entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Timestamp) Unwrap() *Timestamp {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Timestamp is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Timestamp) String() string {
	var builder strings.Builder
	builder.WriteString("Timestamp(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("type=")
	builder.WriteString(t.Type)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(t.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Timestamps is a parsable slice of Timestamp.
type Timestamps []*Timestamp