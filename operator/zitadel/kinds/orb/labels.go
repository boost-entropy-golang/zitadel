package orb

import "github.com/caos/orbos/pkg/labels"

func mustZITADELOperator(binaryVersion *string) *labels.Operator {

	version := "unknown"
	if binaryVersion != nil {
		version = *binaryVersion
	}

	return labels.MustForOperator("ZITADEL", "zitadel.caos.ch", version)
}
