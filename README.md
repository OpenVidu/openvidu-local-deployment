# OpenVidu Local Deployment
Docker compose to run OpenVidu locally for development purposes

## Requirements
On **Windows** and **MacOS**:
- **Docker Desktop**

On **Linux**:
- **Docker**
- **Docker Compose**

## Install

### Windows

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment
.\configure_lan_private_ip_windows.bat
```

### Mac

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment
./configure_lan_private_ip_mac.sh
```

### Linux

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment
./configure_lan_private_ip_linux.sh
```

## Run OpenVidu

```sh
docker compose up
```
