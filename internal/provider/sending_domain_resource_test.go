package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccSendingDomainResourceConfig("mail.myhotel.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mandrill_sending_domain.test", "domain_name", "mail.myhotel.com"),
					resource.TestCheckResourceAttr("mandrill_sending_domain.test", "id", "mail.myhotel.com"),
					resource.TestCheckResourceAttr("mandrill_sending_domain.test", "verify_txt_record", "mandrill_verify.LL4219A0xGM7dKAsgFRa4w"),
					resource.TestCheckResourceAttr("mandrill_sending_domain.test", "spf_valid", "false"),
					resource.TestCheckResourceAttr("mandrill_sending_domain.test", "dkim_valid", "false"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "mandrill_sending_domain.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				// ImportStateVerifyIgnore: []string{"domain_name"},
			},
			// Update and Read testing
			// {
			// 	Config: testAccSendingDomainResourceConfig("two"),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("mandrill_sending_domain.test", "configurable_attribute", "two"),
			// 	),
			// },
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccSendingDomainResourceConfig(domain_name string) string {
	return fmt.Sprintf(`
resource "mandrill_sending_domain" "test" {
  domain_name = %[1]q
}
`, domain_name)
}
