package stratos

import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "time"

  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerSupportGroups() *schema.Resource {
  return &schema.Resource{
    ReadContext: dataSourceServerSupportGroups,
    Schema: map[string]*schema.Schema{},
  }
}

Schema: map[string]*schema.Schema{
  "serverSupportGroups": &schema.Schema{
    Type:     schema.TypeList,
    Computed: true,
    Elem: &schema.Resource{
      Schema: map[string]*schema.Schema{
        "id": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "name": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "sys_id": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "ad_group": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "email": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "chef_org": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "splunk_server_class": &schema.Schema{
          Type:     schema.TypeString,
          Computed: true,
        },
        "approval_group": &schema.Schema{
          Type:     schema.TypeList,
          Computed: true,
          Elem: &schema.Resource{
            Schema: map[string]*schema.Schema{
              "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
            Schema: map[string]*schema.Schema{
              "sys_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
          },
        },
        "vsphere_folder": &schema.Schema{
          Type:     schema.TypeList,
          Computed: true,
          Elem: &schema.Resource{
            Schema: map[string]*schema.Schema{
              "aoc": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
            Schema: map[string]*schema.Schema{
              "bnt": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
          },
        },
        "ou": &schema.Schema{
          Type:     schema.TypeList,
          Computed: true,
          Elem: &schema.Resource{
            Schema: map[string]*schema.Schema{
              "acf": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
            Schema: map[string]*schema.Schema{
              "corp": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
            Schema: map[string]*schema.Schema{
              "io": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
            Schema: map[string]*schema.Schema{
              "direct": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
              },
            },
          },
        },
      },
    },
  },
},

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  c := m.(*hc.Client)

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  supportGroupID := Get("id")

  supportGroup, err := c.GetOrder(supportGroupID)
  if err != nil {
    return diag.FromErr(err)
  }

  supportGroupOUs := flattenOUItemsData(&supportGroup.approval_group)
  if err := d.Set("ou", supportGroupOUs); err != nil {
    return diag.FromErr(err)
  }

  supportGroupAGs := flattenApprovalGroupData(&supportGroup.approval_group)
  if err := d.Set("ou", supportGroupOUs); err != nil {
    return diag.FromErr(err)
  }

  supportGroupVmFolders := flattenVsphereFolderData(&supportGroup.vsphere_folder)
  if err := d.Set("ou", supportGroupVMFolders); err != nil {
    return diag.FromErr(err)
  }

  d.SetId(supportGroupID)

  return diags
}

func flattenOUItemsData(supportGroupOUs *[]hc.ouItem) []interface{} {
  if supportGroupOUs != nil {
    ois := make([]interface{}, len(*supportGroupOUs), len(*supportGroupOUs))

    for i, ouItem := range *supportGroupOUs {
      ou := make(map[string]interface{})

      ou["acf"] = ouItem.SupportGroup.acf
      ou["corp"] = ouItem.SupportGroup.corp
      ou["io"] = ouItem.SupportGroup.io
      ou["direct"] = ouItem.SupportGroup.direct

      ous[i] = ou
    }

    return ous
  }

  return make([]interface{}, 0)
}

func flattenApprovalGroupData(supportGroupAGs *[]hc.approvalGroupItem) []interface{} {
  if supportGroupAGs != nil {
    ois := make([]interface{}, len(*supportGroupAGs), len(*supportGroupAGs))

    for i, approvalGroupItem := range *supportGroupAGs {
      ag := make(map[string]interface{})

      ag["name"] = approvalGroupItem.ApprovalGroup.name
      ag["sys_id"] = approvalGroupItem.ApprovalGroup.sys_id

      ags[i] = ag
    }

    return ags
  }

  return make([]interface{}, 0)
}

func flattenVsphereFolderData(supportGroupVmFolders *[]hc.vsphereFolderItem) []interface{} {
  if supportGroupVmFolders != nil {
    ois := make([]interface{}, len(*supportGroupVmFolders), len(*supportGroupVmFolders))

    for i, vsphereFolderItem := range *supportGroupVmFolders {
      vmf := make(map[string]interface{})

      vmf["aoc"] = vsphereFolderItem.VsphereFolder.aoc
      vmf["bnt"] = vsphereFolderItem.VsphereFolder.bnt

      vmfs[i] = vmf
    }

    return vmfs
  }

  return make([]interface{}, 0)
}