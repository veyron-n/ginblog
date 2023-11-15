# postman_json文件
{
	"info": {
		"_postman_id": "318317f2-ecbb-4a20-a28f-03db804ab3be",
		"name": "ginblog",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "28172129"
	},
	"item": [
		{
			"name": "user",
			"item": [
				{
					"name": "新增用户",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"jack\",\r\n    \"password\": \"123456\",\r\n    \"role\": 1 \r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/user/add",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"user",
								"add"
							]
						}
					},
					"response": []
				},
				{
					"name": "查询用户列表",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/users?username=admin&pageSize=5&pageNum=1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"users"
							],
							"query": [
								{
									"key": "username",
									"value": "admin"
								},
								{
									"key": "pageSize",
									"value": "5"
								},
								{
									"key": "pageNum",
									"value": "1",
									"description": "分页: 当前页码"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "删除用户",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/user/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"user",
								"2"
							]
						}
					},
					"response": []
				},
				{
					"name": "编辑用户",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"helloworld\",\r\n    \"role\": 2\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/user/3",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"user",
								"3"
							]
						}
					},
					"response": []
				},
				{
					"name": "登录",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"admin_0\",\r\n    \"password\": \"123456\"\r\n}   ",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/login",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"login"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "category",
			"item": [
				{
					"name": "新增分类",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"name\": \"golang\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/category/add",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"category",
								"add"
							]
						}
					},
					"response": []
				},
				{
					"name": "查询分类列表",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/categorys?pageSize=5&pageNum=1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"categorys"
							],
							"query": [
								{
									"key": "pageSize",
									"value": "5",
									"description": "分页: 每页个数"
								},
								{
									"key": "pageNum",
									"value": "1",
									"description": "分页: 当前页码"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "删除分类",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/category/1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"category",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "编辑分类",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"name\": \"分类2\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/category/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"category",
								"2"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "article",
			"item": [
				{
					"name": "新增文章",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"title\": \"标题7\",\r\n    \"cid\": 1,\r\n    \"desc\": \"描述7\",\r\n    \"context\": \"正文7\",\r\n    \"img\": \"img7\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/article/add",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"article",
								"add"
							]
						}
					},
					"response": []
				},
				{
					"name": "分页查询文章列表",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/articles?pageSize=5&pageNum=1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"articles"
							],
							"query": [
								{
									"key": "pageSize",
									"value": "5",
									"description": "分页: 每页个数"
								},
								{
									"key": "pageNum",
									"value": "1",
									"description": "分页: 当前页码"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "查询单个文章",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/article/info/1",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"article",
								"info",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "查询分类下的所有文章",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/article/list/1?pageSize=2&pageNum=2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"article",
								"list",
								"1"
							],
							"query": [
								{
									"key": "pageSize",
									"value": "2",
									"description": "分页: 每页个数"
								},
								{
									"key": "pageNum",
									"value": "2",
									"description": "分页: 当前页"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "删除文章",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/article/5",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"article",
								"5"
							]
						}
					},
					"response": []
				},
				{
					"name": "编辑文章",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTA3MTQ1NTUsImlzcyI6ImdpbmJsb2cifQ.re6EpN8TZ9UMgnjOSkbJHxHJO5W0JjLPhVIwWLC5GTU",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"title\": \"golang\",\r\n    \"cid\": 2,\r\n    \"desc\": \"描述123\",\r\n    \"context\": \"正文123\",\r\n    \"img\": \"img123\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/article/2",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"article",
								"2"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "upload",
			"item": [
				{
					"name": "文件上传",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIiLCJleHAiOjE2OTYzMzY2OTAsImlzcyI6ImdpbmJsb2cifQ.q_ltvMRCNnCVmvmsaGuY2RUner1v8XJhcFjXweeF8Uk",
									"type": "string"
								}
							]
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "formdata",
							"formdata": [
								{
									"key": "file",
									"type": "file",
									"src": "/C:/Users/CH/Desktop/img2.jpg"
								}
							]
						},
						"url": {
							"raw": "http://127.0.0.1:3000/api/v1/upload",
							"protocol": "http",
							"host": [
								"127",
								"0",
								"0",
								"1"
							],
							"port": "3000",
							"path": [
								"api",
								"v1",
								"upload"
							]
						}
					},
					"response": []
				}
			]
		}
	]
}
