package database

import (
	"github.com/caos/orbos/pkg/kubernetes"
	"github.com/caos/orbos/pkg/tree"
	"github.com/caos/zitadel/operator/api/core"
	databasev1 "github.com/caos/zitadel/operator/api/database/v1"
	macherrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
)

const (
	Namespace  = "caos-system"
	kind       = "Database"
	apiVersion = "caos.ch/v1"
	Name       = "database"
)

func ReadCrd(k8sClient kubernetes.ClientInt) (*tree.Tree, error) {
	unstruct, err := k8sClient.GetNamespacedCRDResource(databasev1.GroupVersion.Group, databasev1.GroupVersion.Version, kind, Namespace, Name)
	if err != nil {
		if macherrs.IsNotFound(err) || meta.IsNoMatchError(err) {
			return nil, nil
		}
		return nil, err
	}

	return core.UnmarshalUnstructuredSpec(unstruct)
}

func WriteCrd(k8sClient kubernetes.ClientInt, t *tree.Tree) error {

	unstruct, err := core.MarshalToUnstructuredSpec(t)
	if err != nil {
		return err
	}

	unstruct.SetName(Name)
	unstruct.SetNamespace(Namespace)
	unstruct.SetKind(kind)
	unstruct.SetAPIVersion(apiVersion)

	return k8sClient.ApplyNamespacedCRDResource(databasev1.GroupVersion.Group, databasev1.GroupVersion.Version, kind, Namespace, Name, unstruct)
}
