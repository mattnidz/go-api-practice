  pipeline {
        agent any

        stages {
            stage('Build') {
                steps {
                    echo 'Building for None ...'
		    go build
                }
            }
            stage('Test') {
                steps {
                    echo 'Testing for None..'
		    go test
                }
            }
            stage('Deploy to K8s') {
                steps {
                    echo 'Deploying for None...'
                }
          }
    }
    
