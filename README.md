# image-service
图片服务器

## How to use

### Install

#### 克隆项目
> $ git clone https://github.com/yue-best-practices/image-service.git

#### 进入项目文件夹
> $ cd image-service

#### 制作docker镜像

> $ make docker

#### 清理临时文件

> $ make clean

#### 配置服务

修改`docker-compose.yml` 文件，配置服务启动端口`APP_PORT`以及图片存储路径`IMAGE_PATH`

#### 启动服务

> $ make run 

#### 停止服务

> $ make stop

## API

### 上传文件

#### Request

Url: `/upload`

Method: `POST`

Headers:

| key | value |
| --- | --- | 
| Content-Type | `multipart/form-data` |

Body:

| key | type |
| --- | --- |
| `image` | `File` image file (`jpg`/`jepg`/`gif`/`png`)


#### Response

```json
{
  "code": 200,
  "data": {
    "image": "MjAxOS8wMi8xMC9hM2IwMjgxY2IyNGE0ZDA3OGNmNjYzMjBjMGI1NGYzOS5naWY="
  },
  "msg": "Success"
}
```


### 访问图片

#### Request

Url: `/:image`(上传图片接口返回的`image`值)

Method: `GET`