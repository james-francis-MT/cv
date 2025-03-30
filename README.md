# CV Website (Go + HTMX)

🚀 A personal CV website built using **Go, HTMX, and Docker**. This project serves as an online portfolio to showcase projects, experience, and contact details in a lightweight and interactive way.

## Features
- ⚡ **Go Backend** – Fast and efficient web server
- 🌐 **HTMX** – Smooth, JavaScript-free interactivity
- 🎨 **TailwindCSS** – Modern styling
- 🗄 **PostgreSQL** – Database for dynamic content
- 🐳 **Dockerized** – Easy deployment and scaling
- 🔐 **Admin Panel** – Secure project & contact management

## Tech Stack
| Technology    | Purpose          |
|--------------|----------------|
| Go          | Backend API & Templating |
| HTMX        | Frontend Interactivity |
| Tailwind CSS | Styling |
| PostgreSQL  | Database |
| Docker      | Containerization |
| Nginx       | Reverse Proxy (Prod) |

## Getting Started
### 1. Clone the Repository
```sh
git clone https://github.com/james-francis-MT/cv.git
cd cv
```

### 2. Build & Run with Docker
```sh
docker-compose up --build
```

### 3. Open in Browser
- Visit `http://localhost:8000`

## Project Structure
```
cv-website/
│── cmd/                  # Main entry point (main.go)
│── internal/             # Business logic
│   ├── handlers/         # Route handlers
│   ├── models/           # Database models
│   └── templates/        # HTML templates
│── static/               # CSS, images, JS
│── config/               # Config settings
│── Dockerfile            # Docker build instructions
│── docker-compose.yml    # Multi-container setup
│── go.mod                # Go dependencies
│── README.md             # Project documentation
```

## Roadmap
- [x] Setup Go HTTP server
- [x] Dockerize application
- [ ] Static pages
- [ ] Dynamic project showcase
- [ ] Contact form with database storage
- [ ] Admin panel for managing content

## Deployment
TODO

## License
MIT License © 2025 James Francis


