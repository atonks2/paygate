/*
 * Paygate API
 *
 * PayGate is a RESTful API enabling first-party Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transfers to be created without a deep understanding of a full NACHA file specification. First-party transfers initiate at an Originating Depository Financial Institution (ODFI) and are sent off to other Financial Institutions.  A namespace is a value used to isolate models from each other. This can be set to a \"user ID\" from your authentication service or any value your system has to identify.  There are also [admin endpoints](https://moov-io.github.io/paygate/admin/) for back-office operations.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// CreateTransfer These fields are used to initiate a Transfer between two Customer objects and their Accounts.
type CreateTransfer struct {
	Amount      Amount      `json:"amount"`
	Source      Source      `json:"source"`
	Destination Destination `json:"destination"`
	// Brief description of the transaction, that may appear on the receiving entity’s financial statement
	Description string `json:"description"`
	// When set to true this indicates the transfer should be processed the same day if possible.
	SameDay bool `json:"sameDay,omitempty"`
}
