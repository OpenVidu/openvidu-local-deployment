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
docker compose up
```

When the deployment is ready you will see the following message in the logs:

```
readycheck     | ------------------------
readycheck     | OpenVidu is ready!
readycheck     | Open https://openvidu-local.dev:4443/ in your browser
readycheck     | ------------------------
```

## Additional Notes

### Using localhost

- This deployment is configured to use a domain name which points to `127.0.0.1` by default. If you want to use `localhost` instead, you can change the `LOCAL_DOMAIN` variable in the `.env` file.

### Enabling and Disabling TLS
- You can enable and disable TLS by setting `USE_TLS` to `true` or `false` in the `.env` file.

### LAN Access (Optional)

If you want to access the deployment in your LAN for Android or iOS devices, you need to do the following:

1. Get the private IP of your computer in your LAN.
2. Configure your Firewall to allow devices in your LAN to access your computer.

If your IP for example is `192.168.1.10`, the URL of your deployment in your LAN will be `https://192-168-1-10.openvidu-local.dev:4443/`.

### About `openvidu-local.dev`

This is a magic domain name like [nip.io](https://nip.io) or [traefik.me](https://traefik.me), which can resolve to any IP specified as a subdomain. It also offers a wildcard certificates which is automatically used by `caddy-proxy` in the local deployment to provide HTTPS for any subdomain.

This is useful for local development, as you can access your deployment using a domain name instead of an IP address, and you can use HTTPS without having to deal with self-signed certificates, **BUT it is not suitable for production environments.**

### Edge cases:

- Linux: All works just fine
- Windows (Docker desktop):
    It looks like there is a little edge case which we are fighting with WSL + Docker. Looks related with this: https://stackoverflow.com/questions/61629450/webrtc-does-not-work-in-the-modern-browsers-when-a-peer-is-not-behind-a-nat-beca

    The behaviour is the following
    - **Chrome based browsers**: Looks like everything works fine. ICE protocol finds a path to communicate the browser and openvidu.
    - **Firefox browser**:
     The only working candidate in Firefox is filtered, I was able to workaround this limitation with `media.peerconnection.ice.obfuscate_host_addresses=false`.
