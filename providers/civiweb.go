package providers

import (
	"errors"
)

// Civiweb has migrated to mon-vie-via.businessfrance.fr
// The new site requires authentication, so this provider is deprecated.

type civiweb struct {
	latest bool
}

// NewCiviweb returns a deprecated provider
// The site has migrated to mon-vie-via.businessfrance.fr and requires auth
func NewCiviweb() Provider {
	return &civiweb{false}
}

type civiwebLatest struct {
	civiweb
}

// NewCiviwebLatest returns a deprecated provider
func NewCiviwebLatest() Provider {
	return &civiwebLatest{civiweb{true}}
}

func (c *civiweb) RetrieveJobs(fn func(job *Job)) error {
	return errors.New("civiweb provider is deprecated: site migrated to mon-vie-via.businessfrance.fr and requires authentication")
}
