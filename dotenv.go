// Package dotenv provides methods for loading env files.
package dotenv

import (
	"os"
)

// Load takes in any number of env files. Defaults to .env
// It will parse the env files and load them into your environment.
// This will not overwrite existing environment variables.
func Load(filenames ...string) error {
	environmentVariables := make(map[string]string)
	var err error

	// Defaults to .env file
	if len(filenames) == 0 {
		environmentVariables, err = parsefile(".env")

		if err != nil {
			return err
		}
	}

	for _, filename := range filenames {
		environmentVariables, err = parsefile(filename)

		if err != nil {
			return err
		}
	}

	SetEnvironmentVars(environmentVariables, false)

	return nil
}

// LoadWithOverride takes in any number of env files. Defaults to .env
// It will parse the env files and load them into your environment.
// This will overwrite existing environment variables.
func LoadWithOverride(filenames ...string) error {
	environmentVariables := make(map[string]string)
	var err error

	// Defaults to .env file
	if len(filenames) == 0 {
		environmentVariables, err = parsefile(".env")

		if err != nil {
			return err
		}
	}

	for _, filename := range filenames {
		environmentVariables, err = parsefile(filename)

		if err != nil {
			return err
		}
	}

	SetEnvironmentVars(environmentVariables, true)

	return nil
}

func SetEnvironmentVars(envVars map[string]string, overload bool) {
	for key, value := range envVars {
		_, exists := os.LookupEnv(key)

		if overload {
			os.Setenv(key, value)
		} else if !exists {
			os.Setenv(key, value)
		}
	}
}

// Read takes in any number of env files. Defaults to .env
// It will parse the env files and return them in a map[string]string.
func Read(filenames ...string) (map[string]string, error) {
    environmentVariables := make(map[string]string)
	var err error

	// Defaults to .env file
	if len(filenames) == 0 {
		environmentVariables, err = parsefile(".env")

		if err != nil {
			return nil, err
		}
	}

	for _, filename := range filenames {
		environmentVariables, err = parsefile(filename)

		if err != nil {
			return nil, err
		}
	}

	return environmentVariables, nil
}
