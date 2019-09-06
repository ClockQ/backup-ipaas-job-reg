package main

import (
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhChannel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHandler"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhThirdHelper"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	ProjName = "job-reg"
)

func main() {
	//PhEnv.SetEnv()

	_ = os.Setenv("BP_LOG_TIME_FORMAT", "2006-01-02 15:04:05")

	if ok := os.Getenv("PROJECT_NAME"); ok == "" {
		_ = os.Setenv("PROJECT_NAME", ProjName)
	}

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	prefix := os.Getenv("PREFIX")
	writeTimeoutInt, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	writeTimeout := time.Second * time.Duration(writeTimeoutInt)

	//connectResponseTopic := os.Getenv("CONNECT_RESPONSE_TOPIC")
	jobResponseTopic := os.Getenv("JOB_RESPONSE_TOPIC")
	//tmAggResponseTopic := os.Getenv("TMAGG_RESPONSE_TOPIC")

	schemaRepositoryUrl := os.Getenv("BM_KAFKA_SCHEMA_REGISTRY_URL")
	mqttUrl := os.Getenv("MQTT_URL")
	mqttChannel := os.Getenv("MQTT_CHANNEL")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPwd := os.Getenv("REDIS_PWD")

	kh := PhChannel.PhKafkaHelper{}.New(schemaRepositoryUrl)
	mh := PhThirdHelper.PhMqttHelper{}.New(mqttUrl, mqttChannel)
	rh := PhThirdHelper.PhRedisHelper{}.New(redisHost, redisPort, redisPwd)

	// 协程启动 Kafka Consumer
	//go kh.Linster([]string{connectResponseTopic}, &(PhModel.ConnectResponse{}), PhHandler.ConnectResponseHandler(kh, mh, rh))
	go kh.Linster([]string{jobResponseTopic}, &(PhModel.JobResponse{}), PhHandler.JobResponseHandler(kh, mh, rh))
	//go kh.Linster([]string{tmAggResponseTopic}, &(PhModel.TmAggResponse{}), PhHandler.TmAggResponseHandler(kh, mh, rh))

	/// 下面不用管，网上抄的
	// 主动关闭服务器
	addr := ip + ":" + port
	mux := http.NewServeMux()
	mux.HandleFunc(prefix, PhHandler.PhHttpHandler(kh, mh, rh))

	// 一个通知退出的chan
	var server *http.Server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	server = &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		Handler:      mux,
	}
	bpLog := log.NewLogicLoggerBuilder().Build()

	bpLog.Info("Starting httpserver in " + port)
	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			bpLog.Error("Close server:", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			bpLog.Error("Server closed under request")
		} else {
			bpLog.Error("Server closed unexpected", err)
		}
	}
	bpLog.Error("Server exited")
}
