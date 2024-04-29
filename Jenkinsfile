pipeline{
    agent any
    tools {
        go 'Golang'
    }
    stages {
        stage("build") {
            steps {
                script {
                    echo "building go apllication"
                    sh 'go version'
                }
            }
        }

        stage("deploy") {
            steps {
                script {
                    echo "deploy go apllication"
                }
            }
        }
    }
}