'use client';

import ReactMarkdown from 'react-markdown';
import { SearchResponse } from '@/types';
import { ExternalLink, Globe, MessageCircle } from 'lucide-react';

interface SearchResultsProps {
  results: SearchResponse;
  onRelatedQuestionClick: (question: string) => void;
}

export function SearchResults({ results, onRelatedQuestionClick }: SearchResultsProps) {
  return (
    <div className="max-w-4xl mx-auto space-y-6">
      {/* AI Answer */}
      <div className="bg-white dark:bg-gray-800 rounded-xl shadow-lg p-6 border border-gray-100 dark:border-gray-700">
        <div className="flex items-center gap-2 mb-4">
          <MessageCircle className="w-5 h-5 text-primary-500" />
          <h2 className="text-lg font-semibold text-gray-900 dark:text-white">
            AI Answer
          </h2>
        </div>
        <div className="markdown-content prose prose-gray dark:prose-invert max-w-none">
          <ReactMarkdown>{results.answer}</ReactMarkdown>
        </div>
      </div>

      {/* Sources */}
      <div className="bg-white dark:bg-gray-800 rounded-xl shadow-lg p-6 border border-gray-100 dark:border-gray-700">
        <div className="flex items-center gap-2 mb-4">
          <Globe className="w-5 h-5 text-primary-500" />
          <h2 className="text-lg font-semibold text-gray-900 dark:text-white">
            Sources
          </h2>
        </div>
        <div className="space-y-4">
          {results.sources.map((source, index) => (
            <a
              key={index}
              href={source.url}
              target="_blank"
              rel="noopener noreferrer"
              className="block p-4 rounded-lg border border-gray-200 dark:border-gray-600 hover:border-primary-500 hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-all duration-200"
            >
              <div className="flex items-start justify-between gap-4">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <span className="flex-shrink-0 w-6 h-6 flex items-center justify-center bg-primary-100 dark:bg-primary-900/30 text-primary-600 dark:text-primary-400 rounded text-sm font-medium">
                      {index + 1}
                    </span>
                    <h3 className="font-medium text-gray-900 dark:text-white truncate">
                      {source.title}
                    </h3>
                  </div>
                  <p className="text-sm text-gray-600 dark:text-gray-400 line-clamp-2">
                    {source.description}
                  </p>
                  <p className="text-xs text-gray-400 dark:text-gray-500 mt-1 truncate">
                    {source.url}
                  </p>
                </div>
                <ExternalLink className="flex-shrink-0 w-4 h-4 text-gray-400" />
              </div>
            </a>
          ))}
        </div>
      </div>

      {/* Related Questions */}
      {results.related_questions && results.related_questions.length > 0 && (
        <div className="bg-white dark:bg-gray-800 rounded-xl shadow-lg p-6 border border-gray-100 dark:border-gray-700">
          <h2 className="text-lg font-semibold text-gray-900 dark:text-white mb-4">
            Related Questions
          </h2>
          <div className="space-y-2">
            {results.related_questions.map((question, index) => (
              <button
                key={index}
                onClick={() => onRelatedQuestionClick(question)}
                className="w-full text-left p-3 rounded-lg border border-gray-200 dark:border-gray-600 hover:border-primary-500 hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-all duration-200"
              >
                <span className="text-gray-700 dark:text-gray-300">
                  {question}
                </span>
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
