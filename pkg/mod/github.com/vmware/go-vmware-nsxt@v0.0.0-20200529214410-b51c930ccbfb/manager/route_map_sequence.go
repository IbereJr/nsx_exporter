/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type RouteMapSequence struct {

	// Action for the Sequence
	Action string `json:"action"`

	// Match Criteria for the RouteMap Sequence
	MatchCriteria *RouteMapSequenceMatch `json:"match_criteria"`

	// Set Criteria for the RouteMap Sequence
	SetCriteria *RouteMapSequenceSet `json:"set_criteria,omitempty"`
}
