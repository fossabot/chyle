package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antham/envh"
)

func getCustomAPIMandatoryParamsRefs() []struct {
	ref      *string
	keyChain []string
} {
	return []struct {
		ref      *string
		keyChain []string
	}{
		{
			&chyleConfig.DECORATORS.CUSTOMAPI.ENDPOINT.URL,
			[]string{"CHYLE", "DECORATORS", "CUSTOMAPI", "ENDPOINT", "URL"},
		},
		{
			&chyleConfig.DECORATORS.CUSTOMAPI.CREDENTIALS.TOKEN,
			[]string{"CHYLE", "DECORATORS", "CUSTOMAPI", "CREDENTIALS", "TOKEN"},
		},
	}
}

func getCustomAPIFeatureRefs() []*bool {
	return []*bool{
		&chyleConfig.FEATURES.DECORATORS.ENABLED,
		&chyleConfig.FEATURES.DECORATORS.CUSTOMAPI,
	}
}

func getCustomAPICustomValidationFuncs(config *envh.EnvTree) []func() error {
	return []func() error{
		func() error {
			keyChain := []string{"CHYLE", "DECORATORS", "CUSTOMAPI", "ENDPOINT", "URL"}
			URL := config.FindStringUnsecured(keyChain...)

			if !regexp.MustCompile(`{{\s*ID\s*}}`).MatchString(URL) {
				return fmt.Errorf(`ensure you defined a placeholder {{ID}} in URL defined in "%s"`, strings.Join(keyChain, "_"))
			}

			return nil
		},
	}
}

func getCustomAPICustomSettersFuncs() []func(*CHYLE) {
	return []func(*CHYLE){}
}

// customAPIDecoratorConfigurator creates a custom api configurater from apiDecoratorConfigurator
func customAPIDecoratorConfigurator(config *envh.EnvTree) configurater {
	return &apiDecoratorConfigurator{
		config: config,
		apiDecoratorConfig: apiDecoratorConfig{
			"CUSTOMAPIID",
			"CUSTOMAPI",
			&chyleConfig.DECORATORS.CUSTOMAPI.KEYS,
			getCustomAPIMandatoryParamsRefs(),
			getCustomAPIFeatureRefs(),
			getCustomAPICustomValidationFuncs(config),
			getCustomAPICustomSettersFuncs(),
		},
	}
}
