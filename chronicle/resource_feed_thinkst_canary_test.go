package chronicle

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccChronicleFeedThinkstCanary_Basic(t *testing.T) {
	displayName := "test" + randString(40)
	enabled := "true"
	namespace := "test"
	labels := `"test"="test"`
	hostname := randString(5) + ".canary.tools"
	key := "auth_token"
	value := randString(10)

	rootRef := feedThinkstCanaryRef("test")
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckChronicleFeedThinkstCanaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", enabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					resource.TestCheckResourceAttr(rootRef, "details.0.hostname", hostname),
				),
			},
			{
				ResourceName:      rootRef,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"display_name", "state", "details.0.authentication.0.key",
					"details.0.authentication.0.value"},
			},
		},
	})
}

func TestAccChronicleFeedThinkstCanary_UpdateAuth(t *testing.T) {
	displayName := "test" + randString(40)
	enabled := "true"
	namespace := "test"
	labels := `"test"="test"`
	hostname := randString(5) + ".canary.tools"
	key := "auth_token"
	value := randString(10)
	value1 := randString(10)

	rootRef := feedThinkstCanaryRef("test")
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckChronicleFeedThinkstCanaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", enabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					resource.TestCheckResourceAttr(rootRef, "details.0.hostname", hostname),
				),
			},
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname, key, value1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", enabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					testAccCheckChronicleFeedThinkstCanaryAuthUpdated(t, rootRef, key, value1),
				),
			},
			{
				ResourceName:      rootRef,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"display_name", "state", "details.0.authentication.0.key",
					"details.0.authentication.0.value"},
			},
		},
	})
}

func TestAccChronicleFeedThinkstCanary_UpdateEnabled(t *testing.T) {
	displayName := "test" + randString(40)
	enabled := "true"
	notEnabled := "false"
	namespace := "test"
	labels := `"test"="test"`
	hostname := randString(5) + ".canary.tools"
	key := "auth_token"
	value := randString(10)

	rootRef := feedThinkstCanaryRef("test")
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckChronicleFeedThinkstCanaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", enabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					resource.TestCheckResourceAttr(rootRef, "details.0.hostname", hostname),
				),
			},
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, notEnabled, namespace, labels, hostname, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", notEnabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					resource.TestCheckResourceAttr(rootRef, "details.0.hostname", hostname),
				),
			},
			{
				ResourceName:      rootRef,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"display_name", "state", "details.0.authentication.0.key",
					"details.0.authentication.0.value"},
			},
		},
	})
}

func TestAccChronicleThinkstCanary_UpdateHostname(t *testing.T) {
	displayName := "test" + randString(40)
	enabled := "true"
	namespace := "test"
	labels := `"test"="test"`
	hostname := randString(5) + ".canary.tools"
	hostname1 := randString(5) + ".canary.tools"
	key := "auth_token"
	value := randString(10)

	rootRef := feedThinkstCanaryRef("test")
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckChronicleFeedThinkstCanaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", enabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					resource.TestCheckResourceAttr(rootRef, "details.0.hostname", hostname),
				),
			},
			{
				Config: testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname1, key, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChronicleFeedThinkstCanaryExists(rootRef),
					resource.TestCheckResourceAttr(rootRef, "enabled", enabled),
					resource.TestCheckResourceAttr(rootRef, "namespace", namespace),
					resource.TestCheckResourceAttr(rootRef, "details.0.hostname", hostname1),
				),
			},
			{
				ResourceName:      rootRef,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"display_name", "state", "details.0.authentication.0.key",
					"details.0.authentication.0.value"},
			},
		},
	})
}

//nolint:unparam
func testAccCheckChronicleFeedThinkstCanaryAuthUpdated(t *testing.T, n, key, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		if rs.Primary.Attributes["details.0.authentication.0.key"] != key ||
			rs.Primary.Attributes["details.0.authentication.0.value"] != value {
			return fmt.Errorf("key or value differs")
		}

		return nil
	}
}

//nolint:unparam
func testAccCheckChronicleFeedThinkstCanary(displayName, enabled, namespace, labels, hostname, key, value string) string {
	return fmt.Sprintf(
		`resource "chronicle_feed_thinkst_canary" "test" {
			display_name = "%s"
			enabled = %s
			namespace = "%s"
			labels = {
				%s
			}
			details {
				hostname = "%s"
				authentication {
					key = "%s"
					value = "%s"
				}
			}
			}`, displayName, enabled, namespace, labels, hostname, key, value)
}

func testAccCheckChronicleFeedThinkstCanaryExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("ID for %s in state", n)
		}
		return nil
	}
}

func testAccCheckChronicleFeedThinkstCanaryDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "chronicle_feed_thinkst_canary.test" {
			continue
		}

		if rs.Primary.ID != "" {
			return fmt.Errorf("Object %q still exists", rs.Primary.ID)
		}
		return nil
	}
	return nil
}

//nolint:unparam
func feedThinkstCanaryRef(name string) string {
	return fmt.Sprintf("chronicle_feed_thinkst_canary.%v", name)
}
