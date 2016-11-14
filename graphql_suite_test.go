package graphql_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGraphql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphql Suite")
}
