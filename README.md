Start docker

http://localhost:4646/ui/jobs


curl http://127.0.0.1:4646/v1/nodes

sudo ln -s $HOME/.docker/run/docker.sock /var/run/docker.sock

nomad agent -dev -bind 0.0.0.0 -config nomad.hcl


curl -X PUT localhost:3000/services/mypage -d '{"url": "https://pastebin.com/raw/vDzvUbJS", "script": true}'

curl localhost:3000/services/mypage -X POST -d '{"url": "https://pastebin.com/raw/xp02ittK", "script": false}'


nomad node status -verbose
