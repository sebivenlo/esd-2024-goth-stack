FROM golang:1.23

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

ENV PATH="/go/bin:${PATH}"

WORKDIR /workspace
COPY . /workspace

EXPOSE 8080