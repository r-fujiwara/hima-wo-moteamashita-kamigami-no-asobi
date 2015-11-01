# copy paste from http://morizyun.github.io/blog/parallel-process-ruby-gem/
require "net/http"
require 'open-uri'
require 'thread'

url = "http://localhost:8000"

urls = []
100000.times{|i| urls.push(url) }
urls.push(nil)

q = Queue.new
urls.each { |url| q.push(url) }

max_thread = 8

Array.new(max_thread) do |i|
  Thread.new {
    begin
      while url = q.pop(true)
        puts "start request: #{url}\n"
        resp = open(url) rescue next
        puts resp.read
        puts "end request: #{url}\n"
     end
     q.push nil
   end
  }
end.each(&:join)
puts "finish process"
