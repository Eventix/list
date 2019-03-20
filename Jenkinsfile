node {
  stage('Prep') {
    checkout scm
  }

  stage('Tests') {
    parallel 'Test': {
      docker.image('golang-base').inside("-v ${env.WORKSPACE}:/go/src/eventix.io/list") { go ->
        sh "cd /go/src/eventix.io/list; go test -count=1 -v ./... | tee /tmp/testoutput.log"
        sh "cat /tmp/testoutput.log | go-junit-report > /go/src/eventix.io/list/results/report.xml"
      }
      junit 'results/*.xml'
    }, 'Checks': {
      docker.image('golang-base').inside("-v ${env.WORKSPACE}:/go/src/eventix.io/list") { go ->
        sh "cd /go/src/eventix.io/list; staticcheck ./..."
      }
    }
  }
}
