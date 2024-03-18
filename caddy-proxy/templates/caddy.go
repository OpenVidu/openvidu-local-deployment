package templates

type CaddyData struct {
	DomainName  string
	IsLocalhost bool
	IsTLS       bool
}

const CaddyfileTemplate = `
# Minio
{{- if hasSuffix .DomainName ".openvidu-local.dev" }}
http{{if .IsTLS}}s{{end}}://*.openvidu-local.dev:9000, http{{if .IsTLS}}s{{end}}://openvidu-local.dev:9000 {
{{- else }}
http{{if .IsTLS}}s{{end}}://{{.DomainName}}:9000 {
{{- end }}
	{{if .IsTLS}}{{if hasSuffix .DomainName "openvidu-local.dev"}}tls internal {
		get_certificate http https://certs.openvidu-local.dev/caddy.pem
	}{{else}}tls internal{{end}}{{end}}
	reverse_proxy http://minio:9000
}

# General
{{- if hasSuffix .DomainName ".openvidu-local.dev" }}
http{{if .IsTLS}}s{{end}}://*.openvidu-local.dev:4443, http{{if .IsTLS}}s{{end}}://openvidu-local.dev:4443 {
{{- else }}
http{{if .IsTLS}}s{{end}}://{{.DomainName}}:4443 {
{{- end }}
	{{if .IsTLS}}{{if hasSuffix .DomainName "openvidu-local.dev"}}tls internal {
		get_certificate http https://certs.openvidu-local.dev/caddy.pem
	}{{else}}tls internal{{end}}{{end}}

	# Api
	@openvidu path /twirp/* /rtc/* /rtc
	handle @openvidu {
		reverse_proxy http://openvidu:7880
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

	# Default /
	handle_path /* {
		root * /var/www/
		file_server
	}
}

# Your OpenVidu App
{{- if hasSuffix .DomainName ".openvidu-local.dev" }}
http{{if .IsTLS}}s{{end}}://*.openvidu-local.dev:8000, http{{if .IsTLS}}s{{end}}://openvidu-local.dev:8000 {
{{- else }}
http{{if .IsTLS}}s{{end}}://{{.DomainName}}:8000 {
{{- end }}
	{{if .IsTLS}}{{if hasSuffix .DomainName "openvidu-local.dev"}}tls internal {
		get_certificate http https://certs.openvidu-local.dev/caddy.pem
	}{{else}}tls internal{{end}}{{end}}
	handle_errors {
		@502 expression {http.error.status_code} == 502
		rewrite @502 /app502.html
		file_server {
			root /var/www
		}
	}
	reverse_proxy http://host.docker.internal:5442
}

`
