# 🚽 UWaterLoo Locator: Campus Washroom Discovery Platform

A robust and performant application helping UWaterloo students locate and review campus washrooms. Get real-time information about the nearest facilities, complete with crowd-sourced reviews and status updates.

## 🏗 Architecture

The application follows a modern microservices architecture with:

- **Frontend**: React Native mobile application
- **Backend**: Go-based microservices using:
  - Gin web framework for RESTful API
  - CQRS pattern for efficient read/write operations
  - Event-driven architecture for real-time updates
  - Event Sourcing for state management
  - Repository pattern for data access
  - SQLite for persistent storage (easily scalable to PostgreSQL)
- **Location Services**: Google Maps API integration

## 🎨 Design Patterns

- **CQRS (Command Query Responsibility Segregation)**: Separate read and write operations through specialized interfaces
- **Repository Pattern**: Abstract data persistence through WashroomRepository and LocationQueryRepository
- **Event Sourcing**: Track state changes through immutable events using the Event interface
- **Aggregate Pattern**: Encapsulate domain logic and ensure consistency using Aggregate interface
- **Interface Segregation**: Clean separation of concerns in repository and service layers
- **Dependency Injection**: Used throughout services for loose coupling
- **Factory Pattern**: Service and repository creation through New* functions

## 🚀 Key Features

- Real-time washroom location discovery
- Proximity-based search with customizable radius
- Building and floor-level washroom filtering
- Advanced search with amenity filtering
- Crowd-sourced review system
- Real-time status updates and issue reporting
- Cross-platform mobile support

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

## 🔧 Containerization

The application uses Docker for containerization:

- Multi-stage builds for optimized image sizes
- Non-root user execution for enhanced security
- Persistent volume mounts for database storage
- Environment-based configuration

### Running the Service

```bash
# Using Docker Compose
docker-compose up --build

# Or using Docker directly
docker build -t washroom-service ./washroom-data-service
docker run -p 8080:8080 washroom-service
```

## 🌟 Coming Soon

- OAuth2 authentication
- Advanced analytics dashboard
- Push notifications for status updates
- Accessibility features
- Cross-platform desktop support

## 👥 Contributors

- [@JustColdToast](https://github.com/JustColdToast)
- Original concept inspired by SacHacks 2022 project: [Loo-Locator](https://github.com/SippinOnJuiceBox/Loo-Locator-Find-the-nearest-washroom)

## 📝 License

MIT License - See LICENSE file for details

