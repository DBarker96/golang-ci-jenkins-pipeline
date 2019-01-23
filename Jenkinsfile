pipeline {
    agents {
        agent {
            docker {
                image 'golang:1.9.2'            
            }
        }
        agent {
            docker {
            image 'usemtech/nodejs-mocha:latest' 
            }        
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
                sh 'go test -v'
            }
        }
        stage('Deliver') {
            steps {
                sh 'go run sum.go'
            }
        }
    }
}
    
