import * as React from 'react'
import './index.css'
import { createRoot } from 'react-dom/client'

import { TanStackDevtools } from '@tanstack/react-devtools'
import { formDevtoolsPlugin } from '@tanstack/react-form-devtools'
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

import App from './App.tsx'

const rootElement = document.getElementById('root')!

const queryClient = new QueryClient()

createRoot(rootElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <App />
    </QueryClientProvider>

    <TanStackDevtools
      config={{ hideUntilHover: true }}
      plugins={[formDevtoolsPlugin()]}
    />
  </React.StrictMode>,
)