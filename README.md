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

## 3 - run vulnerable API's locally
On the `labs` folders has scripts to build and run vulnerable API's locally.
The purpose is to have local labs to exploit and learn API security.
