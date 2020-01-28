/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Batch struct for Batch
type Batch struct {
	BatchHeader  BatchHeader   `json:"batchHeader,omitempty"`
	EntryDetails []EntryDetail `json:"entryDetails,omitempty"`
	BatchControl BatchControl  `json:"batchControl,omitempty"`
}
