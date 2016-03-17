require 'rake/testtask'
require 'minitest/autorun'
require 'codeclimate-test-reporter'
CodeClimate::TestReporter.start

Rake::TestTask.new do |t|
  t.test_files = FileList['ruby/*/*test.rb']
end
desc 'Run tests associated with ruby exercises'

task default: 'test'
