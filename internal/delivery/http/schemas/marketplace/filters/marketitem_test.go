package filters

import (
	"encoding/json"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego/filters"
	"github.com/legocy-co/legocy/pkg/helpers"
	"net/url"
	"reflect"
	"testing"
)

// TestBindMarketItemAndLegoSetFilters final version
func TestBindMarketItemAndLegoSetFilters(t *testing.T) {
	tests := []struct {
		name       string
		query      url.Values
		wantStruct MarketItemFilterDTO
	}{
		{
			name: "Complete fields",
			query: url.Values{
				"price_gte":                []string{"50.5"},
				"price_lte":                []string{"150.75"},
				"set_state__in":            []string{"new", "used"},
				"location__in":             []string{"store1", "store2"},
				"lego_set[npieces_gte]":    []string{"100"},
				"lego_set[npieces_lte]":    []string{"500"},
				"lego_set[series_id__in]":  []string{"1", "2", "3"},
				"lego_set[set_number__in]": []string{"123", "456"},
				"lego_set[name__ilike]":    []string{"Millennium Falcon"},
			},
			wantStruct: MarketItemFilterDTO{
				PriceGTE:  ptrFloat64(50.5),
				PriceLTE:  ptrFloat64(150.75),
				SetStates: []string{"new", "used"},
				Locations: []string{"store1", "store2"},
				LegoSet: &filters.LegoSetFilterDTO{
					NpiecesGTE: ptrInt(100),
					NpiecesLTE: ptrInt(500),
					SeriesIDs:  ptrSliceInt([]int{1, 2, 3}),
					SetNumbers: ptrSliceInt([]int{123, 456}),
					Name:       ptrString("Millennium Falcon"),
				},
			},
		},
		{
			name: "Missing LegoSet fields",
			query: url.Values{
				"price_gte":     []string{"100.0"},
				"set_state__in": []string{"available,discontinued"},
				"location__in":  []string{"online"},
			},
			wantStruct: MarketItemFilterDTO{
				PriceGTE:  ptrFloat64(100.0),
				SetStates: []string{"available", "discontinued"},
				Locations: []string{"online"},
				LegoSet:   nil,
			},
		},
		{
			name: "Partially filled nested struct",
			query: url.Values{
				"price_gte":             []string{"300.0"},
				"set_state__in":         []string{"new", "used"},
				"location__in":          []string{"online"},
				"lego_set[npieces_gte]": []string{"100"}, // Only npieces_gte is provided for the nested struct
			},
			wantStruct: MarketItemFilterDTO{
				PriceGTE:  ptrFloat64(300.0),
				SetStates: []string{"new", "used"},
				Locations: []string{"online"},
				LegoSet: &filters.LegoSetFilterDTO{
					NpiecesGTE: ptrInt(100),
					NpiecesLTE: nil,
					SeriesIDs:  nil,
					SetNumbers: nil,
					Name:       nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotStruct MarketItemFilterDTO
			helpers.BindQueryParamsToStruct(&gotStruct, tt.query)
			if !compareMarketItemFilterDTO(gotStruct, tt.wantStruct) {

				jsonGotStruct, _ := json.Marshal(gotStruct)
				jsonWantStruct, _ := json.Marshal(tt.wantStruct)

				t.Errorf(
					"%s: BindQueryParamsToStruct() got = %+v, want %+v",
					tt.name, string(jsonGotStruct), string(jsonWantStruct),
				)
			}
		})
	}
}

func ptrFloat64(f float64) *float64 {
	return &f
}

func ptrInt(i int) *int {
	return &i
}

func ptrString(s string) *string {
	return &s
}

func ptrSliceInt(si []int) *[]int {
	return &si
}

func compareMarketItemFilterDTO(got, want MarketItemFilterDTO) bool {
	if !compareFloat64Pointers(got.PriceGTE, want.PriceGTE) || !compareFloat64Pointers(got.PriceLTE, want.PriceLTE) {
		return false
	}
	if !reflect.DeepEqual(got.SetStates, want.SetStates) || !reflect.DeepEqual(got.Locations, want.Locations) {
		return false
	}
	// For LegoSet, compare only if both are non-nil, or both are nil.
	if got.LegoSet != nil && want.LegoSet != nil {
		if !compareLegoSetFilterDTO(*got.LegoSet, *want.LegoSet) {
			return false
		}
	} else if got.LegoSet != want.LegoSet { // Checks if one is nil and the other isn't
		return false
	}
	return true
}

func compareLegoSetFilterDTO(got, want filters.LegoSetFilterDTO) bool {
	if (got.NpiecesGTE == nil) != (want.NpiecesGTE == nil) || (got.NpiecesGTE != nil && *got.NpiecesGTE != *want.NpiecesGTE) {
		return false
	}
	if (got.NpiecesLTE == nil) != (want.NpiecesLTE == nil) || (got.NpiecesLTE != nil && *got.NpiecesLTE != *want.NpiecesLTE) {
		return false
	}
	if !compareOptionalIntSlice(got.SeriesIDs, want.SeriesIDs) {
		return false
	}
	if !compareOptionalIntSlice(got.SetNumbers, want.SetNumbers) {
		return false
	}
	if (got.Name == nil) != (want.Name == nil) || (got.Name != nil && *got.Name != *want.Name) {
		return false
	}
	return true
}

func compareOptionalIntSlice(got, want *[]int) bool {
	if got == nil && want == nil {
		return true
	}
	if (got == nil) != (want == nil) {
		return false
	}
	if got != nil && want != nil {
		return reflect.DeepEqual(*got, *want)
	}
	return false // Shouldn't reach here
}

func compareFloat64Pointers(got, want *float64) bool {
	if (got == nil) != (want == nil) {
		return false // One is nil, the other isn't
	}
	if got != nil && want != nil && *got != *want {
		return false // Both non-nil but different values
	}
	return true
}
