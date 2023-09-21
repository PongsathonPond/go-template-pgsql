//go:build integration
// +build integration

package redis_test

import (
	"testing"

	"idev-cms-service/repository/redis/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
