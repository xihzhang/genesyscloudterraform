package telephony_providers_edges_site

import (
	"context"
	"fmt"

	"github.com/mypurecloud/platform-client-sdk-go/v109/platformclientv2"
)

/*
The genesyscloud_telephony_providers_edges_site_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.

Each proxy implementation:

1.  Should provide a private package level variable that holds a instance of a proxy class.
2.  A New... constructor function  to initialize the proxy object.  This constructor should only be used within
    the proxy.
3.  A get private constructor function that the classes in the package can be used to to retrieve
    the proxy.  This proxy should check to see if the package level proxy instance is nil and
    should initialize it, otherwise it should return the instance
4.  Type definitions for each function that will be used in the proxy.  We use composition here
    so that we can easily provide mocks for testing.
5.  A struct for the proxy that holds an attribute for each function type.
6.  Wrapper methods on each of the elements on the struct.
7.  Function implementations for each function type definition.

*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *siteProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type getAllManagedSitesFunc func(ctx context.Context, p *siteProxy) (*[]platformclientv2.Site, error)
type getAllUnmanagedSitesFunc func(ctx context.Context, p *siteProxy) (*[]platformclientv2.Site, error)
type createSiteFunc func(ctx context.Context, p *siteProxy, site *platformclientv2.Site) (*platformclientv2.Site, error)
type deleteSiteFunc func(ctx context.Context, p *siteProxy, siteId string) (*platformclientv2.APIResponse, error)
type getSiteByIdFunc func(ctx context.Context, p *siteProxy, siteId string) (site *platformclientv2.Site, resp *platformclientv2.APIResponse, err error)
type getSiteIdByNameFunc func(ctx context.Context, p *siteProxy, siteName string, managed bool) (siteId string, retryable bool, err error)
type updateSiteFunc func(ctx context.Context, p *siteProxy, siteId string, site *platformclientv2.Site) (*platformclientv2.Site, *platformclientv2.APIResponse, error)

type createSiteOutboundRouteFunc func(ctx context.Context, p *siteProxy, siteId string, outboundRoute *platformclientv2.Outboundroutebase) (*platformclientv2.Outboundroutebase, error)
type getSiteOutboundRoutesFunc func(ctx context.Context, p *siteProxy, siteId string) (*[]platformclientv2.Outboundroutebase, error)
type updateSiteOutboundRouteFunc func(ctx context.Context, p *siteProxy, siteId string, outboundRouteId string, outboundRoute *platformclientv2.Outboundroutebase) (*platformclientv2.Outboundroutebase, error)
type deleteSiteOutboundRouteFunc func(ctx context.Context, p *siteProxy, siteId string, outboundRouteId string) (*platformclientv2.APIResponse, error)

type getSiteNumberPlansFunc func(ctx context.Context, p *siteProxy, siteId string) (*[]platformclientv2.Numberplan, *platformclientv2.APIResponse, error)
type updateSiteNumberPlansFunc func(ctx context.Context, p *siteProxy, siteId string, numberPlans *[]platformclientv2.Numberplan) (*[]platformclientv2.Numberplan, *platformclientv2.APIResponse, error)

type getLocationFunc func(ctx context.Context, p *siteProxy, locationId string) (*platformclientv2.Locationdefinition, error)

type getTelephonyMediaregionsFunc func(ctx context.Context, p *siteProxy) (*platformclientv2.Mediaregions, error)

// siteProxy contains all of the methods that call genesys cloud APIs.
type siteProxy struct {
	clientConfig *platformclientv2.Configuration
	edgesApi     *platformclientv2.TelephonyProvidersEdgeApi
	locationsApi *platformclientv2.LocationsApi
	telephonyApi *platformclientv2.TelephonyApi

	getAllManagedSitesAttr   getAllManagedSitesFunc
	getAllUnmanagedSitesAttr getAllUnmanagedSitesFunc
	createSiteAttr           createSiteFunc
	deleteSiteAttr           deleteSiteFunc
	getSiteByIdAttr          getSiteByIdFunc
	getSiteIdByNameAttr      getSiteIdByNameFunc
	updateSiteAttr           updateSiteFunc

	createSiteOutboundRouteAttr createSiteOutboundRouteFunc
	getSiteOutboundRoutesAttr   getSiteOutboundRoutesFunc
	updateSiteOutboundRouteAttr updateSiteOutboundRouteFunc
	deleteSiteOutboundRouteAttr deleteSiteOutboundRouteFunc

	getSiteNumberPlansAttr    getSiteNumberPlansFunc
	updateSiteNumberPlansAttr updateSiteNumberPlansFunc

	getLocationAttr getLocationFunc

	getTelephonyMediaregionsAttr getTelephonyMediaregionsFunc
}

// newSiteProxy initializes the Site proxy with all of the data needed to communicate with Genesys Cloud
func newSiteProxy(clientConfig *platformclientv2.Configuration) *siteProxy {
	edgesApi := platformclientv2.NewTelephonyProvidersEdgeApiWithConfig(clientConfig)
	locationsApi := platformclientv2.NewLocationsApiWithConfig(clientConfig)
	telephonyApi := platformclientv2.NewTelephonyApiWithConfig(clientConfig)

	return &siteProxy{
		clientConfig: clientConfig,
		edgesApi:     edgesApi,
		locationsApi: locationsApi,
		telephonyApi: telephonyApi,

		getAllManagedSitesAttr:   getAllManagedSitesFn,
		getAllUnmanagedSitesAttr: getAllUnmanagedSitesFn,
		createSiteAttr:           createSiteFn,
		deleteSiteAttr:           deleteSiteFn,
		getSiteByIdAttr:          getSiteByIdFn,
		getSiteIdByNameAttr:      getSiteIdByNameFn,
		updateSiteAttr:           updateSiteFn,

		createSiteOutboundRouteAttr: createSiteOutboundRouteFn,
		getSiteOutboundRoutesAttr:   getSiteOutboundRoutesFn,
		updateSiteOutboundRouteAttr: updateSiteOutboundRouteFn,
		deleteSiteOutboundRouteAttr: deleteSiteOutboundRouteFn,

		getSiteNumberPlansAttr:    getSiteNumberPlansFn,
		updateSiteNumberPlansAttr: updateSiteNumberPlansFn,

		getLocationAttr: getLocationFn,

		getTelephonyMediaregionsAttr: getTelephonyMediaregionsFn,
	}
}

// getSiteProxy acts as a singleton for the internalProxy.  It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getSiteProxy(clientConfig *platformclientv2.Configuration) *siteProxy {
	if internalProxy == nil {
		internalProxy = newSiteProxy(clientConfig)
	}
	return internalProxy
}

// getAllManagedSitesFunc retrieves all managed Genesys Cloud Sites
func (p *siteProxy) getAllManagedSites(ctx context.Context) (*[]platformclientv2.Site, error) {
	return p.getAllManagedSitesAttr(ctx, p)
}

// getAllUnmanagedSitesFunc retrieves all unmanaged Genesys Cloud Sites
func (p *siteProxy) getAllUnmanagedSites(ctx context.Context) (*[]platformclientv2.Site, error) {
	return p.getAllUnmanagedSitesAttr(ctx, p)
}

// createSiteFunc creates a Genesys Cloud Site
func (p *siteProxy) createSite(ctx context.Context, site *platformclientv2.Site) (*platformclientv2.Site, error) {
	return p.createSiteAttr(ctx, p, site)
}

// deleteSiteFunc deletes a Genesys Cloud Site by ID
func (p *siteProxy) deleteSite(ctx context.Context, siteId string) (*platformclientv2.APIResponse, error) {
	return p.deleteSiteAttr(ctx, p, siteId)
}

// getSiteByIdFunc returns a single Genesys Cloud Site by Id
func (p *siteProxy) getSiteById(ctx context.Context, siteId string) (site *platformclientv2.Site, resp *platformclientv2.APIResponse, err error) {
	return p.getSiteByIdAttr(ctx, p, siteId)
}

// getSiteIdByNameFunc returns a single Genesys Cloud Site by Name
func (p *siteProxy) getSiteIdByName(ctx context.Context, siteName string, managed bool) (siteId string, retryable bool, err error) {
	return p.getSiteIdByNameAttr(ctx, p, siteName, managed)
}

// updateSiteFunc updates a Genesys Cloud Site
func (p *siteProxy) updateSite(ctx context.Context, siteId string, site *platformclientv2.Site) (*platformclientv2.Site, *platformclientv2.APIResponse, error) {
	return p.updateSiteAttr(ctx, p, siteId, site)
}

// createSiteOutboundRouteFunc creates an Outbound Route for a Genesys Cloud Site
func (p *siteProxy) createSiteOutboundRoute(ctx context.Context, siteId string, outboundRoute *platformclientv2.Outboundroutebase) (*platformclientv2.Outboundroutebase, error) {
	return p.createSiteOutboundRouteAttr(ctx, p, siteId, outboundRoute)
}

// getSiteByIdFunc returns a single Outbound Route by Id
func (p *siteProxy) getSiteOutboundRoutes(ctx context.Context, siteId string) (*[]platformclientv2.Outboundroutebase, error) {
	return p.getSiteOutboundRoutesAttr(ctx, p, siteId)
}

// updateSiteFunc updates a Genesys Cloud Outbound Route for a Genesys Cloud Site
func (p *siteProxy) updateSiteOutboundRoute(ctx context.Context, siteId string, outboundRouteId string, outboundRoute *platformclientv2.Outboundroutebase) (*platformclientv2.Outboundroutebase, error) {
	return p.updateSiteOutboundRouteAttr(ctx, p, siteId, outboundRouteId, outboundRoute)
}

// deleteSiteFunc deletes a Genesys Cloud Outbound Route by Id for a Genesys Cloud Site
func (p *siteProxy) deleteSiteOutboundRoute(ctx context.Context, siteId string, outboundRouteId string) (*platformclientv2.APIResponse, error) {
	return p.deleteSiteOutboundRouteAttr(ctx, p, siteId, outboundRouteId)
}

// getSiteNumberPlansFunc retrieves all Number Plans of a Genesys Cloud Sites
func (p *siteProxy) getSiteNumberPlans(ctx context.Context, siteId string) (*[]platformclientv2.Numberplan, *platformclientv2.APIResponse, error) {
	return p.getSiteNumberPlansAttr(ctx, p, siteId)
}

// updateSiteNumberPlansFunc updates the Number Plans for a Genesys Cloud Site
func (p *siteProxy) updateSiteNumberPlans(ctx context.Context, siteId string, numberPlans *[]platformclientv2.Numberplan) (*[]platformclientv2.Numberplan, *platformclientv2.APIResponse, error) {
	return p.updateSiteNumberPlansAttr(ctx, p, siteId, numberPlans)
}

// getLocation retrieves a Genesys Cloud Location by Id
func (p *siteProxy) getLocation(ctx context.Context, locationId string) (*platformclientv2.Locationdefinition, error) {
	return p.getLocationAttr(ctx, p, locationId)
}

func (p *siteProxy) getTelephonyMediaregions(ctx context.Context) (*platformclientv2.Mediaregions, error) {
	return p.getTelephonyMediaregionsAttr(ctx, p)
}

func getAllManagedSitesFn(ctx context.Context, p *siteProxy) (*[]platformclientv2.Site, error) {
	var allManagedSites []platformclientv2.Site

	for pageNum := 1; ; pageNum++ {
		const pageSize = 100
		sites, _, err := p.edgesApi.GetTelephonyProvidersEdgesSites(pageSize, pageNum, "", "", "", "", true)
		if err != nil {
			return nil, err
		}

		if sites.Entities == nil || len(*sites.Entities) == 0 {
			break
		}

		for _, site := range *sites.Entities {
			if site.State != nil && *site.State != "deleted" {
				allManagedSites = append(allManagedSites, site)
			}
		}
	}

	return &allManagedSites, nil
}

func getAllUnmanagedSitesFn(ctx context.Context, p *siteProxy) (*[]platformclientv2.Site, error) {
	var allUnmanagedSites []platformclientv2.Site

	for pageNum := 1; ; pageNum++ {
		const pageSize = 100
		sites, _, err := p.edgesApi.GetTelephonyProvidersEdgesSites(pageSize, pageNum, "", "", "", "", false)
		if err != nil {
			return nil, err
		}

		if sites.Entities == nil || len(*sites.Entities) == 0 {
			break
		}

		for _, site := range *sites.Entities {
			if site.State != nil && *site.State != "deleted" {
				allUnmanagedSites = append(allUnmanagedSites, site)
			}
		}
	}

	return &allUnmanagedSites, nil
}

func createSiteFn(ctx context.Context, p *siteProxy, siteReq *platformclientv2.Site) (*platformclientv2.Site, error) {
	site, _, err := p.edgesApi.PostTelephonyProvidersEdgesSites(*siteReq)
	if err != nil {
		return nil, err
	}

	return site, nil
}

func deleteSiteFn(ctx context.Context, p *siteProxy, siteId string) (*platformclientv2.APIResponse, error) {
	resp, err := p.edgesApi.DeleteTelephonyProvidersEdgesSite(siteId)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func getSiteByIdFn(ctx context.Context, p *siteProxy, siteId string) (*platformclientv2.Site, *platformclientv2.APIResponse, error) {
	site, resp, err := p.edgesApi.GetTelephonyProvidersEdgesSite(siteId)
	if err != nil {
		return nil, resp, err
	}

	return site, resp, nil
}

func getSiteIdByNameFn(ctx context.Context, p *siteProxy, siteName string, managed bool) (siteId string, retryable bool, err error) {
	siteId = ""

	for pageNum := 1; ; pageNum++ {
		const pageSize = 50
		sites, _, err := p.edgesApi.GetTelephonyProvidersEdgesSites(pageSize, pageNum, "", "", siteName, "", managed)
		if err != nil {
			return "", false, err
		}

		if sites.Entities == nil || len(*sites.Entities) == 0 {
			return "", true, fmt.Errorf("no sites found with name %s", siteName)
		}

		for _, site := range *sites.Entities {
			if (site.Name != nil && *site.Name == siteName) && (site.State != nil && *site.State != "deleted") {
				siteId = *site.Id
				break
			}
		}
		if siteId != "" {
			break
		}
	}

	return siteId, false, nil
}

func updateSiteFn(ctx context.Context, p *siteProxy, siteId string, site *platformclientv2.Site) (*platformclientv2.Site, *platformclientv2.APIResponse, error) {
	updatedSite, resp, err := p.edgesApi.PutTelephonyProvidersEdgesSite(siteId, *site)
	if err != nil {
		return nil, resp, err
	}

	return updatedSite, resp, nil
}

func createSiteOutboundRouteFn(ctx context.Context, p *siteProxy, siteId string, outboundRoute *platformclientv2.Outboundroutebase) (*platformclientv2.Outboundroutebase, error) {
	obr, _, err := p.edgesApi.PostTelephonyProvidersEdgesSiteOutboundroutes(siteId, *outboundRoute)
	if err != nil {
		return nil, err
	}

	return obr, nil
}

func getSiteOutboundRoutesFn(ctx context.Context, p *siteProxy, siteId string) (*[]platformclientv2.Outboundroutebase, error) {
	var allOutboundRoutes = []platformclientv2.Outboundroutebase{}

	for pageNum := 1; ; pageNum++ {
		const pageSize = 100
		outboundRoutes, _, err := p.edgesApi.GetTelephonyProvidersEdgesSiteOutboundroutes(siteId, pageSize, pageNum, "", "", "")
		if err != nil {
			return nil, err
		}
		if outboundRoutes.Entities == nil || len(*outboundRoutes.Entities) == 0 {
			break
		}
		allOutboundRoutes = append(allOutboundRoutes, *outboundRoutes.Entities...)
	}

	return &allOutboundRoutes, nil
}

func updateSiteOutboundRouteFn(ctx context.Context, p *siteProxy, siteId string, outboundRouteId string, outboundRoute *platformclientv2.Outboundroutebase) (*platformclientv2.Outboundroutebase, error) {
	obrs, _, err := p.edgesApi.PutTelephonyProvidersEdgesSiteOutboundroute(siteId, outboundRouteId, *outboundRoute)
	if err != nil {
		return nil, err
	}

	return obrs, nil
}

func deleteSiteOutboundRouteFn(ctx context.Context, p *siteProxy, siteId string, outboundRouteId string) (*platformclientv2.APIResponse, error) {
	resp, err := p.edgesApi.DeleteTelephonyProvidersEdgesSiteOutboundroute(siteId, outboundRouteId)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func getSiteNumberPlansFn(ctx context.Context, p *siteProxy, siteId string) (*[]platformclientv2.Numberplan, *platformclientv2.APIResponse, error) {
	numberPlans, resp, err := p.edgesApi.GetTelephonyProvidersEdgesSiteNumberplans(siteId)
	if err != nil {
		return nil, resp, err
	}

	return &numberPlans, resp, nil
}

func updateSiteNumberPlansFn(ctx context.Context, p *siteProxy, siteId string, numberPlansUpdate *[]platformclientv2.Numberplan) (*[]platformclientv2.Numberplan, *platformclientv2.APIResponse, error) {
	numberPlans, resp, err := p.edgesApi.PutTelephonyProvidersEdgesSiteNumberplans(siteId, *numberPlansUpdate)
	if err != nil {
		return nil, resp, err
	}

	return &numberPlans, resp, nil
}

func getLocationFn(ctx context.Context, p *siteProxy, locationId string) (*platformclientv2.Locationdefinition, error) {
	location, _, err := p.locationsApi.GetLocation(locationId, nil)
	if err != nil {
		return nil, err
	}
	if location.EmergencyNumber == nil {
		return nil, fmt.Errorf("location with id %v does not have an emergency number", locationId)
	}

	return location, nil
}

func getTelephonyMediaregionsFn(ctx context.Context, p *siteProxy) (*platformclientv2.Mediaregions, error) {
	telephonyRegions, _, err := p.telephonyApi.GetTelephonyMediaregions()
	if err != nil {
		return nil, err
	}

	return telephonyRegions, nil
}
