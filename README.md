# Go_basic

## เริ่มต้น (Intro)
1.  สร้าง Go Modules ขึ้นมา (go.mod) เอาจัดการ Package 
    ด้วยการสั่ง go mod init pratice_101 ใน terminal
2.  เชื่อมต่อกับ GitHub ผ่าน command line  
        echo "# Go_basic" >> README.md  
        git init  
        git add README.md  
        git commit -m "first commit"  
        git branch -M main  
        git remote add origin https://github.com บลาๆๆ  
        git push -u origin main  
3.  โครงสร้างของคำสั่ง Go 
    - main() ทำงานอันแรกที่ Go จะเรียกใช้ โดยอัตโนมัติ  
    - ขอบเขต ใช้ {} ในการกำหนดขอบเขตของฟังก์ชัน  
    - comment ใช้ // เพื่อใส่ข้อความในโค้ด หรือ /* */ เพื่อใส่ comment แบบหลายบรรทัด

## ชนิดข้อมูลพื้นฐาน (Data Type)

| ชนิดข้อมูล | คำอธิบาย | Zero Value |
| :--- | :--- | :---: |
| `bool` | ค่าทางตรรกศาสตร์ (true/false) | `false` |
| `int` `int8` `int16` `int32` `int64`<br>`uint` `uint8` `uint16` `uint32` `uint64`<br>`uintptr` | ตัวเลขที่ไม่มีจุดทศนิยม | `0` |
| `float32` `float64` | ตัวเลขที่มีจุดทศนิยม | `0` |
| `string` | ชุดข้อความ | `""` |

> **ยิ่งจำนวนของ bit มากเท่าไร แสดงว่าเราสามารถเก็บค่าได้มากเท่านั้น**

## นิยามตัวแปร (Variable)
1.  Manual Type 
  ```
  var name string 
  name = "Zen"
  ```
2. Type Inference
  ```
  name := "Zen"
  ```
