# coding: utf-8
namespace :APNsGoServer do

  desc 'build'
  task :build do
    sh('GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/apnsgoserver')
    sh('GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/apnsgoserver')
  end

  desc "run example"
  task :run_example do
    sh('go run main.go -c config.toml')
  end


end
