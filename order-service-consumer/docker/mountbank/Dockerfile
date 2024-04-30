# Use an official Node.js runtime as a parent image
FROM node:latest

ENV MOUNTEBANK_VERSION=2.8.2
ENV MOUNTEBANK_GRPC_VERSION=0.2.5

# Set the working directory in the container
WORKDIR /usr/src/app

# Install Mountebank globally via npm
RUN npm install --location=global mountebank@${MOUNTEBANK_VERSION} --omit=dev
RUN npm install --location=global @ozonru/mountebank-grpc@${MOUNTEBANK_GRPC_VERSION} --omit=dev

# Expose the port Mountebank will run on
EXPOSE 2525 8090 9090

ENTRYPOINT ["mb"]
# Define the command to run Mountebank when the container starts
CMD ["start", "--protofile" , "/mountbank/protocols.json", "--configfile","/mountebank/imposters.ejs", "--allowInjection"]