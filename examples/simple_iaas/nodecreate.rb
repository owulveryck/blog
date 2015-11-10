require 'msgpack/rpc'
class MyHandler
    def NodeCreate(kind, size, disksize, leasedays, environmenttype, description) 
        print "Creating the node with parameters: ",kind, size, disksize, leasedays, environmenttype, description
        return "ok"
    end
end
svr = MessagePack::RPC::Server.new
svr.listen('0.0.0.0', 18800, MyHandler.new)
svr.run
