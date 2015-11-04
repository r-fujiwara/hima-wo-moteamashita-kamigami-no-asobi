require "net/http"
require "logger"

uri = URI("http://localhost")
port = "8000"
logger = Logger.new(STDOUT)

Net::HTTP.start(uri.host, port) do |http|
  request = Net::HTTP::Get.new uri

  response = http.request request # Net::HTTPResponse object
  puts response
end
