// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"fmt"

	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/protobom/pkg/storage"

	"github.com/protobom/storage/internal/backends/ent"
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

type (
	contactOwnerIDKey      struct{}
	externalReferenceIDKey struct{}
	metadataIDKey          struct{}
	nodeIDKey              struct{}
	nodeListIDKey          struct{}
)

// Store implements the storage.Storer interface.
func (backend *Backend) Store(doc *sbom.Document, opts *storage.StoreOptions) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	if opts == nil {
		opts = &storage.StoreOptions{
			BackendOptions: NewBackendOptions(),
		}
	}

	if _, ok := opts.BackendOptions.(*BackendOptions); !ok {
		return fmt.Errorf("%w", errInvalidEntOptions)
	}

	tx, err := backend.client.Tx(backend.ctx)
	if err != nil {
		return fmt.Errorf("creating transactional client: %w", err)
	}

	backend.ctx = ent.NewTxContext(backend.ctx, tx)

	if err := tx.Document.Create().
		SetID(doc.Metadata.Id).
		OnConflict().
		Ignore().
		Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
		return rollback(tx, fmt.Errorf("ent.Document: %w", err))
	}

	if err := backend.saveMetadata(doc.Metadata); err != nil {
		return rollback(tx, err)
	}

	if err := backend.saveNodeList(doc.NodeList); err != nil {
		return rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}

	return nil
}

func (backend *Backend) saveDocumentTypes(docTypes []*sbom.DocumentType) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)

	for _, dt := range docTypes {
		typeName := documenttype.Type(dt.Type.String())

		newDocType := tx.DocumentType.Create().
			SetNillableType(&typeName).
			SetNillableName(dt.Name).
			SetNillableDescription(dt.Description)

		if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
			newDocType.SetMetadataID(metadataID)
		}

		err := newDocType.OnConflict().Ignore().Exec(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.DocumentType: %w", err)
		}
	}

	return nil
}

func (backend *Backend) saveEdges(edges []*sbom.Edge) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)

	for _, edge := range edges {
		for _, toID := range edge.To {
			newEdgeType := tx.EdgeType.Create().
				SetType(edgetype.Type(edge.Type.String())).
				SetFromID(edge.From).
				SetToID(toID)

			err := newEdgeType.OnConflict().Ignore().Exec(backend.ctx)
			if err != nil && !ent.IsConstraintError(err) {
				return fmt.Errorf("ent.Node: %w", err)
			}
		}
	}

	return nil
}

func (backend *Backend) saveExternalReferences(refs []*sbom.ExternalReference) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)

	for _, ref := range refs {
		newRef := tx.ExternalReference.Create().
			SetURL(ref.Url).
			SetComment(ref.Comment).
			SetAuthority(ref.Authority).
			SetType(externalreference.Type(ref.Type.String()))

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			newRef.SetNodeID(nodeID)
		}

		id, err := newRef.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, externalReferenceIDKey{}, id)

		if err := backend.saveHashesEntries(ref.Hashes); err != nil {
			return err
		}
	}

	return nil
}

func (backend *Backend) saveHashesEntries(hashes map[int32]string) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)
	entries := []*ent.HashesEntryCreate{}

	for alg, content := range hashes {
		algName := sbom.HashAlgorithm(alg).String()

		entry := tx.HashesEntry.Create().
			SetHashAlgorithmType(hashesentry.HashAlgorithmType(algName)).
			SetHashData(content)

		if externalReferenceID, ok := backend.ctx.Value(externalReferenceIDKey{}).(int); ok {
			entry.SetExternalReferenceID(externalReferenceID)
		}

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			entry.SetNodeID(nodeID)
		}

		entries = append(entries, entry)
	}

	if err := tx.HashesEntry.CreateBulk(entries...).
		Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.HashesEntry: %w", err)
	}

	return nil
}

func (backend *Backend) saveIdentifiersEntries(idents map[int32]string) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)
	entries := []*ent.IdentifiersEntryCreate{}

	for typ, value := range idents {
		typeName := sbom.SoftwareIdentifierType(typ).String()

		entry := tx.IdentifiersEntry.Create().
			SetSoftwareIdentifierType(identifiersentry.SoftwareIdentifierType(typeName)).
			SetSoftwareIdentifierValue(value)

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			entry.SetNodeID(nodeID)
		}

		entries = append(entries, entry)
	}

	if err := tx.IdentifiersEntry.CreateBulk(entries...).
		Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.IdentifiersEntry: %w", err)
	}

	return nil
}

func (backend *Backend) saveMetadata(md *sbom.Metadata) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)

	newMetadata := tx.Metadata.Create().
		SetID(md.Id).
		SetDocumentID(md.Id).
		SetVersion(md.Version).
		SetName(md.Name).
		SetComment(md.Comment).
		SetDate(md.Date.AsTime())

	err := newMetadata.OnConflict().Ignore().Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Metadata: %w", err)
	}

	backend.ctx = context.WithValue(backend.ctx, metadataIDKey{}, md.Id)

	if err := backend.savePersons(md.Authors); err != nil {
		return err
	}

	if err := backend.saveDocumentTypes(md.DocumentTypes); err != nil {
		return err
	}

	if err := backend.saveTools(md.Tools); err != nil {
		return err
	}

	return nil
}

func (backend *Backend) saveNodeList(nodeList *sbom.NodeList) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)
	newNodeList := tx.NodeList.Create().
		SetRootElements(nodeList.RootElements)

	if documentID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
		newNodeList.SetDocumentID(documentID)
	}

	id, err := newNodeList.OnConflict().Ignore().ID(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.NodeList: %w", err)
	}

	backend.ctx = context.WithValue(backend.ctx, nodeListIDKey{}, id)

	if err := backend.saveNodes(nodeList.Nodes); err != nil {
		return err
	}

	// Update nodes of this node list with their typed edges.
	if err := backend.saveEdges(nodeList.Edges); err != nil {
		return err
	}

	return nil
}

func (backend *Backend) saveNodes(nodes []*sbom.Node) error { //nolint:cyclop
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	for _, n := range nodes {
		newNode := backend.newNodeCreate(n)

		if err := newNode.OnConflict().Ignore().Exec(backend.ctx); err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.Node: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, nodeIDKey{}, n.Id)

		if err := backend.saveExternalReferences(n.ExternalReferences); err != nil {
			return err
		}

		if err := backend.savePersons(n.Originators); err != nil {
			return err
		}

		if err := backend.savePersons(n.Suppliers); err != nil {
			return err
		}

		if err := backend.savePurposes(n.PrimaryPurpose); err != nil {
			return err
		}

		if err := backend.saveHashesEntries(n.Hashes); err != nil {
			return err
		}

		if err := backend.saveIdentifiersEntries(n.Identifiers); err != nil {
			return err
		}
	}

	return nil
}

func (backend *Backend) savePersons(persons []*sbom.Person) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)

	for _, p := range persons {
		newPerson := tx.Person.Create().
			SetName(p.Name).
			SetEmail(p.Email).
			SetIsOrg(p.IsOrg).
			SetPhone(p.Phone).
			SetURL(p.Url)

		if contactOwnerID, ok := backend.ctx.Value(contactOwnerIDKey{}).(int); ok {
			newPerson.SetContactOwnerID(contactOwnerID)
		}

		if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
			newPerson.SetMetadataID(metadataID)
		}

		id, err := newPerson.OnConflict().Ignore().ID(backend.ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return fmt.Errorf("ent.ExternalReference: %w", err)
		}

		backend.ctx = context.WithValue(backend.ctx, contactOwnerIDKey{}, id)

		if err := backend.savePersons(p.Contacts); err != nil {
			return err
		}
	}

	return nil
}

func (backend *Backend) savePurposes(purposes []sbom.Purpose) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)
	builders := []*ent.PurposeCreate{}

	for idx := range purposes {
		newPurpose := tx.Purpose.Create().
			SetPrimaryPurpose(purpose.PrimaryPurpose(purposes[idx].String()))

		if nodeID, ok := backend.ctx.Value(nodeIDKey{}).(string); ok {
			newPurpose.SetNodeID(nodeID)
		}

		builders = append(builders, newPurpose)
	}

	err := tx.Purpose.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Tool: %w", err)
	}

	return nil
}

func (backend *Backend) saveTools(tools []*sbom.Tool) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx := ent.TxFromContext(backend.ctx)
	builders := []*ent.ToolCreate{}

	for _, t := range tools {
		newTool := tx.Tool.Create().
			SetName(t.Name).
			SetVersion(t.Version).
			SetVendor(t.Vendor)

		if metadataID, ok := backend.ctx.Value(metadataIDKey{}).(string); ok {
			newTool.SetMetadataID(metadataID)
		}

		builders = append(builders, newTool)
	}

	err := tx.Tool.CreateBulk(builders...).
		OnConflict().
		Ignore().
		Exec(backend.ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("ent.Tool: %w", err)
	}

	return nil
}

func (backend *Backend) newNodeCreate(n *sbom.Node) *ent.NodeCreate {
	tx := ent.TxFromContext(backend.ctx)

	newNode := tx.Node.Create().
		SetID(n.Id).
		SetAttribution(n.Attribution).
		SetBuildDate(n.BuildDate.AsTime()).
		SetComment(n.Comment).
		SetCopyright(n.Copyright).
		SetDescription(n.Description).
		SetFileName(n.FileName).
		SetFileTypes(n.FileTypes).
		SetLicenseComments(n.LicenseComments).
		SetLicenseConcluded(n.LicenseConcluded).
		SetLicenses(n.Licenses).
		SetName(n.Name).
		SetReleaseDate(n.ReleaseDate.AsTime()).
		SetSourceInfo(n.SourceInfo).
		SetSummary(n.Summary).
		SetType(node.Type(n.Type.String())).
		SetURLDownload(n.UrlDownload).
		SetURLHome(n.UrlHome).
		SetValidUntilDate(n.ValidUntilDate.AsTime()).
		SetVersion(n.Version)

	if nodeListID, ok := backend.ctx.Value(nodeListIDKey{}).(int); ok {
		newNode.AddNodeListIDs(nodeListID)
	}

	return newNode
}

func rollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		return fmt.Errorf("rolling back transaction: %w", rollbackErr)
	}

	return err
}
