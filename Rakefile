require 'rake/testtask'
require 'minitest/autorun'

Rake::TestTask.new do |t|
  t.test_files = FileList['ruby/*/*test.rb']
end
desc 'Run tests associated with ruby exercises'

task default: 'test'
