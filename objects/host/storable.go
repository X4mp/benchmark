package host

type storableHost struct {
	ID             string `json:"id"`
	IP             string `json:"ip"`
	Hostname       string `json:"hostname"`
	Longitude      int    `json:"longitude"`
	Latitude       int    `json:"latitude"`
	CountryID      string `json:"country_id"`
	RegionID       string `json:"region_id"`
	CityID         string `json:"city_id"`
	OrganizationID string `json:"organization_id"`
}

func createStorableHost(ins Host) *storableHost {

	var countryID string
	if ins.HasCountry() {
		countryID = ins.Country().ID().String()
	}

	var regionID string
	if ins.HasRegion() {
		regionID = ins.Region().ID().String()
	}

	var cityID string
	if ins.HasCity() {
		cityID = ins.City().ID().String()
	}

	var organizationID string
	if ins.HasOrganization() {
		organizationID = ins.Organization().ID().String()
	}

	out := storableHost{
		ID:             ins.ID().String(),
		IP:             ins.IP().String(),
		Hostname:       ins.Hostname(),
		Longitude:      ins.Longitude(),
		Latitude:       ins.Latitude(),
		CountryID:      countryID,
		RegionID:       regionID,
		CityID:         cityID,
		OrganizationID: organizationID,
	}

	return &out
}
