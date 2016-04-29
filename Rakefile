require 'rake/testtask'
require 'codeclimate-test-reporter'
CodeClimate::TestReporter.start

Rake::TestTask.new do |task|
  task.libs << 'ruby'
  task.pattern = 'ruby/*/*test.rb'
end
desc 'Run tests in test folder'

task default: 'test'
