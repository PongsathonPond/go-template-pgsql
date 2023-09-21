package authen_test

import (
	"testing"

	"idev-cms-service/app/authen/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
