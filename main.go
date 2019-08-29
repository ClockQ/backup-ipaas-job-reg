package main

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHandler"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMessage"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// TODO 统统环境变量
var ip = ""
var port = "9213"
var prefix = "/v1.0/job_reg"
var WriteTimeout = time.Second * 4

const (
	LogPath              = "job_reg.log"
	KafkabRokerUrl       = "123.56.179.133:9092"
	SchemaRepositoryUrl  = "http://123.56.179.133:8081"
	KafkaGroup           = "test20190828"
	CaLocation           = "/opt/kafka/pharbers-secrets/snakeoil-ca-1.crt"
	CASignedLocation     = "/opt/kafka/pharbers-secrets/kafkacat-ca1-signed.pem"
	SSLKeyLocation       = "/opt/kafka/pharbers-secrets/kafkacat.client.key"
	SSLPwd               = "pharbers"
	JobRequestTopic      = "cjob-test"
	JobResponseTopic     = "cjob-test2"
	ConnectRequestTopic  = "ConnectRequest"
	ConnectResponseTopic = "ConnectResponse"
	MqttUrl              = "http://59.110.31.215:6542/v0/publish"
	MqttChannel          = "test-qi/"
)

func main() {
	_ = os.Setenv("LOG_PATH", LogPath)
	_ = os.Setenv("BM_KAFKA_BROKER", KafkabRokerUrl)
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", SchemaRepositoryUrl)
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", KafkaGroup)
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", CaLocation)
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", CASignedLocation)
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", SSLKeyLocation)
	_ = os.Setenv("BM_KAFKA_SSL_PASS", SSLPwd)
	_ = os.Setenv("JOB_REQUEST_TOPIC", JobRequestTopic)
	_ = os.Setenv("CONNECT_REQUEST_TOPIC", ConnectRequestTopic)

	if ok := os.Getenv("REG_PORT"); ok != "" {
		port = ok
	}
	if ok := os.Getenv("REG_PREFIX"); ok != "" {
		prefix = ok
	}
	if ok := os.Getenv("LOG_PATH"); ok == "" {
		_ = os.Setenv("LOG_PATH", LogPath)
	}

	mh := PhMessage.PhMqttHandler{}.New(MqttUrl, MqttChannel)
	kh := PhMessage.PhKafkaHandler{}.New(SchemaRepositoryUrl)

	// 协程启动 Kafka Consumer
	go kh.Linster([]string{JobResponseTopic}, &(PhModel.JobResponse{}), PhHandler.JobResponseHandler(mh))
	go kh.Linster([]string{ConnectResponseTopic}, &(PhModel.ConnectResponse{}), PhHandler.ConnectResponseHandler(mh))

	/// 下面不用管，网上抄的
	// 主动关闭服务器
	addr := ip + ":" + port
	mux := http.NewServeMux()
	mux.HandleFunc(prefix, PhHandler.PhHttpHandler(kh, mh))

	// 一个通知退出的chan
	var server *http.Server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	server = &http.Server{
		Addr:         addr,
		WriteTimeout: WriteTimeout,
		Handler:      mux,
	}

	bmlog.StandardLogger().Info("Starting httpserver in " + port)
	log.Println("Starting httpserver in " + port)
	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			bmlog.StandardLogger().Error("Close server:", err)
			log.Fatal("Close server:", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			bmlog.StandardLogger().Error("Server closed under request")
			log.Fatal("Server closed under request")
		} else {
			bmlog.StandardLogger().Error("Server closed unexpected", err)
			log.Fatal("Server closed unexpected", err)
		}
	}
	bmlog.StandardLogger().Error("Server exited")
	log.Fatal("Server exited")
}
