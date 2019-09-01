# ipaas-producer

# 依赖关系
名称 | 依赖 | 介绍
:---: | :---: | :---: 
PhEnv | | 设置环境变量
PhModel | | 定义项目实体，以及各种构建实体的方法
PhThirdHelper | | Redis、MQTT、OSS
PhChannel | PhModel | Kafka
PhPanic | PhModel、PhThirdHelper(mqtt) | 异常处理
PhJobManager | PhModel、PhThirdHelper、PhChannel、PhPanic | 任务注册，任务执行
PhHandler | PhModel、PhThirdHelper、PhChannel、PhPanic、PhJobManager | 业务逻辑拼接
