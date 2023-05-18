---
page_title: "genesyscloud_externalcontacts_contact Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud External Contact
---
# genesyscloud_externalcontacts_contact (Resource)

Genesys Cloud External Contact

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

* [GET /api/v2/externalcontacts/contacts](https://developer.genesys.cloud/commdigital/externalcontacts/externalcontacts-apis#get-api-v2-externalcontacts-contacts)
* [POST /api/v2/externalcontacts/contacts](https://developer.genesys.cloud/commdigital/externalcontacts/externalcontacts-apis#post-api-v2-externalcontacts-contacts)
* [GET /api/v2/externalcontacts/contacts/{contactId}](https://developer.genesys.cloud/commdigital/externalcontacts/externalcontacts-apis#get-api-v2-externalcontacts-contacts--contactId-)
* [PUT /api/v2/externalcontacts/contacts/{contactId}](https://developer.genesys.cloud/commdigital/externalcontacts/externalcontacts-apis#put-api-v2-externalcontacts-contacts--contactId-)
* [DELETE /api/v2/externalcontacts/contacts/{contactId}](https://developer.genesys.cloud/commdigital/externalcontacts/externalcontacts-apis#delete-api-v2-externalcontacts-contacts--contactId-)

## Example Usage

```terraform
resource "genesyscloud_externalcontacts_contact" "contact" {
  first_name  = "jean"
  middle_name = "jacques"
  last_name   = "dupont"
  salutation  = "salutation"
  title       = "genesys staff"
  work_phone {
    display      = "+33 0 00 00 00 00"
    extension    = 4
    accepts_sms  = false
    e164         = "+330000000000"
    country_code = "FR"
  }
  cell_phone {
    display      = "+33 0 00 00 00 01"
    extension    = 4
    accepts_sms  = false
    e164         = "+330000000001"
    country_code = "FR"
  }
  home_phone {
    display      = "+33 0 00 00 00 02"
    extension    = 4
    accepts_sms  = false
    e164         = "+330000000002"
    country_code = "FR"
  }
  other_phone {
    display      = "+33 0 00 00 00 03"
    extension    = 4
    accepts_sms  = false
    e164         = "+330000000003"
    country_code = "FR"
  }
  work_email     = "workEmail@example.com"
  personal_email = "personalEmail@example.com"
  other_email    = "otherEmail@example.com"
  address {
    address1     = "1 rue de la paix"
    address2     = "2 rue de la paix"
    city         = "Paris"
    state        = "île-de-France"
    postal_code  = "75000"
    country_code = "FR"
  }
  twitter_id {
    id          = "1725137533"
    name        = "@KMbappe"
    screen_name = "KMbappe"
  }
  line_id {
    ids {
      user_id = "1234"
    }
    display_name = "lineName"
  }
  whatsapp_id {
    phone_number {
      display      = "+33 0 00 00 00 01"
      extension    = 4
      accepts_sms  = false
      e164         = "+330000000001"
      country_code = "FR"
    }
    display_name = "whatsappName"
  }
  facebook_id {
    ids {
      scoped_id = "myscopeIdUrl"
    }
    display_name = "facebookName"
  }
  survey_opt_out      = false
  external_system_url = "https://systemUrl.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address` (Block List, Max: 1) Contact address. (see [below for nested schema](#nestedblock--address))
- `cell_phone` (Block List, Max: 1) Contact call phone settings. (see [below for nested schema](#nestedblock--cell_phone))
- `external_system_url` (String) Contact external system url.
- `facebook_id` (Block List, Max: 1) Contact facebook account informations. (see [below for nested schema](#nestedblock--facebook_id))
- `first_name` (String) The first name of the contact.
- `home_phone` (Block List, Max: 1) Contact home phone settings. (see [below for nested schema](#nestedblock--home_phone))
- `last_name` (String) The last name of the contact.
- `line_id` (Block List, Max: 1) Contact line account informations. (see [below for nested schema](#nestedblock--line_id))
- `middle_name` (String) The middle name of the contact.
- `other_email` (String) Contact other email.
- `other_phone` (Block List, Max: 1) Contact other phone settings. (see [below for nested schema](#nestedblock--other_phone))
- `personal_email` (String) Contact personal email.
- `salutation` (String) The salutation of the contact.
- `survey_opt_out` (Boolean) Contact survey opt out preference.
- `title` (String) The title of the contact.
- `twitter_id` (Block List, Max: 1) Contact twitter account informations. (see [below for nested schema](#nestedblock--twitter_id))
- `whatsapp_id` (Block List, Max: 1) Contact whatsapp account informations. (see [below for nested schema](#nestedblock--whatsapp_id))
- `work_email` (String) Contact work email.
- `work_phone` (Block List, Max: 1) Contact work phone settings. (see [below for nested schema](#nestedblock--work_phone))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--address"></a>
### Nested Schema for `address`

Optional:

- `address1` (String) Contact address 1.
- `address2` (String) Contact address 2.
- `city` (String) Contact address city.
- `country_code` (String) Contact address country code.
- `postal_code` (String) Contact address postal code.
- `state` (String) Contact address state.


<a id="nestedblock--cell_phone"></a>
### Nested Schema for `cell_phone`

Optional:

- `accepts_sms` (Boolean) If contact accept SMS.
- `country_code` (String) Phone number country code.
- `display` (String) Display string of the phone number.
- `e164` (String) Phone number in e164 format.
- `extension` (Number) Phone extension.


<a id="nestedblock--facebook_id"></a>
### Nested Schema for `facebook_id`

Optional:

- `display_name` (String) Contact whatsapp display name.
- `ids` (Block List) Contact facebook scoped id. (see [below for nested schema](#nestedblock--facebook_id--ids))

<a id="nestedblock--facebook_id--ids"></a>
### Nested Schema for `facebook_id.ids`

Optional:

- `scoped_id` (String) Contact facebook scoped id.



<a id="nestedblock--home_phone"></a>
### Nested Schema for `home_phone`

Optional:

- `accepts_sms` (Boolean) If contact accept SMS.
- `country_code` (String) Phone number country code.
- `display` (String) Display string of the phone number.
- `e164` (String) Phone number in e164 format.
- `extension` (Number) Phone extension.


<a id="nestedblock--line_id"></a>
### Nested Schema for `line_id`

Optional:

- `display_name` (String) Contact line display name.
- `ids` (Block List) Contact line id. (see [below for nested schema](#nestedblock--line_id--ids))

<a id="nestedblock--line_id--ids"></a>
### Nested Schema for `line_id.ids`

Optional:

- `user_id` (String) Contact line id.



<a id="nestedblock--other_phone"></a>
### Nested Schema for `other_phone`

Optional:

- `accepts_sms` (Boolean) If contact accept SMS.
- `country_code` (String) Phone number country code.
- `display` (String) Display string of the phone number.
- `e164` (String) Phone number in e164 format.
- `extension` (Number) Phone extension.


<a id="nestedblock--twitter_id"></a>
### Nested Schema for `twitter_id`

Optional:

- `id` (String) Contact twitter id.
- `name` (String) Contact twitter name.
- `screen_name` (String) Contact twitter screen name.

Read-Only:

- `profile_url` (String) Contact twitter account url.


<a id="nestedblock--whatsapp_id"></a>
### Nested Schema for `whatsapp_id`

Required:

- `display_name` (String) Contact whatsapp display name.
- `phone_number` (Block List, Min: 1) Contact whatsapp phone number. (see [below for nested schema](#nestedblock--whatsapp_id--phone_number))

<a id="nestedblock--whatsapp_id--phone_number"></a>
### Nested Schema for `whatsapp_id.phone_number`

Optional:

- `accepts_sms` (Boolean) If contact accept SMS.
- `country_code` (String) Phone number country code.
- `display` (String) Display string of the phone number.
- `e164` (String) Phone number in e164 format.
- `extension` (Number) Phone extension.



<a id="nestedblock--work_phone"></a>
### Nested Schema for `work_phone`

Optional:

- `accepts_sms` (Boolean) If contact accept SMS.
- `country_code` (String) Phone number country code.
- `display` (String) Display string of the phone number.
- `e164` (String) Phone number in e164 format.
- `extension` (Number) Phone extension.
