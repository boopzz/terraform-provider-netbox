package netbox

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/secrets"
)

func dataNetboxJSONSecretsSecretRolesList() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxJSONSecretsSecretRolesListRead,

		Schema: map[string]*schema.Schema{
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataNetboxJSONSecretsSecretRolesListRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBoxAPI)

	params := secrets.NewSecretsSecretRolesListParams()
	limit := int64(d.Get("limit").(int))
	params.Limit = &limit

	list, err := client.Secrets.SecretsSecretRolesList(params, nil)
	if err != nil {
		return err
	}

	j, _ := json.Marshal(list.Payload.Results)

	d.Set("json", string(j))
	d.SetId("NetboxJSONSecretsSecretRolesList")

	return nil
}