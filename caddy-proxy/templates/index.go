package templates

type IndexData struct {
	OpenViduVersion        string
	LanMode                bool
	HttpUrl                string
	HttpsUrl               string
	WsUrl                  string
	WssUrl                 string
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
        <h1 class="display-4">Welcome to OpenVidu Local Deployment</h1>
        <p class="lead">OpenVidu Version: <strong>{{.OpenViduVersion}}</strong></p>
        <div class="alert alert-warning" role="alert">
            <span>This deployment is only for development purposes.</span>
        </div>
        <hr class="my-4">
        <h2>OpenVidu Server and LiveKit Server API:</h2>
        <ul>
            <li>From this machine:
                <ul>
                    <li><a href="{{.HttpUrl}}" target="_blank">{{.HttpUrl}}</a></li>
                    <li><a href="{{.WsUrl}}" target="_blank">{{.WsUrl}}</a></li>
                </ul>
            </li>
            <li>From other devices in your LAN:
                <ul>
                    <li><a href="{{.HttpsUrl}}" target="_blank">{{.HttpsUrl}}</a></li>
                    <li><a href="{{.WssUrl}}" target="_blank">{{.WssUrl}}</a></li>
                </ul>
            </li>
        </ul>
        <hr class="my-4">
        <h2>Services and passwords:</h2>
        <ul>
            <li><b>OpenVidu API:</b>
                <ul>
                    <li>Username: <code>OPENVIDUAPP</code></li>
                    <li>Password: <code>{{.OpenViduSecret}}</code></li>
                </ul>
            </li>
            <li>LiveKit API:
                <ul>
                    <li>API Key: <code>{{.LiveKitApiKey}}</code></li>
                    <li>API Secret: <code>{{.LiveKitApiSecret}}</code></li>
                </ul>
            </li>
            <li><a href="/minio-console" target="_blank">MinIO</a>
                <ul>
                    <li>Username: <code>{{.MinioAdminKey}}</code></li>
                    <li>Password: <code>{{.MinioAdminSecret}}</code></li>
                </ul>
            </li>
            <li><a href="/dashboard" target="_blank">OpenVidu Dashboard</a>
                <ul>
                    <li>Username: <code>{{.DashboardAdminUsername}}</code></li>
                    <li>Password: <code>{{.DashboardAdminPassword}}</code></li>
                </ul>
            </li>
            <li><a href="/openvidu-call" target="_blank">OpenVidu Call</a></li>
        </ul>
    </div>
</body>

</html>
`
