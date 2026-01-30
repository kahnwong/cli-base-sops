package cli_base_sops

import (
	"fmt"

	cliBase "github.com/kahnwong/cli-base"
	"gopkg.in/yaml.v3"

	"github.com/getsops/sops/v3/decrypt"
)

func decryptSops(path string, format string) ([]byte, error) {
	data, err := decrypt.File(path, format) // format: yaml, txt, etc. Refer to sops docs.
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt sops config at %s: %w", path, err)
	}

	return data, nil
}

func ReadYamlSops[T any](path string) (*T, error) {
	// check if config exists
	path, err := cliBase.CheckIfConfigExists(path)

	if err != nil {
		return nil, fmt.Errorf("config doesn't exist at %s: %w", path, err)
	}

	// decrypt sops
	data, err := decryptSops(path, "yaml")
	if err != nil {
		return nil, err
	}

	// unmarshall
	// ref: <https://stackoverflow.com/a/71955439>
	config := new(T)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config at %s: %w", path, err)
	}

	return config, nil
}
