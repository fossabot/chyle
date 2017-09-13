package builder

import (
	"fmt"
	"github.com/antham/strumt"
)

// NewEnvPrompts creates several prompts at once to populate environments variables
func NewEnvPrompts(configs []EnvConfig, store *Store) []strumt.Prompter {
	results := []strumt.Prompter{}

	for _, config := range configs {
		results = append(results, NewEnvPrompt(config, store))
	}

	return results
}

// NewEnvPrompt creates a prompt to populate an environment variable
func NewEnvPrompt(config EnvConfig, store *Store) strumt.Prompter {
	return &template{
		config.ID,
		config.PromptString,
		func(string) string { return config.NextID },
		func(error) string { return config.ID },
		ParseEnv(config.Env, store),
	}
}

// ParseEnv provides an env parser callback
func ParseEnv(env string, store *Store) func(value string) error {
	return func(value string) error {
		if value == "" {
			return fmt.Errorf("No value given")
		}

		(*store)[env] = value

		return nil
	}
}
