module test-rebrain-repo-lib

go 1.22.2

require (
	utils v1.1.0
	utils/v3 v3.0.0
)

replace (
	utils v1.1.0 => gitlab.rebrainme.com/golang_users_repos/5807/utils.git v1.1.0
	utils/v3 v3.0.0 => gitlab.rebrainme.com/golang_users_repos/5807/utils.git/v3 v3.0.0
)
