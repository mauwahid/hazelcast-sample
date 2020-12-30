package main

import "github.com/hazelcast/hazelcast-go-client"

func main() {
	mapSampleRun()
}


func mapSampleRun() {
	// Start the Hazelcast Client and connect to an already running Hazelcast Cluster on 127.0.0.1
	hz, _ := hazelcast.NewClient()
	// Get the Distributed Map from Cluster.
	mp, _ := hz.GetMap("myDistributedMap")
	//Standard Put and Get.
	mp.Put("key", "value")
	mp.Get("key")
	//Concurrent Map methods, optimistic updating
	mp.PutIfAbsent("somekey", "somevalue")
	mp.ReplaceIfSame("key", "value", "newvalue")
	// Shutdown this hazelcast client

	hz.Shutdown()
}