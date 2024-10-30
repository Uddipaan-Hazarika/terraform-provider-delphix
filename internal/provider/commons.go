package provider

const (
	Pending                string = "PENDING"
	Started                string = "STARTED"
	Timedout               string = "TIMEDOUT"
	Failed                 string = "FAILED"
	Completed              string = "COMPLETED"
	Canceled               string = "CANCELED"
	Abandoned              string = "ABANDONED"
	JOB_STATUS_SLEEP_TIME  int    = 5
	STATUS_POLL_SLEEP_TIME int    = 20
	DLPX                   string = "[DELPHIX] "
	INFO                   string = "[INFO] "
	WARN                   string = "[WARN] "
	ERROR                  string = "[ERROR] "
)

var updatableEnvKeys = map[string]bool{
	"name":               true,
	"cluster_home":       true,
	"connector_port":     true,
	"username":           true,
	"password":           true,
	"description":        true,
	"allow_provisioning": true,
	"is_staging":         true,
	"version":            true,
	"oracle_base":        true,
	"bits":               true,
	"tags":               true,
	"hosts":              true,
}

var isDestructiveEnvUpdate = map[string]bool{
	"name":               false,
	"cluster_home":       true,
	"connector_port":     true,
	"username":           true,
	"password":           true,
	"description":        false,
	"allow_provisioning": false,
	"is_staging":         false,
	"version":            false,
	"oracle_base":        false,
	"bits":               false,
	"tags":               false,
	"hosts":              true,
}

var updatableVdbKeys = map[string]bool{
	"name":                          true,
	"db_username":                   true,
	"db_password":                   true,
	"pre_refresh":                   true,
	"post_refresh":                  true,
	"configure_clone":               true,
	"pre_snapshot":                  true,
	"post_snapshot":                 true,
	"pre_rollback":                  true,
	"post_rollback":                 true,
	"pre_start":                     true,
	"post_start":                    true,
	"pre_stop":                      true,
	"post_stop":                     true,
	"template_id":                   true,
	"pre_script":                    true,
	"post_script":                   true,
	"custom_env_vars":               true,
	"custom_env_files":              true,
	"appdata_source_params":         true,
	"config_params":                 true,
	"cdb_tde_keystore_password":     true,
	"target_vcdb_tde_keystore_path": true,
	"tde_key_identifier":            true,
	"parent_tde_keystore_password":  true,
	"parent_tde_keystore_path":      true,
	"additional_mount_points":       true,
	"cdc_on_provision":              true,
	"environment_user_id":           true,
	"listener_ids":                  true,
	"vdb_restart":                   true,
	"new_dbid":                      true,
	"mount_point":                   true,
	"tags":                          true,
}

var isDestructiveVdbUpdate = map[string]bool{
	"name":                          false,
	"db_username":                   false,
	"db_password":                   false,
	"pre_refresh":                   false,
	"post_refresh":                  false,
	"configure_clone":               false,
	"pre_snapshot":                  false,
	"post_snapshot":                 false,
	"pre_rollback":                  false,
	"post_rollback":                 false,
	"pre_start":                     false,
	"post_start":                    false,
	"pre_stop":                      false,
	"post_stop":                     false,
	"template_id":                   true,
	"pre_script":                    false,
	"post_script":                   false,
	"custom_env_vars":               false,
	"custom_env_files":              false,
	"appdata_source_params":         true,
	"config_params":                 true,
	"cdb_tde_keystore_password":     true,
	"target_vcdb_tde_keystore_path": true,
	"tde_key_identifier":            true,
	"parent_tde_keystore_password":  true,
	"parent_tde_keystore_path":      true,
	"additional_mount_points":       false,
	"cdc_on_provision":              true,
	"environment_user_id":           true,
	"listener_ids":                  false,
	"vdb_restart":                   false,
	"new_dbid":                      false,
	"mount_point":                   true,
	"tags":                          false,
}
