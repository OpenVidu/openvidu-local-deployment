package templates

type IndexData struct {
	OpenViduVersion        string
	DomainName             string
	IsLocalhost            bool
	IsTLS                  bool
	DashboardAdminUsername string
	DashboardAdminPassword string
	MinioAdminKey          string
	MinioAdminSecret       string
	LiveKitApiKey          string
	LiveKitApiSecret       string
	OpenViduSecret         string
}

const IndexTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OpenVidu Local Deployment</title>
    <!-- Bootstrap CSS CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <!-- Custom styles -->
    <style>
        .container {
            padding-top: 50px;
            padding-left: 100px;
            padding-right: 100px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div>
            <h1 class="display-4">Welcome to OpenVidu Local Deployment</h1>
            <p class="lead">OpenVidu Version: <strong>{{ .OpenViduVersion }}</strong></p>
            <div class="alert alert-warning" role="alert">
                <span>This deployment is only for development purposes.</span>
            </div>
            <hr class="my-4">
            <ul>
                <li><strong>OpenVidu API URL:</strong> <a href="http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443" target="_blank">http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443</a>
                    <ul>
                        <li><strong>Secret:</strong> <code>{{ .OpenViduSecret }}</code></li>
                    </ul>
                </li>
				<li><strong>LiveKit API URL:</strong> <a href="http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443" target="_blank">http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443</a>
                    <ul>
                        <li><strong>API Key:</strong> <code>{{ .LiveKitApiKey }}</code></li>
                        <li><strong>API Secret:</strong> <code>{{ .LiveKitApiSecret }}</code></li>
                    </ul>
                </li>
				<li><strong>OpenVidu Dashboard:</strong> <a href="http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443/dashboard" target="_blank">http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443/dashboard</a>
                    <ul>
                        <li><strong>Username:</strong> <code>{{ .DashboardAdminUsername }}</code></li>
                        <li><strong>Password:</strong> <code>{{ .DashboardAdminPassword }}</code></li>
                    </ul>
                </li>
				<li><strong>Minio:</strong> <a href="http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443/minio-console" target="_blank">http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443/minio-console</a>
                    <ul>
                        <li><strong>Minio Admin User:</strong> <code>{{ .MinioAdminKey }}</code></li>
                        <li><strong>Minio Admin Password:</strong> <code>{{ .MinioAdminSecret }}</code></li>
                    </ul>
                </li>
                <li><strong>OpenVidu Call (Default App):</strong> <a href="http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443/openvidu-call" target="_blank">http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:4443/openvidu-call</a></li>
                <hr class="my-4">
                <li><strong>Your App:</strong> <a href="http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:8000" target="_blank">http{{if .IsTLS}}s{{end}}://{{ .DomainName }}:8000</a>:
                <span>If you are developing an application and run it locally at port 5442, you will see your application here, under the same domain and TLS certificate as OpenVidu.</span></li>
            </ul>
            <hr class="my-4">
            {{- if not .IsLocalhost }}
                <p>If you want to access this deployment with <code>http(s)://localhost:4443</code>, just change the <code>LOCAL_DOMAIN</code> variable to <code>localhost</code> in the <code>.env</code> file.</p>
            {{- else }}
                <p>If you want to access this deployment with <code>http(s)://openvidu-local.dev:4443</code>, just change the <code>LOCAL_DOMAIN</code> variable to <code>openvidu-local.dev</code> in the <code>.env</code> file.</p>
            {{- end }}
            {{- if .IsTLS }}
                <p>If you want to disable TLS, just change the <code>USE_TLS</code> variable to <code>false</code> in the <code>.env</code> file.</p>
            {{- else }}
                <p>If you want to enable TLS, just change the <code>USE_TLS</code> variable to <code>true</code> in the <code>.env</code> file.</p>
            {{- end }}
        </div>
    </div>
</body>
</html>`
