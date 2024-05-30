# OpenVidu Local Deployment
Docker Compose files to run OpenVidu locally for development purposes.

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

## OpenVidu PRO (Evaluation Mode)

> OpenVidu PRO can be executed locally in evaluation mode for free for development and testing purposes.
> Some limits apply: maximum 2 concurrent Rooms, 8 Participants per Room, 5 minutes duration per Room.

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
