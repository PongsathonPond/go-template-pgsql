package implement_test

import (
	"testing"

	"idev-cms-service/service/authen/implement/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
