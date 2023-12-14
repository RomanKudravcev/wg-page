include ./backend/.env
include ./frontend/.env

build-backend:
	docker build -t $(BACKEND_IMAGE_NAME) ./backend

deploy-backend:
	docker run -e DB_USERNAME=$(DB_USERNAME) -e DB_PASSWORD=$(DB_PASSWORD) -e DB_IP=$(DB_IP) -e DB_PORT=$(DB_PORT) -e DATABASE=$(DATABASE) -p $(BACKEND_PORT):$(BACKEND_PORT) --name $(BACKEND_CONTAINER_NAME) $(BACKEND_IMAGE_NAME)

build-frontend:
	docker build -t $(FRONTEND_IMAGE_NAME) ./frontend

deploy-frontend:
	docker run -p $(FRONTEND_PORT):$(FRONTEND_PORT) --name $(FRONTEND_CONTAINER_NAME) $(FRONTEND_IMAGE_NAME)