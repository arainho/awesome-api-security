# Awesome API Security 

**WARNING: This is a development branch, use it at your own risk !!!**

This branch is dedicated to experimenting around automatic and manual flows related to security testing for APIs.

## 1 - workdlists bundle
Collect several wordlists and build a bundle.

## 2 - build docker image for api-security-toolbox
The [api-security-toolbox](https://hub.docker.com/r/arainho/api-security-toolbox) is a swiss knife for API security testing. 
This image is based on `Dockerfile`.

The `Dockerfile.multistage` is a different version of api-security-toolbox with:
- multi-stage build to save disk space
- tools and utilities installed as normal user appuser
- linux os utilities installed as root

Please note that the docker images (Dockerfile, Dockerfile.multistage) are in alpha stage.

## 3. run the api-security-toolbox
```bash
docker run -it --rm arainho/api-security-toolbox /bin/bash
```

If you need apicheck tools inside the toolbox you need docker unix socket inside the container
```bash
docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock arainho/api-security-toolbox /bin/bash
```

Then you can install [apicheck](https://bbva.github.io/apicheck/docs) tools
```bash
acp install jwt-checker
acp install acurl
acp install oas-checker
acp install send-to-proxy
acp install apicheck-curl
acp install sensitive-data
acp install replay
acp install openapiv3-lint
acp install openapiv2-lint
acp install oas-checker
```

**warning: Using docker.sock could expose your host within the toolbox container** as stated in this [article](https://www.ctl.io/developers/blog/post/tutorial-understanding-the-security-risks-of-running-docker-containers).

## 4 - run vulnerable API's locally
On the [`labs`](./labs) folders has scripts to build and run vulnerable API's locally.  
The purpose is to have local labs to exploit and learn API security.
