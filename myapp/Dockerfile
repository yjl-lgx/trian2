FROM centos
COPY . ./
EXPOSE 8080 50051
ENV OEM=myapp 
ENTRYPOINT ["./mainsevr", "-stderrthreshold=info"]