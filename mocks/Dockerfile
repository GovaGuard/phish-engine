FROM mokapi/mokapi:0.18.1

RUN apk add openssl

RUN mkdir demo

RUN openssl genrsa -out /demo/ca.key 4096 
RUN openssl req -x509 -new -nodes -key /demo/ca.key -sha256 -days 3650 \
        -out /demo/ca.crt \
        -subj '/CN=localhost/C=AT/ST=Vienna/L=Vienna/O=MyOrg'
        
COPY ./smtp.yaml /demo/

CMD ["--Providers.File.Directory=/demo"]
