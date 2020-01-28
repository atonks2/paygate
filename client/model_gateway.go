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

// Gateway struct for Gateway
type Gateway struct {
	// ID to uniquely identify a gateway
	ID string `json:"ID,omitempty"`
	// Routing Number - four digit Federal Reserve Routing Symbol and the four digit ABA Institution Identifier
	Origin string `json:"origin"`
	// Legal name associated with the origin routing number.
	OriginName string `json:"originName"`
	// Routing Number - four digit Federal Reserve Routing Symbol and the four digit ABA Institution Identifier
	Destination string `json:"destination"`
	// Legal name associated with the destination routing number
	DestinationName string    `json:"destinationName"`
	Created         time.Time `json:"created,omitempty"`
}
