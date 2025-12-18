# Vote-project
# **Vote Project – Favorite Programming Language Voting App**

**Moringa AI Capstone Project**

---

## **1. Project Overview**

This project is a simple voting application built using **Go (Golang)** for the backend and static **HTML/CSS** for the frontend. Users can vote for their favorite programming language from a set of options (Go, Python, JavaScript, Java). The app tracks votes **in memory**, displays results with animated bars, and allows resetting the votes.

**Key Features:**
- Server-side routing using Go.
- In-memory vote counting (no database required).
- Voting page with logos and interactive UI.
- Results page with animated gradient bars.
- Reset functionality to restart voting.
- Fully responsive and styled with gradient backgrounds and hover effects.

---

## **2. System Requirements**

You will need:

- **Operating System:** Linux, macOS, or Windows.  
- **Go Compiler:** Installed from [https://go.dev/dl](https://go.dev/dl)  
- **Text Editor:** VS Code (recommended)  
- **Terminal:** Linux Terminal, PowerShell, or Command Prompt  
- **Browser:** Any modern browser (Chrome, Firefox, Edge)  

---

## **3. Installing Go on Linux**

Open a terminal and run the following commands:

Update package list

sudo apt update

Install Go

sudo apt install golang-go -y

Verify installation

go version


**Optional:** Install the latest Go manually:

wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz

sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
go version


---

## **4. Clone the Repository**

git clone <your-repo-url>
cd vote-project


The project folder structure should look like this:

vote-project/
│
├─ Backend/
│ └─ hello.go
│
├─ Frontend/
│ ├─ vote.html
│ ├─ results.html
│ ├─ style.css
│ └─ images/
│ ├─ go.svg
│ ├─ python.svg
│ ├─ javascript.png
│ └─ java.svg


---

## **5. Running the Project**

1. Open a terminal in the `Backend` folder:

cd Backend
go run hello.go

2. Open your browser and navigate to:
http://localhost:8080


3. Start voting and view results by clicking **"See Results"**.  

4. Reset votes anytime using the **Reset Votes** button.  

---

## **6. AI Prompt Journal (Learning Experiences)**

**Prompt 1:**  
*"Explain Go programming language in simple terms for a beginner."*  
**AI Summary:** Go is easy to learn, fast, and widely used for backend systems.  
**Reflection:** Helped me understand why Go is suitable for server-side routing and simple backend apps.

**Prompt 2:**  
*"Give me a minimal Go example that serves a static HTML page."*  
**AI Summary:** Showed how to use Go's `http` package and `FileServer` to serve HTML/CSS files.  
**Reflection:** Learned server-side routing and how to link frontend and backend.

**Prompt 3:**  
*"Why is my CSS not loading when served from Go?"*  
**AI Summary:** Explained the need for `http.StripPrefix("/static/", fs)` to properly serve static assets.  
**Reflection:** Fixed styling issues for both voting and results pages.

**Prompt 4:**  
*"List common beginner errors when serving static files in Go."*  
**AI Summary:** Errors include incorrect paths, missing StripPrefix, and incorrect URLs.  
**Reflection:** Helped me troubleshoot faster and maintain consistent styling.

**Prompt 5:**  
*"How can I make the voting page visually appealing?"*  
**AI Summary:** Suggested gradient backgrounds, hover animations, and responsive grid layout for options.  
**Reflection:** Implemented dynamic hover effects and gradients for a modern UI.

**Prompt 6:**  
*"How can I show vote results visually?"*  
**AI Summary:** Recommended using progress bars or width-based divs with dynamic colors to represent vote counts.  
**Reflection:** Added animated bars for Go, Python, JavaScript, and Java with matching colors.

**Prompt 7:**  
*"How do I reset the votes safely without a database?"*  
**AI Summary:** Explained clearing an in-memory map and redirecting to the main page.  
**Reflection:** Learned how to implement reset functionality with server-side concurrency safety using a mutex.

**Prompt 8:**  
*"How can I make the vote selection intuitive?"*  
**AI Summary:** Suggested adding clickable labels with images and hiding radio buttons.  
**Reflection:** Improved the voting experience by letting users click images to select options.

**Prompt 9:**  
*"How do I make the app responsive?"*  
**AI Summary:** Explained CSS grid layouts and media queries to adjust the option cards for smaller screens.  
**Reflection:** Ensured mobile-friendly layout for phones and tablets.

**Prompt 10:**  
*"How can I show a more realistic 'live' vote result?"*  
**AI Summary:** Recommended dynamically calculating percentages and animating bars based on vote counts.  
**Reflection:** Implemented animated width bars that grow according to votes.

**Prompt 11:**  
*"How can I make the buttons more engaging?"*  
**AI Summary:** Suggested hover transitions, rounded corners, and gradient backgrounds.  
**Reflection:** Enhanced the "Vote", "See Results", and "Reset" buttons for better UX.

**Prompt 12:**  
*"How can I add multiple voting options for learning purposes?"*  
**AI Summary:** Proposed adding more languages, images, and option labels.  
**Reflection:** Extended project with multiple choices to simulate a real voting scenario.

---

## **7. Concepts Learned**

- Server-side routing in Go.  
- Serving HTML, CSS, and static assets.  
- In-memory data management for votes.  
- Linking backend logic with frontend interactivity.  
- Gradient and animated styling for UI consistency.  
- Responsive design with CSS Grid and media queries.  
- Using AI to guide learning step-by-step.  

---

## **8. References**

- [Official Go Documentation](https://go.dev/doc/)  
- [Go by Example](https://gobyexample.com/)  
- [W3Schools Go Tutorial](https://www.w3schools.com/go)  
- [DigitalOcean Go Guides](https://www.digitalocean.com/community/tags/go)




