package instance

import (
	"testing"

	"github.com/scaleway/scaleway-cli/internal/core"
)

func Test_IpCreate(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		Cmd:      "scw instance ip create",
		Check:    core.TestCheckGolden(),
	}))
}

func Test_IpDelete(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		Cmd:      "scw instance ip delete",
		Check:    core.TestCheckGolden(),
	}))
}

func Test_IpGet(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		Cmd:      "scw instance ip get",
		Check:    core.TestCheckGolden(),
	}))
}

func Test_IpList(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		Cmd:      "scw instance ip list",
		Check:    core.TestCheckGolden(),
	}))
}

func Test_IpUpdate(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		Cmd:      "scw instance ip update",
		Check:    core.TestCheckGolden(),
	}))
}