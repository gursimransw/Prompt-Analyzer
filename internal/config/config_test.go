package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestYamlConfigLoader(t *testing.T) {
	tempDir := t.TempDir()

	configPath := filepath.Join(tempDir, "config.yaml")

	configContent := `
env: dev
detection_rule_library: configs/rules/rules.json
policy_config: configs/policy/policy.json
http_server:
  address: localhost:8000
`

	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test config file: %v", err)
	}

	t.Setenv("CONFIG_PATH", configPath)

	configuration := MustLoad()

	if configuration.HTTPServer.Addr != "localhost:8000" {
		t.Fatalf("expected server address localhost:8000, got %s", configuration.HTTPServer.Addr)
	}

	if configuration.Env != "dev" {
		t.Fatalf("expected env dev, got %s", configuration.Env)
	}

	if configuration.DetectionRuleLibrary != "configs/rules/rules.json" {
		t.Fatalf("expected detection rule library path configs/rules/rules.json, got %s", configuration.DetectionRuleLibrary)
	}

	if configuration.PolicyConfig != "configs/policy/policy.json" {
		t.Fatalf("expected policy config path configs/policy/policy.json, got %s", configuration.PolicyConfig)
	}
}
