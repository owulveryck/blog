deploy:
	hugo -d s3/
test:
	hugo server --buildDrafts --watch
