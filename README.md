## ⚙️ Project Overview (Smart e-Banking System)

**Owned by:** MotionProgramming (Md Abdullah Rajeeb)

---

## 🏦 SEBS (Smart e-Banking System)

### 📌 Introduction

**Smart e-Banking System (SEBS)** is a **Go (Golang)-based web application** designed to provide secure and efficient online banking services. Traditional banking systems are often limited to physical branches or outdated interfaces. SEBS automates banking operations with a fast, modern backend in Go, secure user management, real-time transactions, and an interactive frontend, making banking accessible anytime, anywhere.

---

### 🎯 Objectives

* Store and manage bank account and transaction data securely.
* Enable users to **view balance**, **transfer funds**, and **view transaction history**.
* Allow admins to **manage users, apply interest**, and monitor system activity.
* Reduce manual errors and paperwork through a smart, digital system.
* Provide secure login and **role-based access control** for system integrity.

---

### 🛠 Tools & Technologies

* **Programming Language:** Go (Golang)
* **Frontend:** HTML, CSS, Bootstrap, JavaScript
* **Backend:** Go (`Gin` or `Echo` framework)
* **Database:** SQL (MySQL or PostgreSQL recommended)
* **Authentication:** JWT (JSON Web Tokens)
* **Visualization:** Chart.js
* **Testing:** Postman (for API testing)

---

### ✨ Features

#### 🔑 User Authentication

* Secure login and registration system.
* Supports multiple user roles: Admin and Customer.
* JWT-based session handling for security.

#### 💳 Banking Operations

* View account details and current balance.
* Transfer funds between customer accounts.
* View transaction history with filters (date, amount, account).

#### 💰 Interest Management

* Admin can set interest rates per account.
* Automatically or manually apply interest monthly.
* Logs all interest payments for transparency.

#### 📊 Reporting

* Dashboard with financial statistics and user activity.
* Generate charts for transaction volume, interest paid, etc.
* Downloadable reports (CSV/PDF - optional).

#### ⚡ JavaScript Enhancements

* Form validation for inputs like amount, date, and login.
* Real-time feedback on transfers and balance updates.
* Dynamic dashboard with charts powered by Chart.js.

---

### 🖥 System Design

* **Input:** Users interact with forms for login, transfers, and account views.
* **Process:**

  * Go backend validates inputs, processes logic, and accesses the database using GORM.
  * Business logic applies interest and handles transactions securely.
* **Output:** Data is returned via JSON and rendered dynamically on the frontend using JavaScript.

---

### 📂 Database Design

#### Users Table

* `user_id` (Primary Key)
* `username`
* `password` (hashed)
* `role` (Admin / Customer)

#### Accounts Table

* `account_id` (Primary Key)
* `user_id` (Foreign Key)
* `balance`
* `account_type` (Savings / Checking)
* `interest_rate`

#### Transactions Table

* `transaction_id` (Primary Key)
* `from_account`
* `to_account`
* `amount`
* `timestamp`
* `description`

#### InterestHistory Table

* `interest_id` (Primary Key)
* `account_id`
* `interest_amount`
* `applied_on`

---

### 🚀 Workflow

1. User logs into the system via the frontend.
2. Form inputs are sent to the backend through HTTP POST requests.
3. Backend (written in Go) validates the input, handles business logic, and interacts with the SQL database using GORM.
4. Transactions and interest are logged and updated in the database.
5. The frontend dynamically displays the results and updates the UI using JavaScript and Chart.js.

---

### 📑 Expected Outcomes

* Easy and secure access to digital banking features.
* Accurate and fast transaction processing.
* Reliable interest calculation and application.
* Data-driven insights for admins via reports and analytics.

---

### 📌 How to Run

```bash
# Clone the repository
git clone https://github.com/MotionProgramming/Smart-e-Banking-System.git
cd Smart-e-Banking-System
```

1. Open the project in VS Code or GoLand.
2. Import the database from the `database/sebs.sql` file.
3. Run the Go server:

   ```bash
   go mod tidy
   go run main.go
   ```
4. Open your browser at:
   👉 `http://localhost:8080/`

---

### 📮 API Testing with Postman

* Use the provided Postman Collection (`/postman/SEBS-Collection.json`)
* Test all major endpoints:

  * `POST /login`
  * `POST /register`
  * `GET /accounts`
  * `POST /transfer`
  * `GET /transactions`
  * `POST /apply-interest`

---

### 👨‍💻 User Roles

* **Admin**

  * Create/manage users and accounts
  * Set interest rates
  * View all system transactions
  * Monitor system usage

* **Customer**

  * View personal account and balance
  * Transfer funds
  * View transaction history

---

### 📄 License

This project is for **educational purposes only** and not intended for real banking deployment without security and compliance verification.

---

Would you like the **Postman collection template** or the **API documentation format (Swagger)?** I can generate those next.
