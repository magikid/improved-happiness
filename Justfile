build:
    GOOS=linux GOARCH=amd64 go build -o snake-${CI_COMMIT_SHORT_SHA} .

deploy: build
    scp snake.service /etc/systemd/system/snake.service
    ssh root@45.32.186.123 'systemctl daemon-reload'
    scp snake-${CI_COMMIT_SHORT_SHA} root@45.32.186.123:/root/snakes/
    ssh root@45.32.186.123 'systemctl stop snake.service'
    ssh root@45.32.186.123 ln -fs /root/snakes/snake-${CI_COMMIT_SHORT_SHA} /root/snakes/snake
    ssh root@45.32.186.123 'systemctl start snake.service'
