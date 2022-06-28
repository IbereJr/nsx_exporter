/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

type BackupConfiguration struct {

	// true if automated backup is enabled
	BackupEnabled bool `json:"backup_enabled"`

	// Set when backups should be taken - on a weekly schedule or at regular intervals.
	BackupSchedule *BackupSchedule `json:"backup_schedule,omitempty"`

	// The minimum number of seconds between each upload of the inventory summary to backup server.
	InventorySummaryInterval int64 `json:"inventory_summary_interval"`

	// Passphrase used to encrypt backup files.
	Passphrase string `json:"passphrase,omitempty"`

	// The server to which backups will be sent.
	RemoteFileServer *RemoteFileServer `json:"remote_file_server"`
}
