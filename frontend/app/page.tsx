'use client';

import { useState } from 'react';
import { SearchBar } from '@/components/SearchBar';
import { SearchResults } from '@/components/SearchResults';
import { SearchResponse } from '@/types';
import { search } from '@/lib/api';
import { Sparkles } from 'lucide-react';

export default function Home() {
  const [isLoading, setIsLoading] = useState(false);
  const [results, setResults] = useState<SearchResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSearch = async (query: string) => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await search(query);
      setResults(response);
    } catch (err) {
      setError('Failed to perform search. Please try again.');
      console.error('Search error:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleRelatedQuestion = (question: string) => {
    handleSearch(question);
  };

  return (
    <main className="min-h-screen bg-gradient-to-b from-gray-50 to-white dark:from-gray-900 dark:to-gray-800">
      <div className="container mx-auto px-4 py-8">
        {/* Header */}
        <div className={`transition-all duration-500 ${results ? 'mb-8' : 'mb-16 pt-20'}`}>
          <div className={`text-center ${results ? '' : 'mb-12'}`}>
            <div className="flex items-center justify-center gap-3 mb-4">
              <Sparkles className="w-10 h-10 text-primary-500" />
              <h1 className="text-4xl font-bold text-gray-900 dark:text-white">
                AI Search
              </h1>
            </div>
            {!results && (
              <p className="text-gray-600 dark:text-gray-300 text-lg">
                Get intelligent answers powered by AI
              </p>
            )}
          </div>

          <div className={`max-w-3xl mx-auto ${results ? '' : 'mt-8'}`}>
            <SearchBar onSearch={handleSearch} isLoading={isLoading} />
          </div>
        </div>

        {/* Error State */}
        {error && (
          <div className="max-w-3xl mx-auto mb-8">
            <div className="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
              <p className="text-red-600 dark:text-red-400">{error}</p>
            </div>
          </div>
        )}

        {/* Results */}
        {results && (
          <SearchResults
            results={results}
            onRelatedQuestionClick={handleRelatedQuestion}
          />
        )}

        {/* Empty State */}
        {!results && !isLoading && !error && (
          <div className="max-w-2xl mx-auto text-center">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-12">
              <SuggestionCard
                title="What is machine learning?"
                onClick={() => handleSearch("What is machine learning?")}
              />
              <SuggestionCard
                title="How does photosynthesis work?"
                onClick={() => handleSearch("How does photosynthesis work?")}
              />
              <SuggestionCard
                title="Best practices for React"
                onClick={() => handleSearch("Best practices for React")}
              />
              <SuggestionCard
                title="History of the internet"
                onClick={() => handleSearch("History of the internet")}
              />
            </div>
          </div>
        )}
      </div>
    </main>
  );
}

function SuggestionCard({ title, onClick }: { title: string; onClick: () => void }) {
  return (
    <button
      onClick={onClick}
      className="p-4 text-left bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-primary-500 hover:shadow-md transition-all duration-200"
    >
      <span className="text-gray-700 dark:text-gray-300">{title}</span>
    </button>
  );
}
