package templates

type CaddyData struct {
	LanMode   bool
	LanDomain string
	HttpUrl   string
	HttpsUrl  string
	HttpPort  string
	HttpsPort string
	AppPort   string
}

const CaddyfileTemplate = `
(index) {
	# Default /
	handle_path /* {
		root * /var/www/
		file_server
	}
}
(general_rules) {
	# LiveKit API
	@openvidu path /twirp/* /rtc/* /rtc
	handle @openvidu {
		reverse_proxy http://openvidu:7880
	}

	# OpenVidu v2 API
	@openvidu_v2 path /openvidu/api/* /openvidu/ws/*
	handle @openvidu_v2 {
		reverse_proxy http://openvidu-v2compatibility:4443
	}

	# Minio console
	redir /minio-console /minio-console/
	handle_path /minio-console/* {
		uri strip_prefix /minio-console
		reverse_proxy http://minio:9001
	}

	# OpenVidu Dashboard
	redir /dashboard /dashboard/
	handle_path /dashboard/* {
		rewrite * {path}
		reverse_proxy http://dashboard:5000
	}

	# OpenVidu Call (Default App)
	redir /openvidu-call /openvidu-call/
	handle_path /openvidu-call/* {
		rewrite * {path}
		reverse_proxy http://default-app:5442
	}

}
(custom_app) {
	handle_errors {
		@502 expression {http.error.status_code} == 502
		rewrite @502 /app502.html
		file_server {
			root /var/www
		}
	}
	reverse_proxy http://host.docker.internal:{{ .AppPort }}
}

# Servers
:{{.HttpPort}} {
	import general_rules
	import index
}

{{- if .HttpsUrl }}

{{- if .LanMode }}
{{.HttpsUrl}} {
	{{- if hasSuffix .LanDomain ".openvidu-local.dev" }}
	tls internal {
		get_certificate http https://certs.openvidu-local.dev/caddy.pem
	}
	{{- else }}
	tls internal
	{{- end }}
	import general_rules
	import custom_app
}
{{- else }}
https://*:{{.HttpsPort}} {
	tls internal
	import general_rules
	import custom_app
}
{{- end }}

{{- end}}
`
