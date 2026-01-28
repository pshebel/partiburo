import * as React from 'react'
import './index.css'
import { createRoot } from 'react-dom/client'

// import { TanStackDevtools } from '@tanstack/react-devtools'
// import { formDevtoolsPlugin } from '@tanstack/react-form-devtools'
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

import App from './App.tsx'
console.log("version 0")
const rootElement = document.getElementById('root')!

// const queryClient = new QueryClient()
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // The function receives the failure count and the error object
      retry: (failureCount, error: any) => {
        const status = error?.status || error?.code;
        if (status === 404 || status === 400) {
          return false;
        }

        // 1. Always stop after a certain number of retries for other errors
        if (failureCount >= 3) return false;

        
        return true; 
      },
    },
  },
});

createRoot(rootElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <App />
    </QueryClientProvider>
{/* 
    <TanStackDevtools
      config={{ hideUntilHover: true }}
      plugins={[formDevtoolsPlugin()]}
    /> */}
  </React.StrictMode>,
)