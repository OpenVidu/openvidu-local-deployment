# OpenVidu Local Deployment - Cadddy Proxy

If you want to modify any of the rules at the caddy-proxy container, just build the image again and run the local deployment with the new image.

```bash
docker build --build-arg VERSION=custom -t caddy-proxy .
```
