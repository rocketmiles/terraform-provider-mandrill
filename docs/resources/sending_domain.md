---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mandrill_sending_domain Resource - terraform-provider-mandrill"
subcategory: ""
description: |-
  Sending domain
---

# mandrill_sending_domain (Resource)

Sending domain



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **domain_name** (String) Sending domain name

### Read-Only

- **dkim_valid** (Boolean) Set to true if DKIM DNS configuration is valid
- **id** (String) Sending domain name as ID
- **spf_valid** (Boolean) Set to true if SPF DNS configuration is valid
- **verify_txt_record** (String) The full verify TXT record value to be set on domain zone

