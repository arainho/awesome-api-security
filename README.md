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

## 5. collaboration
For adding new tools of fix broken ones from [ToDo list](TODO.md) use the following procedure:
0. Clone the repository
   ```bash
   git clone -b dev https://github.com/arainho/awesome-api-security
   git checkout -b tool-xyz
   ```
2. Open the Dockerfile.testing file
   - Add your command(s) in the RUN below   
   - uncomment the lines in RUN       
5. Build the image
   ```bash
   docker build -t api-security-toolbox:local -f Dockerfile.testing 
   ```
6. If everything looks good, create a pull request
   ```bash
   git add Dockerfile.testing
   git commit -m "new entry for toll-xyz"
   git push origin tool-xyz
   ```
   
   you can check more information on creating a pull request [here](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request)
   
 7. A maintainer will review the pull request
    - manual review
    - manualy merge to dev
    - add extra lines on Dockerfile and Dockerfile.multistage
    - github actions workflow will run
    - If all looks good your PR will pass 😃
