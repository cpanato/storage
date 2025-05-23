// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package edgetype

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the edgetype type in the database.
	Label = "edge_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProtoMessage holds the string denoting the proto_message field in the database.
	FieldProtoMessage = "proto_message"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldNodeID holds the string denoting the node_id field in the database.
	FieldNodeID = "node_id"
	// FieldToNodeID holds the string denoting the to_node_id field in the database.
	FieldToNodeID = "to_node_id"
	// EdgeFrom holds the string denoting the from edge name in mutations.
	EdgeFrom = "from"
	// EdgeTo holds the string denoting the to edge name in mutations.
	EdgeTo = "to"
	// EdgeDocuments holds the string denoting the documents edge name in mutations.
	EdgeDocuments = "documents"
	// EdgeNodeLists holds the string denoting the node_lists edge name in mutations.
	EdgeNodeLists = "node_lists"
	// Table holds the table name of the edgetype in the database.
	Table = "edge_types"
	// FromTable is the table that holds the from relation/edge.
	FromTable = "edge_types"
	// FromInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	FromInverseTable = "nodes"
	// FromColumn is the table column denoting the from relation/edge.
	FromColumn = "node_id"
	// ToTable is the table that holds the to relation/edge.
	ToTable = "edge_types"
	// ToInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	ToInverseTable = "nodes"
	// ToColumn is the table column denoting the to relation/edge.
	ToColumn = "to_node_id"
	// DocumentsTable is the table that holds the documents relation/edge. The primary key declared below.
	DocumentsTable = "document_edge_types"
	// DocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentsInverseTable = "documents"
	// NodeListsTable is the table that holds the node_lists relation/edge. The primary key declared below.
	NodeListsTable = "node_list_edges"
	// NodeListsInverseTable is the table name for the NodeList entity.
	// It exists in this package in order to avoid circular dependency with the "nodelist" package.
	NodeListsInverseTable = "node_lists"
)

// Columns holds all SQL columns for edgetype fields.
var Columns = []string{
	FieldID,
	FieldProtoMessage,
	FieldType,
	FieldNodeID,
	FieldToNodeID,
}

var (
	// DocumentsPrimaryKey and DocumentsColumn2 are the table columns denoting the
	// primary key for the documents relation (M2M).
	DocumentsPrimaryKey = []string{"document_id", "edge_type_id"}
	// NodeListsPrimaryKey and NodeListsColumn2 are the table columns denoting the
	// primary key for the node_lists relation (M2M).
	NodeListsPrimaryKey = []string{"node_list_id", "edge_type_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/protobom/storage/internal/backends/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeUNKNOWN              Type = "UNKNOWN"
	TypeAmends               Type = "amends"
	TypeAncestor             Type = "ancestor"
	TypeBuildDependency      Type = "buildDependency"
	TypeBuildTool            Type = "buildTool"
	TypeContains             Type = "contains"
	TypeContainedBy          Type = "contained_by"
	TypeCopy                 Type = "copy"
	TypeDataFile             Type = "dataFile"
	TypeDependencyManifest   Type = "dependencyManifest"
	TypeDependsOn            Type = "dependsOn"
	TypeDependencyOf         Type = "dependencyOf"
	TypeDescendant           Type = "descendant"
	TypeDescribes            Type = "describes"
	TypeDescribedBy          Type = "describedBy"
	TypeDevDependency        Type = "devDependency"
	TypeDevTool              Type = "devTool"
	TypeDistributionArtifact Type = "distributionArtifact"
	TypeDocumentation        Type = "documentation"
	TypeDynamicLink          Type = "dynamicLink"
	TypeExample              Type = "example"
	TypeExpandedFromArchive  Type = "expandedFromArchive"
	TypeFileAdded            Type = "fileAdded"
	TypeFileDeleted          Type = "fileDeleted"
	TypeFileModified         Type = "fileModified"
	TypeGenerates            Type = "generates"
	TypeGeneratedFrom        Type = "generatedFrom"
	TypeMetafile             Type = "metafile"
	TypeOptionalComponent    Type = "optionalComponent"
	TypeOptionalDependency   Type = "optionalDependency"
	TypeOther                Type = "other"
	TypePackages             Type = "packages"
	TypePatch                Type = "patch"
	TypePrerequisite         Type = "prerequisite"
	TypePrerequisiteFor      Type = "prerequisiteFor"
	TypeProvidedDependency   Type = "providedDependency"
	TypeRequirementFor       Type = "requirementFor"
	TypeRuntimeDependency    Type = "runtimeDependency"
	TypeSpecificationFor     Type = "specificationFor"
	TypeStaticLink           Type = "staticLink"
	TypeTest                 Type = "test"
	TypeTestCase             Type = "testCase"
	TypeTestDependency       Type = "testDependency"
	TypeTestTool             Type = "testTool"
	TypeVariant              Type = "variant"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeUNKNOWN, TypeAmends, TypeAncestor, TypeBuildDependency, TypeBuildTool, TypeContains, TypeContainedBy, TypeCopy, TypeDataFile, TypeDependencyManifest, TypeDependsOn, TypeDependencyOf, TypeDescendant, TypeDescribes, TypeDescribedBy, TypeDevDependency, TypeDevTool, TypeDistributionArtifact, TypeDocumentation, TypeDynamicLink, TypeExample, TypeExpandedFromArchive, TypeFileAdded, TypeFileDeleted, TypeFileModified, TypeGenerates, TypeGeneratedFrom, TypeMetafile, TypeOptionalComponent, TypeOptionalDependency, TypeOther, TypePackages, TypePatch, TypePrerequisite, TypePrerequisiteFor, TypeProvidedDependency, TypeRequirementFor, TypeRuntimeDependency, TypeSpecificationFor, TypeStaticLink, TypeTest, TypeTestCase, TypeTestDependency, TypeTestTool, TypeVariant:
		return nil
	default:
		return fmt.Errorf("edgetype: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the EdgeType queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByNodeID orders the results by the node_id field.
func ByNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNodeID, opts...).ToFunc()
}

// ByToNodeID orders the results by the to_node_id field.
func ByToNodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToNodeID, opts...).ToFunc()
}

// ByFromField orders the results by from field.
func ByFromField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromStep(), sql.OrderByField(field, opts...))
	}
}

// ByToField orders the results by to field.
func ByToField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToStep(), sql.OrderByField(field, opts...))
	}
}

// ByDocumentsCount orders the results by documents count.
func ByDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentsStep(), opts...)
	}
}

// ByDocuments orders the results by documents terms.
func ByDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodeListsCount orders the results by node_lists count.
func ByNodeListsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodeListsStep(), opts...)
	}
}

// ByNodeLists orders the results by node_lists terms.
func ByNodeLists(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeListsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFromStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, FromTable, FromColumn),
	)
}
func newToStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ToTable, ToColumn),
	)
}
func newDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DocumentsTable, DocumentsPrimaryKey...),
	)
}
func newNodeListsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeListsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, NodeListsTable, NodeListsPrimaryKey...),
	)
}
