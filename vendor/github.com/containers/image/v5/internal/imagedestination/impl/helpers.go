package impl

import (
	"github.com/containers/image/v5/internal/manifest"
	"github.com/containers/image/v5/internal/private"
)

// OriginalCandidateMatchesTryReusingBlobOptions returns true if the original blob passed to TryReusingBlobWithOptions
// is acceptable based on opts.
func OriginalCandidateMatchesTryReusingBlobOptions(opts private.TryReusingBlobOptions) bool {
	return manifest.CandidateCompressionMatchesReuseConditions(manifest.ReuseConditions{
		PossibleManifestFormats: opts.PossibleManifestFormats,
		RequiredCompression:     opts.RequiredCompression,
	}, opts.OriginalCompression)
}
