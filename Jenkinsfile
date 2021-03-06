pipeline {
    agent {
        docker {
            image 'golang'
        }
    }
    environment { 
        GIT_COMMITTER_NAME = 'jenkins'
        GIT_COMMITTER_EMAIL = 'jenkins@localhost'
        GIT_AUTHOR_NAME = 'jenkins'
        GIT_AUTHOR_EMAIL = 'jenkins@localhost'
        GOPATH = "${WORKSPACE}"
    }
    stages {
        stage('Test') {
            steps {
                sh '''
                    cd ${WORKSPACE}
                    go get github.com/golang/lint/golint
                    go get github.com/tebeka/go2xunit
                    go get github.com/YaleSpinup/eventreporter/...
                    ./bin/golint src/github.com/YaleSpinup/eventreporter/.. > lint.txt
                    go test -v $(go list github.com/YaleSpinup/eventreporter/... | grep -v /vendor/) | ./bin/go2xunit -output tests.xml
                '''
            }
            post {
                success {
                    stash includes: 'lint.txt,test.xml', name: 'reaperTests'
                }
            }
        }
    }
    options {
        buildDiscarder(logRotator(numToKeepStr:'3'))
        timeout(time: 60, unit: 'MINUTES')
    }
}
