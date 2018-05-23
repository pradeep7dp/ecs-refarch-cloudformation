# Start from a small base
FROM golang:onbuild

# Our application requires no privileges
# so run it with a non-root user
ADD users /etc/passwd
USER nobody

# Our application runs on port 8001
# so allow hosts to bind to that port
EXPOSE 8001

# Add our application binary
ADD ../ /app

# Run our application!
ENTRYPOINT [ "/app" ]
