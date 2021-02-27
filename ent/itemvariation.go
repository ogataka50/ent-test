// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemvariation"
)

// ItemVariation is the model entity for the ItemVariation schema.
type ItemVariation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VariantName holds the value of the "variant_name" field.
	VariantName string `json:"variant_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemVariationQuery when eager-loading is set.
	Edges               ItemVariationEdges `json:"edges"`
	item_item_variation *int
}

// ItemVariationEdges holds the relations/edges for other nodes in the graph.
type ItemVariationEdges struct {
	// OriginItem holds the value of the origin_item edge.
	OriginItem *Item `json:"origin_item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OriginItemOrErr returns the OriginItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemVariationEdges) OriginItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.OriginItem == nil {
			// The edge origin_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.OriginItem, nil
	}
	return nil, &NotLoadedError{edge: "origin_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ItemVariation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case itemvariation.FieldID:
			values[i] = &sql.NullInt64{}
		case itemvariation.FieldVariantName:
			values[i] = &sql.NullString{}
		case itemvariation.ForeignKeys[0]: // item_item_variation
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ItemVariation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ItemVariation fields.
func (iv *ItemVariation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case itemvariation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			iv.ID = int(value.Int64)
		case itemvariation.FieldVariantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field variant_name", values[i])
			} else if value.Valid {
				iv.VariantName = value.String
			}
		case itemvariation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field item_item_variation", value)
			} else if value.Valid {
				iv.item_item_variation = new(int)
				*iv.item_item_variation = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOriginItem queries the "origin_item" edge of the ItemVariation entity.
func (iv *ItemVariation) QueryOriginItem() *ItemQuery {
	return (&ItemVariationClient{config: iv.config}).QueryOriginItem(iv)
}

// Update returns a builder for updating this ItemVariation.
// Note that you need to call ItemVariation.Unwrap() before calling this method if this ItemVariation
// was returned from a transaction, and the transaction was committed or rolled back.
func (iv *ItemVariation) Update() *ItemVariationUpdateOne {
	return (&ItemVariationClient{config: iv.config}).UpdateOne(iv)
}

// Unwrap unwraps the ItemVariation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (iv *ItemVariation) Unwrap() *ItemVariation {
	tx, ok := iv.config.driver.(*txDriver)
	if !ok {
		panic("ent: ItemVariation is not a transactional entity")
	}
	iv.config.driver = tx.drv
	return iv
}

// String implements the fmt.Stringer.
func (iv *ItemVariation) String() string {
	var builder strings.Builder
	builder.WriteString("ItemVariation(")
	builder.WriteString(fmt.Sprintf("id=%v", iv.ID))
	builder.WriteString(", variant_name=")
	builder.WriteString(iv.VariantName)
	builder.WriteByte(')')
	return builder.String()
}

// ItemVariations is a parsable slice of ItemVariation.
type ItemVariations []*ItemVariation

func (iv ItemVariations) config(cfg config) {
	for _i := range iv {
		iv[_i].config = cfg
	}
}