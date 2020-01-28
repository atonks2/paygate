/*
 * Paygate Admin API
 *
 * Paygate is ...
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CutoffTime struct for CutoffTime
type CutoffTime struct {
	// 24-hour timestamp for last processing minute
	Cutoff float32 `json:"cutoff"`
	// IANA timezone name for cutoff time
	Location string `json:"location"`
}
