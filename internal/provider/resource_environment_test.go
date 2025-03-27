package provider

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var env_name = "test-acc-name"

func TestAccEnvironment_positive(t *testing.T) {
	engineId := os.Getenv("ACC_ENV_ENGINE_ID")
	username := os.Getenv("ACC_ENV_USERNAME")
	password := os.Getenv("ACC_ENV_PASSWORD")
	hostname := os.Getenv("ACC_ENV_HOSTNAME")
	toolkitPath := os.Getenv("ACC_ENV_TOOLKIT_PATH")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccEnvPreCheck(t, engineId, username, password, hostname, toolkitPath) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEnvDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctEnvConfigBasic(engineId, username, password, hostname, toolkitPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctEnvResourceExists("delphix_environment.new_env", engineId),
					resource.TestCheckResourceAttr("delphix_environment.new_env", "name", env_name)),
			},
		},
	})
}

func TestAccEnvironment_update_positive(t *testing.T) {
	engineId := os.Getenv("ACC_ENV_ENGINE_ID")
	username := os.Getenv("ACC_ENV_USERNAME")
	password := os.Getenv("ACC_ENV_PASSWORD")
	hostname := os.Getenv("ACC_ENV_HOSTNAME")
	toolkitPath := os.Getenv("ACC_ENV_TOOLKIT_PATH")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccEnvPreCheck(t, engineId, username, password, hostname, toolkitPath) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEnvDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctEnvConfigBasic(engineId, username, password, hostname, toolkitPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctEnvResourceExists("delphix_environment.new_env", engineId),
					resource.TestCheckResourceAttr("delphix_environment.new_env", "name", env_name)),
			},
			{
				// positive env update case
				Config: testAccEnvUpdatePositive(engineId, username, password, hostname, toolkitPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("delphix_environment.new_env", "name", "updated-name")),
			},
		},
	})
}

func TestAccEnvironment_update_negative(t *testing.T) {
	engineId := os.Getenv("ACC_ENV_ENGINE_ID")
	username := os.Getenv("ACC_ENV_USERNAME")
	password := os.Getenv("ACC_ENV_PASSWORD")
	hostname := os.Getenv("ACC_ENV_HOSTNAME")
	toolkitPath := os.Getenv("ACC_ENV_TOOLKIT_PATH")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccEnvPreCheck(t, engineId, username, password, hostname, toolkitPath) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEnvDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDctEnvConfigBasic(engineId, username, password, hostname, toolkitPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDctEnvResourceExists("delphix_environment.new_env", engineId),
					resource.TestCheckResourceAttr("delphix_environment.new_env", "name", env_name)),
			},
			{
				// negative update test case
				Config:      testAccEnvUpdateNegative(engineId, username, password, "updated-hostname", toolkitPath),
				ExpectError: regexp.MustCompile("Error running apply: exit status 1"),
			},
		},
	})
}

func testAccEnvPreCheck(t *testing.T, engineId string, username string, password string, hostname string, toolkitPath string) {
	testAccPreCheck(t)
	if engineId == "" {
		t.Fatal("ACC_ENV_ENGINE_ID must be set for env acceptance tests")
	}
	if username == "" {
		t.Fatal("ACC_ENV_USERNAME must be set for env acceptance tests")
	}
	if password == "" {
		t.Fatal("ACC_ENV_PASSWORD must be set for env acceptance tests")
	}
	if hostname == "" {
		t.Fatal("ACC_ENV_HOSTNAME must be set for env acceptance tests")
	}
	if toolkitPath == "" {
		t.Fatal("ACC_ENV_TOOLKIT_PATH must be set for env acceptance tests")
	}
}

func escape(s string) string {
	// Escape backslash or terraform interprets it as a special character
	return strings.ReplaceAll(s, "\\", "\\\\")
}

func testAccCheckDctEnvConfigBasic(engineId string, username string, password string, hostname string, toolkitPath string) string {
	return fmt.Sprintf(`
	resource "delphix_environment" "new_env" {
		engine_id = %s
		os_type = "UNIX"
		username = "%s"
		password = "%s"
		name = "%s"
		hosts {
			hostname = "%s"
			toolkit_path = "%s"
		}
	}
	`, engineId, escape(username), escape(password), env_name, escape(hostname), escape(toolkitPath))
}

func testAccCheckDctEnvResourceExists(n string, engineId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		EnvId := rs.Primary.ID
		if EnvId == "" {
			return fmt.Errorf("No EnvID set")
		}

		client := testAccProvider.Meta().(*apiClient).client
		res, _, err := client.EnvironmentsAPI.GetEnvironmentById(context.Background(), EnvId).Execute()
		if err != nil {
			return err
		}

		dctEngineId := res.GetEngineId()
		if dctEngineId != engineId {
			return fmt.Errorf("dctEngineId %s does not match provided engineID %s", dctEngineId, engineId)
		}

		return nil
	}
}

func testAccCheckEnvDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "delphix_environment" {
			continue
		}

		EnvId := rs.Primary.ID

		_, httpResp, _ := client.EnvironmentsAPI.GetEnvironmentById(context.Background(), EnvId).Execute()
		if httpResp == nil {
			return fmt.Errorf("Environment has not been deleted")
		}

		if httpResp.StatusCode != 404 {
			return fmt.Errorf("Exepcted a 404 Not Found for a deleted Environment but got %d", httpResp.StatusCode)
		}
	}

	return nil
}

func testAccEnvUpdatePositive(engineId string, username string, password string, hostname string, toolkitPath string) string {
	return fmt.Sprintf(`
	resource "delphix_environment" "new_env" {
		engine_id = %s
		os_type = "UNIX"
		username = "%s"
		password = "%s"
		name = "updated-name"
		hosts {
			hostname = "%s"
			toolkit_path = "%s"
		}
	}
	`, engineId, escape(username), escape(password), escape(hostname), escape(toolkitPath))
}

func testAccEnvUpdateNegative(engineId string, username string, password string, hostname string, toolkitPath string) string {
	return fmt.Sprintf(`
	resource "delphix_environment" "new_env" {
		engine_id = %s
		os_type = "UNIX"
		username = "%s"
		password = "%s"
		name = "%s"
		hosts {
			hostname = "%s"
			toolkit_path = "%s"
		}
	}
	`, engineId, escape(username), escape(password), env_name, escape(hostname), escape(toolkitPath))
}
