# OpenVidu Local Deployment
Docker compose to run OpenVidu locally for development purposes

## Requirements
On **Windows** and **MacOS**:
- **Docker Desktop**

On **Linux**:
- **Docker**
- **Docker Compose**

## How to run

```sh
git clone https://github.com/OpenVidu/openvidu-local-deployment
cd openvidu-local-deployment
docker compose up
```

When the deployment is ready you will see the following message in the logs:

```
readycheck     | ------------------------
readycheck     | OpenVidu is ready!
readycheck     | Open http://localhost:4443/ in your browser
readycheck     | ------------------------
```

## Additional Notes

### LAN Access (Optional)

If you want to access the deployment in your LAN for Android or iOS devices, you need to do the following:

1. Get the private IP of your computer in your LAN.
2. Configure your Firewall to allow devices in your LAN to access your computer.
3. Change `LOCAL_DOMAIN` in the `.env` file to have the IP of your computer in your LAN.

    If your IP for example is `192.168.1.10`, `LOCAL_DOMAIN` should be `192-168-1-10.openvidu-local.dev`.

### About `openvidu-local.dev`

When you develop WebRTC applications, you require a secure context (HTTPS) to access the camera and microphone. This is a requirement of the WebRTC standard.

With the aim of making it easier to develop with OpenVidu, we provide a magic domain name `openvidu-local.dev` which can resolve to any IP specified as a subdomain and it offers a valid wildcard certificate for HTTPS. It is similar to [nip.io](https://nip.io) or [traefik.me](https://traefik.me) or [localtls](https://github.com/Corollarium/localtls).

But take into account that this is just as secure as a HTTP connection, so it is not suitable for production environments.

### Edge cases:

- Linux: All works just fine
- Windows (Docker desktop):
    It looks like there is a little edge case which we are fighting with WSL + Docker. Looks related with this: https://stackoverflow.com/questions/61629450/webrtc-does-not-work-in-the-modern-browsers-when-a-peer-is-not-behind-a-nat-beca

    The behaviour is the following
    - **Chrome based browsers**: Looks like everything works fine. ICE protocol finds a path to communicate the browser and openvidu.
    - **Firefox browser**:
     The only working candidate in Firefox is filtered, I was able to workaround this limitation with `media.peerconnection.ice.obfuscate_host_addresses=false`.
