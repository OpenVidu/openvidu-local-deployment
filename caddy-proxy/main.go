package main

import (
	"bytes"
	"fmt"
	"html/template"
	"local-caddy-generate/templates"
	"os"
	"strconv"
	"strings"
)

type TemplateData any

var indexData = &templates.IndexData{}
var caddyData = &templates.CaddyData{}

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

	rawApp502, err := GenerateTemplate(templates.App502Template, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = WriteStringToFile("app502.html", rawApp502)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func Initialize() error {
	version := os.Getenv("VERSION")
	if version == "" {
		return fmt.Errorf("VERSION is not set")
	}

	localDomain := os.Getenv("LOCAL_DOMAIN")
	if localDomain == "" {
		return fmt.Errorf("LOCAL_DOMAIN is not set")
	}

	rawUseTLS := os.Getenv("USE_TLS")
	if rawUseTLS == "" {
		return fmt.Errorf("USE_TLS is not set")
	}
	useTLS, err := strconv.ParseBool(rawUseTLS)
	if err != nil {
		return fmt.Errorf("USE_TLS is not a boolean")
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
		DomainName:             localDomain,
		IsLocalhost:            localDomain == "localhost",
		IsTLS:                  useTLS,
		LiveKitApiKey:          livekitApiKey,
		LiveKitApiSecret:       livekitApiSecret,
		OpenViduSecret:         openviduSecret,
		DashboardAdminUsername: dashboadAdminUsername,
		DashboardAdminPassword: dashboardAdminPassword,
		MinioAdminKey:          minioAccessKey,
		MinioAdminSecret:       minioSecretKey,
	}

	caddyData = &templates.CaddyData{
		DomainName:  localDomain,
		IsLocalhost: localDomain == "localhost",
		IsTLS:       useTLS,
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
