# OpenVidu Local Deployment
Docker compose to run OpenVidu locally for development purposes

## Requirements
On **Windows** and **MacOS**:
- **Docker Desktop**

On **Linux**:
- **Docker**
- **Docker Compose**

---

## OpenVidu COMMUNITY

### Install OpenVidu COMMUNITY

#### Windows

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment/community
.\configure_lan_private_ip_windows.bat
```

#### Mac

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment/community
./configure_lan_private_ip_mac.sh
```

#### Linux

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment/community
./configure_lan_private_ip_linux.sh
```

### Run OpenVidu COMMUNITY

```sh
docker compose up
```

---

## OpenVidu PRO

### Install OpenVidu PRO

#### Windows

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment/pro
.\configure_lan_private_ip_windows.bat
```

#### Mac

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment/pro
./configure_lan_private_ip_mac.sh
```

#### Linux

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment/pro
./configure_lan_private_ip_linux.sh
```

### Run OpenVidu PRO

```sh
docker compose up
```