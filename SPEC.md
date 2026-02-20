# SPEC.md: bsalbilla06.com

## Project Overview
A lightweight, fun, and functional personal website for `bsalbilla06.com`. The site uses a unique "hand-drawn" aesthetic, appearing as pencil on graph paper.

## Tech Stack (GoTH)
- **Backend:** Go (Standard Library preferred)
- **Frontend:** HTMX
- **Styling:** Tailwind CSS (Custom fonts and borders for hand-drawn look)
- **Templating:** Go `html/template`
- **Deployment:** Docker, AWS (EC2, RDS, S3, Cognito)

## Project Structure (Standard Go Layout)
- `/cmd/server/`: Main application entry point.
- `/internal/`: Private application and library code.
    - `/handlers/`: HTTP request handlers.
    - `/middleware/`: HTMX and Auth middleware.
    - `/models/`: Data structures.
- `/web/`:
    - `/template/`: Go HTML templates.
    - `/static/`: CSS (Tailwind), JS, and hand-drawn assets.
- `/configs/`: Configuration files.
- `/build/`: Dockerfile and deployment scripts.

## Core Features & Pages
1. **Landing/Home Page:** Introduction with a personal touch.
2. **About Me Page:** Bio and interests.
3. **Experience Page:** Career history and projects.
4. **Contact Page:** Links to social media pages, email, and phone number.

## Design System
- **Background:** Graph paper pattern (CSS grid or SVG).
- **Typography:** Pencil-style handwritten font (e.g., "Architects Daughter" or similar).
- **UI Components:**
    - **Cards:** White "paper" background with subtle drop shadows and "tape" (rotated rectangles) at the corners.
    - **Navbar:** Handwritten text; active page circled with a "red crayon" effect (SVG overlay).
    - **Assets:** Pencil-drawn icons and sketches (placeholders initially).

## Technical Requirements
- **HTMX:** Used for seamless page transitions and form submissions without full page reloads.
- **Security:** 
    - No hardcoded secrets (use `os.Getenv`).
    - AWS Cognito for authentication if an admin dashboard is added.
    - Principle of Least Privilege for AWS IAM roles.
- **Storage:** RDS for structured data (e.g., experience entries), S3 for images/assets.
- **Containerization:** Multistage Dockerfile for optimized Go builds.

## Deployment Pipeline
1. Build Go binary and Tailwind CSS.
2. Build Docker image.
3. Push to AWS ECR (optional) and deploy to EC2.
4. Configure RDS and S3 access.
