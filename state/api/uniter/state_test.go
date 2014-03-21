// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter_test

import (
	apitesting "launchpad.net/juju-core/state/api/testing"

	gc "launchpad.net/gocheck"
)

type stateSuite struct {
	uniterSuite
	*apitesting.APIAddresserTests
	*apitesting.EnvironWatcherTests
}

var _ = gc.Suite(&stateSuite{})

func (s *stateSuite) SetUpTest(c *gc.C) {
	s.uniterSuite.SetUpTest(c)
	s.APIAddresserTests = apitesting.NewAPIAddresserTests(s.BackingState, s.uniter)
	s.EnvironWatcherTests = apitesting.NewEnvironWatcherTests(s.uniter, s.BackingState, apitesting.NoSecrets)
}

func (s *stateSuite) TestProviderType(c *gc.C) {
	cfg, err := s.State.EnvironConfig()
	c.Assert(err, gc.IsNil)

	providerType, err := s.uniter.ProviderType()
	c.Assert(err, gc.IsNil)
	c.Assert(providerType, gc.DeepEquals, cfg.Type())
}
