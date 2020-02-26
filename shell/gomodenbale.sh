go env -w GO111MODULE=on #이거 켜면 go get 으로 가저오는게 모두 mod 하위로 들어감. src 하위로 go get 가져올땐 끄고 가져오길.
go env -w GOPROXY=direct
go env -w GOSUMDB=off