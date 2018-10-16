# Kube Test

Test project for learning Kubernetes

    docker build . -t kube-test
    docker run -p 11011:11011 -d kube-test

Then test the service by:

    curl localhost:11011/ping

If you got the response `{"message":"pong"}`, the server is successfully launched!