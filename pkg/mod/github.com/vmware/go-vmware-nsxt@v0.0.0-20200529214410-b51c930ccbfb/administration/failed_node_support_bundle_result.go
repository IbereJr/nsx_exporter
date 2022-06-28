/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

type FailedNodeSupportBundleResult struct {

	// Error code
	ErrorCode string `json:"error_code,omitempty"`

	// Error message
	ErrorMessage string `json:"error_message,omitempty"`

	// Display name of node
	NodeDisplayName string `json:"node_display_name,omitempty"`

	// UUID of node
	NodeId string `json:"node_id,omitempty"`
}
