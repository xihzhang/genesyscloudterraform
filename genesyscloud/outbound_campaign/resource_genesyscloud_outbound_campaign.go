package outbound_campaign

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v115/platformclientv2"
	"log"
	gcloud "terraform-provider-genesyscloud/genesyscloud"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"
	"time"
)

/*
The resource_genesyscloud_outbound_campaign.go contains all of the methods that perform the core logic for a resource.
*/

// getAllAuthOutboundCampaign retrieves all of the outbound campaign via Terraform in the Genesys Cloud and is used for the exporter
func getAllAuthOutboundCampaign(ctx context.Context, clientConfig *platformclientv2.Configuration) (resourceExporter.ResourceIDMetaMap, diag.Diagnostics) {
	resources := make(resourceExporter.ResourceIDMetaMap)
	proxy := getOutboundCampaignProxy(clientConfig)

	campaigns, err := proxy.getAllOutboundCampaign(ctx)
	if err != nil {
		return nil, diag.Errorf("Failed to get campaigns: %s", err)
	}

	for _, campaign := range *campaigns {
		if *campaign.CampaignStatus != "off" && *campaign.CampaignStatus != "on" {
			*campaign.CampaignStatus = "off"
		}
		resources[*campaign.Id] = &resourceExporter.ResourceMeta{Name: *campaign.Name}
	}

	return resources, nil
}

// createOutboundCampaign is used by the outbound_campaign resource to create Genesys cloud outbound campaign
func createOutboundCampaign(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*gcloud.ProviderMeta).ClientConfig
	proxy := newOutboundCampaignProxy(sdkConfig)
	campaignStatus := d.Get("campaign_status").(string)

	campaign := getOutboundCampaignFromResourceData(d)

	// Create campaign
	log.Printf("Creating Outbound Campaign %s", *campaign.Name)
	outboundCampaign, err := proxy.createOutboundCampaign(ctx, &campaign)
	if err != nil {
		return diag.Errorf("Failed to create Outbound Campaign %s: %s", *campaign.Name, err)
	}

	d.SetId(*outboundCampaign.Id)

	// Campaigns can be enabled after creation
	if campaignStatus == "on" {
		d.Set("campaign_status", campaignStatus)
		diag := updateOutboundCampaignStatus(ctx, d, proxy, *outboundCampaign, campaignStatus)
		if diag != nil {
			return diag
		}
	}

	log.Printf("Created Outbound Campaign %s %s", *outboundCampaign.Name, *outboundCampaign.Id)

	return readOutboundCampaign(ctx, d, meta)
}

// readOutboundCampaign is used by the outbound_campaign resource to read an outbound campaign from genesys cloud
func readOutboundCampaign(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*gcloud.ProviderMeta).ClientConfig
	proxy := newOutboundCampaignProxy(sdkConfig)

	log.Printf("Reading Outbound Campaign %s", d.Id())

	return gcloud.WithRetriesForRead(ctx, d, func() *retry.RetryError {
		campaign, resp, getErr := proxy.getOutboundCampaignById(ctx, d.Id())
		if getErr != nil {
			if gcloud.IsStatus404(resp) {
				return retry.RetryableError(fmt.Errorf("Failed to read Outbound Campaign %s: %s", d.Id(), getErr))
			}
			return retry.NonRetryableError(fmt.Errorf("Failed to read Outbound Campaign %s: %s", d.Id(), getErr))
		}

		if *campaign.CampaignStatus == "stopping" {
			return retry.RetryableError(fmt.Errorf("Outbound Campaign still stopping %s", d.Id()))
		}

		//cc := consistency_checker.NewConsistencyCheck(ctx, d, meta, ResourceOutboundCampaign())

		resourcedata.SetNillableValue(d, "name", campaign.Name)
		resourcedata.SetNillableReference(d, "contact_list_id", campaign.ContactList)
		resourcedata.SetNillableReference(d, "queue_id", campaign.Queue)
		resourcedata.SetNillableValue(d, "dialing_mode", campaign.DialingMode)
		resourcedata.SetNillableReference(d, "script_id", campaign.Script)
		resourcedata.SetNillableReference(d, "edge_group_id", campaign.EdgeGroup)
		resourcedata.SetNillableReference(d, "site_id", campaign.Site)
		resourcedata.SetNillableValue(d, "campaign_status", campaign.CampaignStatus)
		if campaign.PhoneColumns != nil {
			d.Set("phone_columns", flattenSdkoutboundcampaignPhonecolumnSlice(*campaign.PhoneColumns))
		}
		resourcedata.SetNillableValue(d, "abandon_rate", campaign.AbandonRate)
		if campaign.DncLists != nil {
			d.Set("dnc_list_ids", gcloud.SdkDomainEntityRefArrToList(*campaign.DncLists))
		}
		resourcedata.SetNillableReference(d, "callable_time_set_id", campaign.CallableTimeSet)
		resourcedata.SetNillableReference(d, "call_analysis_response_set_id", campaign.CallAnalysisResponseSet)
		resourcedata.SetNillableValue(d, "caller_name", campaign.CallerName)
		resourcedata.SetNillableValue(d, "caller_address", campaign.CallerAddress)
		resourcedata.SetNillableValue(d, "outbound_line_count", campaign.OutboundLineCount)
		if campaign.RuleSets != nil {
			d.Set("rule_set_ids", gcloud.SdkDomainEntityRefArrToList(*campaign.RuleSets))
		}
		resourcedata.SetNillableValue(d, "skip_preview_disabled", campaign.SkipPreviewDisabled)
		resourcedata.SetNillableValue(d, "preview_time_out_seconds", campaign.PreviewTimeOutSeconds)
		resourcedata.SetNillableValue(d, "always_running", campaign.AlwaysRunning)
		if campaign.ContactSorts != nil {
			d.Set("contact_sorts", flattenSdkoutboundcampaignContactsortSlice(*campaign.ContactSorts))
		}
		resourcedata.SetNillableValue(d, "no_answer_timeout", campaign.NoAnswerTimeout)
		resourcedata.SetNillableValue(d, "call_analysis_language", campaign.CallAnalysisLanguage)
		resourcedata.SetNillableValue(d, "priority", campaign.Priority)
		if campaign.ContactListFilters != nil {
			d.Set("contact_list_filter_ids", gcloud.SdkDomainEntityRefArrToList(*campaign.ContactListFilters))
		}
		resourcedata.SetNillableReference(d, "division_id", campaign.Division)
		if campaign.DynamicContactQueueingSettings != nil {
			d.Set("dynamic_contact_queueing_settings", flattenSdkDynamicContactQueueingSettings(*campaign.DynamicContactQueueingSettings))
		}

		log.Printf("Read Outbound Campaign %s %s", d.Id(), *campaign.Name)
		//return cc.CheckState()
		return nil
	})
}

// Function to format JSON response - Go
func formatJSON(input any) string {
	output, err := json.MarshalIndent(input, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	return string(output)
}

// updateOutboundCampaign is used by the outbound_campaign resource to update an outbound campaign in Genesys Cloud
func updateOutboundCampaign(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*gcloud.ProviderMeta).ClientConfig
	proxy := newOutboundCampaignProxy(sdkConfig)
	campaignStatus := d.Get("campaign_status").(string)

	campaign := getOutboundCampaignFromResourceData(d)

	log.Printf("Updating Outbound Campaign %s", *campaign.Name)
	campaignSdk, err := proxy.updateOutboundCampaign(ctx, d.Id(), &campaign)
	if err != nil {
		return diag.Errorf("Failed to update campaign %s", err)
	}

	// Check if Campaign Status needs updated
	diagErr := updateOutboundCampaignStatus(ctx, d, proxy, *campaignSdk, campaignStatus)
	if diagErr != nil {
		return diagErr
	}

	log.Printf("Updated Outbound Campaign %s", *campaign.Name)
	return readOutboundCampaign(ctx, d, meta)
}

// deleteOutboundCampaign is used by the outbound_campaign resource to delete an outbound campaign from Genesys cloud
func deleteOutboundCampaign(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*gcloud.ProviderMeta).ClientConfig
	proxy := newOutboundCampaignProxy(sdkConfig)

	campaignStatus := d.Get("campaign_status").(string)

	// Campaigns have to be turned off before they can be deleted
	if campaignStatus != "off" {
		diagErr := gcloud.RetryWhen(gcloud.IsStatus400, func() (*platformclientv2.APIResponse, diag.Diagnostics) {
			log.Printf("Turning off Outbound Campaign before deletion")
			outboundCampaign, resp, getErr := proxy.getOutboundCampaignById(ctx, d.Id())
			if getErr != nil {
				return resp, diag.Errorf("Failed to read Outbound Campaign %s: %s", d.Id(), getErr)
			}
			outboundCampaign.CampaignStatus = platformclientv2.String("off")
			// Handles updating the campaign based on what is set in ResourceData.campaign_status
			diagErr := updateOutboundCampaignStatus(ctx, d, proxy, *outboundCampaign, campaignStatus)
			if diagErr != nil {
				return resp, diagErr
			}
			return resp, nil
		})
		if diagErr != nil {
			return diagErr
		}
	}
	_, err := proxy.deleteOutboundCampaign(ctx, d.Id())
	if err != nil {
		return diag.Errorf("Failed to delete campaign %s: %s", d.Id(), err)
	}

	return gcloud.WithRetries(ctx, 30*time.Second, func() *retry.RetryError {
		_, resp, err := proxy.getOutboundCampaignById(ctx, d.Id())
		if err != nil {
			if gcloud.IsStatus404(resp) {
				// Outbound Campaign deleted
				log.Printf("Deleted Outbound Campaign %s", d.Id())
				return nil
			}
			return retry.NonRetryableError(fmt.Errorf("Error deleting Outbound Campaign %s: %s", d.Id(), err))
		}

		return retry.RetryableError(fmt.Errorf("Outbound Campaign %s still exists", d.Id()))
	})
}