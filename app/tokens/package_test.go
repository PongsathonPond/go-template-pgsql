package tokens_test

import (
	"testing"

	"idev-cms-service/app/tokens/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
