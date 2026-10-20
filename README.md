# Groupie Tracker

Groupie Tracker is a web application that allows users to explore and track their favorite music artists. The app provides detailed information about artists, including their creation date, first album, members, concert locations, and dates. It is designed to offer an interactive and visually appealing experience for music enthusiasts.

---

## Table of Contents

- [About the Project](#about-the-project)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributors](#contributors)
- [License](#license)

---

## About the Project

Groupie Tracker is a school project developed by students of Zone01 Rouen Normandie. The goal of the project is to create a functional and visually appealing web application that fetches and displays data about music artists using an external API.

The application is built with **Go** for the backend and uses **HTML**, **CSS**, and **JavaScript** for the frontend. It emphasizes clean design, responsive layouts, and efficient data handling.

---

## Features

- **Artist Cards**: Display a grid of artists with their names and images.
- **Detailed Artist Information**: View detailed information about an artist, including:
  - Creation date
  - First album
  - Members
  - Concert locations and dates
- **Interactive Popups**: Open artist details in a popup without reloading the page.
- **Error Handling**: Custom error pages for better user experience.
- **Responsive Design**: Optimized for desktop, tablet, and mobile devices.

---

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS, JavaScript
- **API**: [Groupie Trackers API](https://groupietrackers.herokuapp.com/)
- **Templates**: Go's `text/template` package
- **Styling**: CSS with responsive design principles

---

## Getting Started

Follow these steps to set up and run the project locally.

### Prerequisites

- Go 1.20+ installed on your machine
- A web browser (e.g., Chrome, Firefox)
- Internet connection to fetch data from the API

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/groupie-tracker.git
   cd groupie-tracker
   ```

2. Install dependencies (if any).

3. Run the application:

   ```bash
    go run main.go
   ```

4. Open your browser and navigate to:

   ```bash
    http://localhost:8080
   ```

## Usage

1. Home Page:

    - Displays a grid of artist cards.
    - Click on a card to view detailed information about the artist.

2. Artist Details:

    - View information such as creation date, first album, members, and concert details.
    - Use the close button or click outside the popup to return to the home page.

3. Error Handling:

    - If an artist is not found or an error occurs, a custom error page will be displayed.

## Project Structure

```Text
groupie-tracker/
├── assets/                 # Static assets (CSS, images, videos)
│   ├── css/
│   ├── images/
│   └── videos/
├── handlers/               # Go handlers for HTTP requests
│   ├── index.go
│   ├── artist.go
│   └── httpclient.go
├── models/                 # Data models
│   ├── artists.go
│   └── relations.go
├── templates/              # HTML templates
│   ├── index.html
│   ├── artist.html
│   └── error.html
├── main.go                 # Entry point of the application
└── README.md               # Project documentation
```

## Contributors

This project was developed by:

- Nathan L.
- Killian H.
- Johann M.

Special thanks to Zone01 Rouen Normandie for their guidance and support.

## License

This project is licensed under the MIT License. You are free to use, modify, and distribute this project as long as proper credit is given.

## Acknowledgments

- Zone01 Rouen Normandie for providing the learning environment.
- Groupie Trackers API for the data used in this project.
- Open-source libraries and tools that made this project possible.
