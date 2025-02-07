# 🚽 UWaterLoo Locator: Find Your Nearest Washroom

A cozy project built by me to help UWaterloo students find and review campus washrooms. Whether you’re in a rush or just curious, this app gives you real-time updates on washroom locations and reviews from people just like you.

## 🏗 How It's Built

I put together a modern mix of tech that’s both powerful and fun to work with:

- **Frontend**: A React Native app that feels right at home on your phone.
- **Backend**: Go-powered microservices with:
  - Gin for a straightforward RESTful API
  - A CQRS setup to keep things neat and efficient
  - Event-driven magic for those real-time alerts
  - Event Sourcing to track every little change
  - A Repository pattern to keep data tidy
  - SQLite for storage (with an easy path to PostgreSQL if needed)
- **Location Services**: Easily find places via the Google Maps API

## 🎨 Cool Design Choices

I kept things simple but smart:

- **CQRS**: Separating commands and queries so everything runs smooth.
- **Repository Pattern**: Abstracting data details away.
- **Event Sourcing**: Every change gets its own moment in time.
- **Aggregate Pattern and Interface Segregation**: Keeping the code clear and focused.
- **Dependency Injection & Factory Pattern**: For a flexible, easy-to-test setup

## 🚀 Features You’ll Love

- Quick, real-time washroom searches.
- Customizable search radius to fit your pace.
- Filters for building and floor levels.
- Advanced search options including amenities.
- Community reviews so you know what to expect.
- Up-to-date status and issue reporting.
- Cross-platform support that keeps you connected.

## 🛠 Technical Stack

### Backend
- Go 1.23.6
- Gin Web Framework
- SQLite (with support for PostgreSQL migration)
- Event Sourcing
- Clean Architecture principles
- Docker containerization
- Docker Compose for orchestration

### Infrastructure
- Docker multi-stage builds for optimized images
- Alpine-based containers for minimal footprint
- Container health checks
- Docker volumes for persistent storage
- Docker networks for service isolation

### Frontend
- React Native
- Google Maps SDK
- TypeScript
- Native device location services

## 🔧 Running the App

A few quick commands and you're ready to go:

```bash
# Fire up using Docker Compose
docker-compose up --build

# Or do it manually
docker build -t washroom-service ./washroom-data-service
docker run -p 8080:8080 washroom-service
```

## 🌟 What’s Next?

- OAuth2 authentication for extra security.
- Real-time push alerts.
- More accessibility features.
- Desktop support on the horizon.
