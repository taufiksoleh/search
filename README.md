# AI Search Engine

An AI-powered search engine similar to Perplexity, built with a Go backend and Next.js frontend.

## Features

- **AI-Powered Answers**: Get intelligent, summarized answers to your queries
- **Source Citations**: View the sources used to generate answers
- **Related Questions**: Discover related topics to explore
- **Web Scraping**: Automatically fetches and analyzes content from search results
- **Modern UI**: Clean, responsive design with dark mode support

## Architecture

```
search/
├── backend/          # Go backend server
│   ├── handlers/     # HTTP request handlers
│   ├── services/     # Business logic services
│   ├── models/       # Data models
│   └── main.go       # Entry point
├── frontend/         # Next.js frontend
│   ├── app/          # App router pages
│   ├── components/   # React components
│   ├── lib/          # Utility functions
│   └── types/        # TypeScript types
└── README.md
```

## Prerequisites

- Go 1.21+
- Node.js 18+
- npm or yarn

## Setup

### Backend

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Copy the environment file and configure your API keys:
   ```bash
   cp .env.example .env
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the server:
   ```bash
   make run
   # or
   go run main.go
   ```

The backend will start on `http://localhost:8080`

### Frontend

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Copy the environment file:
   ```bash
   cp .env.example .env.local
   ```

4. Run the development server:
   ```bash
   npm run dev
   ```

The frontend will start on `http://localhost:3000`

## API Configuration

### Search APIs (choose one)

- **Brave Search API**: Get your API key from [Brave Search API](https://brave.com/search/api/)
- **SerpAPI**: Get your API key from [SerpAPI](https://serpapi.com/)

### AI APIs (choose one)

- **OpenAI**: Get your API key from [OpenAI Platform](https://platform.openai.com/)
- **Anthropic Claude**: Get your API key from [Anthropic Console](https://console.anthropic.com/)

## API Endpoints

### POST /api/search

Search and get AI-generated answers.

**Request:**
```json
{
  "query": "What is machine learning?"
}
```

**Response:**
```json
{
  "query": "What is machine learning?",
  "answer": "Machine learning is a subset of artificial intelligence...",
  "sources": [
    {
      "title": "Machine Learning - Wikipedia",
      "url": "https://en.wikipedia.org/wiki/Machine_learning",
      "description": "Machine learning is a field of study..."
    }
  ],
  "related_questions": [
    "What are types of machine learning?",
    "How does deep learning differ from machine learning?"
  ]
}
```

### GET /api/health

Health check endpoint.

## Docker Deployment

### Quick Start with Docker Compose

1. Copy the environment file and add your API keys:
   ```bash
   cp .env.example .env
   ```

2. Build and run with Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the application:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

### Production Deployment with Nginx

```bash
docker-compose --profile production up --build
```

This will also start an Nginx reverse proxy on port 80.

### Individual Container Builds

**Backend:**
```bash
cd backend
docker build -t search-backend .
docker run -p 8080:8080 --env-file ../.env search-backend
```

**Frontend:**
```bash
cd frontend
docker build -t search-frontend .
docker run -p 3000:3000 -e NEXT_PUBLIC_API_URL=http://localhost:8080 search-frontend
```

## Deployment Options

### Recommended Platforms

| Platform | Type | Best For |
|----------|------|----------|
| **Railway** | PaaS | Easiest deployment, auto-scaling |
| **Render** | PaaS | Free tier available, simple setup |
| **Fly.io** | PaaS | Global edge deployment |
| **DigitalOcean App Platform** | PaaS | Simple container deployment |
| **AWS ECS/Fargate** | Cloud | Enterprise, full control |
| **Google Cloud Run** | Cloud | Serverless containers |
| **Azure Container Apps** | Cloud | Microsoft ecosystem |
| **Kubernetes** | Self-hosted | Maximum control, complex setup |

### Railway Deployment

1. Connect your GitHub repository
2. Add environment variables in Railway dashboard
3. Railway auto-detects Dockerfiles and deploys

### Render Deployment

1. Create a new "Blueprint" from your repo
2. Render uses `render.yaml` or auto-detects services
3. Add environment variables in dashboard

### Fly.io Deployment

```bash
# Install flyctl
curl -L https://fly.io/install.sh | sh

# Deploy backend
cd backend
fly launch
fly secrets set OPENAI_API_KEY=your-key

# Deploy frontend
cd frontend
fly launch
```

## Development

### Demo Mode

The application includes mock data functionality that works without API keys configured. This is useful for development and testing.

### Adding New Features

1. **New Search Providers**: Add new search methods in `backend/services/search.go`
2. **New AI Providers**: Add new AI integrations in `backend/services/ai.go`
3. **UI Components**: Add new components in `frontend/components/`

## License

MIT
