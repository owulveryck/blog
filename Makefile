deploy:
	cd ../owulveryck.github.io
	git pull
	cd ../blog
	hugo -t bootstrap -d ../owulveryck.github.io 
	cd ../owulveryck.github.io
	git add .
	git commit -m"Deploiement"
test:
	hugo server -t bootstrap --buildDrafts --watch
