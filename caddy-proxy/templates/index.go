package templates

type IndexData struct {
	OpenViduVersion        string
	LanMode                bool
	HttpUrl                string
	HttpsUrl               string
	DashboardAdminUsername string
	DashboardAdminPassword string
	MinioAdminKey          string
	MinioAdminSecret       string
	LiveKitApiKey          string
	LiveKitApiSecret       string
	OpenViduSecret         string
}

const IndexTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OpenVidu Local Deployment</title>
    <!-- Bootstrap CSS CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" />
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
            {{- if .HttpsUrl }}

            <h2>HTTPS URLs</h2>
            {{- if .LanMode }}
            <div class="alert alert-info" role="alert">
                <span>You can access from any device in your local network using the following URLs:</span>
            </div>
            {{- end}}
            <ul>
                <li><strong>OpenVidu and Livekit API: </strong><a href="{{ .HttpsUrl }}"
                        target="_blank">{{ .HttpsUrl }}</a></li>
                <li><strong>OpenVidu Dashboard: </strong><a href="{{ .HttpsUrl }}/dashboard"
                        target="_blank">{{ .HttpsUrl }}/dashboard</a></li>
                <li><strong>Minio Console: </strong><a href="{{ .HttpsUrl }}/minio-console"
                        target="_blank">{{ .HttpsUrl }}/minio-console</a></li>
                <li><strong>OpenVidu Call: </strong><a href="{{ .HttpsUrl }}/openvidu-call"
                        target="_blank">{{ .HttpsUrl }}/openvidu-call</a></li>
                <li><strong>Your App: </strong><span>Any App you deploy at port 5442 will be available here: </span>
                    <ul>
                        <li><a href="{{ .HttpsUrl }}"
                                target="_blank">{{ .HttpsUrl }}</a></li>
                    </ul>
                </li>
            </ul>
            <hr class="my-4">

            {{- end }}
            <h2>HTTP URLs</h2>
            <ul>
                <li><strong>OpenVidu and Livekit API: </strong><a href="{{ .HttpUrl }}"
                        target="_blank">{{ .HttpUrl }}</a></li>
                <li><strong>OpenVidu Dashboard: </strong><a href="{{ .HttpUrl }}/dashboard"
                        target="_blank">{{ .HttpUrl }}/dashboard</a></li>
                <li><strong>Minio Console: </strong><a href="{{ .HttpUrl }}/minio-console"
                        target="_blank">{{ .HttpUrl }}/minio-console</a></li>
                {{- if not .HttpsUrl }}
                <li><strong>OpenVidu Call: </strong><a href="{{ .HttpUrl }}/openvidu-call"
                        target="_blank">{{ .HttpUrl }}/openvidu-call</a></li>
                {{- end }}
            </ul>
            <hr class="my-4">
            <!-- Section with Credentials -->
            <h2>Credentials</h2>
            <ul>
                <li><strong>OpenVidu (Basic auth):</strong>
                    <ul>
                        <li>Username: <code>OPENVIDUAPP</code></li>
                        <li>Password: <code>{{ .OpenViduSecret }}</code></li>
                    </ul>
                </li>
                <li><strong>LiveKit API:</strong>
                    <ul>
                        <li>API Key: <code>{{ .LiveKitApiKey }}</code></li>
                        <li>API Secret: <code>{{ .LiveKitApiSecret }}</code></li>
                    </ul>
                </li>
                <li><strong>OpenVidu Dashboard: </strong>
                    <ul>
                        <li>Username: <code>{{ .DashboardAdminUsername }}</code></li>
                        <li>Password: <code>{{ .DashboardAdminPassword }}</code></li>
                    </ul>
                </li>
                <li><strong>Minio: </strong>
                    <ul>
                        <li>User: <code>{{ .MinioAdminKey }}</code></li>
                        <li>Password: <code>{{ .MinioAdminSecret }}</code></li>
                    </ul>
                </li>
            </ul>
        </div>
</body>
</html>
`
