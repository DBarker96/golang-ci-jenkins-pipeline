pipeline {
    agent {
        docker {
            image 'golang:1.9.2'
        }
    }
    environment {
        CI = 'true'
    }
    stages {
        stage('Build') {
            steps {     
                sh 'go version'          
                sh 'go build'   
                } 
            
        }
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
        stage('Deliver') {
            steps {
                sh 'go run sum.go'
            }
        }
    }
}
    
