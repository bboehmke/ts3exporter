package api

import "fmt"

type Server struct {
	VirtualserverAutostart          int    `json:"virtualserver_autostart,string"`
	VirtualserverClientsonline      int    `json:"virtualserver_clientsonline,string"`
	VirtualserverID                 int    `json:"virtualserver_id,string"`
	VirtualserverMachineID          string `json:"virtualserver_machine_id"`
	VirtualserverMaxclients         int    `json:"virtualserver_maxclients,string"`
	VirtualserverName               string `json:"virtualserver_name"`
	VirtualserverPort               int    `json:"virtualserver_port,string"`
	VirtualserverQueryclientsonline int    `json:"virtualserver_queryclientsonline,string"`
	VirtualserverStatus             string `json:"virtualserver_status"`
	VirtualserverUptime             int    `json:"virtualserver_uptime,string"`
}

func (c *TS3HttpClient) ServerList() ([]Server, error) {
	var servers []Server
	return servers, c.request("serverlist", &servers)
}

type ServerInfo struct {
	ConnectionBandwidthReceivedLastMinuteTotal          int     `json:"connection_bandwidth_received_last_minute_total,string"`
	ConnectionBandwidthReceivedLastSecondTotal          int     `json:"connection_bandwidth_received_last_second_total,string"`
	ConnectionBandwidthSentLastMinuteTotal              int     `json:"connection_bandwidth_sent_last_minute_total,string"`
	ConnectionBandwidthSentLastSecondTotal              int     `json:"connection_bandwidth_sent_last_second_total,string"`
	ConnectionBytesReceivedControl                      int     `json:"connection_bytes_received_control,string"`
	ConnectionBytesReceivedKeepalive                    int     `json:"connection_bytes_received_keepalive,string"`
	ConnectionBytesReceivedSpeech                       int     `json:"connection_bytes_received_speech,string"`
	ConnectionBytesReceivedTotal                        int     `json:"connection_bytes_received_total,string"`
	ConnectionBytesSentControl                          int     `json:"connection_bytes_sent_control,string"`
	ConnectionBytesSentKeepalive                        int     `json:"connection_bytes_sent_keepalive,string"`
	ConnectionBytesSentSpeech                           int     `json:"connection_bytes_sent_speech,string"`
	ConnectionBytesSentTotal                            int     `json:"connection_bytes_sent_total,string"`
	ConnectionFiletransferBandwidthReceived             int     `json:"connection_filetransfer_bandwidth_received,string"`
	ConnectionFiletransferBandwidthSent                 int     `json:"connection_filetransfer_bandwidth_sent,string"`
	ConnectionFiletransferBytesReceivedTotal            int     `json:"connection_filetransfer_bytes_received_total,string"`
	ConnectionFiletransferBytesSentTotal                int     `json:"connection_filetransfer_bytes_sent_total,string"`
	ConnectionPacketsReceivedControl                    int     `json:"connection_packets_received_control,string"`
	ConnectionPacketsReceivedKeepalive                  int     `json:"connection_packets_received_keepalive,string"`
	ConnectionPacketsReceivedSpeech                     int     `json:"connection_packets_received_speech,string"`
	ConnectionPacketsReceivedTotal                      int     `json:"connection_packets_received_total,string"`
	ConnectionPacketsSentControl                        int     `json:"connection_packets_sent_control,string"`
	ConnectionPacketsSentKeepalive                      int     `json:"connection_packets_sent_keepalive,string"`
	ConnectionPacketsSentSpeech                         int     `json:"connection_packets_sent_speech,string"`
	ConnectionPacketsSentTotal                          int     `json:"connection_packets_sent_total,string"`
	VirtualserverAntifloodPointsNeededCommandBlock      int     `json:"virtualserver_antiflood_points_needed_command_block,string"`
	VirtualserverAntifloodPointsNeededIPBlock           int     `json:"virtualserver_antiflood_points_needed_ip_block,string"`
	VirtualserverAntifloodPointsNeededPluginBlock       int     `json:"virtualserver_antiflood_points_needed_plugin_block,string"`
	VirtualserverAntifloodPointsTickReduce              int     `json:"virtualserver_antiflood_points_tick_reduce,string"`
	VirtualserverAskForPrivilegekey                     int     `json:"virtualserver_ask_for_privilegekey,string"`
	VirtualserverAutostart                              int     `json:"virtualserver_autostart,string"`
	VirtualserverChannelTempDeleteDelayDefault          int     `json:"virtualserver_channel_temp_delete_delay_default,string"`
	VirtualserverChannelsOnline                         int     `json:"virtualserver_channelsonline,string"`
	VirtualserverClientConnections                      int     `json:"virtualserver_client_connections,string"`
	VirtualserverClientsonline                          int     `json:"virtualserver_clientsonline,string"`
	VirtualserverCodecEncryptionMode                    int     `json:"virtualserver_codec_encryption_mode,string"`
	VirtualserverComplainAutoBanCount                   int     `json:"virtualserver_complain_autoban_count,string"`
	VirtualserverComplainAutoBanTime                    int     `json:"virtualserver_complain_autoban_time,string"`
	VirtualserverComplainRemoveTime                     int     `json:"virtualserver_complain_remove_time,string"`
	VirtualserverCreated                                int     `json:"virtualserver_created,string"`
	VirtualserverDefaultChannelAdminGroup               int     `json:"virtualserver_default_channel_admin_group,string"`
	VirtualserverDefaultChannelGroup                    int     `json:"virtualserver_default_channel_group,string"`
	VirtualserverDefaultServerGroup                     int     `json:"virtualserver_default_server_group,string"`
	VirtualserverDownloadQuota                          uint    `json:"virtualserver_download_quota,string"`
	VirtualserverFilebase                               string  `json:"virtualserver_filebase"`
	VirtualserverFlagPassword                           int     `json:"virtualserver_flag_password,string"`
	VirtualserverHostbannerGfxInterval                  int     `json:"virtualserver_hostbanner_gfx_interval,string"`
	VirtualserverHostbannerGfxURL                       string  `json:"virtualserver_hostbanner_gfx_url"`
	VirtualserverHostbannerMode                         int     `json:"virtualserver_hostbanner_mode,string"`
	VirtualserverHostbannerURL                          string  `json:"virtualserver_hostbanner_url"`
	VirtualserverHostbuttonGfxURL                       string  `json:"virtualserver_hostbutton_gfx_url"`
	VirtualserverHostbuttonTooltip                      string  `json:"virtualserver_hostbutton_tooltip"`
	VirtualserverHostbuttonURL                          string  `json:"virtualserver_hostbutton_url"`
	VirtualserverHostmessage                            string  `json:"virtualserver_hostmessage"`
	VirtualserverHostmessageMode                        int     `json:"virtualserver_hostmessage_mode,string"`
	VirtualserverIconID                                 int     `json:"virtualserver_icon_id,string"`
	VirtualserverID                                     int     `json:"virtualserver_id,string"`
	VirtualserverIP                                     string  `json:"virtualserver_ip"`
	VirtualserverLogChannel                             int     `json:"virtualserver_log_channel,string"`
	VirtualserverLogClient                              int     `json:"virtualserver_log_client,string"`
	VirtualserverLogFiletransfer                        int     `json:"virtualserver_log_filetransfer,string"`
	VirtualserverLogPermissions                         int     `json:"virtualserver_log_permissions,string"`
	VirtualserverLogQuery                               int     `json:"virtualserver_log_query,string"`
	VirtualserverLogServer                              int     `json:"virtualserver_log_server,string"`
	VirtualserverMachineID                              string  `json:"virtualserver_machine_id"`
	VirtualserverMaxDownloadTotalBandwidth              uint    `json:"virtualserver_max_download_total_bandwidth,string"`
	VirtualserverMaxUploadTotalBandwidth                uint    `json:"virtualserver_max_upload_total_bandwidth,string"`
	VirtualserverMaxclients                             int     `json:"virtualserver_maxclients,string"`
	VirtualserverMinAndroidVersion                      int     `json:"virtualserver_min_android_version,string"`
	VirtualserverMinClientVersion                       int     `json:"virtualserver_min_client_version,string"`
	VirtualserverMinClientsInChannelBeforeForcedSilence int     `json:"virtualserver_min_clients_in_channel_before_forced_silence,string"`
	VirtualserverMinIosVersion                          int     `json:"virtualserver_min_ios_version,string"`
	VirtualserverMonthBytesDownloaded                   int     `json:"virtualserver_month_bytes_downloaded,string"`
	VirtualserverMonthBytesUploaded                     int     `json:"virtualserver_month_bytes_uploaded,string"`
	VirtualserverName                                   string  `json:"virtualserver_name"`
	VirtualserverNamePhonetic                           string  `json:"virtualserver_name_phonetic"`
	VirtualserverNeededIdentitySecurityLevel            int     `json:"virtualserver_needed_identity_security_level,string"`
	VirtualserverNickname                               string  `json:"virtualserver_nickname"`
	VirtualserverPassword                               string  `json:"virtualserver_password"`
	VirtualserverPlatform                               string  `json:"virtualserver_platform"`
	VirtualserverPort                                   int     `json:"virtualserver_port,string"`
	VirtualserverPrioritySpeakerDimmModificator         float64 `json:"virtualserver_priority_speaker_dimm_modificator,string"`
	VirtualserverQueryClientConnections                 int     `json:"virtualserver_query_client_connections,string"`
	VirtualserverQueryclientsonline                     int     `json:"virtualserver_queryclientsonline,string"`
	VirtualserverReservedSlots                          int     `json:"virtualserver_reserved_slots,string"`
	VirtualserverStatus                                 string  `json:"virtualserver_status"`
	VirtualserverTotalBytesDownloaded                   int     `json:"virtualserver_total_bytes_downloaded,string"`
	VirtualserverTotalBytesUploaded                     int     `json:"virtualserver_total_bytes_uploaded,string"`
	VirtualserverTotalPacketlossControl                 float64 `json:"virtualserver_total_packetloss_control,string"`
	VirtualserverTotalPacketlossKeepalive               float64 `json:"virtualserver_total_packetloss_keepalive,string"`
	VirtualserverTotalPacketlossSpeech                  float64 `json:"virtualserver_total_packetloss_speech,string"`
	VirtualserverTotalPacketlossTotal                   float64 `json:"virtualserver_total_packetloss_total,string"`
	VirtualserverTotalPing                              float64 `json:"virtualserver_total_ping,string"`
	VirtualserverUniqueIdentifier                       string  `json:"virtualserver_unique_identifier"`
	VirtualserverUploadQuota                            uint    `json:"virtualserver_upload_quota,string"`
	VirtualserverUptime                                 int     `json:"virtualserver_uptime,string"`
	VirtualserverVersion                                string  `json:"virtualserver_version"`
	VirtualserverWeblistEnabled                         int     `json:"virtualserver_weblist_enabled,string"`
	VirtualserverWelcomemessage                         string  `json:"virtualserver_welcomemessage"`
}

func (c *TS3HttpClient) ServerInfo(id int) (*ServerInfo, error) {
	var serverInfos []ServerInfo

	err := c.request(fmt.Sprintf("%d/serverinfo", id), &serverInfos)
	if err != nil {
		return nil, err
	}

	return &serverInfos[0], nil
}