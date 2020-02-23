node {
    def app
    def dockerfile = '2.0.1'
    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("polla:${dockerfile}")
    }
    stage('Clean up basura') {
        sh 'docker rmi $(docker images --filter "dangling=true" -q --no-trunc)'
    }
}
