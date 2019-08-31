package main

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHandler"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func setEnv() {
	const (
		Ip           = ""
		Port         = "9213"
		Prefix       = "/v1.0/job_reg"
		LogPath      = "job_reg.log"
		WriteTimeout = "4"

		JobRequestTopic      = "cjob-request"
		JobResponseTopic     = "cjob-response"
		ConnectRequestTopic  = "ConnectRequest"
		ConnectResponseTopic = "ConnectResponse"
		TmAggRequestTopic    = "TmAggRequest"
		TmAggResponseTopic   = "TmAggResponse"

		MqttUrl     = "http://59.110.31.215:6542/v0/publish"
		MqttChannel = "test-qi/"

		RedisHost = "59.110.31.215"
		RedisPort = "6378"
		RedisPwd  = ""

		KafkabRokerUrl      = "123.56.179.133:9092"
		SchemaRepositoryUrl = "http://123.56.179.133:8081"
		KafkaGroup          = "job-reg"
		CaLocation          = "/opt/kafka/pharbers-secrets/snakeoil-ca-1.crt"
		CASignedLocation    = "/opt/kafka/pharbers-secrets/kafkacat-ca1-signed.pem"
		SSLKeyLocation      = "/opt/kafka/pharbers-secrets/kafkacat.client.key"
		SSLPwd              = "pharbers"
	)

	_ = os.Setenv("IP", Ip)
	_ = os.Setenv("PORT", Port)
	_ = os.Setenv("PREFIX", Prefix)
	_ = os.Setenv("LOG_PATH", LogPath)
	_ = os.Setenv("WRITE_TIMEOUT", WriteTimeout)

	_ = os.Setenv("CONNECT_REQUEST_TOPIC", ConnectRequestTopic)
	_ = os.Setenv("CONNECT_RESPONSE_TOPIC", ConnectResponseTopic)
	_ = os.Setenv("JOB_REQUEST_TOPIC", JobRequestTopic)
	_ = os.Setenv("JOB_RESPONSE_TOPIC", JobResponseTopic)
	_ = os.Setenv("TMAGG_REQUEST_TOPIC", TmAggRequestTopic)
	_ = os.Setenv("TMAGG_RESPONSE_TOPIC", TmAggResponseTopic)

	_ = os.Setenv("MQTT_URL", MqttUrl)
	_ = os.Setenv("MQTT_CHANNEL", MqttChannel)

	_ = os.Setenv("REDIS_HOST", RedisHost)
	_ = os.Setenv("REDIS_PORT", RedisPort)
	_ = os.Setenv("REDIS_PWD", RedisPwd)

	_ = os.Setenv("BM_KAFKA_BROKER", KafkabRokerUrl)
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", SchemaRepositoryUrl)
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", KafkaGroup)
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", CaLocation)
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", CASignedLocation)
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", SSLKeyLocation)
	_ = os.Setenv("BM_KAFKA_SSL_PASS", SSLPwd)
}

func main() {
	setEnv()

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	prefix := os.Getenv("PREFIX")
	writeTimeoutInt, _ := strconv.Atoi(os.Getenv("JOB_REQUEST_TOPIC"))
	writeTimeout := time.Second * time.Duration(writeTimeoutInt)

	jobResponseTopic := os.Getenv("JOB_RESPONSE_TOPIC")
	connectResponseTopic := os.Getenv("CONNECT_RESPONSE_TOPIC")
	tmAggResponseTopic := os.Getenv("TMAGG_RESPONSE_TOPIC")

	schemaRepositoryUrl := os.Getenv("BM_KAFKA_SCHEMA_REGISTRY_URL")
	mqttUrl := os.Getenv("MQTT_URL")
	mqttChannel := os.Getenv("MQTT_CHANNEL")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPwd := os.Getenv("REDIS_PWD")

	kh := PhHelper.PhKafkaHelper{}.New(schemaRepositoryUrl)
	mh := PhMqttHelper.PhMqttHelper{}.New(mqttUrl, mqttChannel)
	rh := PhHelper.PhRedisHelper{}.New(redisHost, redisPort, redisPwd)

	// 协程启动 Kafka Consumer
	go kh.Linster([]string{connectResponseTopic}, &(PhModel.ConnectResponse{}), PhHandler.ConnectResponseHandler(kh, mh, rh), mh)
	go kh.Linster([]string{jobResponseTopic}, &(PhModel.JobResponse{}), PhHandler.JobResponseHandler(kh, mh, rh), mh)
	go kh.Linster([]string{tmAggResponseTopic}, &(PhModel.TmAggResponse{}), PhHandler.TmAggResponseHandler(kh, mh, rh), mh)

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
