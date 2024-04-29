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
                    echo "This is $BRANCH_NAME branch"
                }
            }
        }

        stage("deploy") {
            when {
                expression {
                    BRANCH_NAME == 'main'
                }
            }
            steps {
                script {
                    echo "deploy go apllication"
                }
            }
        }
    }
}