FROM centos
COPY . ./
EXPOSE 8080 50051
ENV OEM=myapp 
ENV VER=3
ENTRYPOINT ["./mainsevr", "-stderrthreshold=info"]