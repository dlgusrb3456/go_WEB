FROM golang:1.20

# Set destination for COPY
WORKDIR ./src
# Download Go modules, copy all sub files in specific folder.
COPY app/ ./app
COPY model/ ./model
COPY public/ ./public

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY secret.go ./app
RUN go mod download
RUN go get -u github.com/dlgusrb3456/get_UUID

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# Build, to use sqlite3 CGO_ENABLED=1
RUN CGO_ENABLED=1 GOOS=linux go build -o todo main.go 

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 3000

# Run
CMD [ "./todo" ]