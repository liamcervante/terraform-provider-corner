package provider1

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser1(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configResourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("tf5muxprovider_user1.example", "name", regexp.MustCompile("^Exam")),
				),
			},
		},
	})
}

const configResourceBasic = `
resource "tf5muxprovider_user1" "example" {
  age   = 200
  email = "example@example.com"
  name  = "Example Name"
}
`
