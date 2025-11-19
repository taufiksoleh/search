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

## Development

### Demo Mode

The application includes mock data functionality that works without API keys configured. This is useful for development and testing.

### Adding New Features

1. **New Search Providers**: Add new search methods in `backend/services/search.go`
2. **New AI Providers**: Add new AI integrations in `backend/services/ai.go`
3. **UI Components**: Add new components in `frontend/components/`

## License

MIT
