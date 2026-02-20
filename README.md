# 🎓 Student API with Gin Framework

 REST API สำหรับจัดการข้อมูลนักเรียน พัฒนาด้วยภาษา **Go** โดยใช้ **Gin Framework** และเก็บข้อมูลในฐานข้อมูล **SQLite** 
 ระบบถูกออกแบบตามหลัก Layered Architecture (Handler, Service, Repository, Model) เพื่อความง่ายต่อการบำรุงรักษา

## 📦 Project Structure

```
go-api-gin/
├─ main.go
├─ models/
├─ repositories/
├─ services/
├─ handlers/
├─ config/
└─ students.db
```

## 🚀 How to Run

### Prerequisites
---
1.  **Go (Golang):** เวอร์ชัน 1.20 ขึ้นไป
    * 👉 https://go.dev/dl
2.  **Git:** สำหรับใช้คำสั่ง clone โปรเจกต์
3.  **SQLite3:** สำหรับจัดการไฟล์ฐานข้อมูล `students.db`
4.  **Postman** สำหรับทดสอบ API 

---

## 🚀 ขั้นตอนการเริ่มทำงาน (Getting Started)

1.  **Clone โปรเจกต์:**
    ```bash
    git clone https://github.com/NgerNgaEiiei/go-api-gin-lab.git
    cd go-api-gin-lab
    ```

2.  **Install dependencies**
    ```bash
    go mod tidy
    ```

3.  **Run the server**
    ```bash
    go run main.go
    ```
    Server will start at: http://localhost:8080

## 🧪 การทดสอบ API ด้วย Postman (Testing Guide)

### Get ready before Test API
**Postman**
1. **ตั้งค่า URL**: `http://localhost:8080`
2. **ตั้งค่า Header** (สำหรับ POST และ PUT):
   - Key: `Content-Type`
   - Value: `application/json`

---
### 1. Get All Students
**ดึงข้อมูลนักศึกษาทั้งหมด**

```http
GET /students
```
**How to test API with Postman:**
1. เลือก Method `GET`
2. ใส่ URL: `http://localhost:8080/students`
3. กด `Send`

**Response ที่คาดหวัง:**
```json
{
  "id": "6609650152",
  "name": "NgerNga Eiei",
  "major": "Computer Science",
  "gpa": 3.00
}
```

---
### 2. Get Student by ID
**ดึงข้อมูลนักศึกษาด้วย ID**

```http
GET /students/:id
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `GET`
2. ใส่ URL: `http://localhost:8080/students/6609650152`
3. กด `Send`

**Response สำเร็จ (200 OK):**
```json
{
  "id": "6609650152",
  "name": "NgerNga Eiei",
  "major": "Computer Science",
  "gpa": 3.00
}
```

**Response กรณีที่ไม่มีข้อมูลนักศึกษาหมายเลขนั้น:**
```json
{
  "error": "Student not found"
}
```

---

### 3. Create Student 
**เพิ่มข้อมูลนักศึกษาใหม่**

```http
POST /students
```

 **ขั้นตอนใน Postman:**
1. เลือก Method: `POST`
2. ใส่ URL: `http://localhost:8080/students`
3. ไปที่แท็บ `Headers`
   - Key: `Content-Type`
   - Value: `application/json`
4. ไปที่แท็บ `Body`
   - เลือก `raw`
   - เลือกประเภท `JSON`
5. ใส่ JSON ด้านล่าง
**ใส่ Request Body:**
```json
{
  "id": "6609650999",
  "name": "สวัสดี คนไทย",
  "major": "ชอบใส่ใจ",
  "gpa": 4
}
```
6. กด `Send`

**Response สำเร็จ (201 Created):**
```json
{
  "id": "6609650999",
  "name": "สวัสดี คนไทย",
  "major": "ชอบใส่ใจ",
  "gpa": 4
}
```


**Validation Errors (400 Bad Request):**

**Response กรณีไม่ใส่ ID :**
```json
{
  "error": "id must not be empty"
}
```

**Response กรณีไม่ใส่ Name :**
```json
{
  "error": "name must not be empty"
}
```

**Response กรณี GPA ไม่อยู่ในช่วง 0-4:**
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

---

### 4. Update Student ✅
**แก้ไขข้อมูลนักศึกษาที่มีอยู่แล้ว**

```http
PUT /students/:id
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `PUT`
2. ใส่ URL: `http://localhost:8080/students/6609650152`
3. ไปที่แท็บ `Headers`
   - Key: `Content-Type`
   - Value: `application/json`
4. ไปที่แท็บ `Body`
   - เลือก `raw`
   - เลือกประเภท `JSON`
5. ใส่ JSON ด้านล่าง
**Request Body:**
```json
{
  "id": "6609650152",
  "name": "เงอะงะ ง้องแง้ง",
  "major": "ลูกไก่",
  "gpa": 4.0
}
```
6. กด `Send`

**Response สำเร็จ (200 OK):**
```json
{
  "id": "6609650152",
  "name": "เงอะงะ ง้องแง้ง",
  "major": "ลูกไก่",
  "gpa": 4.0
}
```

**Error Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

**Validation Errors (400 Bad Request):**

**Response กรณีไม่ใส่ ID :**
```json
{
  "error": "id must not be empty"
}
```

**Response กรณีไม่ใส่ Name :**
```json
{
  "error": "name must not be empty"
}
```

**Response กรณี GPA ไม่อยู่ในช่วง 0-4:**
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

### 5. Delete Student ✅
**ลบข้อมูลนักศึกษา (ต้องแน่ใจว่ามี ID นี้อยู่ในฐานข้อมูล)**

```http
DELETE /students/:id
```

**ขั้นตอนใน Postman:**
1. เลือก Method: `DELETE`
2. ใส่ URL: `http://localhost:8080/students/6609650999`
3. กด `Send`

**Response สำเร็จ (204 No Content):**
- ไม่มี Response Body
- HTTP Status: `204 No Content`

**Error Response (404 Not Found):**
```json
{
  "error": "Student not found"
}
```

---

### 6. **Test Validation Errors:**

#### 🧪 Test Case: Empty ID

ขั้นตอนใน Postman:
> 1. เลือก Method: `POST`
> 2. URL: `http://localhost:8080/students`
> 3. Headers: `Content-Type: application/json`
> 4. Body (raw JSON):

```json
{
  "id": "",
  "name": "John Doe",
  "major": "Computer Science",
  "gpa": 3.50
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "id must not be empty"
}
```

#### 🧪 Test Case: Empty Name

**Body:**
```json
{
  "id": "6609650999",
  "name": "",
  "major": "Test",
  "gpa": 3.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "name must not be empty"
}
```

---

#### 🧪 Test Case: GPA > 4.00

**Body:**
```json
{
  "id": "6609650999",
  "name": "Test",
  "major": "Test",
  "gpa": 5.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

---

#### 🧪 Test Case: GPA < 0.00

**Body:**
```json
{
  "id": "6609650999",
  "name": "Test",
  "major": "Test",
  "gpa": -1.0
}
```

**Expected Response (400 Bad Request):**
```json
{
  "error": "gpa must be between 0.00 and 4.00"
}
```

---







   







