# build from scratch
FROM golang:1.19

# workign dir
WORKDIR /test

# before anything else add the dirs

# add the mod dir

#         hacky work around bull
ADD mod /test
RUN mkdir mod
# icba understanding why tf it just doesn't want to play ball when trying to add the dirs
# so for now this is what we will do, yey :)
RUN mv get  mod


# copy the templates
COPY templates /test
RUN mkdir templates
RUN mv html templates
RUN mv css templates

# fuckin do all the other stuff
COPY go.* /test
COPY go.mod /test
# copy the go files
COPY *.go /test



# make sure all dependecnies are installed
#RUN go mod download

# compile the program
RUN go build
#RUN CGO_ENABLED=0 GOOS=linux go build -o /main

#COPY main ./
RUN ls -a
# expose the port 8080 to the network
EXPOSE 8080

# run the program
ENTRYPOINT ["/test/webz"]

# RUNNIG
# BUILD -> sudo docker build --tag *name* .
# RUN   -> sudo docker run --publish 8080:8080 *name*