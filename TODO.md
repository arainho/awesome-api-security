# TODO's list

[] racepwn 
  - fix RUN lines for [racepwn](https://github.com/racepwn/racepwn)

[] restler-fuzzer
   - fix RUN lines for [restler-fuzzer](https://github.com/microsoft/restler-fuzzer)
   - or use oficial [Dockerfile](https://github.com/microsoft/restler-fuzzer/blob/main/docker/Dockerfile) in our multi-stage build
      
[] openapi_security_scanner
   - fix RUN lines for [openapi_security_scanner](https://github.com/ngalongc/openapi_security_scanner)


For adding new tools of fix the above use the following procedure:
0. git clone -b dev https://github.com/arainho/awesome-api-security   
1. Add your command(s) in the RUN below   
2. uncomment the lines in RUN       
3. docker build -t api-security-toolbox:local -f Dockerfile.testing 
 
