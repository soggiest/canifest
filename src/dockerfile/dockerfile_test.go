package dockerfile

import (
	"testing"
)

func TestExists(t *testing.T) {
	if exists("/etc") == false {
		// t.Errorf("File exists, but function returned false...")
		t.Fatalf("File exists, but function returned false...")
	}	
}
