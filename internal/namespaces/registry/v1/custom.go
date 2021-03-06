package registry

import (
	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-cli/internal/human"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
)

// GetCommands returns registry commands.
//
// This function:
// - Gets the generated commands
// - Register handwritten marshalers
// - Apply handwritten overrides (of Command.Run)
func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()
	cmds.Merge(core.NewCommands(
		registryLoginCommand(),
		registryLogoutCommand(),
		registryDockerHelperGetCommand(),
		registryDockerHelperEraseCommand(),
		registryDockerHelperListCommand(),
		registryDockerHelperStoreCommand(),
		registryInstallDockerHelperCommand(),
	))

	cmds.MustFind("registry", "tag", "get").Override(tagGetBuilder)
	cmds.MustFind("registry", "tag", "list").Override(tagListBuilder)

	cmds.MustFind("registry", "image", "get").Override(imageGetBuilder)
	cmds.MustFind("registry", "image", "list").Override(imageListBuilder)

	human.RegisterMarshalerFunc(registry.NamespaceStatus(0), human.EnumMarshalFunc(namespaceStatusMarshalSpecs))
	human.RegisterMarshalerFunc(registry.ImageStatus(0), human.EnumMarshalFunc(imageStatusMarshalSpecs))
	human.RegisterMarshalerFunc(registry.TagStatus(0), human.EnumMarshalFunc(tagStatusMarshalSpecs))

	return cmds
}
