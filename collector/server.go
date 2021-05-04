package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"ts3exporter/api"
)

var serverInfoLabels = []string{"virtualserver"}

type ServerInfo struct {
	api *api.TS3HttpClient

	ClientsOnline             *prometheus.Desc
	QueryClientsOnline        *prometheus.Desc
	Online                    *prometheus.Desc
	MaxClients                *prometheus.Desc
	Uptime                    *prometheus.Desc
	ChannelsOnline            *prometheus.Desc
	MaxDownloadTotalBandwidth *prometheus.Desc
	MaxUploadTotalBandwidth   *prometheus.Desc
	ClientsConnections        *prometheus.Desc
	QueryClientsConnections   *prometheus.Desc

	FileTransferBytesSentTotal     *prometheus.Desc
	FileTransferBytesReceivedTotal *prometheus.Desc

	ControlBytesSendTotal     *prometheus.Desc
	ControlBytesReceivedTotal *prometheus.Desc

	SpeechBytesSendTotal     *prometheus.Desc
	SpeechBytesReceivedTotal *prometheus.Desc

	KeepAliveBytesSendTotal     *prometheus.Desc
	KeepAliveBytesReceivedTotal *prometheus.Desc

	BytesSendTotal     *prometheus.Desc
	BytesReceivedTotal *prometheus.Desc
}

func NewServerInfo(api *api.TS3HttpClient) *ServerInfo {
	return &ServerInfo{
		api: api,

		ClientsOnline:      			prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "clients_online"), "number of currently online clients", serverInfoLabels, nil),
		QueryClientsOnline:             prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "query_clients_online"), "number of currently online query clients", serverInfoLabels, nil),
		Online:                         prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "online"), "is the virtual server online", serverInfoLabels, nil),
		MaxClients:                     prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "max_clients"), "maximal number of allowed clients", serverInfoLabels, nil),
		Uptime:                         prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "uptime"), "uptime of the virtual server", serverInfoLabels, nil),
		ChannelsOnline:                 prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "channels_online"), "number of online channels", serverInfoLabels, nil),
		MaxDownloadTotalBandwidth:      prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "download_bandwidth_bytes_max"), "maximal bandwidth available for downloads", serverInfoLabels, nil),
		MaxUploadTotalBandwidth:        prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "upload_bandwidth_bytes_max"), "maximal bandwidth available for uploads", serverInfoLabels, nil),
		ClientsConnections:             prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "client_connections"), "currently established client connections", serverInfoLabels, nil),
		QueryClientsConnections:        prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "query_client_connections"), "currently established query client connections", serverInfoLabels, nil),
		FileTransferBytesSentTotal:     prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "file_transfer_bytes_sent_total"), "total sent bytes for file transfers", serverInfoLabels, nil),
		FileTransferBytesReceivedTotal: prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "file_transfer_bytes_received_total"), "total received bytes for file transfers", serverInfoLabels, nil),
		ControlBytesSendTotal:          prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "control_bytes_sent_total"), "total sent bytes for control traffic", serverInfoLabels, nil),
		ControlBytesReceivedTotal:      prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "control_bytes_received_total"), "total received bytes for control traffic", serverInfoLabels, nil),
		SpeechBytesSendTotal:           prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "speech_bytes_sent_total"), "total sent bytes for speech traffic", serverInfoLabels, nil),
		SpeechBytesReceivedTotal:       prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "speech_bytes_received_total"), "total received bytes for speech traffic", serverInfoLabels, nil),
		KeepAliveBytesSendTotal:        prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "keepalive_bytes_sent_total"), "total send bytes for keepalive traffic", serverInfoLabels, nil),
		KeepAliveBytesReceivedTotal:    prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "keepalive_bytes_received_total"), "total received bytes for keepalive traffic", serverInfoLabels, nil),
		BytesSendTotal:                 prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "bytes_send_total"), "total send bytes", serverInfoLabels, nil),
		BytesReceivedTotal:             prometheus.NewDesc(prometheus.BuildFQName("ts3", "serverinfo", "bytes_received_total"), "total received bytes", serverInfoLabels, nil),
	}
}

func (i *ServerInfo) Describe(desc chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(i, desc)
}

func (i *ServerInfo) Collect(metric chan<- prometheus.Metric) {
	servers, err := i.api.ServerList()
	if err != nil {
		log.Printf("failed to list servers: %v", err)
	}

	for _, server := range servers {
		serverInfo, err := i.api.ServerInfo(server.VirtualserverID)
		if err != nil {
			log.Printf("failed to get server info: %v", err)
		}

		metric <- prometheus.MustNewConstMetric(i.ClientsOnline, prometheus.GaugeValue, float64(serverInfo.VirtualserverClientsonline), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.QueryClientsOnline, prometheus.GaugeValue, float64(serverInfo.VirtualserverQueryclientsonline), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.MaxClients, prometheus.GaugeValue, float64(serverInfo.VirtualserverMaxclients), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.Uptime, prometheus.CounterValue, float64(serverInfo.VirtualserverUptime), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.ChannelsOnline, prometheus.GaugeValue, float64(serverInfo.VirtualserverChannelsOnline), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.MaxDownloadTotalBandwidth, prometheus.GaugeValue, float64(serverInfo.VirtualserverMaxDownloadTotalBandwidth), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.MaxUploadTotalBandwidth, prometheus.GaugeValue, float64(serverInfo.VirtualserverMaxUploadTotalBandwidth), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.ClientsConnections, prometheus.GaugeValue, float64(serverInfo.VirtualserverClientConnections), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.QueryClientsConnections, prometheus.GaugeValue, float64(serverInfo.VirtualserverQueryClientConnections), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.FileTransferBytesSentTotal, prometheus.CounterValue, float64(serverInfo.ConnectionFiletransferBytesSentTotal), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.FileTransferBytesReceivedTotal, prometheus.CounterValue, float64(serverInfo.ConnectionFiletransferBytesReceivedTotal), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.ControlBytesSendTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesSentControl), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.ControlBytesReceivedTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesReceivedControl), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.SpeechBytesSendTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesSentSpeech), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.SpeechBytesReceivedTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesReceivedSpeech), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.KeepAliveBytesSendTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesSentKeepalive), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.KeepAliveBytesReceivedTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesReceivedKeepalive), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.BytesSendTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesSentTotal), server.VirtualserverName)
		metric <- prometheus.MustNewConstMetric(i.BytesReceivedTotal, prometheus.CounterValue, float64(serverInfo.ConnectionBytesReceivedTotal), server.VirtualserverName)

		var status float64
		if serverInfo.VirtualserverStatus == "online" {
			status = 1.0
		}
		metric <- prometheus.MustNewConstMetric(i.Online, prometheus.GaugeValue, status, server.VirtualserverName)
	}
}
