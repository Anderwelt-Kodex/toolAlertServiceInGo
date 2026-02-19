package toolservicealert

import (
	"testing"
)

func TestCallingToolServiceAlert(t *testing.T) {

	t.Run("should produce a warning", func(t *testing.T) {

		val, _ := CallingToolServiceAlert("tool001", 3)

		if val != 1 {
			t.Errorf("something wrong (1)")
		}

	})
	t.Run("should produce no warning", func(t *testing.T) {

		val, _ := CallingToolServiceAlert("tool002", 7)

		if val != 0 {
			t.Errorf("something wrong (2)")
		}
	})

	t.Run("should produce a critical warning", func(t *testing.T) {

		val, _ := CallingToolServiceAlert("tool003", 5)

		if val != 2 {
			t.Errorf("something wrong (3)")
		}

	})

}

func TestCalculateToolRuntime(t *testing.T) {

	t.Run("should produce a number", func(t *testing.T) {
		val, _ := CalculateToolRuntime(7, 3)

		if val != 10 {
			t.Errorf("something wrong (10)")
		}
	})
	t.Run("should produce a number", func(t *testing.T) {
		val, _ := CalculateToolRuntime(1, 22)

		if val != 23 {
			t.Errorf("something wrong (23)")
		}
	})
	t.Run("should produce a number", func(t *testing.T) {
		val, _ := CalculateToolRuntime(5, 1)

		if val != 6 {
			t.Errorf("something wrong (6)")
		}
	})
}

func TestGetToolInformation(t *testing.T) {

	t.Run("should get struct toolInformation for ToolId", func(t *testing.T) {

		val, _ := GetToolInformation("tool003")

		if val.toolId != "tool003" {
			t.Errorf("toolId not ok")
		}
		if val.toolCumulativeRuntime != 1 {
			t.Errorf("cumRuntime not ok")
		}
		if val.maxServiceThreshold != 5 {
			t.Errorf("maxServiceThreshold not ok")
		}

	})
}


