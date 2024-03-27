image_name="web"

docker build . -t $image_name
if [ $? -eq 0 ]; then
	docker run -d -p 81:80 $image_name
	docker run -d -p 82:80 $image_name
	docker run -d -p 83:80 $image_name
	docker run -d -p 84:80 $image_name
	docker run -d -p 85:80 $image_name
else
	echo "error on build image"
fi
docker ps
