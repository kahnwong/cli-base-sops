package cli_base_sops

import (
	cliBase "github.com/kahnwong/cli-base"
	"gopkg.in/yaml.v3"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/rs/zerolog/log"
)

func decryptSops(path string, format string) []byte {
	data, err := decrypt.File(path, format) // format: yaml, txt, etc. Refer to sops docs.
	if err != nil {
		log.Fatal().Msgf("Failed to decrypt sops config at: %s", path)
	}

	return data
}

func ReadYamlSops[T any](path string) *T {
	// check if config exists
	path, err := cliBase.CheckIfConfigExists(path)

	if err != nil {
		log.Fatal().Msgf("Config doesn't exist at: %s", path)
	}

	// decrypt sops
	data := decryptSops(path, "yaml")

	// unmarshall
	// ref: <https://stackoverflow.com/a/71955439>
	config := new(T)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal().Msgf("Error unmarshalling config: %s", path)
	}

	return config
}
