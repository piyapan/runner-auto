FROM gitlab/gitlab-runner
ADD runner-auto /usr/bin/runner-auto
ENTRYPOINT ["/usr/bin/runner-auto"]
