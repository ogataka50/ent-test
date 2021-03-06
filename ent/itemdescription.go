// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemdescription"
)

// ItemDescription is the model entity for the ItemDescription schema.
type ItemDescription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemDescriptionQuery when eager-loading is set.
	Edges                 ItemDescriptionEdges `json:"edges"`
	item_item_description *int
}

// ItemDescriptionEdges holds the relations/edges for other nodes in the graph.
type ItemDescriptionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Item `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemDescriptionEdges) OwnerOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ItemDescription) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case itemdescription.FieldID:
			values[i] = &sql.NullInt64{}
		case itemdescription.FieldDescription:
			values[i] = &sql.NullString{}
		case itemdescription.ForeignKeys[0]: // item_item_description
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ItemDescription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ItemDescription fields.
func (id *ItemDescription) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case itemdescription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			id.ID = int(value.Int64)
		case itemdescription.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				id.Description = value.String
			}
		case itemdescription.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field item_item_description", value)
			} else if value.Valid {
				id.item_item_description = new(int)
				*id.item_item_description = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the ItemDescription entity.
func (id *ItemDescription) QueryOwner() *ItemQuery {
	return (&ItemDescriptionClient{config: id.config}).QueryOwner(id)
}

// Update returns a builder for updating this ItemDescription.
// Note that you need to call ItemDescription.Unwrap() before calling this method if this ItemDescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (id *ItemDescription) Update() *ItemDescriptionUpdateOne {
	return (&ItemDescriptionClient{config: id.config}).UpdateOne(id)
}

// Unwrap unwraps the ItemDescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (id *ItemDescription) Unwrap() *ItemDescription {
	tx, ok := id.config.driver.(*txDriver)
	if !ok {
		panic("ent: ItemDescription is not a transactional entity")
	}
	id.config.driver = tx.drv
	return id
}

// String implements the fmt.Stringer.
func (id *ItemDescription) String() string {
	var builder strings.Builder
	builder.WriteString("ItemDescription(")
	builder.WriteString(fmt.Sprintf("id=%v", id.ID))
	builder.WriteString(", description=")
	builder.WriteString(id.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ItemDescriptions is a parsable slice of ItemDescription.
type ItemDescriptions []*ItemDescription

func (id ItemDescriptions) config(cfg config) {
	for _i := range id {
		id[_i].config = cfg
	}
}
