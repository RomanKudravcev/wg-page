pipeline {
    agent any

    environment {
        DB_USERNAME_CREDENTIAL = credentials('DB_USERNAME_CREDENTIAL')
        DB_PASSWORD_CREDENTIAL = credentials('DB_PASSWORD_CREDENTIAL')
        DB_IP_CREDENTIAL = credentials('DB_IP_CREDENTIAL')
        DB_PORT_CREDENTIAL = credentials('DB_PORT_CREDENTIAL')
        DATABASE_CREDENTIAL = credentials('DATABASE_CREDENTIAL')

        BACKEND_IMAGE_NAME = "backend"
        BACKEND_CONTAINER_NAME = "backend_application"
        BACKEND_PORT = "8080"

        FRONTEND_IMAGE_NAME=frontend
        FRONTEND_CONTAINER_NAME=frontend_application
        FRONTEND_PORT=80
    }

 stages {
        stage('Build Backend') {
            steps {
                script {
                    echo "Building backend..."
                    sh "make build-backend"
                }
            }
        }

        stage('Deploy Backend') {
            steps {
                script {
                     withCredentials([
                        string(credentialsId: 'DB_USERNAME_CREDENTIAL', variable: 'DB_USERNAME'),
                        string(credentialsId: 'DB_PASSWORD_CREDENTIAL', variable: 'DB_PASSWORD'),
                        string(credentialsId: 'DB_IP_CREDENTIAL', variable: 'DB_IP'),
                        string(credentialsId: 'DB_PORT_CREDENTIAL', variable: 'DB_PORT'),
                        string(credentialsId: 'DATABASE_CREDENTIAL', variable: 'DATABASE')
                    ]) {
                        echo "Deploying backend..."
                        sh "make deploy-backend"
                    }
                }
            }
        }

        stage('Build Frontend') {
            steps {
                script {
                    echo "Building frontend..."
                    sh "make build-frontend"
                }
            }
        }

        stage('Deploy Frontend') {
            steps {
                script {
                    echo "Deploying frontend..."
                    sh "make deploy-frontend"
                }
            }
        }
    }
}