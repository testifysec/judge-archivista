// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/testifysec/archivist/ent/attestation"
	"github.com/testifysec/archivist/ent/attestationcollection"
)

// Attestation is the model entity for the Attestation schema.
type Attestation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttestationQuery when eager-loading is set.
	Edges                               AttestationEdges `json:"edges"`
	attestation_collection_attestations *int
}

// AttestationEdges holds the relations/edges for other nodes in the graph.
type AttestationEdges struct {
	// AttestationCollection holds the value of the attestation_collection edge.
	AttestationCollection *AttestationCollection `json:"attestation_collection,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AttestationCollectionOrErr returns the AttestationCollection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttestationEdges) AttestationCollectionOrErr() (*AttestationCollection, error) {
	if e.loadedTypes[0] {
		if e.AttestationCollection == nil {
			// The edge attestation_collection was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: attestationcollection.Label}
		}
		return e.AttestationCollection, nil
	}
	return nil, &NotLoadedError{edge: "attestation_collection"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attestation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case attestation.FieldID:
			values[i] = new(sql.NullInt64)
		case attestation.FieldType:
			values[i] = new(sql.NullString)
		case attestation.ForeignKeys[0]: // attestation_collection_attestations
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Attestation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attestation fields.
func (a *Attestation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attestation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case attestation.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case attestation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field attestation_collection_attestations", value)
			} else if value.Valid {
				a.attestation_collection_attestations = new(int)
				*a.attestation_collection_attestations = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAttestationCollection queries the "attestation_collection" edge of the Attestation entity.
func (a *Attestation) QueryAttestationCollection() *AttestationCollectionQuery {
	return (&AttestationClient{config: a.config}).QueryAttestationCollection(a)
}

// Update returns a builder for updating this Attestation.
// Note that you need to call Attestation.Unwrap() before calling this method if this Attestation
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attestation) Update() *AttestationUpdateOne {
	return (&AttestationClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Attestation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attestation) Unwrap() *Attestation {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attestation is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attestation) String() string {
	var builder strings.Builder
	builder.WriteString("Attestation(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", type=")
	builder.WriteString(a.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Attestations is a parsable slice of Attestation.
type Attestations []*Attestation

func (a Attestations) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}