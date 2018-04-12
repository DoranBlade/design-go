package simpleFactory

import "testing"

func TestFactory(t *testing.T) {
	runner := CreateRunner("dog")
	runner.run()
}

