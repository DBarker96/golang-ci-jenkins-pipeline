pipeline {
    agents none
    environment {
        CI = 'true'
    }
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.9.2'            
                }
            }
            steps {     
                sh 'go version'          
                sh 'go build'   
            }            
        }
        stage('Go Tests') {
            agent {
                docker {
                    image 'golang:1.9.2'            
                }
            }
            steps {
                sh 'go test -v'
            }
        }
        stage('Mocha Tests') {
            agent {
                docker {
                    image 'usemtech/nodejs-mocha:latest'
                }
            }
        }
        stage('Deliver') {
            agent {
                docker {
                    image 'golang:1.9.2'            
                }
            }
            steps {
                sh 'go run sum.go'
            }
        }
    }
}
    
