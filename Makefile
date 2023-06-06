docker-build:
	docker build -t jokenpo .
docker-run: 
	docker run -d -p 8000:8000 jokenpo
run: docker-build docker-run