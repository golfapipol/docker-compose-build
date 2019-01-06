# Docker compose build

ถ้าเราใส่ build ลงใน docker-compose แล้วเราสั่ง

```bash
docker-compose build
```

จะทำให้ทุก service ที่มี tag build ถูก build
เหมาะสำหรับทำงานในเครื่องตัวเอง จะได้ไม่ต้อง build มือ

```yaml
stub-service:
    container_name: stub-webservice
    image: stub-service:latest # <--- อาจจะต้องหา version มาอ้างอิงแต่ถ้าในเครื่องตัวเองใช้ tag latest ก็ได้
    build:
      context: stubby # <--- folder ที่ทำการสั่ง docker build 
      dockerfile: Dockerfile_stub # <--- Dockerfile กรณีมีหลาย dockerfile ต้องเขียนระบุ
```

จากนั้นเราก็สั่ง up ได้เลย เพราะ image พร้อมหมดแล้ว
```bash
docker-compose up
```