export interface SearchResult {
  title: string;
  url: string;
  description: string;
  content?: string;
}

export interface SearchResponse {
  query: string;
  answer: string;
  sources: SearchResult[];
  related_questions?: string[];
}

export interface SearchRequest {
  query: string;
}
