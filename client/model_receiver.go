/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// Receiver struct for Receiver
type Receiver struct {
	// Receiver ID
	ID string `json:"ID,omitempty"`
	// The receivers email address
	Email string `json:"email,omitempty"`
	// The depository account to be used by default per transfer. ID must be a valid Receiver Depository account
	DefaultDepository string `json:"defaultDepository,omitempty"`
	// Defines the status of the Receiver
	Status string `json:"status,omitempty"`
	// optional object required for Know Your Customer (KYC) validation of this Originator
	BirthDate time.Time `json:"birthDate,omitempty"`
	Address   Address   `json:"address,omitempty"`
	// Optional ID when Originator data was created against Moov's Customers service
	CustomerID string `json:"customerID,omitempty"`
	// Additional meta data to be used for display only
	Metadata string    `json:"metadata,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Updated  time.Time `json:"updated,omitempty"`
}
