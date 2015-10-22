deploy:
	hugo -t uno -d ../owulveryck.github.io 
test:
	hugo server -t uno --buildDrafts --watch
