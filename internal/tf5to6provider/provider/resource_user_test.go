package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("tf5to6provider_user.example", "name", regexp.MustCompile("^Exam")),
				),
			},
		},
	})
}

const configResourceBasic = `
resource "tf5to6provider_user" "example" {
  age   = 200
  email = "example@example.com"
  name  = "Example Name"
}
`
