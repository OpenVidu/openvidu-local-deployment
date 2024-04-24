package main

import (
	"bytes"
	"fmt"
	"local-caddy-generate/templates"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type TemplateData any

var indexData = &templates.IndexData{}
var caddyData = &templates.CaddyData{}
var app502ClientData = &templates.App502Data{}
var app502ServerData = &templates.App502Data{}

func main() {
	err := Initialize()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawIndex, err := GenerateTemplate(templates.IndexTemplate, indexData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = WriteStringToFile("index.html", rawIndex)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawCaddyfile, err := GenerateTemplate(templates.CaddyfileTemplate, caddyData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = WriteStringToFile("Caddyfile", rawCaddyfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawAppClient502, err := GenerateTemplate(templates.App502Template, app502ClientData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = WriteStringToFile("app502client.html", rawAppClient502)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawAppServer502, err := GenerateTemplate(templates.App502Template, app502ServerData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = WriteStringToFile("app502server.html", rawAppServer502)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func Initialize() error {
	// OpenVidu && LiveKit API
	httpPort := 7880
	httpsPort := 7443

	// Http ports
	appClientPort := 5080
	appClientServer := 6080

	// Https ports
	httpsAppClientPort := 5443
	httpsAppServerPort := 6443

	version := os.Getenv("VERSION")
	if version == "" {
		return fmt.Errorf("VERSION is not set")
	}

	rawUseHTTPS := os.Getenv("USE_HTTPS")
	if rawUseHTTPS == "" {
		rawUseHTTPS = "false"
	}
	useTLS, err := strconv.ParseBool(rawUseHTTPS)
	if err != nil {
		return fmt.Errorf("USE_HTTPS is not a boolean")
	}

	lanMode := os.Getenv("LAN_MODE")
	if lanMode == "" {
		lanMode = "false"
	}

	lanPrivateIP := os.Getenv("LAN_PRIVATE_IP")
	if lanPrivateIP == "" {
		return fmt.Errorf("LAN_PRIVATE_IP is not set")
	}

	lanDomain := os.Getenv("LAN_DOMAIN")
	if lanDomain == "" {
		lanDomain = "openvidu-local.dev"
	}

	if lanPrivateIP != "none" && lanDomain == "openvidu-local.dev" {
		ipDashes := strings.ReplaceAll(lanPrivateIP, ".", "-")
		lanDomain = fmt.Sprintf("%s.%s", ipDashes, lanDomain)
	}

	httpUrl := fmt.Sprintf("http://localhost:%d", httpPort)
	httpsUrl := ""
	wsUrl := fmt.Sprintf("ws://localhost:%d", httpPort)
	wssUrl := ""
	httpsAppClientUrl := ""
	httpsAppServerUrl := ""
	if useTLS {
		httpsUrl = fmt.Sprintf("https://localhost:%d", httpsPort)
		wssUrl = fmt.Sprintf("wss://localhost:%d", httpsPort)
		httpsAppClientUrl = fmt.Sprintf("https://localhost:%d", httpsAppClientPort)
		httpsAppServerUrl = fmt.Sprintf("https://localhost:%d", httpsAppServerPort)
		if lanMode == "true" {
			httpsUrl = fmt.Sprintf("https://%s:%d", lanDomain, httpsPort)
			wssUrl = fmt.Sprintf("wss://%s:%d", lanDomain, httpsPort)
			httpsAppClientUrl = fmt.Sprintf("https://%s:%d", lanDomain, httpsAppClientPort)
			httpsAppServerUrl = fmt.Sprintf("https://%s:%d", lanDomain, httpsAppServerPort)
		}
	}

	livekitApiKey := os.Getenv("LIVEKIT_API_KEY")
	if livekitApiKey == "" {
		return fmt.Errorf("LIVEKIT_API_KEY is not set")
	}
	livekitApiSecret := os.Getenv("LIVEKIT_API_SECRET")
	if livekitApiSecret == "" {
		return fmt.Errorf("LIVEKIT_API_SECRET is not set")
	}
	openviduSecret := os.Getenv("OPENVIDU_SHIM_SECRET")
	if openviduSecret == "" {
		return fmt.Errorf("OPENVIDU_SHIM_SECRET is not set")
	}
	dashboadAdminUsername := os.Getenv("DASHBOARD_ADMIN_USERNAME")
	if dashboadAdminUsername == "" {
		return fmt.Errorf("DASHBOARD_ADMIN_USERNAME is not set")
	}
	dashboardAdminPassword := os.Getenv("DASHBOARD_ADMIN_PASSWORD")
	if dashboardAdminPassword == "" {
		return fmt.Errorf("DASHBOARD_ADMIN_PASSWORD is not set")
	}
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	if minioAccessKey == "" {
		return fmt.Errorf("MINIO_ACCESS_KEY is not set")
	}
	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	if minioSecretKey == "" {
		return fmt.Errorf("MINIO_SECRET_KEY is not set")
	}

	indexData = &templates.IndexData{
		OpenViduVersion:        version,
		LanMode:                lanMode == "true",
		HttpUrl:                httpUrl,
		HttpsUrl:               httpsUrl,
		WsUrl:                  wsUrl,
		WssUrl:                 wssUrl,
		LiveKitApiKey:          livekitApiKey,
		LiveKitApiSecret:       livekitApiSecret,
		OpenViduSecret:         openviduSecret,
		DashboardAdminUsername: dashboadAdminUsername,
		DashboardAdminPassword: dashboardAdminPassword,
		MinioAdminKey:          minioAccessKey,
		MinioAdminSecret:       minioSecretKey,
	}

	caddyData = &templates.CaddyData{
		LanMode:   lanMode == "true",
		LanDomain: lanDomain,
		// Main OpenVidu and LiveKit API ports
		HttpPort:  strconv.Itoa(httpPort),
		HttpsPort: strconv.Itoa(httpsPort),
		// Main OpenVidu and LiveKit API URLs
		HttpUrl:  httpUrl,
		HttpsUrl: httpsUrl,
		// Tutorial ports
		AppClientPort:      strconv.Itoa(appClientPort),
		AppServerPort:      strconv.Itoa(appClientServer),
		HttpsAppClientPort: strconv.Itoa(httpsAppClientPort),
		HttpsAppServerPort: strconv.Itoa(httpsAppServerPort),
		// Tutorial URLs
		HttpsAppClientUrl: httpsAppClientUrl,
		HttpsAppServerUrl: httpsAppServerUrl,
	}

	app502ClientData = &templates.App502Data{
		Title:   "Application Client Not Found",
		Message: fmt.Sprintf("Run your Application Client at port <b>%d</b> and you will see it here", appClientPort),
	}

	app502ServerData = &templates.App502Data{
		Title:   "Application Server Not Found",
		Message: fmt.Sprintf("Run your Application Server at port <b>%d</b> and you will see it here", appClientServer),
	}

	return nil

}

func GenerateTemplate(templateString string, data TemplateData) (string, error) {
	funcs := map[string]any{
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix}

	tmpl, err := template.New("template").Funcs(funcs).Parse(templateString)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WriteStringToFile(filePath string, data string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
