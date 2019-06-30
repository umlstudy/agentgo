
# 맵 선언

func test() {
    var serverInfoMap = map[string]ServerInfo{}
	serverInfoMap["mysvr"] = ServerInfo{"mysvr", "mysvr", []ResourceStatus{
		ResourceStatus{"cpu", 1, 100, "cpu", 1},
		ResourceStatus{"mem", 1, 100, "memory", 1},
		ResourceStatus{"di1", 1, 100, "disk1", 1},
		ResourceStatus{"di2", 1, 100, "disk2", 1},
	}}
}