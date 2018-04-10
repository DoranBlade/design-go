package factory

import "testing"

func TestFactory(t *testing.T) {
	runner := CreateRunner("dog")
	runner.run()
}

