deploy:
	cd ../owulveryck.github.io
	git pull
	cd ../blog
	hugo -d ../owulveryck.github.io 
	cd ../owulveryck.github.io
	git add .
	git commit -m"Deploiement"
test:
	hugo server --buildDrafts --watch
